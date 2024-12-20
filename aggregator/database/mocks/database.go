// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nuffle-Labs/nffl/aggregator/database (interfaces: Databaser)
//
// Generated by this command:
//
//	mockgen -destination=./mocks/database.go -package=mocks github.com/Nuffle-Labs/nffl/aggregator/database Databaser
//
// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	models "github.com/Nuffle-Labs/nffl/aggregator/database/models"
	messages "github.com/Nuffle-Labs/nffl/core/types/messages"
	prometheus "github.com/prometheus/client_golang/prometheus"
	gomock "go.uber.org/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockDatabaser is a mock of Databaser interface.
type MockDatabaser struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaserMockRecorder
}

// MockDatabaserMockRecorder is the mock recorder for MockDatabaser.
type MockDatabaserMockRecorder struct {
	mock *MockDatabaser
}

// NewMockDatabaser creates a new mock instance.
func NewMockDatabaser(ctrl *gomock.Controller) *MockDatabaser {
	mock := &MockDatabaser{ctrl: ctrl}
	mock.recorder = &MockDatabaserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabaser) EXPECT() *MockDatabaserMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDatabaser) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatabaserMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDatabaser)(nil).Close))
}

// DB mocks base method.
func (m *MockDatabaser) DB() *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DB")
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// DB indicates an expected call of DB.
func (mr *MockDatabaserMockRecorder) DB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DB", reflect.TypeOf((*MockDatabaser)(nil).DB))
}

// EnableMetrics mocks base method.
func (m *MockDatabaser) EnableMetrics(arg0 *prometheus.Registry) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableMetrics", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableMetrics indicates an expected call of EnableMetrics.
func (mr *MockDatabaserMockRecorder) EnableMetrics(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableMetrics", reflect.TypeOf((*MockDatabaser)(nil).EnableMetrics), arg0)
}

// FetchCheckpointMessages mocks base method.
func (m *MockDatabaser) FetchCheckpointMessages(arg0, arg1 uint64) (*messages.CheckpointMessages, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchCheckpointMessages", arg0, arg1)
	ret0, _ := ret[0].(*messages.CheckpointMessages)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCheckpointMessages indicates an expected call of FetchCheckpointMessages.
func (mr *MockDatabaserMockRecorder) FetchCheckpointMessages(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCheckpointMessages", reflect.TypeOf((*MockDatabaser)(nil).FetchCheckpointMessages), arg0, arg1)
}

// FetchOperatorSetUpdate mocks base method.
func (m *MockDatabaser) FetchOperatorSetUpdate(arg0 uint64) (*messages.OperatorSetUpdateMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOperatorSetUpdate", arg0)
	ret0, _ := ret[0].(*messages.OperatorSetUpdateMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOperatorSetUpdate indicates an expected call of FetchOperatorSetUpdate.
func (mr *MockDatabaserMockRecorder) FetchOperatorSetUpdate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOperatorSetUpdate", reflect.TypeOf((*MockDatabaser)(nil).FetchOperatorSetUpdate), arg0)
}

// FetchOperatorSetUpdateAggregation mocks base method.
func (m *MockDatabaser) FetchOperatorSetUpdateAggregation(arg0 uint64) (*messages.MessageBlsAggregation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchOperatorSetUpdateAggregation", arg0)
	ret0, _ := ret[0].(*messages.MessageBlsAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchOperatorSetUpdateAggregation indicates an expected call of FetchOperatorSetUpdateAggregation.
func (mr *MockDatabaserMockRecorder) FetchOperatorSetUpdateAggregation(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchOperatorSetUpdateAggregation", reflect.TypeOf((*MockDatabaser)(nil).FetchOperatorSetUpdateAggregation), arg0)
}

// FetchStateRootUpdate mocks base method.
func (m *MockDatabaser) FetchStateRootUpdate(arg0 uint32, arg1 uint64) (*messages.StateRootUpdateMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStateRootUpdate", arg0, arg1)
	ret0, _ := ret[0].(*messages.StateRootUpdateMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchStateRootUpdate indicates an expected call of FetchStateRootUpdate.
func (mr *MockDatabaserMockRecorder) FetchStateRootUpdate(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStateRootUpdate", reflect.TypeOf((*MockDatabaser)(nil).FetchStateRootUpdate), arg0, arg1)
}

// FetchStateRootUpdateAggregation mocks base method.
func (m *MockDatabaser) FetchStateRootUpdateAggregation(arg0 uint32, arg1 uint64) (*messages.MessageBlsAggregation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchStateRootUpdateAggregation", arg0, arg1)
	ret0, _ := ret[0].(*messages.MessageBlsAggregation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchStateRootUpdateAggregation indicates an expected call of FetchStateRootUpdateAggregation.
func (mr *MockDatabaserMockRecorder) FetchStateRootUpdateAggregation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchStateRootUpdateAggregation", reflect.TypeOf((*MockDatabaser)(nil).FetchStateRootUpdateAggregation), arg0, arg1)
}

// StoreOperatorSetUpdate mocks base method.
func (m *MockDatabaser) StoreOperatorSetUpdate(arg0 messages.OperatorSetUpdateMessage) (*models.OperatorSetUpdateMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreOperatorSetUpdate", arg0)
	ret0, _ := ret[0].(*models.OperatorSetUpdateMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreOperatorSetUpdate indicates an expected call of StoreOperatorSetUpdate.
func (mr *MockDatabaserMockRecorder) StoreOperatorSetUpdate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreOperatorSetUpdate", reflect.TypeOf((*MockDatabaser)(nil).StoreOperatorSetUpdate), arg0)
}

// StoreOperatorSetUpdateAggregation mocks base method.
func (m *MockDatabaser) StoreOperatorSetUpdateAggregation(arg0 *models.OperatorSetUpdateMessage, arg1 messages.MessageBlsAggregation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreOperatorSetUpdateAggregation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreOperatorSetUpdateAggregation indicates an expected call of StoreOperatorSetUpdateAggregation.
func (mr *MockDatabaserMockRecorder) StoreOperatorSetUpdateAggregation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreOperatorSetUpdateAggregation", reflect.TypeOf((*MockDatabaser)(nil).StoreOperatorSetUpdateAggregation), arg0, arg1)
}

// StoreStateRootUpdate mocks base method.
func (m *MockDatabaser) StoreStateRootUpdate(arg0 messages.StateRootUpdateMessage) (*models.StateRootUpdateMessage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreStateRootUpdate", arg0)
	ret0, _ := ret[0].(*models.StateRootUpdateMessage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StoreStateRootUpdate indicates an expected call of StoreStateRootUpdate.
func (mr *MockDatabaserMockRecorder) StoreStateRootUpdate(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreStateRootUpdate", reflect.TypeOf((*MockDatabaser)(nil).StoreStateRootUpdate), arg0)
}

// StoreStateRootUpdateAggregation mocks base method.
func (m *MockDatabaser) StoreStateRootUpdateAggregation(arg0 *models.StateRootUpdateMessage, arg1 messages.MessageBlsAggregation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreStateRootUpdateAggregation", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreStateRootUpdateAggregation indicates an expected call of StoreStateRootUpdateAggregation.
func (mr *MockDatabaserMockRecorder) StoreStateRootUpdateAggregation(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreStateRootUpdateAggregation", reflect.TypeOf((*MockDatabaser)(nil).StoreStateRootUpdateAggregation), arg0, arg1)
}
