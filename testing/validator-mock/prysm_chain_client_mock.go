// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/v5/validator/client/iface (interfaces: PrysmChainClient)
//
// Generated by this command:
//
//	mockgen -package=validator_mock -destination=testing/validator-mock/prysm_chain_client_mock.go github.com/prysmaticlabs/prysm/v5/validator/client/iface PrysmChainClient
//

// Package validator_mock is a generated GoMock package.
package validator_mock

import (
	context "context"
	reflect "reflect"

	validator "github.com/Dorol-Chain/Prysm-drl/v5/consensus-types/validator"
	iface "github.com/Dorol-Chain/Prysm-drl/v5/validator/client/iface"
	gomock "go.uber.org/mock/gomock"
)

// MockPrysmChainClient is a mock of PrysmChainClient interface.
type MockPrysmChainClient struct {
	ctrl     *gomock.Controller
	recorder *MockPrysmChainClientMockRecorder
}

// MockPrysmChainClientMockRecorder is the mock recorder for MockPrysmChainClient.
type MockPrysmChainClientMockRecorder struct {
	mock *MockPrysmChainClient
}

// NewMockPrysmChainClient creates a new mock instance.
func NewMockPrysmChainClient(ctrl *gomock.Controller) *MockPrysmChainClient {
	mock := &MockPrysmChainClient{ctrl: ctrl}
	mock.recorder = &MockPrysmChainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPrysmChainClient) EXPECT() *MockPrysmChainClientMockRecorder {
	return m.recorder
}

// ValidatorCount mocks base method.
func (m *MockPrysmChainClient) ValidatorCount(arg0 context.Context, arg1 string, arg2 []validator.Status) ([]iface.ValidatorCount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorCount", arg0, arg1, arg2)
	ret0, _ := ret[0].([]iface.ValidatorCount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorCount indicates an expected call of ValidatorCount.
func (mr *MockPrysmChainClientMockRecorder) ValidatorCount(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorCount", reflect.TypeOf((*MockPrysmChainClient)(nil).ValidatorCount), arg0, arg1, arg2)
}