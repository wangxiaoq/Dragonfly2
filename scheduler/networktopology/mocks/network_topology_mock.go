// Code generated by MockGen. DO NOT EDIT.
// Source: network_topology.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	networktopology "d7y.io/dragonfly/v2/scheduler/networktopology"
	gomock "github.com/golang/mock/gomock"
)

// MockNetworkTopology is a mock of NetworkTopology interface.
type MockNetworkTopology struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkTopologyMockRecorder
}

// MockNetworkTopologyMockRecorder is the mock recorder for MockNetworkTopology.
type MockNetworkTopologyMockRecorder struct {
	mock *MockNetworkTopology
}

// NewMockNetworkTopology creates a new mock instance.
func NewMockNetworkTopology(ctrl *gomock.Controller) *MockNetworkTopology {
	mock := &MockNetworkTopology{ctrl: ctrl}
	mock.recorder = &MockNetworkTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkTopology) EXPECT() *MockNetworkTopologyMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockNetworkTopology) Delete(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockNetworkTopologyMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNetworkTopology)(nil).Delete), arg0, arg1)
}

// Has mocks base method.
func (m *MockNetworkTopology) Has(arg0, arg1 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockNetworkTopologyMockRecorder) Has(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockNetworkTopology)(nil).Has), arg0, arg1)
}

// ProbedCount mocks base method.
func (m *MockNetworkTopology) ProbedCount(arg0 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProbedCount", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProbedCount indicates an expected call of ProbedCount.
func (mr *MockNetworkTopologyMockRecorder) ProbedCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProbedCount", reflect.TypeOf((*MockNetworkTopology)(nil).ProbedCount), arg0)
}

// Probes mocks base method.
func (m *MockNetworkTopology) Probes(arg0, arg1 string) networktopology.Probes {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Probes", arg0, arg1)
	ret0, _ := ret[0].(networktopology.Probes)
	return ret0
}

// Probes indicates an expected call of Probes.
func (mr *MockNetworkTopologyMockRecorder) Probes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Probes", reflect.TypeOf((*MockNetworkTopology)(nil).Probes), arg0, arg1)
}

// Store mocks base method.
func (m *MockNetworkTopology) Store(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Store", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Store indicates an expected call of Store.
func (mr *MockNetworkTopologyMockRecorder) Store(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Store", reflect.TypeOf((*MockNetworkTopology)(nil).Store), arg0, arg1)
}
