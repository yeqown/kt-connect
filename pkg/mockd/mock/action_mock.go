// Code generated by MockGen. DO NOT EDIT.
// Source: ../kt/command/types.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	options "github.com/alibaba/kt-connect/pkg/kt/options"
	gomock "github.com/golang/mock/gomock"
)

// MockActionInterface is a mock of ActionInterface interface
type MockActionInterface struct {
	ctrl     *gomock.Controller
	recorder *MockActionInterfaceMockRecorder
}

// MockActionInterfaceMockRecorder is the mock recorder for MockActionInterface
type MockActionInterfaceMockRecorder struct {
	mock *MockActionInterface
}

// NewMockActionInterface creates a new mock instance
func NewMockActionInterface(ctrl *gomock.Controller) *MockActionInterface {
	mock := &MockActionInterface{ctrl: ctrl}
	mock.recorder = &MockActionInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockActionInterface) EXPECT() *MockActionInterfaceMockRecorder {
	return m.recorder
}

// OpenDashboard mocks base method
func (m *MockActionInterface) OpenDashboard(options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenDashboard", options)
	ret0, _ := ret[0].(error)
	return ret0
}

// OpenDashboard indicates an expected call of OpenDashboard
func (mr *MockActionInterfaceMockRecorder) OpenDashboard(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenDashboard", reflect.TypeOf((*MockActionInterface)(nil).OpenDashboard), options)
}

// Connect mocks base method
func (m *MockActionInterface) Connect(options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connect", options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Connect indicates an expected call of Connect
func (mr *MockActionInterfaceMockRecorder) Connect(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connect", reflect.TypeOf((*MockActionInterface)(nil).Connect), options)
}

// Check mocks base method
func (m *MockActionInterface) Check(options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Check", options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Check indicates an expected call of Check
func (mr *MockActionInterfaceMockRecorder) Check(options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockActionInterface)(nil).Check), options)
}

// Run mocks base method
func (m *MockActionInterface) Run(service string, options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", service, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockActionInterfaceMockRecorder) Run(service, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockActionInterface)(nil).Run), service, options)
}

// Exchange mocks base method
func (m *MockActionInterface) Exchange(service string, options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exchange", service, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Exchange indicates an expected call of Exchange
func (mr *MockActionInterfaceMockRecorder) Exchange(service, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exchange", reflect.TypeOf((*MockActionInterface)(nil).Exchange), service, options)
}

// Mesh mocks base method
func (m *MockActionInterface) Mesh(service string, options *options.DaemonOptions) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mesh", service, options)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mesh indicates an expected call of Mesh
func (mr *MockActionInterfaceMockRecorder) Mesh(service, options interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mesh", reflect.TypeOf((*MockActionInterface)(nil).Mesh), service, options)
}

// ApplyDashboard mocks base method
func (m *MockActionInterface) ApplyDashboard() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyDashboard")
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyDashboard indicates an expected call of ApplyDashboard
func (mr *MockActionInterfaceMockRecorder) ApplyDashboard() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyDashboard", reflect.TypeOf((*MockActionInterface)(nil).ApplyDashboard))
}