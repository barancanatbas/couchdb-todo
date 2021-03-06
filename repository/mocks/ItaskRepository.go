// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	model "couchdb/model"

	mock "github.com/stretchr/testify/mock"
)

// ItaskRepository is an autogenerated mock type for the ItaskRepository type
type ItaskRepository struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *ItaskRepository) Close() {
	_m.Called()
}

// Create provides a mock function with given fields: task
func (_m *ItaskRepository) Create(task model.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *ItaskRepository) Delete(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: id
func (_m *ItaskRepository) Get(id string) model.Task {
	ret := _m.Called(id)

	var r0 model.Task
	if rf, ok := ret.Get(0).(func(string) model.Task); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Task)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *ItaskRepository) GetAll() []model.Task {
	ret := _m.Called()

	var r0 []model.Task
	if rf, ok := ret.Get(0).(func() []model.Task); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Task)
		}
	}

	return r0
}

// Update provides a mock function with given fields: task
func (_m *ItaskRepository) Update(task model.Task) error {
	ret := _m.Called(task)

	var r0 error
	if rf, ok := ret.Get(0).(func(model.Task) error); ok {
		r0 = rf(task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
