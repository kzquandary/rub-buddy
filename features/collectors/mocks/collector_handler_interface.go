// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// CollectorHandlerInterface is an autogenerated mock type for the CollectorHandlerInterface type
type CollectorHandlerInterface struct {
	mock.Mock
}

// GetCollector provides a mock function with given fields:
func (_m *CollectorHandlerInterface) GetCollector() echo.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetCollector")
	}

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// Login provides a mock function with given fields:
func (_m *CollectorHandlerInterface) Login() echo.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// Register provides a mock function with given fields:
func (_m *CollectorHandlerInterface) Register() echo.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// UpdateCollector provides a mock function with given fields:
func (_m *CollectorHandlerInterface) UpdateCollector() echo.HandlerFunc {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for UpdateCollector")
	}

	var r0 echo.HandlerFunc
	if rf, ok := ret.Get(0).(func() echo.HandlerFunc); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(echo.HandlerFunc)
		}
	}

	return r0
}

// NewCollectorHandlerInterface creates a new instance of CollectorHandlerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCollectorHandlerInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *CollectorHandlerInterface {
	mock := &CollectorHandlerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}