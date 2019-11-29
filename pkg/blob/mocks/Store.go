// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import blob "github.com/applike/gosoline/pkg/blob"
import mock "github.com/stretchr/testify/mock"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// Copy provides a mock function with given fields: batch
func (_m *Store) Copy(batch blob.CopyBatch) {
	_m.Called(batch)
}

// CopyOne provides a mock function with given fields: obj
func (_m *Store) CopyOne(obj *blob.CopyObject) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.CopyObject) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Read provides a mock function with given fields: batch
func (_m *Store) Read(batch blob.Batch) {
	_m.Called(batch)
}

// ReadOne provides a mock function with given fields: obj
func (_m *Store) ReadOne(obj *blob.Object) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.Object) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Write provides a mock function with given fields: batch
func (_m *Store) Write(batch blob.Batch) {
	_m.Called(batch)
}

// WriteOne provides a mock function with given fields: obj
func (_m *Store) WriteOne(obj *blob.Object) error {
	ret := _m.Called(obj)

	var r0 error
	if rf, ok := ret.Get(0).(func(*blob.Object) error); ok {
		r0 = rf(obj)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
