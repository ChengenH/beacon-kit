// Code generated by mockery v2.49.0. DO NOT EDIT.

package mocks

import (
	math "github.com/berachain/beacon-kit/primitives/math"
	mock "github.com/stretchr/testify/mock"

	transition "github.com/berachain/beacon-kit/primitives/transition"
)

// StateProcessor is an autogenerated mock type for the StateProcessor type
type StateProcessor[BeaconStateT any] struct {
	mock.Mock
}

type StateProcessor_Expecter[BeaconStateT any] struct {
	mock *mock.Mock
}

func (_m *StateProcessor[BeaconStateT]) EXPECT() *StateProcessor_Expecter[BeaconStateT] {
	return &StateProcessor_Expecter[BeaconStateT]{mock: &_m.Mock}
}

// ProcessSlots provides a mock function with given fields: _a0, _a1
func (_m *StateProcessor[BeaconStateT]) ProcessSlots(_a0 BeaconStateT, _a1 math.U64) (transition.ValidatorUpdates, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for ProcessSlots")
	}

	var r0 transition.ValidatorUpdates
	var r1 error
	if rf, ok := ret.Get(0).(func(BeaconStateT, math.U64) (transition.ValidatorUpdates, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(BeaconStateT, math.U64) transition.ValidatorUpdates); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(transition.ValidatorUpdates)
		}
	}

	if rf, ok := ret.Get(1).(func(BeaconStateT, math.U64) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateProcessor_ProcessSlots_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessSlots'
type StateProcessor_ProcessSlots_Call[BeaconStateT any] struct {
	*mock.Call
}

// ProcessSlots is a helper method to define mock.On call
//   - _a0 BeaconStateT
//   - _a1 math.U64
func (_e *StateProcessor_Expecter[BeaconStateT]) ProcessSlots(_a0 interface{}, _a1 interface{}) *StateProcessor_ProcessSlots_Call[BeaconStateT] {
	return &StateProcessor_ProcessSlots_Call[BeaconStateT]{Call: _e.mock.On("ProcessSlots", _a0, _a1)}
}

func (_c *StateProcessor_ProcessSlots_Call[BeaconStateT]) Run(run func(_a0 BeaconStateT, _a1 math.U64)) *StateProcessor_ProcessSlots_Call[BeaconStateT] {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(BeaconStateT), args[1].(math.U64))
	})
	return _c
}

func (_c *StateProcessor_ProcessSlots_Call[BeaconStateT]) Return(_a0 transition.ValidatorUpdates, _a1 error) *StateProcessor_ProcessSlots_Call[BeaconStateT] {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *StateProcessor_ProcessSlots_Call[BeaconStateT]) RunAndReturn(run func(BeaconStateT, math.U64) (transition.ValidatorUpdates, error)) *StateProcessor_ProcessSlots_Call[BeaconStateT] {
	_c.Call.Return(run)
	return _c
}

// NewStateProcessor creates a new instance of StateProcessor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateProcessor[BeaconStateT any](t interface {
	mock.TestingT
	Cleanup(func())
}) *StateProcessor[BeaconStateT] {
	mock := &StateProcessor[BeaconStateT]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}