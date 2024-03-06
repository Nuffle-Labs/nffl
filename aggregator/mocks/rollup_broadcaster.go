// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/near-sffl/aggregator (interfaces: RollupBroadcasterer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/rollup_broadcaster.go -package=mocks github.com/NethermindEth/near-sffl/aggregator RollupBroadcasterer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	contractSFFLRegistryRollup "github.com/NethermindEth/near-sffl/contracts/bindings/SFFLRegistryRollup"
	gomock "go.uber.org/mock/gomock"
)

// MockRollupBroadcasterer is a mock of RollupBroadcasterer interface.
type MockRollupBroadcasterer struct {
	ctrl     *gomock.Controller
	recorder *MockRollupBroadcastererMockRecorder
}

// MockRollupBroadcastererMockRecorder is the mock recorder for MockRollupBroadcasterer.
type MockRollupBroadcastererMockRecorder struct {
	mock *MockRollupBroadcasterer
}

// NewMockRollupBroadcasterer creates a new mock instance.
func NewMockRollupBroadcasterer(ctrl *gomock.Controller) *MockRollupBroadcasterer {
	mock := &MockRollupBroadcasterer{ctrl: ctrl}
	mock.recorder = &MockRollupBroadcastererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRollupBroadcasterer) EXPECT() *MockRollupBroadcastererMockRecorder {
	return m.recorder
}

// BroadcastOperatorSetUpdate mocks base method.
func (m *MockRollupBroadcasterer) BroadcastOperatorSetUpdate(arg0 context.Context, arg1 contractSFFLRegistryRollup.OperatorSetUpdateMessage, arg2 contractSFFLRegistryRollup.OperatorsSignatureInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BroadcastOperatorSetUpdate", arg0, arg1, arg2)
}

// BroadcastOperatorSetUpdate indicates an expected call of BroadcastOperatorSetUpdate.
func (mr *MockRollupBroadcastererMockRecorder) BroadcastOperatorSetUpdate(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BroadcastOperatorSetUpdate", reflect.TypeOf((*MockRollupBroadcasterer)(nil).BroadcastOperatorSetUpdate), arg0, arg1, arg2)
}

// GetErrorChan mocks base method.
func (m *MockRollupBroadcasterer) GetErrorChan() <-chan error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetErrorChan")
	ret0, _ := ret[0].(<-chan error)
	return ret0
}

// GetErrorChan indicates an expected call of GetErrorChan.
func (mr *MockRollupBroadcastererMockRecorder) GetErrorChan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetErrorChan", reflect.TypeOf((*MockRollupBroadcasterer)(nil).GetErrorChan))
}
