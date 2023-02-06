// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"

	onfleet "github.com/onfleet/cq-source-onfleet/internal/go-onfleet/onfleet"
)

// ClientInterface is an autogenerated mock type for the ClientInterface type
type ClientInterface struct {
	mock.Mock
}

// Admins provides a mock function with given fields:
func (_m *ClientInterface) Admins() onfleet.AdminsServiceInterface {
	ret := _m.Called()

	var r0 onfleet.AdminsServiceInterface
	if rf, ok := ret.Get(0).(func() onfleet.AdminsServiceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(onfleet.AdminsServiceInterface)
		}
	}

	return r0
}

// Do provides a mock function with given fields: ctx, req, v
func (_m *ClientInterface) Do(ctx context.Context, req *http.Request, v interface{}) error {
	ret := _m.Called(ctx, req, v)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request, interface{}) error); ok {
		r0 = rf(ctx, req, v)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Hubs provides a mock function with given fields:
func (_m *ClientInterface) Hubs() onfleet.HubsServiceInterface {
	ret := _m.Called()

	var r0 onfleet.HubsServiceInterface
	if rf, ok := ret.Get(0).(func() onfleet.HubsServiceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(onfleet.HubsServiceInterface)
		}
	}

	return r0
}

// Init provides a mock function with given fields: apiKey
func (_m *ClientInterface) Init(apiKey string) {
	_m.Called(apiKey)
}

// NewRequest provides a mock function with given fields: method, path, body
func (_m *ClientInterface) NewRequest(method string, path string, body interface{}) (*http.Request, error) {
	ret := _m.Called(method, path, body)

	var r0 *http.Request
	if rf, ok := ret.Get(0).(func(string, string, interface{}) *http.Request); ok {
		r0 = rf(method, path, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Request)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, interface{}) error); ok {
		r1 = rf(method, path, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Organization provides a mock function with given fields:
func (_m *ClientInterface) Organization() onfleet.OrganizationServiceInterface {
	ret := _m.Called()

	var r0 onfleet.OrganizationServiceInterface
	if rf, ok := ret.Get(0).(func() onfleet.OrganizationServiceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(onfleet.OrganizationServiceInterface)
		}
	}

	return r0
}

// Tasks provides a mock function with given fields:
func (_m *ClientInterface) Tasks() onfleet.TasksServiceInterface {
	ret := _m.Called()

	var r0 onfleet.TasksServiceInterface
	if rf, ok := ret.Get(0).(func() onfleet.TasksServiceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(onfleet.TasksServiceInterface)
		}
	}

	return r0
}

// Teams provides a mock function with given fields:
func (_m *ClientInterface) Teams() onfleet.TeamsServiceInterface {
	ret := _m.Called()

	var r0 onfleet.TeamsServiceInterface
	if rf, ok := ret.Get(0).(func() onfleet.TeamsServiceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(onfleet.TeamsServiceInterface)
		}
	}

	return r0
}

// WithBaseURL provides a mock function with given fields: baseURL
func (_m *ClientInterface) WithBaseURL(baseURL string) {
	_m.Called(baseURL)
}

// Workers provides a mock function with given fields:
func (_m *ClientInterface) Workers() onfleet.WorkersServiceInterface {
	ret := _m.Called()

	var r0 onfleet.WorkersServiceInterface
	if rf, ok := ret.Get(0).(func() onfleet.WorkersServiceInterface); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(onfleet.WorkersServiceInterface)
		}
	}

	return r0
}

type mockConstructorTestingTNewClientInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewClientInterface creates a new instance of ClientInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewClientInterface(t mockConstructorTestingTNewClientInterface) *ClientInterface {
	mock := &ClientInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
