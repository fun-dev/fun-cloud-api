// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/container/domain/container/ContainerRepository.go

// Package mock_container is a generated GoMock package.
package mock_container

import (
	container "github.com/fun-dev/fun-cloud-api/internal/container/domain/container"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
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

// GetAllByUserID mocks base method
func (m *MockRepository) GetAllByUserID(userID string) ([]*container.Container, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllByUserID", userID)
	ret0, _ := ret[0].([]*container.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllByUserID indicates an expected call of GetAllByUserID
func (mr *MockRepositoryMockRecorder) GetAllByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllByUserID", reflect.TypeOf((*MockRepository)(nil).GetAllByUserID), userID)
}

// Create mocks base method
func (m *MockRepository) Create(userID, imageName string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userID, imageName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Create indicates an expected call of Create
func (mr *MockRepositoryMockRecorder) Create(userID, imageName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockRepository)(nil).Create), userID, imageName)
}

// DeleteByContainerID mocks base method
func (m *MockRepository) DeleteByContainerID(userID, containerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteByContainerID", userID, containerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteByContainerID indicates an expected call of DeleteByContainerID
func (mr *MockRepositoryMockRecorder) DeleteByContainerID(userID, containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteByContainerID", reflect.TypeOf((*MockRepository)(nil).DeleteByContainerID), userID, containerID)
}

// GetDeploymentManifestByContainerID mocks base method
func (m *MockRepository) GetDeploymentManifestByContainerID(containerID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeploymentManifestByContainerID", containerID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeploymentManifestByContainerID indicates an expected call of GetDeploymentManifestByContainerID
func (mr *MockRepositoryMockRecorder) GetDeploymentManifestByContainerID(containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeploymentManifestByContainerID", reflect.TypeOf((*MockRepository)(nil).GetDeploymentManifestByContainerID), containerID)
}

// SaveDeploymentManifestByContainerID mocks base method
func (m *MockRepository) SaveDeploymentManifestByContainerID(containerID, manifest string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveDeploymentManifestByContainerID", containerID, manifest)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveDeploymentManifestByContainerID indicates an expected call of SaveDeploymentManifestByContainerID
func (mr *MockRepositoryMockRecorder) SaveDeploymentManifestByContainerID(containerID, manifest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveDeploymentManifestByContainerID", reflect.TypeOf((*MockRepository)(nil).SaveDeploymentManifestByContainerID), containerID, manifest)
}

// DeleteDeploymentManifestByContainerID mocks base method
func (m *MockRepository) DeleteDeploymentManifestByContainerID(containerID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDeploymentManifestByContainerID", containerID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDeploymentManifestByContainerID indicates an expected call of DeleteDeploymentManifestByContainerID
func (mr *MockRepositoryMockRecorder) DeleteDeploymentManifestByContainerID(containerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDeploymentManifestByContainerID", reflect.TypeOf((*MockRepository)(nil).DeleteDeploymentManifestByContainerID), containerID)
}
