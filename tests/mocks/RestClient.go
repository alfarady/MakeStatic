// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import (
	context "context"

	rest_client "github.com/alfarady/makestatic/internal/rest_client"
	mock "github.com/stretchr/testify/mock"
)

// RestClient is an autogenerated mock type for the RestClient type
type RestClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, req
func (_m *RestClient) Get(ctx context.Context, req rest_client.RestRequest) (rest_client.RestResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 rest_client.RestResponse
	if rf, ok := ret.Get(0).(func(context.Context, rest_client.RestRequest) rest_client.RestResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(rest_client.RestResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, rest_client.RestRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: ctx, req
func (_m *RestClient) Patch(ctx context.Context, req rest_client.RestRequest) (rest_client.RestResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 rest_client.RestResponse
	if rf, ok := ret.Get(0).(func(context.Context, rest_client.RestRequest) rest_client.RestResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(rest_client.RestResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, rest_client.RestRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: ctx, req
func (_m *RestClient) Post(ctx context.Context, req rest_client.RestRequest) (rest_client.RestResponse, error) {
	ret := _m.Called(ctx, req)

	var r0 rest_client.RestResponse
	if rf, ok := ret.Get(0).(func(context.Context, rest_client.RestRequest) rest_client.RestResponse); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(rest_client.RestResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, rest_client.RestRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRestClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewRestClient creates a new instance of RestClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRestClient(t mockConstructorTestingTNewRestClient) *RestClient {
	mock := &RestClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}