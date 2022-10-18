// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/internal/contracts (interfaces: RemoteExecutor)

// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	contracts "github.com/duc-cnzj/mars/internal/contracts"
	gomock "github.com/golang/mock/gomock"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	remotecommand "k8s.io/client-go/tools/remotecommand"
)

// MockRemoteExecutor is a mock of RemoteExecutor interface.
type MockRemoteExecutor struct {
	ctrl     *gomock.Controller
	recorder *MockRemoteExecutorMockRecorder
}

// MockRemoteExecutorMockRecorder is the mock recorder for MockRemoteExecutor.
type MockRemoteExecutorMockRecorder struct {
	mock *MockRemoteExecutor
}

// NewMockRemoteExecutor creates a new mock instance.
func NewMockRemoteExecutor(ctrl *gomock.Controller) *MockRemoteExecutor {
	mock := &MockRemoteExecutor{ctrl: ctrl}
	mock.recorder = &MockRemoteExecutorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRemoteExecutor) EXPECT() *MockRemoteExecutorMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockRemoteExecutor) Execute(arg0 kubernetes.Interface, arg1 *rest.Config, arg2 io.Reader, arg3, arg4 io.Writer, arg5 bool, arg6 remotecommand.TerminalSizeQueue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockRemoteExecutorMockRecorder) Execute(arg0, arg1, arg2, arg3, arg4, arg5, arg6 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockRemoteExecutor)(nil).Execute), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// WithCommand mocks base method.
func (m *MockRemoteExecutor) WithCommand(arg0 []string) contracts.RemoteExecutor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithCommand", arg0)
	ret0, _ := ret[0].(contracts.RemoteExecutor)
	return ret0
}

// WithCommand indicates an expected call of WithCommand.
func (mr *MockRemoteExecutorMockRecorder) WithCommand(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithCommand", reflect.TypeOf((*MockRemoteExecutor)(nil).WithCommand), arg0)
}

// WithContainer mocks base method.
func (m *MockRemoteExecutor) WithContainer(arg0, arg1, arg2 string) contracts.RemoteExecutor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithContainer", arg0, arg1, arg2)
	ret0, _ := ret[0].(contracts.RemoteExecutor)
	return ret0
}

// WithContainer indicates an expected call of WithContainer.
func (mr *MockRemoteExecutorMockRecorder) WithContainer(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithContainer", reflect.TypeOf((*MockRemoteExecutor)(nil).WithContainer), arg0, arg1, arg2)
}

// WithMethod mocks base method.
func (m *MockRemoteExecutor) WithMethod(arg0 string) contracts.RemoteExecutor {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithMethod", arg0)
	ret0, _ := ret[0].(contracts.RemoteExecutor)
	return ret0
}

// WithMethod indicates an expected call of WithMethod.
func (mr *MockRemoteExecutorMockRecorder) WithMethod(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithMethod", reflect.TypeOf((*MockRemoteExecutor)(nil).WithMethod), arg0)
}