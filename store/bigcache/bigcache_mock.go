// Code generated by MockGen. DO NOT EDIT.
// Source: store/bigcache/bigcache.go

// Package bigcache is a generated GoMock package.
package bigcache

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBigcacheClientInterface is a mock of BigcacheClientInterface interface.
type MockBigcacheClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockBigcacheClientInterfaceMockRecorder
}

// MockBigcacheClientInterfaceMockRecorder is the mock recorder for MockBigcacheClientInterface.
type MockBigcacheClientInterfaceMockRecorder struct {
	mock *MockBigcacheClientInterface
}

// NewMockBigcacheClientInterface creates a new mock instance.
func NewMockBigcacheClientInterface(ctrl *gomock.Controller) *MockBigcacheClientInterface {
	mock := &MockBigcacheClientInterface{ctrl: ctrl}
	mock.recorder = &MockBigcacheClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBigcacheClientInterface) EXPECT() *MockBigcacheClientInterfaceMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBigcacheClientInterface) Delete(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBigcacheClientInterfaceMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBigcacheClientInterface)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockBigcacheClientInterface) Get(key string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBigcacheClientInterfaceMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBigcacheClientInterface)(nil).Get), key)
}

// Reset mocks base method.
func (m *MockBigcacheClientInterface) Reset() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reset")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reset indicates an expected call of Reset.
func (mr *MockBigcacheClientInterfaceMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockBigcacheClientInterface)(nil).Reset))
}

// Set mocks base method.
func (m *MockBigcacheClientInterface) Set(key string, entry []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", key, entry)
	ret0, _ := ret[0].(error)
	return ret0
}

// Set indicates an expected call of Set.
func (mr *MockBigcacheClientInterfaceMockRecorder) Set(key, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockBigcacheClientInterface)(nil).Set), key, entry)
}
