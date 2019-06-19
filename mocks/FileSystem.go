// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import vfs "github.com/c2fo/vfs/v5"

// FileSystem is an autogenerated mock type for the FileSystem type
type FileSystem struct {
	mock.Mock
}

// Name provides a mock function with given fields:
func (_m *FileSystem) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewFile provides a mock function with given fields: volume, absFilePath
func (_m *FileSystem) NewFile(volume string, absFilePath string) (vfs.File, error) {
	ret := _m.Called(volume, absFilePath)

	var r0 vfs.File
	if rf, ok := ret.Get(0).(func(string, string) vfs.File); ok {
		r0 = rf(volume, absFilePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vfs.File)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(volume, absFilePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLocation provides a mock function with given fields: volume, absLocPath
func (_m *FileSystem) NewLocation(volume string, absLocPath string) (vfs.Location, error) {
	ret := _m.Called(volume, absLocPath)

	var r0 vfs.Location
	if rf, ok := ret.Get(0).(func(string, string) vfs.Location); ok {
		r0 = rf(volume, absLocPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vfs.Location)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(volume, absLocPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Retry provides a mock function with given fields:
func (_m *FileSystem) Retry() vfs.Retry {
	ret := _m.Called()

	var r0 vfs.Retry
	if rf, ok := ret.Get(0).(func() vfs.Retry); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(vfs.Retry)
		}
	}

	return r0
}

// Scheme provides a mock function with given fields:
func (_m *FileSystem) Scheme() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
