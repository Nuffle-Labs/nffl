// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/NethermindEth/near-sffl/aggregator (interfaces: TaskBlsAggregationService)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/task_blsagg.go -package=mocks github.com/NethermindEth/near-sffl/aggregator TaskBlsAggregationService
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"
	time "time"

	bls "github.com/Layr-Labs/eigensdk-go/crypto/bls"
	types "github.com/Layr-Labs/eigensdk-go/types"
	types0 "github.com/NethermindEth/near-sffl/aggregator/types"
	gomock "go.uber.org/mock/gomock"
)

// MockTaskBlsAggregationService is a mock of TaskBlsAggregationService interface.
type MockTaskBlsAggregationService struct {
	ctrl     *gomock.Controller
	recorder *MockTaskBlsAggregationServiceMockRecorder
}

// MockTaskBlsAggregationServiceMockRecorder is the mock recorder for MockTaskBlsAggregationService.
type MockTaskBlsAggregationServiceMockRecorder struct {
	mock *MockTaskBlsAggregationService
}

// NewMockTaskBlsAggregationService creates a new mock instance.
func NewMockTaskBlsAggregationService(ctrl *gomock.Controller) *MockTaskBlsAggregationService {
	mock := &MockTaskBlsAggregationService{ctrl: ctrl}
	mock.recorder = &MockTaskBlsAggregationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskBlsAggregationService) EXPECT() *MockTaskBlsAggregationServiceMockRecorder {
	return m.recorder
}

// GetResponseChannel mocks base method.
func (m *MockTaskBlsAggregationService) GetResponseChannel() <-chan types0.TaskBlsAggregationServiceResponse {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponseChannel")
	ret0, _ := ret[0].(<-chan types0.TaskBlsAggregationServiceResponse)
	return ret0
}

// GetResponseChannel indicates an expected call of GetResponseChannel.
func (mr *MockTaskBlsAggregationServiceMockRecorder) GetResponseChannel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponseChannel", reflect.TypeOf((*MockTaskBlsAggregationService)(nil).GetResponseChannel))
}

// InitializeNewTask mocks base method.
func (m *MockTaskBlsAggregationService) InitializeNewTask(arg0, arg1 uint32, arg2 types.QuorumNums, arg3 types.QuorumThresholdPercentages, arg4 time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitializeNewTask", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// InitializeNewTask indicates an expected call of InitializeNewTask.
func (mr *MockTaskBlsAggregationServiceMockRecorder) InitializeNewTask(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitializeNewTask", reflect.TypeOf((*MockTaskBlsAggregationService)(nil).InitializeNewTask), arg0, arg1, arg2, arg3, arg4)
}

// ProcessNewSignature mocks base method.
func (m *MockTaskBlsAggregationService) ProcessNewSignature(arg0 context.Context, arg1 uint32, arg2 types.Bytes32, arg3 *bls.Signature, arg4 types.Bytes32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessNewSignature", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessNewSignature indicates an expected call of ProcessNewSignature.
func (mr *MockTaskBlsAggregationServiceMockRecorder) ProcessNewSignature(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessNewSignature", reflect.TypeOf((*MockTaskBlsAggregationService)(nil).ProcessNewSignature), arg0, arg1, arg2, arg3, arg4)
}
