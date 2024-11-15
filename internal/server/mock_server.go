// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/v5/internal/server (interfaces: HttpServer,GrpcServerImp)
//
// Generated by this command:
//
//	mockgen -destination ./mock_server.go -package server github.com/duc-cnzj/mars/v5/internal/server HttpServer,GrpcServerImp
//

// Package server is a generated GoMock package.
package server

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockHttpServer is a mock of HttpServer interface.
type MockHttpServer struct {
	ctrl     *gomock.Controller
	recorder *MockHttpServerMockRecorder
}

// MockHttpServerMockRecorder is the mock recorder for MockHttpServer.
type MockHttpServerMockRecorder struct {
	mock *MockHttpServer
}

// NewMockHttpServer creates a new mock instance.
func NewMockHttpServer(ctrl *gomock.Controller) *MockHttpServer {
	mock := &MockHttpServer{ctrl: ctrl}
	mock.recorder = &MockHttpServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpServer) EXPECT() *MockHttpServerMockRecorder {
	return m.recorder
}

// ListenAndServe mocks base method.
func (m *MockHttpServer) ListenAndServe() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListenAndServe")
	ret0, _ := ret[0].(error)
	return ret0
}

// ListenAndServe indicates an expected call of ListenAndServe.
func (mr *MockHttpServerMockRecorder) ListenAndServe() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListenAndServe", reflect.TypeOf((*MockHttpServer)(nil).ListenAndServe))
}

// Shutdown mocks base method.
func (m *MockHttpServer) Shutdown(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shutdown", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Shutdown indicates an expected call of Shutdown.
func (mr *MockHttpServerMockRecorder) Shutdown(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shutdown", reflect.TypeOf((*MockHttpServer)(nil).Shutdown), arg0)
}

// MockGrpcServerImp is a mock of GrpcServerImp interface.
type MockGrpcServerImp struct {
	ctrl     *gomock.Controller
	recorder *MockGrpcServerImpMockRecorder
}

// MockGrpcServerImpMockRecorder is the mock recorder for MockGrpcServerImp.
type MockGrpcServerImpMockRecorder struct {
	mock *MockGrpcServerImp
}

// NewMockGrpcServerImp creates a new mock instance.
func NewMockGrpcServerImp(ctrl *gomock.Controller) *MockGrpcServerImp {
	mock := &MockGrpcServerImp{ctrl: ctrl}
	mock.recorder = &MockGrpcServerImpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGrpcServerImp) EXPECT() *MockGrpcServerImpMockRecorder {
	return m.recorder
}

// GracefulStop mocks base method.
func (m *MockGrpcServerImp) GracefulStop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GracefulStop")
}

// GracefulStop indicates an expected call of GracefulStop.
func (mr *MockGrpcServerImpMockRecorder) GracefulStop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GracefulStop", reflect.TypeOf((*MockGrpcServerImp)(nil).GracefulStop))
}

// Serve mocks base method.
func (m *MockGrpcServerImp) Serve(arg0 net.Listener) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serve", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Serve indicates an expected call of Serve.
func (mr *MockGrpcServerImpMockRecorder) Serve(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serve", reflect.TypeOf((*MockGrpcServerImp)(nil).Serve), arg0)
}
