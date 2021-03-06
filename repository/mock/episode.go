// Code generated by MockGen. DO NOT EDIT.
// Source: episode.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "github.com/bickyeric/arumba/model"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	reflect "reflect"
)

// MockIEpisode is a mock of IEpisode interface
type MockIEpisode struct {
	ctrl     *gomock.Controller
	recorder *MockIEpisodeMockRecorder
}

// MockIEpisodeMockRecorder is the mock recorder for MockIEpisode
type MockIEpisodeMockRecorder struct {
	mock *MockIEpisode
}

// NewMockIEpisode creates a new mock instance
func NewMockIEpisode(ctrl *gomock.Controller) *MockIEpisode {
	mock := &MockIEpisode{ctrl: ctrl}
	mock.recorder = &MockIEpisodeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEpisode) EXPECT() *MockIEpisodeMockRecorder {
	return m.recorder
}

// FindByNo mocks base method
func (m *MockIEpisode) FindByNo(comicID primitive.ObjectID, no int) (*model.Episode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByNo", comicID, no)
	ret0, _ := ret[0].(*model.Episode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByNo indicates an expected call of FindByNo
func (mr *MockIEpisodeMockRecorder) FindByNo(comicID, no interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByNo", reflect.TypeOf((*MockIEpisode)(nil).FindByNo), comicID, no)
}

// FindAll mocks base method
func (m *MockIEpisode) FindAll(ctx context.Context, comicID primitive.ObjectID, first, offset int) ([]model.Episode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", ctx, comicID, first, offset)
	ret0, _ := ret[0].([]model.Episode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAll indicates an expected call of FindAll
func (mr *MockIEpisodeMockRecorder) FindAll(ctx, comicID, first, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockIEpisode)(nil).FindAll), ctx, comicID, first, offset)
}

// Insert mocks base method
func (m *MockIEpisode) Insert(arg0 *model.Episode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIEpisodeMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIEpisode)(nil).Insert), arg0)
}

// CreateIndex mocks base method
func (m *MockIEpisode) CreateIndex(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndex indicates an expected call of CreateIndex
func (mr *MockIEpisodeMockRecorder) CreateIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockIEpisode)(nil).CreateIndex), arg0)
}
