// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/database/interface.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddOTPClient mocks base method.
func (m *MockStorage) AddOTPClient(ctx context.Context, logger *logrus.Logger, extID, phone string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOTPClient", ctx, logger, extID, phone)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOTPClient indicates an expected call of AddOTPClient.
func (mr *MockStorageMockRecorder) AddOTPClient(ctx, logger, extID, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOTPClient", reflect.TypeOf((*MockStorage)(nil).AddOTPClient), ctx, logger, extID, phone)
}

// FindIDByPhone mocks base method.
func (m *MockStorage) FindIDByPhone(ctx context.Context, logger *logrus.Logger, phone string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindIDByPhone", ctx, logger, phone)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindIDByPhone indicates an expected call of FindIDByPhone.
func (mr *MockStorageMockRecorder) FindIDByPhone(ctx, logger, phone interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindIDByPhone", reflect.TypeOf((*MockStorage)(nil).FindIDByPhone), ctx, logger, phone)
}

// GetExternalID mocks base method.
func (m *MockStorage) GetExternalID(ctx context.Context, logger *logrus.Logger, otpKeys string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExternalID", ctx, logger, otpKeys)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExternalID indicates an expected call of GetExternalID.
func (mr *MockStorageMockRecorder) GetExternalID(ctx, logger, otpKeys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExternalID", reflect.TypeOf((*MockStorage)(nil).GetExternalID), ctx, logger, otpKeys)
}

// UpdateOTP mocks base method.
func (m *MockStorage) UpdateOTP(ctx context.Context, logger *logrus.Logger, otpKey, extID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOTP", ctx, logger, otpKey, extID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOTP indicates an expected call of UpdateOTP.
func (mr *MockStorageMockRecorder) UpdateOTP(ctx, logger, otpKey, extID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOTP", reflect.TypeOf((*MockStorage)(nil).UpdateOTP), ctx, logger, otpKey, extID)
}
