// Code generated by MockGen. DO NOT EDIT.
// Source: auth.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	auth "github.com/matt-hoiland/blueprints/internal/auth"
	common "github.com/stretchr/gomniauth/common"
	objx "github.com/stretchr/objx"
)

// MockOAuthAdapter is a mock of OAuthAdapter interface.
type MockOAuthAdapter struct {
	ctrl     *gomock.Controller
	recorder *MockOAuthAdapterMockRecorder
}

// MockOAuthAdapterMockRecorder is the mock recorder for MockOAuthAdapter.
type MockOAuthAdapterMockRecorder struct {
	mock *MockOAuthAdapter
}

// NewMockOAuthAdapter creates a new mock instance.
func NewMockOAuthAdapter(ctrl *gomock.Controller) *MockOAuthAdapter {
	mock := &MockOAuthAdapter{ctrl: ctrl}
	mock.recorder = &MockOAuthAdapterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOAuthAdapter) EXPECT() *MockOAuthAdapterMockRecorder {
	return m.recorder
}

// Provider mocks base method.
func (m *MockOAuthAdapter) Provider(arg0 string) (auth.Provider, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provider", arg0)
	ret0, _ := ret[0].(auth.Provider)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provider indicates an expected call of Provider.
func (mr *MockOAuthAdapterMockRecorder) Provider(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provider", reflect.TypeOf((*MockOAuthAdapter)(nil).Provider), arg0)
}

// MockProvider is a mock of Provider interface.
type MockProvider struct {
	ctrl     *gomock.Controller
	recorder *MockProviderMockRecorder
}

// MockProviderMockRecorder is the mock recorder for MockProvider.
type MockProviderMockRecorder struct {
	mock *MockProvider
}

// NewMockProvider creates a new mock instance.
func NewMockProvider(ctrl *gomock.Controller) *MockProvider {
	mock := &MockProvider{ctrl: ctrl}
	mock.recorder = &MockProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProvider) EXPECT() *MockProviderMockRecorder {
	return m.recorder
}

// CompleteAuth mocks base method.
func (m *MockProvider) CompleteAuth(data objx.Map) (*common.Credentials, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CompleteAuth", data)
	ret0, _ := ret[0].(*common.Credentials)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteAuth indicates an expected call of CompleteAuth.
func (mr *MockProviderMockRecorder) CompleteAuth(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteAuth", reflect.TypeOf((*MockProvider)(nil).CompleteAuth), data)
}

// GetBeginAuthURL mocks base method.
func (m *MockProvider) GetBeginAuthURL(arg0 *common.State, arg1 objx.Map) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBeginAuthURL", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBeginAuthURL indicates an expected call of GetBeginAuthURL.
func (mr *MockProviderMockRecorder) GetBeginAuthURL(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBeginAuthURL", reflect.TypeOf((*MockProvider)(nil).GetBeginAuthURL), arg0, arg1)
}

// GetUser mocks base method.
func (m *MockProvider) GetUser(arg0 *common.Credentials) (common.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", arg0)
	ret0, _ := ret[0].(common.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockProviderMockRecorder) GetUser(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockProvider)(nil).GetUser), arg0)
}
