// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DataMapper is an autogenerated mock type for the DataMapper type
type DataMapper struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *DataMapper) Delete(_a0 ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Insert provides a mock function with given fields: _a0
func (_m *DataMapper) Insert(_a0 ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *DataMapper) Update(_a0 ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, _a0...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(_a0...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}