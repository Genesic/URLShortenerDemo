// Code generated by MockGen. DO NOT EDIT.
// Source: module.go

// Package shortenController is a generated GoMock package.
package shortenController

import (
	errors "URLShortenerDemo/pkg/errors"
	database "URLShortenerDemo/service/database"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	gorm "gorm.io/gorm"
)

// MockIUrl is a mock of IUrl interface.
type MockIUrl struct {
	ctrl     *gomock.Controller
	recorder *MockIUrlMockRecorder
}

// MockIUrlMockRecorder is the mock recorder for MockIUrl.
type MockIUrlMockRecorder struct {
	mock *MockIUrl
}

// NewMockIUrl creates a new mock instance.
func NewMockIUrl(ctrl *gomock.Controller) *MockIUrl {
	mock := &MockIUrl{ctrl: ctrl}
	mock.recorder = &MockIUrlMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIUrl) EXPECT() *MockIUrlMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockIUrl) Create(tx *gorm.DB, url string, expiredAt time.Time) (*database.Url, *errors.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", tx, url, expiredAt)
	ret0, _ := ret[0].(*database.Url)
	ret1, _ := ret[1].(*errors.ServiceError)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockIUrlMockRecorder) Create(tx, url, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockIUrl)(nil).Create), tx, url, expiredAt)
}

// MockIHash is a mock of IHash interface.
type MockIHash struct {
	ctrl     *gomock.Controller
	recorder *MockIHashMockRecorder
}

// MockIHashMockRecorder is the mock recorder for MockIHash.
type MockIHashMockRecorder struct {
	mock *MockIHash
}

// NewMockIHash creates a new mock instance.
func NewMockIHash(ctrl *gomock.Controller) *MockIHash {
	mock := &MockIHash{ctrl: ctrl}
	mock.recorder = &MockIHashMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIHash) EXPECT() *MockIHashMockRecorder {
	return m.recorder
}

// IDtoShortenID mocks base method.
func (m *MockIHash) IDtoUrlID(id uint) (string, *errors.ServiceError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IDtoUrlID", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(*errors.ServiceError)
	return ret0, ret1
}

// IDtoShortenID indicates an expected call of IDtoShortenID.
func (mr *MockIHashMockRecorder) IDtoShortenID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IDtoUrlID", reflect.TypeOf((*MockIHash)(nil).IDtoUrlID), id)
}
