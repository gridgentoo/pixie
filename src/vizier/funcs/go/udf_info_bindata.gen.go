// Code generated for package funcs by go-bindata DO NOT EDIT. (@generated)
// sources:
// bazel-out/k8-fastbuild/bin/src/vizier/funcs/data/udf.pb
package funcs

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _srcVizierFuncsDataUdfPb = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x58\xcb\x8e\xdc\x4a\x19\x96\xdd\x76\xf5\xe4\xef\xe9\x9e\xbe\x64\x72\x7a\x9c\xe4\x1c\xc7\x08\x29\x84\x05\x23\x04\x6c\x58\x45\x4a\x08\x88\x9b\x50\x72\xd8\xb0\xb0\x6a\xec\x9a\xee\x22\x76\xd9\xb1\xcb\x39\x1d\x24\x56\x9c\x49\xb6\x07\x76\xec\xf2\x00\x3c\x02\x0f\x82\x78\x02\xd8\xf3\x00\xa8\xaa\x5c\x76\x95\xdb\x93\xc5\xd9\xf5\xf7\xfd\x97\xfa\x6f\x65\xff\x6e\x38\x87\x59\xfc\x3a\x27\x98\xd5\xf1\x35\xe5\x81\xeb\xbb\xa1\xff\xd8\x81\x05\xf8\x49\xd1\x30\x1e\x38\x4e\xe8\x5a\xd8\x1d\x60\x6f\x80\xfd\x01\x46\x12\x9f\xc2\x24\xc7\x07\x6d\xad\x91\x17\x7a\x06\x42\x21\x7a\xec\xc0\x1c\x3c\x11\x8e\x38\xd8\x33\xa1\x6b\xc3\xde\x94\x32\xd3\x2d\x3d\x92\x29\xb7\x4b\xb8\xf3\xa6\xc1\x8c\xd3\x8c\xd4\x42\xdf\x1f\x30\x5e\xe8\xc3\x1c\x50\x8d\xf3\x32\x23\x22\x0b\x5f\x38\xa8\x9b\x5c\x04\xe2\x28\x77\x12\x75\x47\x49\x24\x8f\x5a\x3f\x84\x7b\x31\x61\x49\x91\x92\xb8\x26\x8c\x13\x96\x90\xb8\xa4\x24\x51\x8e\xd6\x6b\x58\xc4\xe4\x40\x92\x78\x5f\xd4\x9c\xe1\x9c\x84\xfe\xfa\x02\x56\xba\xee\x94\x5d\x93\x4a\xd8\x04\xae\xef\x87\xee\xfa\x1e\x9c\xc5\x9c\x1c\x78\x4c\xf2\x2b\x92\xa6\x94\xed\x94\x9b\x19\x4c\x70\x9a\x06\xae\xeb\x86\x6e\x0f\xbc\xd0\xeb\x01\x0a\x51\x07\x3c\xd7\x90\x78\xa6\x9a\x6f\x7a\x43\x6e\x88\xd6\x1b\x98\xe1\xb2\xac\x8a\xc3\xf3\x37\x0d\xce\xa4\xba\xb3\x3e\x01\x0f\xd7\x34\x55\x87\x5d\x51\xd6\x9f\xac\x00\x32\x80\x67\x4a\xa4\xcb\x0e\xa8\x98\xbc\x84\xd0\x4c\x4e\xcb\xfa\x11\x04\x49\xc1\x38\xa6\x8c\x54\x31\x4d\x63\x5e\xc4\x35\xc7\x15\x8f\x39\xcd\x65\xc5\x90\x28\xe8\x88\x0a\x6f\x6a\x55\x89\x10\x2e\x8e\xc5\x45\x69\x38\xf8\x2e\x7c\xde\x6b\x88\x92\x0b\x1d\xd3\x46\x39\xfa\x0e\x3c\x38\x56\x1b\x06\x13\xc1\xfd\x31\x25\xeb\xbc\x25\x9c\xb4\x3a\xb5\xac\xaf\xb3\x5e\x00\x4a\xe9\x5b\x9a\x12\x59\x36\xcf\xc2\x9e\x8d\xbd\x81\x5c\x36\x6b\x0e\x3e\x51\xcd\x10\x03\x68\x42\xd7\x82\xae\x2d\x75\x07\x52\xcf\x82\x93\x89\x05\x3d\x5b\xd9\xb3\x95\x65\x1a\x3d\x44\x28\x74\xd6\xa7\xe0\x5d\x53\x96\xb6\x93\x7a\x0a\xfe\x75\x56\x14\x95\xea\xeb\x06\x66\xbb\x8a\x60\x4e\xaa\x57\x7b\xcc\xda\x50\x86\xa4\x3c\x63\x48\xca\x93\xb6\xb0\x34\xc8\xe7\x46\x3a\xa3\x12\xe9\x68\x54\x72\xbb\x37\x99\xc2\x06\x4e\xf7\xe4\x20\x7a\x88\xeb\x84\x52\x35\x08\x73\x40\x94\xbd\x25\x95\x7a\xc0\x09\x1d\x5a\x0a\x95\xb2\x48\xbb\x61\x99\x03\xca\x08\xdb\xf1\xbd\x7c\xc8\x89\x96\x67\xa4\xae\x8d\x5c\x2d\x46\xc6\x67\x31\x32\xae\x73\x98\x6b\xc6\x4c\xf1\x98\x96\xf6\xc7\xf4\x2d\x4e\x64\x66\x6b\x80\xac\xd8\xd1\x04\x67\x4f\x45\x8b\xe4\xdc\x0c\x38\x79\xd8\xaa\xe3\x7e\x53\xc8\x27\xfc\x11\xa5\xb4\xee\xb4\xd4\x6f\xab\xd6\x99\x4d\x49\x5f\x0b\x40\x79\x91\x36\x59\xd1\x3e\x1d\x4c\x8c\x6c\x8c\x06\x72\x84\x54\x15\xf3\x26\xe3\xb4\xcc\xde\xb5\x1e\x6c\x46\x5c\x05\x8b\x91\x97\xc5\x66\xd4\x75\x41\x8c\xec\x30\x27\xaa\x81\x06\x6c\x3d\xb0\x82\x3f\x37\xee\x93\xcd\xb8\x43\xc6\x3d\xd2\x71\x8f\x75\xbc\x21\xe3\x1d\xe9\x78\x47\x3a\xb2\x81\x16\x23\x7b\x77\x06\x27\xac\xce\x8a\xe2\x75\x53\xea\x69\xf3\xcb\xac\x49\x5e\xb7\x8f\xeb\x73\x98\x4b\x18\x5f\x67\x05\xe6\x3f\xf9\x91\xa4\x3d\x71\x93\x14\x4d\x99\x26\xdd\x75\x00\x1b\x35\xb7\x62\x82\xc5\x03\xab\x2e\xb1\x7e\x17\xd9\x32\xf1\xc6\x12\x0a\x4a\x76\x01\xeb\x5e\x26\x7e\xf5\xa2\xfb\x70\xb7\x17\xd5\xa4\x7a\x4b\x13\xd2\x5d\x8c\x87\x70\xef\x58\x78\x9b\xed\xe0\xe1\x6a\xc5\x33\x78\xa8\x3e\x80\x73\x1d\xc5\x48\x26\x6d\xb4\x5a\x6a\xde\xd4\x2d\xac\x4c\x51\x45\x70\xfa\x4e\x48\x1c\x1d\x6a\xf7\x14\x1f\x64\xf2\x05\x6c\xc7\xc4\x7d\x2e\x43\xfb\x41\x36\x83\x98\xcc\x77\xd6\x20\x99\x41\xaa\x73\xf0\xab\xa2\x61\xea\xb5\xed\x8b\x57\x64\x1f\xda\x78\x24\x21\x5c\x98\xe4\x48\x81\x0c\x27\xb7\xe5\x7b\x17\x4e\x6b\x5e\xd1\x32\x2e\x2b\x72\x4d\x0f\xed\xac\xad\xe1\x4e\xdd\x5c\x09\x01\xdb\x05\x13\xdf\x15\x11\x2d\xe1\xa4\x6e\xae\x78\x85\x13\xde\xdf\x53\x83\x69\x6f\x99\xc1\xa0\xa1\x8e\xbe\xb9\x06\x73\x64\x85\x8e\x3c\x23\xf5\x14\x99\xf2\x22\x2b\xbe\x22\x95\x8a\x5b\xe2\xa6\x2c\x35\x9e\x81\xc7\x2b\x9a\x2b\xb0\x81\xd3\xa6\x54\x65\x13\xfb\x4b\xe0\x4c\xd4\x52\xa5\xc9\x24\x4f\x33\xca\x88\xe0\xe5\x70\x76\xbc\xb5\x1e\x4c\x54\xb7\x8f\x85\xaa\x05\x42\xfc\x19\x2c\xb5\x58\x6f\x76\x4a\xb0\x85\x95\x16\x18\x3d\x19\x4a\xfa\xcb\x37\x91\x45\x9f\x69\x49\xa9\x83\x3e\x87\x45\xc7\xb5\xe3\x3d\x38\xb7\xbf\xa5\x42\x60\x24\x29\x04\x6f\x8a\x5a\xf1\x17\xb0\x36\x79\x3d\x96\x03\x91\x39\x1c\x83\xd2\xd8\xe3\x37\x91\xcf\xa3\x2e\xb4\x76\x50\x04\x1d\xfc\xd5\x81\xc5\x0b\xc2\x9f\xee\x08\xe3\x2f\xe5\x31\x5b\x14\xfd\x19\x16\x70\x82\x05\x15\xd3\x74\x39\x09\x1d\x00\xb5\x5a\x2e\xdd\x50\x7c\x2a\x9c\xe8\xea\x2d\xfd\x50\x6c\xea\x40\xcb\x18\xa7\x69\x45\xea\x5a\x32\x2b\x98\x29\x6b\x11\x38\xd1\x54\x22\x5f\xef\xf2\x0e\x2d\x51\xe8\xc0\x67\xb0\xca\x70\xcd\xe3\x3d\xc1\x15\xbf\x22\x98\xc7\xac\x16\x07\x04\x7f\x00\x78\x41\xf8\xcb\x64\x4f\x72\x2c\xc2\xf9\xa5\x38\x82\xe3\xab\x4c\xe5\xd3\xf9\x2b\xb2\x26\x67\x63\x14\x7f\x57\x0e\xa9\x94\xd4\x89\xa0\x82\x3f\xc2\xe6\x05\xe1\xaf\x2a\x9c\x90\xb2\xa0\x46\xd2\xbf\x83\x0d\xcc\x79\xc7\x1b\x99\x77\x27\xcc\xc0\xef\x33\x3a\x05\xa4\xfa\x22\xd1\x06\xe6\x45\xc3\xcb\x86\xc7\x32\x50\x49\x06\x3f\x95\x89\x7c\xf9\xec\xe9\xaf\x68\xcd\xb7\x28\xfa\xbe\xe5\x6d\x05\xb3\x8a\xf0\xa6\x32\xe2\x15\x65\xae\x76\xb6\xf1\xcf\xbe\x8d\xf1\x97\x30\x93\xc6\xaf\xb4\xf5\x33\xcb\x7a\x01\x27\xe2\xfb\xa6\xe1\x45\x25\xf1\x19\xdc\xa1\x8c\xf2\x58\xdb\xc3\x39\x9c\xb5\xe9\x54\x24\xc3\x9c\x16\x4c\xba\x7d\x0c\xd3\xdf\x93\xaa\xa6\x05\xdb\xa2\xe8\xbe\x35\x15\x73\x98\xbe\x55\x22\xa9\xf9\x43\x98\xc7\xcf\xc8\x55\xb3\xfb\xf5\x33\x51\x61\xb2\x75\xa2\x2f\x2c\xfd\x15\xcc\x52\x21\xef\x67\x24\xf8\x31\x2c\x95\xcd\x4b\x8e\x93\xd7\xb2\x45\x63\x66\xb5\x90\xc6\xb2\x53\xd2\xec\xef\x0e\x9c\x29\xbb\x57\xa2\xf4\xbf\x60\xd7\xc5\xd6\x89\xfe\x62\x0f\xad\x99\xfd\x09\xb8\x2d\xbb\x81\xf9\x15\xe6\xc9\x9e\xd4\x62\x7e\x89\x22\xcf\xe1\x4c\x93\xe4\x50\xd2\x8a\x74\x27\xb3\x26\x8f\x5b\x91\x76\x5a\xd3\x3f\x11\xf9\xfb\x2e\x2c\x72\x7c\x50\xdd\x8f\x35\x1b\xfc\x00\x56\xf1\xcf\x09\x2e\x5f\x54\xc5\x57\x7c\x2f\xf3\xaa\xb7\x4e\xb4\x1d\xc6\xb6\x27\xb8\x94\xc9\x7c\x0f\x66\xd2\xe0\xa5\xfc\xae\xfd\xa4\xea\x63\x00\xa5\xca\x31\xff\xa4\xd3\x68\x21\x3f\x1d\x9f\x78\x37\xd3\x9b\xe9\xe5\xcd\xb4\xc3\x1f\xa6\x1f\xa6\x97\x1f\x7a\xfc\x7e\xfa\x7e\x7a\xf9\xbe\xc7\x1f\xd1\x47\x74\xf9\x11\x45\xed\x87\xe5\x25\x89\x3e\xbf\xed\x3b\xef\x89\xe3\x5c\xfe\xdb\x8b\x56\xfa\xa3\x48\x18\xdf\x48\xe7\xcb\x8e\x9a\x38\x37\xda\x7d\x8e\x0f\x4b\x27\x72\xbb\x70\x5a\xfc\xf5\xf4\xf2\x6b\x13\x77\xe1\xb5\xf8\xfd\xc0\x5e\x85\x77\xa6\xfe\x63\xe8\x1d\x1a\x84\xf2\x68\x10\xca\xa5\x41\x28\x9f\x06\xa1\x9c\x2e\xe4\xbf\x11\x76\x90\x94\xd9\x41\x52\x66\x07\x49\x99\x1d\x24\x35\xfc\xdd\x1f\xdd\xf1\x44\xd9\xfe\xe9\x0f\x84\xfa\x3d\x23\x84\xff\x9b\x44\xc1\xd8\x96\x27\x64\xdf\x4c\x44\x3f\xc6\x37\x39\x21\xff\x87\x1b\x3d\xbc\x65\x23\x6b\xcf\x7d\x74\xfb\xfa\xf4\xc4\xfd\x66\x22\x3c\x04\x63\x1b\x92\x30\xff\xdb\x24\xba\x6b\xfc\x23\xd3\x56\xea\xbf\xd3\x68\x63\xb3\x8e\x73\xf9\x9f\x69\x14\x7d\x6a\x41\x6a\x83\x7d\xf4\x89\x0d\xa9\x0d\x78\xd3\x6f\x1a\xfd\x38\x5b\xa4\x9e\x69\x8b\xd4\x83\x6d\x91\x7a\xba\x17\xf2\xaf\x21\xab\x93\x2d\xee\x3a\xd9\x62\xa5\x3f\x58\x57\x44\x64\xf2\x62\x8c\x6f\x1e\x42\xfc\x2f\x2f\xba\x18\xd9\x30\xda\x9c\x2e\x46\x56\x8c\xb6\xf5\xdb\xe3\xc5\xa1\x6d\x7c\x30\xb6\x21\xb4\x6d\x79\x30\xbe\x07\xa8\x2a\xff\x3f\x00\x00\xff\xff\xd3\xc2\xc0\xd9\x41\x14\x00\x00")

func srcVizierFuncsDataUdfPbBytes() ([]byte, error) {
	return bindataRead(
		_srcVizierFuncsDataUdfPb,
		"src/vizier/funcs/data/udf.pb",
	)
}

func srcVizierFuncsDataUdfPb() (*asset, error) {
	bytes, err := srcVizierFuncsDataUdfPbBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "src/vizier/funcs/data/udf.pb", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"src/vizier/funcs/data/udf.pb": srcVizierFuncsDataUdfPb,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//
//	data/
//	  foo.txt
//	  img/
//	    a.png
//	    b.png
//
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"src": &bintree{nil, map[string]*bintree{
		"vizier": &bintree{nil, map[string]*bintree{
			"funcs": &bintree{nil, map[string]*bintree{
				"data": &bintree{nil, map[string]*bintree{
					"udf.pb": &bintree{srcVizierFuncsDataUdfPb, map[string]*bintree{}},
				}},
			}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
