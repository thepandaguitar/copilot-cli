// Code generated by MockGen. DO NOT EDIT.
// Source: ./pkg/archer/env.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	archer "github.com/aws/PRIVATE-amazon-ecs-archer/pkg/archer"
	gomock "github.com/golang/mock/gomock"
)

// MockEnvironmentStore is a mock of EnvironmentStore interface
type MockEnvironmentStore struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentStoreMockRecorder
}

// MockEnvironmentStoreMockRecorder is the mock recorder for MockEnvironmentStore
type MockEnvironmentStoreMockRecorder struct {
	mock *MockEnvironmentStore
}

// NewMockEnvironmentStore creates a new mock instance
func NewMockEnvironmentStore(ctrl *gomock.Controller) *MockEnvironmentStore {
	mock := &MockEnvironmentStore{ctrl: ctrl}
	mock.recorder = &MockEnvironmentStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentStore) EXPECT() *MockEnvironmentStoreMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockEnvironmentStore) ListEnvironments(projectName string) ([]*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", projectName)
	ret0, _ := ret[0].([]*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockEnvironmentStoreMockRecorder) ListEnvironments(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockEnvironmentStore)(nil).ListEnvironments), projectName)
}

// GetEnvironment mocks base method
func (m *MockEnvironmentStore) GetEnvironment(projectName, environmentName string) (*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment", projectName, environmentName)
	ret0, _ := ret[0].(*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment
func (mr *MockEnvironmentStoreMockRecorder) GetEnvironment(projectName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockEnvironmentStore)(nil).GetEnvironment), projectName, environmentName)
}

// CreateEnvironment mocks base method
func (m *MockEnvironmentStore) CreateEnvironment(environment *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvironment", environment)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvironment indicates an expected call of CreateEnvironment
func (mr *MockEnvironmentStoreMockRecorder) CreateEnvironment(environment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockEnvironmentStore)(nil).CreateEnvironment), environment)
}

// MockEnvironmentLister is a mock of EnvironmentLister interface
type MockEnvironmentLister struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentListerMockRecorder
}

// MockEnvironmentListerMockRecorder is the mock recorder for MockEnvironmentLister
type MockEnvironmentListerMockRecorder struct {
	mock *MockEnvironmentLister
}

// NewMockEnvironmentLister creates a new mock instance
func NewMockEnvironmentLister(ctrl *gomock.Controller) *MockEnvironmentLister {
	mock := &MockEnvironmentLister{ctrl: ctrl}
	mock.recorder = &MockEnvironmentListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentLister) EXPECT() *MockEnvironmentListerMockRecorder {
	return m.recorder
}

// Execute mocks base method
func (m *MockEnvironmentLister) ListEnvironments(projectName string) ([]*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", projectName)
	ret0, _ := ret[0].([]*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute
func (mr *MockEnvironmentListerMockRecorder) ListEnvironments(projectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockEnvironmentLister)(nil).ListEnvironments), projectName)
}

// MockEnvironmentGetter is a mock of EnvironmentGetter interface
type MockEnvironmentGetter struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentGetterMockRecorder
}

// MockEnvironmentGetterMockRecorder is the mock recorder for MockEnvironmentGetter
type MockEnvironmentGetterMockRecorder struct {
	mock *MockEnvironmentGetter
}

// NewMockEnvironmentGetter creates a new mock instance
func NewMockEnvironmentGetter(ctrl *gomock.Controller) *MockEnvironmentGetter {
	mock := &MockEnvironmentGetter{ctrl: ctrl}
	mock.recorder = &MockEnvironmentGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentGetter) EXPECT() *MockEnvironmentGetterMockRecorder {
	return m.recorder
}

// GetEnvironment mocks base method
func (m *MockEnvironmentGetter) GetEnvironment(projectName, environmentName string) (*archer.Environment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEnvironment", projectName, environmentName)
	ret0, _ := ret[0].(*archer.Environment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironment indicates an expected call of GetEnvironment
func (mr *MockEnvironmentGetterMockRecorder) GetEnvironment(projectName, environmentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironment", reflect.TypeOf((*MockEnvironmentGetter)(nil).GetEnvironment), projectName, environmentName)
}

// MockEnvironmentCreator is a mock of EnvironmentCreator interface
type MockEnvironmentCreator struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentCreatorMockRecorder
}

// MockEnvironmentCreatorMockRecorder is the mock recorder for MockEnvironmentCreator
type MockEnvironmentCreatorMockRecorder struct {
	mock *MockEnvironmentCreator
}

// NewMockEnvironmentCreator creates a new mock instance
func NewMockEnvironmentCreator(ctrl *gomock.Controller) *MockEnvironmentCreator {
	mock := &MockEnvironmentCreator{ctrl: ctrl}
	mock.recorder = &MockEnvironmentCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentCreator) EXPECT() *MockEnvironmentCreatorMockRecorder {
	return m.recorder
}

// CreateEnvironment mocks base method
func (m *MockEnvironmentCreator) CreateEnvironment(environment *archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEnvironment", environment)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEnvironment indicates an expected call of CreateEnvironment
func (mr *MockEnvironmentCreatorMockRecorder) CreateEnvironment(environment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironment", reflect.TypeOf((*MockEnvironmentCreator)(nil).CreateEnvironment), environment)
}

// MockEnvironmentDeployer is a mock of EnvironmentDeployer interface
type MockEnvironmentDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockEnvironmentDeployerMockRecorder
}

// MockEnvironmentDeployerMockRecorder is the mock recorder for MockEnvironmentDeployer
type MockEnvironmentDeployerMockRecorder struct {
	mock *MockEnvironmentDeployer
}

// NewMockEnvironmentDeployer creates a new mock instance
func NewMockEnvironmentDeployer(ctrl *gomock.Controller) *MockEnvironmentDeployer {
	mock := &MockEnvironmentDeployer{ctrl: ctrl}
	mock.recorder = &MockEnvironmentDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEnvironmentDeployer) EXPECT() *MockEnvironmentDeployerMockRecorder {
	return m.recorder
}

// DeployEnvironment mocks base method
func (m *MockEnvironmentDeployer) DeployEnvironment(env archer.Environment, includeLoadBalancer bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeployEnvironment", env, includeLoadBalancer)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeployEnvironment indicates an expected call of DeployEnvironment
func (mr *MockEnvironmentDeployerMockRecorder) DeployEnvironment(env, includeLoadBalancer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeployEnvironment", reflect.TypeOf((*MockEnvironmentDeployer)(nil).DeployEnvironment), env, includeLoadBalancer)
}

// Wait mocks base method
func (m *MockEnvironmentDeployer) Wait(env archer.Environment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", env)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait
func (mr *MockEnvironmentDeployerMockRecorder) Wait(env interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockEnvironmentDeployer)(nil).Wait), env)
}
