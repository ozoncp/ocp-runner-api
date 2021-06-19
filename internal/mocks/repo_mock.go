// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ozoncp/ocp-runner-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/ozoncp/ocp-runner-api/internal/models"
	ocp_runner_api "github.com/ozoncp/ocp-runner-api/pkg/ocp-runner-api"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddRunner mocks base method.
func (m *MockRepo) AddRunner(arg0 context.Context, arg1 *models.Runner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRunner", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRunner indicates an expected call of AddRunner.
func (mr *MockRepoMockRecorder) AddRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRunner", reflect.TypeOf((*MockRepo)(nil).AddRunner), arg0, arg1)
}

// AddRunners mocks base method.
func (m *MockRepo) AddRunners(arg0 context.Context, arg1 []*models.Runner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRunners", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRunners indicates an expected call of AddRunners.
func (mr *MockRepoMockRecorder) AddRunners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRunners", reflect.TypeOf((*MockRepo)(nil).AddRunners), arg0, arg1)
}

// Close mocks base method.
func (m *MockRepo) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockRepoMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockRepo)(nil).Close))
}

// DescribeRunner mocks base method.
func (m *MockRepo) DescribeRunner(arg0 context.Context, arg1 string) (*models.Runner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeRunner", arg0, arg1)
	ret0, _ := ret[0].(*models.Runner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeRunner indicates an expected call of DescribeRunner.
func (mr *MockRepoMockRecorder) DescribeRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeRunner", reflect.TypeOf((*MockRepo)(nil).DescribeRunner), arg0, arg1)
}

// ListRunners mocks base method.
func (m *MockRepo) ListRunners(arg0 context.Context, arg1 *ocp_runner_api.ListFiltersRequest) ([]*models.Runner, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRunners", arg0, arg1)
	ret0, _ := ret[0].([]*models.Runner)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRunners indicates an expected call of ListRunners.
func (mr *MockRepoMockRecorder) ListRunners(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRunners", reflect.TypeOf((*MockRepo)(nil).ListRunners), arg0, arg1)
}

// RemoveRunner mocks base method.
func (m *MockRepo) RemoveRunner(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRunner", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveRunner indicates an expected call of RemoveRunner.
func (mr *MockRepoMockRecorder) RemoveRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRunner", reflect.TypeOf((*MockRepo)(nil).RemoveRunner), arg0, arg1)
}

// UpdateRunner mocks base method.
func (m *MockRepo) UpdateRunner(arg0 context.Context, arg1 *models.Runner) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRunner", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateRunner indicates an expected call of UpdateRunner.
func (mr *MockRepoMockRecorder) UpdateRunner(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRunner", reflect.TypeOf((*MockRepo)(nil).UpdateRunner), arg0, arg1)
}
