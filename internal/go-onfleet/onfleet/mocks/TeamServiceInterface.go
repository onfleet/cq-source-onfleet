// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	context "context"

	onfleet "github.com/onfleet/cq-source-onfleet/internal/go-onfleet/onfleet"
	mock "github.com/stretchr/testify/mock"
)

// TeamServiceInterface is an autogenerated mock type for the TeamServiceInterface type
type TeamServiceInterface struct {
	mock.Mock
}

// List provides a mock function with given fields: ctx
func (_m *TeamServiceInterface) List(ctx context.Context) ([]onfleet.Team, error) {
	ret := _m.Called(ctx)

	var r0 []onfleet.Team
	if rf, ok := ret.Get(0).(func(context.Context) []onfleet.Team); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]onfleet.Team)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTeamServiceInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewTeamServiceInterface creates a new instance of TeamServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTeamServiceInterface(t mockConstructorTestingTNewTeamServiceInterface) *TeamServiceInterface {
	mock := &TeamServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
