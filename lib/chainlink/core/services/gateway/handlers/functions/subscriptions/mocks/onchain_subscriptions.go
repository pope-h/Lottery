// Code generated by mockery v2.42.2. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	mock "github.com/stretchr/testify/mock"
)

// OnchainSubscriptions is an autogenerated mock type for the OnchainSubscriptions type
type OnchainSubscriptions struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *OnchainSubscriptions) Close() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Close")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMaxUserBalance provides a mock function with given fields: _a0
func (_m *OnchainSubscriptions) GetMaxUserBalance(_a0 common.Address) (*big.Int, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetMaxUserBalance")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(common.Address) (*big.Int, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(common.Address) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(common.Address) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Start provides a mock function with given fields: _a0
func (_m *OnchainSubscriptions) Start(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Start")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewOnchainSubscriptions creates a new instance of OnchainSubscriptions. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOnchainSubscriptions(t interface {
	mock.TestingT
	Cleanup(func())
}) *OnchainSubscriptions {
	mock := &OnchainSubscriptions{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
