// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/clients/s3/interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	sts "github.com/aws/aws-sdk-go-v2/service/sts"
	gomock "github.com/golang/mock/gomock"
)

// Mocks3ManagerAPI is a mock of s3ManagerAPI interface.
type Mocks3ManagerAPI struct {
	ctrl     *gomock.Controller
	recorder *Mocks3ManagerAPIMockRecorder
}

// Mocks3ManagerAPIMockRecorder is the mock recorder for Mocks3ManagerAPI.
type Mocks3ManagerAPIMockRecorder struct {
	mock *Mocks3ManagerAPI
}

// NewMocks3ManagerAPI creates a new mock instance.
func NewMocks3ManagerAPI(ctrl *gomock.Controller) *Mocks3ManagerAPI {
	mock := &Mocks3ManagerAPI{ctrl: ctrl}
	mock.recorder = &Mocks3ManagerAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3ManagerAPI) EXPECT() *Mocks3ManagerAPIMockRecorder {
	return m.recorder
}

// Upload mocks base method.
func (m *Mocks3ManagerAPI) Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, input}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Upload", varargs...)
	ret0, _ := ret[0].(*manager.UploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upload indicates an expected call of Upload.
func (mr *Mocks3ManagerAPIMockRecorder) Upload(ctx, input interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, input}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*Mocks3ManagerAPI)(nil).Upload), varargs...)
}

// Mocks3API is a mock of s3API interface.
type Mocks3API struct {
	ctrl     *gomock.Controller
	recorder *Mocks3APIMockRecorder
}

// Mocks3APIMockRecorder is the mock recorder for Mocks3API.
type Mocks3APIMockRecorder struct {
	mock *Mocks3API
}

// NewMocks3API creates a new mock instance.
func NewMocks3API(ctrl *gomock.Controller) *Mocks3API {
	mock := &Mocks3API{ctrl: ctrl}
	mock.recorder = &Mocks3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mocks3API) EXPECT() *Mocks3APIMockRecorder {
	return m.recorder
}

// MockstsAPI is a mock of stsAPI interface.
type MockstsAPI struct {
	ctrl     *gomock.Controller
	recorder *MockstsAPIMockRecorder
}

// MockstsAPIMockRecorder is the mock recorder for MockstsAPI.
type MockstsAPIMockRecorder struct {
	mock *MockstsAPI
}

// NewMockstsAPI creates a new mock instance.
func NewMockstsAPI(ctrl *gomock.Controller) *MockstsAPI {
	mock := &MockstsAPI{ctrl: ctrl}
	mock.recorder = &MockstsAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockstsAPI) EXPECT() *MockstsAPIMockRecorder {
	return m.recorder
}

// GetCallerIdentity mocks base method.
func (m *MockstsAPI) GetCallerIdentity(ctx context.Context, params *sts.GetCallerIdentityInput, optFns ...func(*sts.Options)) (*sts.GetCallerIdentityOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, params}
	for _, a := range optFns {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCallerIdentity", varargs...)
	ret0, _ := ret[0].(*sts.GetCallerIdentityOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCallerIdentity indicates an expected call of GetCallerIdentity.
func (mr *MockstsAPIMockRecorder) GetCallerIdentity(ctx, params interface{}, optFns ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, params}, optFns...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCallerIdentity", reflect.TypeOf((*MockstsAPI)(nil).GetCallerIdentity), varargs...)
}