// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import stream "github.com/applike/gosoline/pkg/stream"

// Input is an autogenerated mock type for the Input type
type Input struct {
	mock.Mock
}

// Data provides a mock function with given fields:
func (_m *Input) Data() chan *stream.Message {
	ret := _m.Called()

	var r0 chan *stream.Message
	if rf, ok := ret.Get(0).(func() chan *stream.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *stream.Message)
		}
	}

	return r0
}

// Run provides a mock function with given fields:
func (_m *Input) Run() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stop provides a mock function with given fields:
func (_m *Input) Stop() {
	_m.Called()
}