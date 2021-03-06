// Code generated by MockGen. DO NOT EDIT.
// Source: src/server/repositories/repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "github.com/walmartdigital/katalog/domain"
	repositories "github.com/walmartdigital/katalog/server/repositories"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateResource mocks base method
func (m *MockRepository) CreateResource(obj interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResource", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateResource indicates an expected call of CreateResource
func (mr *MockRepositoryMockRecorder) CreateResource(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResource", reflect.TypeOf((*MockRepository)(nil).CreateResource), obj)
}

// UpdateResource mocks base method
func (m *MockRepository) UpdateResource(obj interface{}) (*domain.Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResource", obj)
	ret0, _ := ret[0].(*domain.Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateResource indicates an expected call of UpdateResource
func (mr *MockRepositoryMockRecorder) UpdateResource(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResource", reflect.TypeOf((*MockRepository)(nil).UpdateResource), obj)
}

// DeleteResource mocks base method
func (m *MockRepository) DeleteResource(obj interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResource", obj)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResource indicates an expected call of DeleteResource
func (mr *MockRepositoryMockRecorder) DeleteResource(obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResource", reflect.TypeOf((*MockRepository)(nil).DeleteResource), obj)
}

// GetAllResources mocks base method
func (m *MockRepository) GetAllResources() ([]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllResources")
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllResources indicates an expected call of GetAllResources
func (mr *MockRepositoryMockRecorder) GetAllResources() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResources", reflect.TypeOf((*MockRepository)(nil).GetAllResources))
}

// GetResource mocks base method
func (m *MockRepository) GetResource(id string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResource", id)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResource indicates an expected call of GetResource
func (mr *MockRepositoryMockRecorder) GetResource(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResource", reflect.TypeOf((*MockRepository)(nil).GetResource), id)
}

// MockRepositoryFactory is a mock of RepositoryFactory interface
type MockRepositoryFactory struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryFactoryMockRecorder
}

// MockRepositoryFactoryMockRecorder is the mock recorder for MockRepositoryFactory
type MockRepositoryFactoryMockRecorder struct {
	mock *MockRepositoryFactory
}

// NewMockRepositoryFactory creates a new mock instance
func NewMockRepositoryFactory(ctrl *gomock.Controller) *MockRepositoryFactory {
	mock := &MockRepositoryFactory{ctrl: ctrl}
	mock.recorder = &MockRepositoryFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepositoryFactory) EXPECT() *MockRepositoryFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockRepositoryFactory) Create() repositories.Repository {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create")
	ret0, _ := ret[0].(repositories.Repository)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockRepositoryFactoryMockRecorder) Create() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepositoryFactory)(nil).Create))
}
