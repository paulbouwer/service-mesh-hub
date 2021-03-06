// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mock_vm_validation is a generated GoMock package.
package mock_vm_validation

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.zephyr.solo.io/v1alpha1"
	v1alpha10 "github.com/solo-io/service-mesh-hub/pkg/api/networking.zephyr.solo.io/v1alpha1"
)

// MockVirtualMeshFinder is a mock of VirtualMeshFinder interface.
type MockVirtualMeshFinder struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshFinderMockRecorder
}

// MockVirtualMeshFinderMockRecorder is the mock recorder for MockVirtualMeshFinder.
type MockVirtualMeshFinderMockRecorder struct {
	mock *MockVirtualMeshFinder
}

// NewMockVirtualMeshFinder creates a new mock instance.
func NewMockVirtualMeshFinder(ctrl *gomock.Controller) *MockVirtualMeshFinder {
	mock := &MockVirtualMeshFinder{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshFinderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshFinder) EXPECT() *MockVirtualMeshFinderMockRecorder {
	return m.recorder
}

// GetMeshesForVirtualMesh mocks base method.
func (m *MockVirtualMeshFinder) GetMeshesForVirtualMesh(ctx context.Context, vm *v1alpha10.VirtualMesh) ([]*v1alpha1.Mesh, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMeshesForVirtualMesh", ctx, vm)
	ret0, _ := ret[0].([]*v1alpha1.Mesh)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMeshesForVirtualMesh indicates an expected call of GetMeshesForVirtualMesh.
func (mr *MockVirtualMeshFinderMockRecorder) GetMeshesForVirtualMesh(ctx, vm interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMeshesForVirtualMesh", reflect.TypeOf((*MockVirtualMeshFinder)(nil).GetMeshesForVirtualMesh), ctx, vm)
}
