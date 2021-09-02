// Code generated by MockGen. DO NOT EDIT.
// Source: db/querier.go

// Package db is a generated GoMock package.
package db

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockQuerier is a mock of Querier interface.
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for MockQuerier.
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// NewMockQuerier creates a new mock instance.
func NewMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// CreatePost mocks base method.
func (m *MockQuerier) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePost", ctx, arg)
	ret0, _ := ret[0].(Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePost indicates an expected call of CreatePost.
func (mr *MockQuerierMockRecorder) CreatePost(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePost", reflect.TypeOf((*MockQuerier)(nil).CreatePost), ctx, arg)
}

// CreateUser mocks base method.
func (m *MockQuerier) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", ctx, arg)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockQuerierMockRecorder) CreateUser(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockQuerier)(nil).CreateUser), ctx, arg)
}

// DeletePostByIDs mocks base method.
func (m *MockQuerier) DeletePostByIDs(ctx context.Context, arg DeletePostByIDsParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePostByIDs", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePostByIDs indicates an expected call of DeletePostByIDs.
func (mr *MockQuerierMockRecorder) DeletePostByIDs(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePostByIDs", reflect.TypeOf((*MockQuerier)(nil).DeletePostByIDs), ctx, arg)
}

// FindPostByIDs mocks base method.
func (m *MockQuerier) FindPostByIDs(ctx context.Context, arg FindPostByIDsParams) (Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPostByIDs", ctx, arg)
	ret0, _ := ret[0].(Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPostByIDs indicates an expected call of FindPostByIDs.
func (mr *MockQuerierMockRecorder) FindPostByIDs(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPostByIDs", reflect.TypeOf((*MockQuerier)(nil).FindPostByIDs), ctx, arg)
}

// FindPostsByAuthor mocks base method.
func (m *MockQuerier) FindPostsByAuthor(ctx context.Context, authorID int64) ([]Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPostsByAuthor", ctx, authorID)
	ret0, _ := ret[0].([]Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPostsByAuthor indicates an expected call of FindPostsByAuthor.
func (mr *MockQuerierMockRecorder) FindPostsByAuthor(ctx, authorID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPostsByAuthor", reflect.TypeOf((*MockQuerier)(nil).FindPostsByAuthor), ctx, authorID)
}

// FindUserByEmail mocks base method.
func (m *MockQuerier) FindUserByEmail(ctx context.Context, lower string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", ctx, lower)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockQuerierMockRecorder) FindUserByEmail(ctx, lower interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockQuerier)(nil).FindUserByEmail), ctx, lower)
}

// FindUserByVerificationCode mocks base method.
func (m *MockQuerier) FindUserByVerificationCode(ctx context.Context, verification string) (User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByVerificationCode", ctx, verification)
	ret0, _ := ret[0].(User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByVerificationCode indicates an expected call of FindUserByVerificationCode.
func (mr *MockQuerierMockRecorder) FindUserByVerificationCode(ctx, verification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByVerificationCode", reflect.TypeOf((*MockQuerier)(nil).FindUserByVerificationCode), ctx, verification)
}

// UpdatePost mocks base method.
func (m *MockQuerier) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePost", ctx, arg)
	ret0, _ := ret[0].(Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePost indicates an expected call of UpdatePost.
func (mr *MockQuerierMockRecorder) UpdatePost(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePost", reflect.TypeOf((*MockQuerier)(nil).UpdatePost), ctx, arg)
}

// UpdateUserStatus mocks base method.
func (m *MockQuerier) UpdateUserStatus(ctx context.Context, arg UpdateUserStatusParams) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUserStatus", ctx, arg)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserStatus indicates an expected call of UpdateUserStatus.
func (mr *MockQuerierMockRecorder) UpdateUserStatus(ctx, arg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserStatus", reflect.TypeOf((*MockQuerier)(nil).UpdateUserStatus), ctx, arg)
}
