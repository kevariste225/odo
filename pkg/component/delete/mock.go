// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/component/delete/interface.go

// Package delete is a generated GoMock package.
package delete

import (
	reflect "reflect"

	parser "github.com/devfile/library/pkg/devfile/parser"
	gomock "github.com/golang/mock/gomock"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// ListClusterResourcesToDelete mocks base method
func (m *MockClient) ListClusterResourcesToDelete(componentName, namespace string) ([]unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterResourcesToDelete", componentName, namespace)
	ret0, _ := ret[0].([]unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterResourcesToDelete indicates an expected call of ListClusterResourcesToDelete
func (mr *MockClientMockRecorder) ListClusterResourcesToDelete(componentName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterResourcesToDelete", reflect.TypeOf((*MockClient)(nil).ListClusterResourcesToDelete), componentName, namespace)
}

// DeleteResources mocks base method.
func (m *MockClient) DeleteResources(arg0 []unstructured.Unstructured) []unstructured.Unstructured {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResources", arg0)
	ret0, _ := ret[0].([]unstructured.Unstructured)
	return ret0
}

// DeleteResources indicates an expected call of DeleteResources.
func (mr *MockClientMockRecorder) DeleteResources(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResources", reflect.TypeOf((*MockClient)(nil).DeleteResources), arg0)
}

// ExecutePreStopEvents mocks base method.
func (m *MockClient) ExecutePreStopEvents(devfileObj parser.DevfileObj, appName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecutePreStopEvents", devfileObj, appName)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecutePreStopEvents indicates an expected call of ExecutePreStopEvents.
func (mr *MockClientMockRecorder) ExecutePreStopEvents(devfileObj, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecutePreStopEvents", reflect.TypeOf((*MockClient)(nil).ExecutePreStopEvents), devfileObj, appName)
}

// ListResourcesToDeleteFromDevfile mocks base method.
func (m *MockClient) ListResourcesToDeleteFromDevfile(devfileObj parser.DevfileObj, appName string) (bool, []unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListResourcesToDeleteFromDevfile", devfileObj, appName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]unstructured.Unstructured)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListResourcesToDeleteFromDevfile indicates an expected call of ListResourcesToDeleteFromDevfile.
func (mr *MockClientMockRecorder) ListResourcesToDeleteFromDevfile(devfileObj, appName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListResourcesToDeleteFromDevfile", reflect.TypeOf((*MockClient)(nil).ListResourcesToDeleteFromDevfile), devfileObj, appName)
}
