// Code generated by MockGen. DO NOT EDIT.
// Source: hellomock.go

// Package mock_gomocktest is a generated GoMock package.
package mock_gomocktest

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTakler is a mock of Takler interface
type MockTakler struct {
	ctrl     *gomock.Controller
	recorder *MockTaklerMockRecorder
}

// MockTaklerMockRecorder is the mock recorder for MockTakler
type MockTaklerMockRecorder struct {
	mock *MockTakler
}

// NewMockTakler creates a new mock instance
func NewMockTakler(ctrl *gomock.Controller) *MockTakler {
	mock := &MockTakler{ctrl: ctrl}
	mock.recorder = &MockTaklerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTakler) EXPECT() *MockTaklerMockRecorder {
	return m.recorder
}

// SayHello mocks base method
func (m *MockTakler) SayHello(word string) string {
	ret := m.ctrl.Call(m, "SayHello", word)
	ret0, _ := ret[0].(string)
	return ret0
}

// SayHello indicates an expected call of SayHello
func (mr *MockTaklerMockRecorder) SayHello(word interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SayHello", reflect.TypeOf((*MockTakler)(nil).SayHello), word)
}