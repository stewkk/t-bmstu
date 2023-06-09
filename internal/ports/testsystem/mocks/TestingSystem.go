// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	testsystem "github.com/stewkk/t-bmstu/internal/ports/testsystem"
	mock "github.com/stretchr/testify/mock"
)

// TestingSystem is an autogenerated mock type for the TestingSystem type
type TestingSystem struct {
	mock.Mock
}

type TestingSystem_Expecter struct {
	mock *mock.Mock
}

func (_m *TestingSystem) EXPECT() *TestingSystem_Expecter {
	return &TestingSystem_Expecter{mock: &_m.Mock}
}

// SubmitSolution provides a mock function with given fields: _a0
func (_m *TestingSystem) SubmitSolution(_a0 *testsystem.SubmissionCreateRequest) (*testsystem.SubmissionCreateResponse, error) {
	ret := _m.Called(_a0)

	var r0 *testsystem.SubmissionCreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(*testsystem.SubmissionCreateRequest) (*testsystem.SubmissionCreateResponse, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*testsystem.SubmissionCreateRequest) *testsystem.SubmissionCreateResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*testsystem.SubmissionCreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(*testsystem.SubmissionCreateRequest) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TestingSystem_SubmitSolution_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubmitSolution'
type TestingSystem_SubmitSolution_Call struct {
	*mock.Call
}

// SubmitSolution is a helper method to define mock.On call
//   - _a0 *testsystem.SubmissionCreateRequest
func (_e *TestingSystem_Expecter) SubmitSolution(_a0 interface{}) *TestingSystem_SubmitSolution_Call {
	return &TestingSystem_SubmitSolution_Call{Call: _e.mock.On("SubmitSolution", _a0)}
}

func (_c *TestingSystem_SubmitSolution_Call) Run(run func(_a0 *testsystem.SubmissionCreateRequest)) *TestingSystem_SubmitSolution_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*testsystem.SubmissionCreateRequest))
	})
	return _c
}

func (_c *TestingSystem_SubmitSolution_Call) Return(_a0 *testsystem.SubmissionCreateResponse, _a1 error) *TestingSystem_SubmitSolution_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *TestingSystem_SubmitSolution_Call) RunAndReturn(run func(*testsystem.SubmissionCreateRequest) (*testsystem.SubmissionCreateResponse, error)) *TestingSystem_SubmitSolution_Call {
	_c.Call.Return(run)
	return _c
}

// NewTestingSystem creates a new instance of TestingSystem. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTestingSystem(t interface {
	mock.TestingT
	Cleanup(func())
}) *TestingSystem {
	mock := &TestingSystem{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
