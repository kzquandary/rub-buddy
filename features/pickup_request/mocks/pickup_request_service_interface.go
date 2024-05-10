// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	pickuprequest "rub_buddy/features/pickup_request"

	mock "github.com/stretchr/testify/mock"
)

// PickupRequestServiceInterface is an autogenerated mock type for the PickupRequestServiceInterface type
type PickupRequestServiceInterface struct {
	mock.Mock
}

// CreatePickupRequest provides a mock function with given fields: newData
func (_m *PickupRequestServiceInterface) CreatePickupRequest(newData pickuprequest.PickupRequest) (*pickuprequest.PickupRequest, error) {
	ret := _m.Called(newData)

	if len(ret) == 0 {
		panic("no return value specified for CreatePickupRequest")
	}

	var r0 *pickuprequest.PickupRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(pickuprequest.PickupRequest) (*pickuprequest.PickupRequest, error)); ok {
		return rf(newData)
	}
	if rf, ok := ret.Get(0).(func(pickuprequest.PickupRequest) *pickuprequest.PickupRequest); ok {
		r0 = rf(newData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pickuprequest.PickupRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(pickuprequest.PickupRequest) error); ok {
		r1 = rf(newData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePickupRequestByID provides a mock function with given fields: id, userID
func (_m *PickupRequestServiceInterface) DeletePickupRequestByID(id int, userID uint) error {
	ret := _m.Called(id, userID)

	if len(ret) == 0 {
		panic("no return value specified for DeletePickupRequestByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int, uint) error); ok {
		r0 = rf(id, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllPickupRequest provides a mock function with given fields:
func (_m *PickupRequestServiceInterface) GetAllPickupRequest() ([]pickuprequest.PickupRequestInfo, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllPickupRequest")
	}

	var r0 []pickuprequest.PickupRequestInfo
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]pickuprequest.PickupRequestInfo, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []pickuprequest.PickupRequestInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]pickuprequest.PickupRequestInfo)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPickupRequestByID provides a mock function with given fields: id
func (_m *PickupRequestServiceInterface) GetPickupRequestByID(id int) (pickuprequest.PickupRequestInfo, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetPickupRequestByID")
	}

	var r0 pickuprequest.PickupRequestInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (pickuprequest.PickupRequestInfo, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) pickuprequest.PickupRequestInfo); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(pickuprequest.PickupRequestInfo)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPickupRequestServiceInterface creates a new instance of PickupRequestServiceInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPickupRequestServiceInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *PickupRequestServiceInterface {
	mock := &PickupRequestServiceInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
