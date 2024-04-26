// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/duc-cnzj/mars/v4/internal/plugins (interfaces: GitServer)
//
// Generated by this command:
//
//	mockgen -destination ../mock/mock_git_server.go -package mock github.com/duc-cnzj/mars/v4/internal/plugins GitServer
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	contracts "github.com/duc-cnzj/mars/v4/internal/contracts"
	gomock "go.uber.org/mock/gomock"
)

// MockGitServer is a mock of GitServer interface.
type MockGitServer struct {
	ctrl     *gomock.Controller
	recorder *MockGitServerMockRecorder
}

// MockGitServerMockRecorder is the mock recorder for MockGitServer.
type MockGitServerMockRecorder struct {
	mock *MockGitServer
}

// NewMockGitServer creates a new mock instance.
func NewMockGitServer(ctrl *gomock.Controller) *MockGitServer {
	mock := &MockGitServer{ctrl: ctrl}
	mock.recorder = &MockGitServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGitServer) EXPECT() *MockGitServerMockRecorder {
	return m.recorder
}

// AllBranches mocks base method.
func (m *MockGitServer) AllBranches(arg0 string) ([]contracts.BranchInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllBranches", arg0)
	ret0, _ := ret[0].([]contracts.BranchInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllBranches indicates an expected call of AllBranches.
func (mr *MockGitServerMockRecorder) AllBranches(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllBranches", reflect.TypeOf((*MockGitServer)(nil).AllBranches), arg0)
}

// AllProjects mocks base method.
func (m *MockGitServer) AllProjects() ([]contracts.ProjectInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllProjects")
	ret0, _ := ret[0].([]contracts.ProjectInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllProjects indicates an expected call of AllProjects.
func (mr *MockGitServerMockRecorder) AllProjects() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllProjects", reflect.TypeOf((*MockGitServer)(nil).AllProjects))
}

// Destroy mocks base method.
func (m *MockGitServer) Destroy() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy")
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockGitServerMockRecorder) Destroy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockGitServer)(nil).Destroy))
}

// GetCommit mocks base method.
func (m *MockGitServer) GetCommit(arg0, arg1 string) (contracts.CommitInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommit", arg0, arg1)
	ret0, _ := ret[0].(contracts.CommitInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommit indicates an expected call of GetCommit.
func (mr *MockGitServerMockRecorder) GetCommit(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommit", reflect.TypeOf((*MockGitServer)(nil).GetCommit), arg0, arg1)
}

// GetCommitPipeline mocks base method.
func (m *MockGitServer) GetCommitPipeline(arg0, arg1, arg2 string) (contracts.PipelineInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitPipeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(contracts.PipelineInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitPipeline indicates an expected call of GetCommitPipeline.
func (mr *MockGitServerMockRecorder) GetCommitPipeline(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitPipeline", reflect.TypeOf((*MockGitServer)(nil).GetCommitPipeline), arg0, arg1, arg2)
}

// GetDirectoryFilesWithBranch mocks base method.
func (m *MockGitServer) GetDirectoryFilesWithBranch(arg0, arg1, arg2 string, arg3 bool) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirectoryFilesWithBranch", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirectoryFilesWithBranch indicates an expected call of GetDirectoryFilesWithBranch.
func (mr *MockGitServerMockRecorder) GetDirectoryFilesWithBranch(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirectoryFilesWithBranch", reflect.TypeOf((*MockGitServer)(nil).GetDirectoryFilesWithBranch), arg0, arg1, arg2, arg3)
}

// GetDirectoryFilesWithSha mocks base method.
func (m *MockGitServer) GetDirectoryFilesWithSha(arg0, arg1, arg2 string, arg3 bool) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDirectoryFilesWithSha", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDirectoryFilesWithSha indicates an expected call of GetDirectoryFilesWithSha.
func (mr *MockGitServerMockRecorder) GetDirectoryFilesWithSha(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDirectoryFilesWithSha", reflect.TypeOf((*MockGitServer)(nil).GetDirectoryFilesWithSha), arg0, arg1, arg2, arg3)
}

// GetFileContentWithBranch mocks base method.
func (m *MockGitServer) GetFileContentWithBranch(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileContentWithBranch", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileContentWithBranch indicates an expected call of GetFileContentWithBranch.
func (mr *MockGitServerMockRecorder) GetFileContentWithBranch(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileContentWithBranch", reflect.TypeOf((*MockGitServer)(nil).GetFileContentWithBranch), arg0, arg1, arg2)
}

// GetFileContentWithSha mocks base method.
func (m *MockGitServer) GetFileContentWithSha(arg0, arg1, arg2 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFileContentWithSha", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFileContentWithSha indicates an expected call of GetFileContentWithSha.
func (mr *MockGitServerMockRecorder) GetFileContentWithSha(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFileContentWithSha", reflect.TypeOf((*MockGitServer)(nil).GetFileContentWithSha), arg0, arg1, arg2)
}

// GetProject mocks base method.
func (m *MockGitServer) GetProject(arg0 string) (contracts.ProjectInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", arg0)
	ret0, _ := ret[0].(contracts.ProjectInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockGitServerMockRecorder) GetProject(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockGitServer)(nil).GetProject), arg0)
}

// Initialize mocks base method.
func (m *MockGitServer) Initialize(arg0 map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Initialize", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Initialize indicates an expected call of Initialize.
func (mr *MockGitServerMockRecorder) Initialize(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Initialize", reflect.TypeOf((*MockGitServer)(nil).Initialize), arg0)
}

// ListBranches mocks base method.
func (m *MockGitServer) ListBranches(arg0 string, arg1, arg2 int) (contracts.ListBranchResponseInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBranches", arg0, arg1, arg2)
	ret0, _ := ret[0].(contracts.ListBranchResponseInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBranches indicates an expected call of ListBranches.
func (mr *MockGitServerMockRecorder) ListBranches(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBranches", reflect.TypeOf((*MockGitServer)(nil).ListBranches), arg0, arg1, arg2)
}

// ListCommits mocks base method.
func (m *MockGitServer) ListCommits(arg0, arg1 string) ([]contracts.CommitInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListCommits", arg0, arg1)
	ret0, _ := ret[0].([]contracts.CommitInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListCommits indicates an expected call of ListCommits.
func (mr *MockGitServerMockRecorder) ListCommits(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListCommits", reflect.TypeOf((*MockGitServer)(nil).ListCommits), arg0, arg1)
}

// ListProjects mocks base method.
func (m *MockGitServer) ListProjects(arg0, arg1 int) (contracts.ListProjectResponseInterface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProjects", arg0, arg1)
	ret0, _ := ret[0].(contracts.ListProjectResponseInterface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockGitServerMockRecorder) ListProjects(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockGitServer)(nil).ListProjects), arg0, arg1)
}

// Name mocks base method.
func (m *MockGitServer) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockGitServerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockGitServer)(nil).Name))
}
