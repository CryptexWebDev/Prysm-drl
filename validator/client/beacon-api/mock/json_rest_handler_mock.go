// Code generated by MockGen. DO NOT EDIT.
// Source: validator/client/beacon-api/json_rest_handler.go
//
// Generated by this command:
//
//	mockgen -package=mock -source=validator/client/beacon-api/json_rest_handler.go -destination=validator/client/beacon-api/mock/json_rest_handler_mock.go
//

// Package mock is a generated GoMock package.
package mock

import (
	bytes "bytes"
	context "context"
	http "net/http"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockJsonRestHandler is a mock of JsonRestHandler interface.
type MockJsonRestHandler struct {
	ctrl     *gomock.Controller
	recorder *MockJsonRestHandlerMockRecorder
}

// MockJsonRestHandlerMockRecorder is the mock recorder for MockJsonRestHandler.
type MockJsonRestHandlerMockRecorder struct {
	mock *MockJsonRestHandler
}

// NewMockJsonRestHandler creates a new mock instance.
func NewMockJsonRestHandler(ctrl *gomock.Controller) *MockJsonRestHandler {
	mock := &MockJsonRestHandler{ctrl: ctrl}
	mock.recorder = &MockJsonRestHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockJsonRestHandler) EXPECT() *MockJsonRestHandlerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockJsonRestHandler) Get(ctx context.Context, endpoint string, resp any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", ctx, endpoint, resp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockJsonRestHandlerMockRecorder) Get(ctx, endpoint, resp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockJsonRestHandler)(nil).Get), ctx, endpoint, resp)
}

// Host mocks base method.
func (m *MockJsonRestHandler) Host() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPHost")
	ret0, _ := ret[0].(string)
	return ret0
}

// Host indicates an expected call of Host.
func (mr *MockJsonRestHandlerMockRecorder) Host() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPHost", reflect.TypeOf((*MockJsonRestHandler)(nil).Host))
}

// HttpClient mocks base method.
func (m *MockJsonRestHandler) HttpClient() *http.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HttpClient")
	ret0, _ := ret[0].(*http.Client)
	return ret0
}

// HttpClient indicates an expected call of HttpClient.
func (mr *MockJsonRestHandlerMockRecorder) HttpClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HttpClient", reflect.TypeOf((*MockJsonRestHandler)(nil).HttpClient))
}

// Post mocks base method.
func (m *MockJsonRestHandler) Post(ctx context.Context, endpoint string, headers map[string]string, data *bytes.Buffer, resp any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", ctx, endpoint, headers, data, resp)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockJsonRestHandlerMockRecorder) Post(ctx, endpoint, headers, data, resp any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockJsonRestHandler)(nil).Post), ctx, endpoint, headers, data, resp)
}

// SetHost mocks base method.
func (m *MockJsonRestHandler) SetHost(host string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetHost", host)
}

// SetHost indicates an expected call of SetHost.
func (mr *MockJsonRestHandlerMockRecorder) SetHost(host any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHost", reflect.TypeOf((*MockJsonRestHandler)(nil).SetHost), host)
}