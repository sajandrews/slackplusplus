// Code generated by MockGen. DO NOT EDIT.
// Source: plusplus.go

// Package plusplus is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	slack "github.com/nlopes/slack"
)

// MockSlacker is a mock of Slacker interface
type MockSlacker struct {
	ctrl     *gomock.Controller
	recorder *MockSlackerMockRecorder
}

// MockSlackerMockRecorder is the mock recorder for MockSlacker
type MockSlackerMockRecorder struct {
	mock *MockSlacker
}

// NewMockSlacker creates a new mock instance
func NewMockSlacker(ctrl *gomock.Controller) *MockSlacker {
	mock := &MockSlacker{ctrl: ctrl}
	mock.recorder = &MockSlackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSlacker) EXPECT() *MockSlackerMockRecorder {
	return m.recorder
}

// GetInfo mocks base method
func (m *MockSlacker) GetInfo() *slack.Info {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInfo")
	ret0, _ := ret[0].(*slack.Info)
	return ret0
}

// GetInfo indicates an expected call of GetInfo
func (mr *MockSlackerMockRecorder) GetInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInfo", reflect.TypeOf((*MockSlacker)(nil).GetInfo))
}

// GetUserInfo mocks base method
func (m *MockSlacker) GetUserInfo(arg0 string) (*slack.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserInfo", arg0)
	ret0, _ := ret[0].(*slack.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserInfo indicates an expected call of GetUserInfo
func (mr *MockSlackerMockRecorder) GetUserInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserInfo", reflect.TypeOf((*MockSlacker)(nil).GetUserInfo), arg0)
}

// SendMessage mocks base method
func (m *MockSlacker) SendMessage(arg0 *slack.OutgoingMessage) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SendMessage", arg0)
}

// SendMessage indicates an expected call of SendMessage
func (mr *MockSlackerMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockSlacker)(nil).SendMessage), arg0)
}

// NewOutgoingMessage mocks base method
func (m *MockSlacker) NewOutgoingMessage(arg0, arg1 string, arg2 ...slack.RTMsgOption) *slack.OutgoingMessage {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NewOutgoingMessage", varargs...)
	ret0, _ := ret[0].(*slack.OutgoingMessage)
	return ret0
}

// NewOutgoingMessage indicates an expected call of NewOutgoingMessage
func (mr *MockSlackerMockRecorder) NewOutgoingMessage(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewOutgoingMessage", reflect.TypeOf((*MockSlacker)(nil).NewOutgoingMessage), varargs...)
}

// OpenIMChannel mocks base method
func (m *MockSlacker) OpenIMChannel(arg0 string) (bool, bool, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenIMChannel", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// OpenIMChannel indicates an expected call of OpenIMChannel
func (mr *MockSlackerMockRecorder) OpenIMChannel(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenIMChannel", reflect.TypeOf((*MockSlacker)(nil).OpenIMChannel), arg0)
}