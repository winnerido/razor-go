// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// PathInterface is an autogenerated mock type for the PathInterface type
type PathInterface struct {
	mock.Mock
}

// GetCommitDataFileName provides a mock function with given fields: address
func (_m *PathInterface) GetCommitDataFileName(address string) (string, error) {
	ret := _m.Called(address)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigFilePath provides a mock function with given fields:
func (_m *PathInterface) GetConfigFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultPath provides a mock function with given fields:
func (_m *PathInterface) GetDefaultPath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDisputeDataFileName provides a mock function with given fields: address
func (_m *PathInterface) GetDisputeDataFileName(address string) (string, error) {
	ret := _m.Called(address)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobFilePath provides a mock function with given fields:
func (_m *PathInterface) GetJobFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLogFilePath provides a mock function with given fields: fileName
func (_m *PathInterface) GetLogFilePath(fileName string) (string, error) {
	ret := _m.Called(fileName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProposeDataFileName provides a mock function with given fields: address
func (_m *PathInterface) GetProposeDataFileName(address string) (string, error) {
	ret := _m.Called(address)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
