// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/platform/interface.go

// Package platform is a generated GoMock package.
package platform

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
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

// ExecCMDInContainer mocks base method.
func (m *MockClient) ExecCMDInContainer(ctx context.Context, containerName, podName string, cmd []string, stdout, stderr io.Writer, stdin io.Reader, tty bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecCMDInContainer", ctx, containerName, podName, cmd, stdout, stderr, stdin, tty)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecCMDInContainer indicates an expected call of ExecCMDInContainer.
func (mr *MockClientMockRecorder) ExecCMDInContainer(ctx, containerName, podName, cmd, stdout, stderr, stdin, tty interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecCMDInContainer", reflect.TypeOf((*MockClient)(nil).ExecCMDInContainer), ctx, containerName, podName, cmd, stdout, stderr, stdin, tty)
}

// GetAllPodsInNamespaceMatchingSelector mocks base method.
func (m *MockClient) GetAllPodsInNamespaceMatchingSelector(selector, ns string) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPodsInNamespaceMatchingSelector", selector, ns)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPodsInNamespaceMatchingSelector indicates an expected call of GetAllPodsInNamespaceMatchingSelector.
func (mr *MockClientMockRecorder) GetAllPodsInNamespaceMatchingSelector(selector, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPodsInNamespaceMatchingSelector", reflect.TypeOf((*MockClient)(nil).GetAllPodsInNamespaceMatchingSelector), selector, ns)
}

// GetAllResourcesFromSelector mocks base method.
func (m *MockClient) GetAllResourcesFromSelector(selector, ns string) ([]unstructured.Unstructured, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllResourcesFromSelector", selector, ns)
	ret0, _ := ret[0].([]unstructured.Unstructured)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllResourcesFromSelector indicates an expected call of GetAllResourcesFromSelector.
func (mr *MockClientMockRecorder) GetAllResourcesFromSelector(selector, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResourcesFromSelector", reflect.TypeOf((*MockClient)(nil).GetAllResourcesFromSelector), selector, ns)
}

// GetPodLogs mocks base method.
func (m *MockClient) GetPodLogs(podName, containerName string, followLog bool) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodLogs", podName, containerName, followLog)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodLogs indicates an expected call of GetPodLogs.
func (mr *MockClientMockRecorder) GetPodLogs(podName, containerName, followLog interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodLogs", reflect.TypeOf((*MockClient)(nil).GetPodLogs), podName, containerName, followLog)
}

// GetPodUsingComponentName mocks base method.
func (m *MockClient) GetPodUsingComponentName(componentName string) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodUsingComponentName", componentName)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodUsingComponentName indicates an expected call of GetPodUsingComponentName.
func (mr *MockClientMockRecorder) GetPodUsingComponentName(componentName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodUsingComponentName", reflect.TypeOf((*MockClient)(nil).GetPodUsingComponentName), componentName)
}

// GetPodsMatchingSelector mocks base method.
func (m *MockClient) GetPodsMatchingSelector(selector string) (*v1.PodList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPodsMatchingSelector", selector)
	ret0, _ := ret[0].(*v1.PodList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPodsMatchingSelector indicates an expected call of GetPodsMatchingSelector.
func (mr *MockClientMockRecorder) GetPodsMatchingSelector(selector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPodsMatchingSelector", reflect.TypeOf((*MockClient)(nil).GetPodsMatchingSelector), selector)
}

// GetRunningPodFromSelector mocks base method.
func (m *MockClient) GetRunningPodFromSelector(selector string) (*v1.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRunningPodFromSelector", selector)
	ret0, _ := ret[0].(*v1.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRunningPodFromSelector indicates an expected call of GetRunningPodFromSelector.
func (mr *MockClientMockRecorder) GetRunningPodFromSelector(selector interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRunningPodFromSelector", reflect.TypeOf((*MockClient)(nil).GetRunningPodFromSelector), selector)
}
