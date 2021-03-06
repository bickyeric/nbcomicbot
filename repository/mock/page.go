// Code generated by MockGen. DO NOT EDIT.
// Source: page.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	model "github.com/bickyeric/arumba/model"
	gomock "github.com/golang/mock/gomock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
	reflect "reflect"
)

// MockIPage is a mock of IPage interface
type MockIPage struct {
	ctrl     *gomock.Controller
	recorder *MockIPageMockRecorder
}

// MockIPageMockRecorder is the mock recorder for MockIPage
type MockIPageMockRecorder struct {
	mock *MockIPage
}

// NewMockIPage creates a new mock instance
func NewMockIPage(ctrl *gomock.Controller) *MockIPage {
	mock := &MockIPage{ctrl: ctrl}
	mock.recorder = &MockIPageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIPage) EXPECT() *MockIPageMockRecorder {
	return m.recorder
}

// FindByEpisodeSource mocks base method
func (m *MockIPage) FindByEpisodeSource(episodeID, sourceID primitive.ObjectID) (model.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEpisodeSource", episodeID, sourceID)
	ret0, _ := ret[0].(model.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEpisodeSource indicates an expected call of FindByEpisodeSource
func (mr *MockIPageMockRecorder) FindByEpisodeSource(episodeID, sourceID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEpisodeSource", reflect.TypeOf((*MockIPage)(nil).FindByEpisodeSource), episodeID, sourceID)
}

// FindByEpisode mocks base method
func (m *MockIPage) FindByEpisode(ctx context.Context, episodeID primitive.ObjectID, first, offset int) ([]model.Page, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByEpisode", ctx, episodeID, first, offset)
	ret0, _ := ret[0].([]model.Page)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByEpisode indicates an expected call of FindByEpisode
func (mr *MockIPageMockRecorder) FindByEpisode(ctx, episodeID, first, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByEpisode", reflect.TypeOf((*MockIPage)(nil).FindByEpisode), ctx, episodeID, first, offset)
}

// Insert mocks base method
func (m *MockIPage) Insert(arg0 *model.Page) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockIPageMockRecorder) Insert(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockIPage)(nil).Insert), arg0)
}

// CreateIndex mocks base method
func (m *MockIPage) CreateIndex(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateIndex", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateIndex indicates an expected call of CreateIndex
func (mr *MockIPageMockRecorder) CreateIndex(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateIndex", reflect.TypeOf((*MockIPage)(nil).CreateIndex), arg0)
}
