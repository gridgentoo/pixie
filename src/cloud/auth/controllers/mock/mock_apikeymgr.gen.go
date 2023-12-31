// Code generated by MockGen. DO NOT EDIT.
// Source: server.go

// Package mock_controllers is a generated GoMock package.
package mock_controllers

import (
	context "context"
	reflect "reflect"

	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	controllers "px.dev/pixie/src/cloud/auth/controllers"
)

// MockAPIKeyMgr is a mock of APIKeyMgr interface.
type MockAPIKeyMgr struct {
	ctrl     *gomock.Controller
	recorder *MockAPIKeyMgrMockRecorder
}

// MockAPIKeyMgrMockRecorder is the mock recorder for MockAPIKeyMgr.
type MockAPIKeyMgrMockRecorder struct {
	mock *MockAPIKeyMgr
}

// NewMockAPIKeyMgr creates a new mock instance.
func NewMockAPIKeyMgr(ctrl *gomock.Controller) *MockAPIKeyMgr {
	mock := &MockAPIKeyMgr{ctrl: ctrl}
	mock.recorder = &MockAPIKeyMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAPIKeyMgr) EXPECT() *MockAPIKeyMgrMockRecorder {
	return m.recorder
}

// FetchOrgUserIDUsingAPIKey mocks base method.
func (m *MockAPIKeyMgr) FetchOrgUserIDUsingAPIKey(ctx context.Context, key string) (uuid.UUID, uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOrgUserIDUsingAPIKey", ctx, key)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(uuid.UUID)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FetchOrgUserIDUsingAPIKey indicates an expected call of FetchOrgUserIDUsingAPIKey.
func (mr *MockAPIKeyMgrMockRecorder) FetchOrgUserIDUsingAPIKey(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOrgUserIDUsingAPIKey", reflect.TypeOf((*MockAPIKeyMgr)(nil).FetchOrgUserIDUsingAPIKey), ctx, key)
}

// MockAuthProvider is a mock of AuthProvider interface.
type MockAuthProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAuthProviderMockRecorder
}

// MockAuthProviderMockRecorder is the mock recorder for MockAuthProvider.
type MockAuthProviderMockRecorder struct {
	mock *MockAuthProvider
}

// NewMockAuthProvider creates a new mock instance.
func NewMockAuthProvider(ctrl *gomock.Controller) *MockAuthProvider {
	mock := &MockAuthProvider{ctrl: ctrl}
	mock.recorder = &MockAuthProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthProvider) EXPECT() *MockAuthProviderMockRecorder {
	return m.recorder
}

// CreateIdentity mocks base method.
func (m *MockAuthProvider) CreateIdentity(email string) (*controllers.CreateIdentityResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIdentity", email)
	ret0, _ := ret[0].(*controllers.CreateIdentityResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateIdentity indicates an expected call of CreateIdentity.
func (mr *MockAuthProviderMockRecorder) CreateIdentity(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIdentity", reflect.TypeOf((*MockAuthProvider)(nil).CreateIdentity), email)
}

// CreateInviteLink mocks base method.
func (m *MockAuthProvider) CreateInviteLink(authProviderID string) (*controllers.CreateInviteLinkResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateInviteLink", authProviderID)
	ret0, _ := ret[0].(*controllers.CreateInviteLinkResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateInviteLink indicates an expected call of CreateInviteLink.
func (mr *MockAuthProviderMockRecorder) CreateInviteLink(authProviderID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateInviteLink", reflect.TypeOf((*MockAuthProvider)(nil).CreateInviteLink), authProviderID)
}

// GetUserInfoFromAccessToken mocks base method.
func (m *MockAuthProvider) GetUserInfoFromAccessToken(accessToken string) (*controllers.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfoFromAccessToken", accessToken)
	ret0, _ := ret[0].(*controllers.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfoFromAccessToken indicates an expected call of GetUserInfoFromAccessToken.
func (mr *MockAuthProviderMockRecorder) GetUserInfoFromAccessToken(accessToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfoFromAccessToken", reflect.TypeOf((*MockAuthProvider)(nil).GetUserInfoFromAccessToken), accessToken)
}
