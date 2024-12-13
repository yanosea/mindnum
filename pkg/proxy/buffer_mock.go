// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/proxy/buffer.go
//
// Generated by this command:
//
//	mockgen -source=./pkg/proxy/buffer.go -destination=./pkg/proxy/buffer_mock.go -package=proxy
//

// Package proxy is a generated GoMock package.
package proxy

import (
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockBuffer is a mock of Buffer interface.
type MockBuffer struct {
	ctrl     *gomock.Controller
	recorder *MockBufferMockRecorder
}

// MockBufferMockRecorder is the mock recorder for MockBuffer.
type MockBufferMockRecorder struct {
	mock *MockBuffer
}

// NewMockBuffer creates a new mock instance.
func NewMockBuffer(ctrl *gomock.Controller) *MockBuffer {
	mock := &MockBuffer{ctrl: ctrl}
	mock.recorder = &MockBufferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBuffer) EXPECT() *MockBufferMockRecorder {
	return m.recorder
}

// ReadFrom mocks base method.
func (m *MockBuffer) ReadFrom(r io.Reader) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadFrom", r)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadFrom indicates an expected call of ReadFrom.
func (mr *MockBufferMockRecorder) ReadFrom(r any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadFrom", reflect.TypeOf((*MockBuffer)(nil).ReadFrom), r)
}

// Reset mocks base method.
func (m *MockBuffer) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *MockBufferMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*MockBuffer)(nil).Reset))
}

// String mocks base method.
func (m *MockBuffer) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockBufferMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockBuffer)(nil).String))
}
