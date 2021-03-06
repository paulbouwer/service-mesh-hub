// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1 (interfaces: VirtualMeshCertificateSigningRequestClient)

// Package mock_zephyr_security_clients is a generated GoMock package.
package mock_zephyr_security_clients

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/security.zephyr.solo.io/v1alpha1"
	types "k8s.io/apimachinery/pkg/types"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

// MockVirtualMeshCertificateSigningRequestClient is a mock of VirtualMeshCertificateSigningRequestClient interface.
type MockVirtualMeshCertificateSigningRequestClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMeshCertificateSigningRequestClientMockRecorder
}

// MockVirtualMeshCertificateSigningRequestClientMockRecorder is the mock recorder for MockVirtualMeshCertificateSigningRequestClient.
type MockVirtualMeshCertificateSigningRequestClientMockRecorder struct {
	mock *MockVirtualMeshCertificateSigningRequestClient
}

// NewMockVirtualMeshCertificateSigningRequestClient creates a new mock instance.
func NewMockVirtualMeshCertificateSigningRequestClient(ctrl *gomock.Controller) *MockVirtualMeshCertificateSigningRequestClient {
	mock := &MockVirtualMeshCertificateSigningRequestClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMeshCertificateSigningRequestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMeshCertificateSigningRequestClient) EXPECT() *MockVirtualMeshCertificateSigningRequestClientMockRecorder {
	return m.recorder
}

// CreateVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) CreateVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 *v1alpha1.VirtualMeshCertificateSigningRequest, arg2 ...client.CreateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateVirtualMeshCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateVirtualMeshCertificateSigningRequest indicates an expected call of CreateVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) CreateVirtualMeshCertificateSigningRequest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).CreateVirtualMeshCertificateSigningRequest), varargs...)
}

// DeleteAllOfVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) DeleteAllOfVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 ...client.DeleteAllOfOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteAllOfVirtualMeshCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllOfVirtualMeshCertificateSigningRequest indicates an expected call of DeleteAllOfVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) DeleteAllOfVirtualMeshCertificateSigningRequest(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllOfVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).DeleteAllOfVirtualMeshCertificateSigningRequest), varargs...)
}

// DeleteVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) DeleteVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 types.NamespacedName, arg2 ...client.DeleteOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteVirtualMeshCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVirtualMeshCertificateSigningRequest indicates an expected call of DeleteVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) DeleteVirtualMeshCertificateSigningRequest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).DeleteVirtualMeshCertificateSigningRequest), varargs...)
}

// GetVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) GetVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 types.NamespacedName) (*v1alpha1.VirtualMeshCertificateSigningRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVirtualMeshCertificateSigningRequest", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.VirtualMeshCertificateSigningRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVirtualMeshCertificateSigningRequest indicates an expected call of GetVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) GetVirtualMeshCertificateSigningRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).GetVirtualMeshCertificateSigningRequest), arg0, arg1)
}

// ListVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) ListVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 ...client.ListOption) (*v1alpha1.VirtualMeshCertificateSigningRequestList, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListVirtualMeshCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(*v1alpha1.VirtualMeshCertificateSigningRequestList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVirtualMeshCertificateSigningRequest indicates an expected call of ListVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) ListVirtualMeshCertificateSigningRequest(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).ListVirtualMeshCertificateSigningRequest), varargs...)
}

// PatchVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) PatchVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 *v1alpha1.VirtualMeshCertificateSigningRequest, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchVirtualMeshCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchVirtualMeshCertificateSigningRequest indicates an expected call of PatchVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) PatchVirtualMeshCertificateSigningRequest(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).PatchVirtualMeshCertificateSigningRequest), varargs...)
}

// PatchVirtualMeshCertificateSigningRequestStatus mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) PatchVirtualMeshCertificateSigningRequestStatus(arg0 context.Context, arg1 *v1alpha1.VirtualMeshCertificateSigningRequest, arg2 client.Patch, arg3 ...client.PatchOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PatchVirtualMeshCertificateSigningRequestStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchVirtualMeshCertificateSigningRequestStatus indicates an expected call of PatchVirtualMeshCertificateSigningRequestStatus.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) PatchVirtualMeshCertificateSigningRequestStatus(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchVirtualMeshCertificateSigningRequestStatus", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).PatchVirtualMeshCertificateSigningRequestStatus), varargs...)
}

// UpdateVirtualMeshCertificateSigningRequest mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) UpdateVirtualMeshCertificateSigningRequest(arg0 context.Context, arg1 *v1alpha1.VirtualMeshCertificateSigningRequest, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateVirtualMeshCertificateSigningRequest", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualMeshCertificateSigningRequest indicates an expected call of UpdateVirtualMeshCertificateSigningRequest.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) UpdateVirtualMeshCertificateSigningRequest(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualMeshCertificateSigningRequest", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).UpdateVirtualMeshCertificateSigningRequest), varargs...)
}

// UpdateVirtualMeshCertificateSigningRequestStatus mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) UpdateVirtualMeshCertificateSigningRequestStatus(arg0 context.Context, arg1 *v1alpha1.VirtualMeshCertificateSigningRequest, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateVirtualMeshCertificateSigningRequestStatus", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVirtualMeshCertificateSigningRequestStatus indicates an expected call of UpdateVirtualMeshCertificateSigningRequestStatus.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) UpdateVirtualMeshCertificateSigningRequestStatus(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVirtualMeshCertificateSigningRequestStatus", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).UpdateVirtualMeshCertificateSigningRequestStatus), varargs...)
}

// UpsertVirtualMeshCertificateSigningRequestSpec mocks base method.
func (m *MockVirtualMeshCertificateSigningRequestClient) UpsertVirtualMeshCertificateSigningRequestSpec(arg0 context.Context, arg1 *v1alpha1.VirtualMeshCertificateSigningRequest, arg2 ...client.UpdateOption) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpsertVirtualMeshCertificateSigningRequestSpec", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertVirtualMeshCertificateSigningRequestSpec indicates an expected call of UpsertVirtualMeshCertificateSigningRequestSpec.
func (mr *MockVirtualMeshCertificateSigningRequestClientMockRecorder) UpsertVirtualMeshCertificateSigningRequestSpec(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertVirtualMeshCertificateSigningRequestSpec", reflect.TypeOf((*MockVirtualMeshCertificateSigningRequestClient)(nil).UpsertVirtualMeshCertificateSigningRequestSpec), varargs...)
}
