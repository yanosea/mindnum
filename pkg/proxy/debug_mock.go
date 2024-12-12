// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/proxy/debug.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/proxy/debug.go -destination=./pkg/proxy/debug_mock.go -package=proxy
//

// Package proxy is a generated GoMock package.
package proxy

import (
	reflect "reflect"
	debug "runtime/debug"

	gomock "go.uber.org/mock/gomock"
)

// MockDebug is a mock of Debug interface.
type MockDebug struct {
	ctrl     *gomock.Controller
	recorder *MockDebugMockRecorder
}

// MockDebugMockRecorder is the mock recorder for MockDebug.
type MockDebugMockRecorder struct {
	mock *MockDebug
}

// NewMockDebug creates a new mock instance.
func NewMockDebug(ctrl *gomock.Controller) *MockDebug {
	mock := &MockDebug{ctrl: ctrl}
	mock.recorder = &MockDebugMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDebug) EXPECT() *MockDebugMockRecorder {
	return m.recorder
}

// ReadBuildInfo mocks base method.
func (m *MockDebug) ReadBuildInfo() (*debug.BuildInfo, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadBuildInfo")
	ret0, _ := ret[0].(*debug.BuildInfo)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// ReadBuildInfo indicates an expected call of ReadBuildInfo.
func (mr *MockDebugMockRecorder) ReadBuildInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadBuildInfo", reflect.TypeOf((*MockDebug)(nil).ReadBuildInfo))
}