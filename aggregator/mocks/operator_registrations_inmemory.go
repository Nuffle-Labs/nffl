// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/near-sffl/aggregator (interfaces: OperatorRegistrationsService)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/operator_registrations_inmemory.go -package=mocks github.com/NethermindEth/near-sffl/aggregator OperatorRegistrationsService
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/Layr-Labs/eigensdk-go/types"
	common "github.com/ethereum/go-ethereum/common"
	gomock "go.uber.org/mock/gomock"
)

// MockOperatorRegistrationsService is a mock of OperatorRegistrationsService interface.
type MockOperatorRegistrationsService struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorRegistrationsServiceMockRecorder
}

// MockOperatorRegistrationsServiceMockRecorder is the mock recorder for MockOperatorRegistrationsService.
type MockOperatorRegistrationsServiceMockRecorder struct {
	mock *MockOperatorRegistrationsService
}

// NewMockOperatorRegistrationsService creates a new mock instance.
func NewMockOperatorRegistrationsService(ctrl *gomock.Controller) *MockOperatorRegistrationsService {
	mock := &MockOperatorRegistrationsService{ctrl: ctrl}
	mock.recorder = &MockOperatorRegistrationsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorRegistrationsService) EXPECT() *MockOperatorRegistrationsServiceMockRecorder {
	return m.recorder
}

// GetOperatorPubkeys mocks base method.
func (m *MockOperatorRegistrationsService) GetOperatorPubkeys(arg0 context.Context, arg1 common.Address) (types.OperatorPubkeys, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorPubkeys", arg0, arg1)
	ret0, _ := ret[0].(types.OperatorPubkeys)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetOperatorPubkeys indicates an expected call of GetOperatorPubkeys.
func (mr *MockOperatorRegistrationsServiceMockRecorder) GetOperatorPubkeys(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorPubkeys", reflect.TypeOf((*MockOperatorRegistrationsService)(nil).GetOperatorPubkeys), arg0, arg1)
}

// GetOperatorPubkeysById mocks base method.
func (m *MockOperatorRegistrationsService) GetOperatorPubkeysById(arg0 context.Context, arg1 types.Bytes32) (types.OperatorPubkeys, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOperatorPubkeysById", arg0, arg1)
	ret0, _ := ret[0].(types.OperatorPubkeys)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetOperatorPubkeysById indicates an expected call of GetOperatorPubkeysById.
func (mr *MockOperatorRegistrationsServiceMockRecorder) GetOperatorPubkeysById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOperatorPubkeysById", reflect.TypeOf((*MockOperatorRegistrationsService)(nil).GetOperatorPubkeysById), arg0, arg1)
}