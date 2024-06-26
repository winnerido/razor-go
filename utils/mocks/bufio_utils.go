// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	bufio "bufio"
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// BufioUtils is an autogenerated mock type for the BufioUtils type
type BufioUtils struct {
	mock.Mock
}

// NewScanner provides a mock function with given fields: r
func (_m *BufioUtils) NewScanner(r io.Reader) *bufio.Scanner {
	ret := _m.Called(r)

	var r0 *bufio.Scanner
	if rf, ok := ret.Get(0).(func(io.Reader) *bufio.Scanner); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bufio.Scanner)
		}
	}

	return r0
}
