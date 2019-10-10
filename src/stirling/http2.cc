#include "src/stirling/http2.h"

#include <arpa/inet.h>
extern "C" {
#include <nghttp2/nghttp2_frame.h>
#include <nghttp2/nghttp2_hd.h>
#include <nghttp2/nghttp2_helper.h>
}

#include <algorithm>
#include <map>
#include <utility>
#include <variant>
#include <vector>

#include "absl/strings/str_cat.h"
#include "src/common/base/error.h"
#include "src/common/base/status.h"
#include "src/common/grpcutils/utils.h"
#include "src/stirling/bcc_bpf_interface/grpc.h"
#include "src/stirling/grpc.h"

namespace pl {
namespace stirling {
namespace http2 {

using ::pl::grpc::MethodInputOutput;
using ::pl::grpc::MethodPath;
using ::pl::grpc::ServiceDescriptorDatabase;
using ::pl::stirling::grpc::kGRPCMessageHeaderSizeInBytes;

namespace {

bool IsPadded(const nghttp2_frame_hd& hd) { return hd.flags & NGHTTP2_FLAG_PADDED; }
bool IsEndHeaders(const nghttp2_frame_hd& hd) { return hd.flags & NGHTTP2_FLAG_END_HEADERS; }
bool IsEndStream(const nghttp2_frame_hd& hd) { return hd.flags & NGHTTP2_FLAG_END_STREAM; }

void UnpackData(const uint8_t* buf, Frame* f) {
  nghttp2_data& frame = f->frame.data;
  size_t frame_body_length = frame.hd.length;

  DCHECK(frame.hd.type == NGHTTP2_DATA) << "Must be a DATA frame, got: " << frame.hd.type;

  if (IsPadded(frame.hd)) {
    frame.padlen = *buf;
    ++buf;
    --frame_body_length;
  }

  if (frame_body_length >= frame.padlen) {
    f->u8payload.assign(buf, frame_body_length - frame.padlen);
  } else {
    LOG(DFATAL) << absl::Substitute(
        "Pad length cannot be larger than frame body, padlen: $0 bodylen: $1", frame.padlen,
        frame_body_length);
    f->u8payload.clear();
  }
}

void UnpackHeaders(const uint8_t* buf, Frame* f) {
  nghttp2_headers& frame = f->frame.headers;
  size_t frame_body_length = frame.hd.length;

  DCHECK(frame.hd.type == NGHTTP2_HEADERS) << "Must be a HEADERS frame, got: " << frame.hd.type;

  if (IsPadded(frame.hd)) {
    frame.padlen = *buf;
    ++buf;
    --frame_body_length;
  }

  nghttp2_frame_unpack_headers_payload(&frame, buf);

  const size_t offset = nghttp2_frame_headers_payload_nv_offset(&frame);
  buf += offset;
  frame_body_length -= offset;

  if (frame_body_length >= frame.padlen) {
    f->u8payload.assign(buf, frame_body_length - frame.padlen);
  } else {
    LOG(DFATAL) << absl::Substitute(
        "Pad length cannot be larger than frame body, padlen: $0 bodylen: $1", frame.padlen,
        frame_body_length);
    f->u8payload.clear();
  }
}

void UnpackContinuation(const uint8_t* buf, Frame* frame) {
  frame->u8payload.assign(buf, frame->frame.hd.length);
}

}  // namespace

/**
 * @brief Inflates a complete header block. If the input header block fragment, then the results are
 * undefined.
 *
 * This code follows: https://github.com/nghttp2/nghttp2/blob/master/examples/deflate.c.
 */
ParseState InflateHeaderBlock(nghttp2_hd_inflater* inflater, u8string_view buf, NVMap* nv_map) {
  // TODO(yzhao): Experiment continuous parsing of multiple header block fragments from different
  // HTTP2 streams.
  constexpr bool is_final = true;
  for (;;) {
    int inflate_flags = 0;
    nghttp2_nv nv = {};

    const int rv =
        nghttp2_hd_inflate_hd2(inflater, &nv, &inflate_flags, buf.data(), buf.size(), is_final);
    if (rv < 0) {
      return ParseState::kInvalid;
    }
    buf.remove_prefix(rv);

    if (inflate_flags & NGHTTP2_HD_INFLATE_EMIT) {
      nv_map->insert(
          std::make_pair(std::string(reinterpret_cast<const char*>(nv.name), nv.namelen),
                         std::string(reinterpret_cast<const char*>(nv.value), nv.valuelen)));
    }
    if (inflate_flags & NGHTTP2_HD_INFLATE_FINAL) {
      nghttp2_hd_inflate_end_headers(inflater);
      break;
    }
    if ((inflate_flags & NGHTTP2_HD_INFLATE_EMIT) == 0 && buf.empty()) {
      break;
    }
  }
  return ParseState::kSuccess;
}

std::string_view FrameTypeName(uint8_t type) {
  switch (type) {
    case NGHTTP2_DATA:
      return "DATA";
    case NGHTTP2_HEADERS:
      return "HEADERS";
    case NGHTTP2_PRIORITY:
      return "PRIORITY";
    case NGHTTP2_RST_STREAM:
      return "RST_STREAM";
    case NGHTTP2_SETTINGS:
      return "SETTINGS";
    case NGHTTP2_PUSH_PROMISE:
      return "PUSH_PROMISE";
    case NGHTTP2_PING:
      return "PING";
    case NGHTTP2_GOAWAY:
      return "GOAWAY";
    case NGHTTP2_WINDOW_UPDATE:
      return "WINDOW_UPDATE";
    case NGHTTP2_CONTINUATION:
      return "CONTINUATION";
    case NGHTTP2_ALTSVC:
      return "ALTSVC";
    case NGHTTP2_ORIGIN:
      return "ORIGIN";
    default:
      return "unknown";
  }
}

// frame{} zero initialize the member, which is needed to make sure default value is sensible.
Frame::Frame() : frame{} {}
Frame::~Frame() {
  if (frame.hd.type == NGHTTP2_HEADERS) {
    DCHECK(frame.headers.nva == nullptr);
    DCHECK_EQ(frame.headers.nvlen, 0u);
  }
}

ParseState UnpackFrame(std::string_view* buf, Frame* frame) {
  if (buf->size() < kFrameHeaderSizeInBytes) {
    return ParseState::kNeedsMoreData;
  }

  const uint8_t* u8_buf = reinterpret_cast<const uint8_t*>(buf->data());
  nghttp2_frame_unpack_frame_hd(&frame->frame.hd, u8_buf);
  // HTTP2 spec specifies a maximal frame size. Any frames that is larger than that should be
  // considered a parse failure, which usually is caused by losing track of the frame boundary.
  constexpr int kMaxFrameSize = (1 << 24) - 1;
  if (frame->frame.hd.length > kMaxFrameSize) {
    return ParseState::kInvalid;
  }

  if (buf->size() < kFrameHeaderSizeInBytes + frame->frame.hd.length) {
    return ParseState::kNeedsMoreData;
  }

  buf->remove_prefix(kFrameHeaderSizeInBytes + frame->frame.hd.length);
  u8_buf += kFrameHeaderSizeInBytes;

  const uint8_t type = frame->frame.hd.type;
  switch (type) {
    case NGHTTP2_DATA:
      UnpackData(u8_buf, frame);
      break;
    case NGHTTP2_HEADERS:
      UnpackHeaders(u8_buf, frame);
      break;
    case NGHTTP2_CONTINUATION:
      UnpackContinuation(u8_buf, frame);
      break;
    default:
      VLOG(1) << "Ignoring frame that is not DATA, HEADERS, or CONTINUATION, got: "
              << FrameTypeName(type);
      return ParseState::kIgnored;
  }
  frame->creation_timestamp = std::chrono::steady_clock::now();
  return ParseState::kSuccess;
}

namespace {

ParseState CheckGRPCMessage(std::string_view buf) {
  if (buf.size() < kGRPCMessageHeaderSizeInBytes) {
    return ParseState::kNeedsMoreData;
  }
  const uint32_t len = nghttp2_get_uint32(reinterpret_cast<const uint8_t*>(buf.data()) + 1);
  if (buf.size() < kGRPCMessageHeaderSizeInBytes + len) {
    return ParseState::kNeedsMoreData;
  }
  return ParseState::kSuccess;
}

}  // namespace

// Stitches header block frames, by appending all payload of the CONTINUATION frames of
// a header block into the block's first HEADERS frame. Also inflates the resultant header blocks.
//
// TODO(yzhao): If there is ever header parse failure because of missing data from any of
// the header blocks, all followup parsings are busted. We can record this state, and skip
// parsing for the rest of the stream.
//
// TODO(yzhao): If the header block parse failure is because of application logic errors,
// HTTP2 runtime would reset stream. We now ignore all such signalling frames.
// If needed, we need to handle all those.
void StitchAndInflateHeaderBlocks(nghttp2_hd_inflater* inflater, std::deque<Frame>* frames) {
  Frame* first_headers_frame = nullptr;

  auto reset_when_lost_frame_sync = [&first_headers_frame]() {
    first_headers_frame->headers_parse_state = ParseState::kInvalid;
    first_headers_frame->frame_sync_state = ParseState::kInvalid;
    first_headers_frame = nullptr;
  };

  auto coerce_continuation_to_headers = [&first_headers_frame](Frame& frame) {
    frame.frame.hd.type = NGHTTP2_HEADERS;
    frame.headers_parse_state = ParseState::kInvalid;
    frame.frame_sync_state = ParseState::kInvalid;

    first_headers_frame = &frame;
    if (IsEndHeaders(frame.frame.hd)) {
      first_headers_frame = nullptr;
    }
  };

  for (Frame& frame : *frames) {
    auto frame_type = frame.frame.hd.type;
    if (frame_type == NGHTTP2_HEADERS) {
      if (frame.headers_parse_state != ParseState::kUnknown) {
        // This HEADERS frame is already processed.
        continue;
      }
      if (first_headers_frame != nullptr) {
        LOG(WARNING) << absl::Substitute(
            "HEADERS frame follows another HEADERS frame. "
            "Indicates lost frame boundary. Stream ID: $0.",
            frame.frame.hd.stream_id);
        reset_when_lost_frame_sync();
      }
      first_headers_frame = &frame;
      if (IsEndHeaders(frame.frame.hd)) {
        frame.headers_parse_state = InflateHeaderBlock(inflater, frame.u8payload, &frame.headers);
        first_headers_frame = nullptr;
      }
      continue;
    }
    if (frame_type == NGHTTP2_CONTINUATION) {
      if (frame.consumed) {
        // This CONTINUATION frame is already processed.
        continue;
      }
      if (first_headers_frame == nullptr) {
        LOG(WARNING) << absl::Substitute(
            "There was no HEADERS frame before this CONTINUATION frame. "
            "Indicates lost frame boundary. This frame is coerced into a HEADERS frame. "
            "Stream ID: $0.",
            frame.frame.hd.stream_id);
        coerce_continuation_to_headers(frame);
        continue;
      }
      if (first_headers_frame->frame.hd.stream_id != frame.frame.hd.stream_id) {
        LOG(WARNING) << absl::Substitute(
            "CONTINUATION frame stream ID: $0 is different from "
            "the first HEADERS frame stream ID: $1.",
            frame.frame.hd.stream_id, first_headers_frame->frame.hd.stream_id);
        reset_when_lost_frame_sync();
        coerce_continuation_to_headers(frame);
        continue;
      }
      // CONTINUATION frame can be discarded after being appended to the first HEADERS frame.
      frame.consumed = true;
      // TODO(yzhao): Keep string_view of the frames' payload, and join them lastly.
      first_headers_frame->u8payload.append(frame.u8payload);
      if (IsEndHeaders(frame.frame.hd)) {
        if (first_headers_frame->headers_parse_state == ParseState::kUnknown) {
          // Only parse the header block if it has not been parsed yet.
          first_headers_frame->headers_parse_state = InflateHeaderBlock(
              inflater, first_headers_frame->u8payload, &first_headers_frame->headers);
        }
        first_headers_frame = nullptr;
      }
      continue;
    }
    if (frame_type == NGHTTP2_DATA) {
      if (first_headers_frame != nullptr) {
        LOG(WARNING) << absl::Substitute(
            "DATA frame follows a HEADERS frame. Indicates lost frame boundary. Stream ID: $0.",
            frame.frame.hd.stream_id);
        reset_when_lost_frame_sync();
      }
      continue;
    }
  }
}

/**
 * @brief Given a list of frames for one stream, stitches them together into gRPC request and
 * response messages. Also mark consumed frames, so the caller can destroy them afterwards.
 */
// TODO(yzhao): Turn this into a class that parse frames one by one.
ParseState StitchGRPCMessageFrames(const std::vector<const Frame*>& frames,
                                   std::vector<GRPCMessage>* msgs) {
  size_t data_block_size = 0;
  std::vector<const Frame*> data_block_frames;

  GRPCMessage msg;

  auto handle_end_stream = [&data_block_size, &data_block_frames, &msg, msgs]() {
    // HTTP2 spec allows the following frames sequence:
    // HEADERS frame, no END_HEADERS flag, with END_STREAM flag
    // 0 or more CONTINUATION frames, no END_HEADERS flag
    // CONTINUATION frame, with END_HEADERS flag
    //
    // Fortunately gRPC only set END_STREAM and END_HEADERS together, and never do the above
    // sequence.
    //
    // TODO(yzhao): Consider handle the above case.

    // Join the messages scattered in the frames.
    msg.message.reserve(data_block_size);
    for (auto* f : data_block_frames) {
      msg.message.append(reinterpret_cast<const char*>(f->u8payload.data()), f->u8payload.size());
    }
    msg.frames.insert(msg.frames.end(), data_block_frames.begin(), data_block_frames.end());
    data_block_frames.clear();
    data_block_size = 0;
    msg.parse_state = CheckGRPCMessage(msg.message);
    msgs->emplace_back(std::move(msg));
  };

  for (const Frame* f : frames) {
    const uint8_t type = f->frame.hd.type;
    switch (type) {
      case NGHTTP2_DATA:
        data_block_size += f->u8payload.size();
        data_block_frames.push_back(f);
        // gRPC request EOS (end-of-stream) is indicated by END_STREAM flag on the last DATA frame.
        // This also indicates this message is a request. Keep appending data, and then export when
        // END_STREAM is seen.
        if (IsEndStream(f->frame.hd)) {
          msg.type = MessageType::kRequest;
          msg.timestamp_ns = f->timestamp_ns;
          handle_end_stream();
        }
        break;
      case NGHTTP2_HEADERS:
        if (f->frame_sync_state == ParseState::kInvalid) {
          return ParseState::kInvalid;
        }
        // All headers frames are considered complete header blocks.
        msg.frames.push_back(f);
        // A gRPC response has 2 header blocks, so we need to merge them.
        msg.headers.insert(f->headers.begin(), f->headers.end());
        if (msg.headers_parse_state != ParseState::kUnknown ||
            f->headers_parse_state != ParseState::kSuccess) {
          // Only assign new state for initialization or an error state.
          msg.headers_parse_state = f->headers_parse_state;
        }

        // gRPC response EOS (end-of-stream) is indicated by END_STREAM flag on the last HEADERS
        // frame. This also indicates this message is a response.
        // No CONTINUATION frame will be used.
        if (IsEndStream(f->frame.hd)) {
          msg.type = MessageType::kResponse;
          msg.timestamp_ns = f->timestamp_ns;
          handle_end_stream();
        }
        break;
      default:
        LOG(DFATAL) << "This function does not accept any frame types other than "
                       "DATA, HEADERS.";
        break;
    }
  }
  return ParseState::kSuccess;
}

ParseState StitchFramesToGRPCMessages(const std::deque<Frame>& frames,
                                      std::map<uint32_t, GRPCMessage>* stream_msgs) {
  std::map<uint32_t, std::vector<const Frame*>> stream_frames;

  // Collect frames for each stream.
  for (const Frame& f : frames) {
    if (f.frame.hd.type == NGHTTP2_HEADERS && f.headers_parse_state == ParseState::kUnknown) {
      // Stop at the first unprocessed HEADERS frame, which indicates that the header block started
      // at this frame is not complete yet.
      break;
    }
    if (f.consumed) {
      // Skip consumed frames.
      continue;
    }
    stream_frames[f.frame.hd.stream_id].push_back(&f);
  }
  for (auto& [stream_id, frame_ptrs] : stream_frames) {
    std::vector<GRPCMessage> msgs;
    ParseState state = StitchGRPCMessageFrames(frame_ptrs, &msgs);
    if (state != ParseState::kSuccess) {
      return state;
    }
    if (msgs.empty()) {
      continue;
    }
    if (msgs.size() > 1) {
      return ParseState::kInvalid;
    }
    stream_msgs->emplace(stream_id, std::move(msgs[0]));
  }
  return ParseState::kSuccess;
}

std::vector<GRPCReqResp> MatchGRPCReqResp(std::map<uint32_t, GRPCMessage> reqs,
                                          std::map<uint32_t, GRPCMessage> resps) {
  std::vector<GRPCReqResp> res;

  // Treat 2 maps as 2 lists ordered according to stream ID, then merge on common stream IDs.
  // It probably will be faster than naive iteration and search.
  for (auto req_iter = reqs.begin(), resp_iter = resps.begin();
       req_iter != reqs.end() && resp_iter != resps.end();) {
    const uint32_t req_stream_id = req_iter->first;
    const uint32_t resp_stream_id = resp_iter->first;

    GRPCMessage& stream_req = req_iter->second;
    GRPCMessage& stream_resp = resp_iter->second;
    if (req_stream_id == resp_stream_id) {
      LOG_IF(DFATAL, stream_req.type == stream_resp.type)
          << "gRPC messages from two streams should be different, got the same type: "
          << static_cast<int>(stream_req.type);
      res.push_back(GRPCReqResp{std::move(stream_req), std::move(stream_resp)});
      ++req_iter;
      ++resp_iter;
    } else if (req_stream_id < resp_stream_id) {
      ++req_iter;
    } else {
      ++resp_iter;
    }
  }
  return res;
}

MethodInputOutput GetProtobufMessages(const GRPCMessage& req, ServiceDescriptorDatabase* db) {
  const char kPathHeader[] = ":path";
  auto iter = req.headers.find(kPathHeader);
  if (iter == req.headers.end()) {
    // No call path specified, bail out.
    return {};
  }
  return db->GetMethodInputOutput(MethodPath(iter->second));
}

ParseState DecodeInteger(u8string_view* buf, size_t prefix, uint32_t* res) {
  // Mandates that the input buffer includes all bytes of the encode integer. That means a few
  // cases handled by nghttp2_hd_decode_length() are not needed.

  // 'shift' is used to indicate how many bits are already decoded in the previous calls of
  // nghttp2_hd_decode_length.
  size_t shift = 0;
  // 'finished' is used to indicate the completion of the decoding.
  int finished = 0;
  // 'initial' is the decoded partial value of the previous calls. Note this cannot replace 'shift'
  // as the leading 0's will confuse the shift size.
  const size_t initial = 0;
  ssize_t rv = nghttp2_hd_decode_length(res, &shift, &finished, initial, shift,
                                        const_cast<uint8_t*>(buf->data()),
                                        const_cast<uint8_t*>(buf->data()) + buf->size(), prefix);
  if (rv == -1) {
    return ParseState::kInvalid;
  }
  if (finished != 1) {
    return ParseState::kNeedsMoreData;
  }
  buf->remove_prefix(rv);
  return ParseState::kSuccess;
}

constexpr uint8_t kMaskTableSizeUpdate = 0xE0;
constexpr uint8_t kBitsTableSizeUpdate = 0x20;
constexpr size_t kPrefixTableSizeUpdate = 5;

constexpr uint8_t kMaskIndexed = 0x80;
constexpr uint8_t kBitsIndexed = 0x80;
constexpr size_t kPrefixIndexed = 7;

constexpr uint8_t kMaskLiteralIndexing = 0xC0;
constexpr uint8_t kBitsLiteralIndexing = 0x40;
constexpr size_t kPrefixLiteralIndexing = 6;

// NoIndexing subsumes "without indexing" and "never indexing" as referred in HPACK spec.
// Where they do share the same bit pattern as the following names suggest,
// except kBitsLiteralNeverIndexing.
constexpr uint8_t kMaskLiteralNoIndexing = 0xF0;
constexpr uint8_t kBitsLiteralNoIndexing = 0x00;
constexpr uint8_t kBitsLiteralNeverIndexing = 0x10;
constexpr size_t kPrefixLiteralNoIndexing = 4;

// Prefix of a length field.
constexpr size_t kPrefixLength = 7;

// Decode length and get the rest of the field as data.
ParseState DecodeLengthPrefixedData(u8string_view* buf, u8string_view* res) {
  uint32_t len = 0;
  const ParseState state = DecodeInteger(buf, kPrefixLength, &len);
  if (state != ParseState::kSuccess) {
    return state;
  }

  if (buf->size() < len) {
    return ParseState::kNeedsMoreData;
  }
  *res = buf->substr(0, len);
  buf->remove_prefix(len);
  return ParseState::kSuccess;
}

inline bool IsHuffEncoded(uint8_t first_byte) {
  // The huffman encoded bit is the leftest bit.
  return (first_byte & (1 << 7)) != 0;
}

template <typename NameValueType>
ParseState DecodeLengthPrefixedNameValue(u8string_view* buf, bool new_name,
                                         size_t name_index_prefix, NameValueType* res) {
  std::variant<uint32_t, u8string_view> name_variant;
  if (new_name) {
    // All 0 in the lower 6 bits indicate a new name, instead of an indexed one.
    buf->remove_prefix(1);
    if (buf->empty()) {
      return ParseState::kNeedsMoreData;
    }
    res->is_name_huff_encoded = IsHuffEncoded(buf->front());
    u8string_view name;
    const ParseState state = DecodeLengthPrefixedData(buf, &name);
    if (state != ParseState::kSuccess) {
      return state;
    }
    name_variant = name;
  } else {
    uint32_t name_index = 0;
    const ParseState state = DecodeInteger(buf, name_index_prefix, &name_index);
    if (state != ParseState::kSuccess) {
      return state;
    }
    name_variant = name_index;
  }
  res->is_value_huff_encoded = IsHuffEncoded(buf->front());
  u8string_view value;
  const ParseState state = DecodeLengthPrefixedData(buf, &value);
  if (state != ParseState::kSuccess) {
    return state;
  }
  res->name = name_variant;
  res->value = value;
  return ParseState::kSuccess;
}

// An alternative approach is to patch nghttp2 so that it can output encoded header fields without
// early termination. We prefer this, as the understanding of the protocol details is required for
// the protobuf inference.
//
// See https://http2.github.io/http2-spec/compression.html#detailed.format for more details on
// HPACK header field format.
ParseState ParseHeaderBlock(u8string_view* buf, std::vector<HeaderField>* res) {
  while (!buf->empty()) {
    if ((buf->front() & kMaskTableSizeUpdate) == kBitsTableSizeUpdate) {
      uint32_t v = 0;
      const ParseState state = DecodeInteger(buf, kPrefixTableSizeUpdate, &v);
      if (state != ParseState::kSuccess) {
        return state;
      }
      res->emplace_back(TableSizeUpdate{v});
    } else if ((buf->front() & kMaskIndexed) == kBitsIndexed) {
      uint32_t v = 0;
      const ParseState state = DecodeInteger(buf, kPrefixIndexed, &v);
      if (state != ParseState::kSuccess) {
        return state;
      }
      res->emplace_back(IndexedHeaderField{v});
    } else if ((buf->front() & kMaskLiteralIndexing) == kBitsLiteralIndexing) {
      LiteralHeaderField field;
      field.update_dynamic_table = true;
      ParseState state = DecodeLengthPrefixedNameValue(buf, buf->front() == kBitsLiteralIndexing,
                                                       kPrefixLiteralIndexing, &field);
      if (state != ParseState::kSuccess) {
        return state;
      }
      res->emplace_back(field);
    } else if ((buf->front() & kMaskLiteralNoIndexing) == kBitsLiteralNoIndexing ||
               (buf->front() & kMaskLiteralNoIndexing) == kBitsLiteralNeverIndexing) {
      LiteralHeaderField field;
      field.update_dynamic_table = false;
      ParseState state = DecodeLengthPrefixedNameValue(
          buf, buf->front() == kBitsLiteralNoIndexing || buf->front() == kBitsLiteralNeverIndexing,
          kPrefixLiteralNoIndexing, &field);
      if (state != ParseState::kSuccess) {
        return state;
      }
      res->emplace_back(field);
    }
  }
  return ParseState::kSuccess;
}

}  // namespace http2
}  // namespace stirling
}  // namespace pl
