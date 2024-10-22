// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nuffle-Labs/nffl/core/chainio (interfaces: AvsSubscriberer)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/avs_subscriber.go -package=mocks github.com/Nuffle-Labs/nffl/core/chainio AvsSubscriberer
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	contractBLSApkRegistry "github.com/Layr-Labs/eigensdk-go/contracts/bindings/BLSApkRegistry"
	contractRegistryCoordinator "github.com/Layr-Labs/eigensdk-go/contracts/bindings/RegistryCoordinator"
	contractSFFLOperatorSetUpdateRegistry "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLOperatorSetUpdateRegistry"
	contractSFFLTaskManager "github.com/Nuffle-Labs/nffl/contracts/bindings/SFFLTaskManager"
	types "github.com/ethereum/go-ethereum/core/types"
	event "github.com/ethereum/go-ethereum/event"
	gomock "go.uber.org/mock/gomock"
)

// MockAvsSubscriberer is a mock of AvsSubscriberer interface.
type MockAvsSubscriberer struct {
	ctrl     *gomock.Controller
	recorder *MockAvsSubscribererMockRecorder
}

// MockAvsSubscribererMockRecorder is the mock recorder for MockAvsSubscriberer.
type MockAvsSubscribererMockRecorder struct {
	mock *MockAvsSubscriberer
}

// NewMockAvsSubscriberer creates a new mock instance.
func NewMockAvsSubscriberer(ctrl *gomock.Controller) *MockAvsSubscriberer {
	mock := &MockAvsSubscriberer{ctrl: ctrl}
	mock.recorder = &MockAvsSubscribererMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvsSubscriberer) EXPECT() *MockAvsSubscribererMockRecorder {
	return m.recorder
}

// ParseCheckpointTaskResponded mocks base method.
func (m *MockAvsSubscriberer) ParseCheckpointTaskResponded(arg0 types.Log) (*contractSFFLTaskManager.ContractSFFLTaskManagerCheckpointTaskResponded, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseCheckpointTaskResponded", arg0)
	ret0, _ := ret[0].(*contractSFFLTaskManager.ContractSFFLTaskManagerCheckpointTaskResponded)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParseCheckpointTaskResponded indicates an expected call of ParseCheckpointTaskResponded.
func (mr *MockAvsSubscribererMockRecorder) ParseCheckpointTaskResponded(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseCheckpointTaskResponded", reflect.TypeOf((*MockAvsSubscriberer)(nil).ParseCheckpointTaskResponded), arg0)
}

// SubscribeToNewPubkeyRegistrations mocks base method.
func (m *MockAvsSubscriberer) SubscribeToNewPubkeyRegistrations() (chan *contractBLSApkRegistry.ContractBLSApkRegistryNewPubkeyRegistration, event.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewPubkeyRegistrations")
	ret0, _ := ret[0].(chan *contractBLSApkRegistry.ContractBLSApkRegistryNewPubkeyRegistration)
	ret1, _ := ret[1].(event.Subscription)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToNewPubkeyRegistrations indicates an expected call of SubscribeToNewPubkeyRegistrations.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToNewPubkeyRegistrations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewPubkeyRegistrations", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToNewPubkeyRegistrations))
}

// SubscribeToNewTasks mocks base method.
func (m *MockAvsSubscriberer) SubscribeToNewTasks(arg0 chan *contractSFFLTaskManager.ContractSFFLTaskManagerCheckpointTaskCreated) (event.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToNewTasks", arg0)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToNewTasks indicates an expected call of SubscribeToNewTasks.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToNewTasks(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToNewTasks", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToNewTasks), arg0)
}

// SubscribeToOperatorSetUpdates mocks base method.
func (m *MockAvsSubscriberer) SubscribeToOperatorSetUpdates(arg0 chan *contractSFFLOperatorSetUpdateRegistry.ContractSFFLOperatorSetUpdateRegistryOperatorSetUpdatedAtBlock) (event.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToOperatorSetUpdates", arg0)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToOperatorSetUpdates indicates an expected call of SubscribeToOperatorSetUpdates.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToOperatorSetUpdates(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToOperatorSetUpdates", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToOperatorSetUpdates), arg0)
}

// SubscribeToOperatorSocketUpdates mocks base method.
func (m *MockAvsSubscriberer) SubscribeToOperatorSocketUpdates() (chan *contractRegistryCoordinator.ContractRegistryCoordinatorOperatorSocketUpdate, event.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToOperatorSocketUpdates")
	ret0, _ := ret[0].(chan *contractRegistryCoordinator.ContractRegistryCoordinatorOperatorSocketUpdate)
	ret1, _ := ret[1].(event.Subscription)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SubscribeToOperatorSocketUpdates indicates an expected call of SubscribeToOperatorSocketUpdates.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToOperatorSocketUpdates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToOperatorSocketUpdates", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToOperatorSocketUpdates))
}

// SubscribeToTaskResponses mocks base method.
func (m *MockAvsSubscriberer) SubscribeToTaskResponses(arg0 chan *contractSFFLTaskManager.ContractSFFLTaskManagerCheckpointTaskResponded) (event.Subscription, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeToTaskResponses", arg0)
	ret0, _ := ret[0].(event.Subscription)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubscribeToTaskResponses indicates an expected call of SubscribeToTaskResponses.
func (mr *MockAvsSubscribererMockRecorder) SubscribeToTaskResponses(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeToTaskResponses", reflect.TypeOf((*MockAvsSubscriberer)(nil).SubscribeToTaskResponses), arg0)
}
