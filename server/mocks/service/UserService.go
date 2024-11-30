// Code generated by mockery v2.49.1. DO NOT EDIT.

package mocks

import (
	context "context"
	request "simple-dashboard-server/api/request"

	mock "github.com/stretchr/testify/mock"

	response "simple-dashboard-server/api/response"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// EditUserLogin provides a mock function with given fields: ctx, req
func (_m *UserService) EditUserLogin(ctx context.Context, req request.EditUserRequest) error {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for EditUserLogin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, request.EditUserRequest) error); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetProfileUserLogin provides a mock function with given fields: ctx, userId
func (_m *UserService) GetProfileUserLogin(ctx context.Context, userId string) (response.UserResponse, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileUserLogin")
	}

	var r0 response.UserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (response.UserResponse, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) response.UserResponse); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Get(0).(response.UserResponse)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
