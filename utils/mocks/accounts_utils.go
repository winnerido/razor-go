// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ecdsa "crypto/ecdsa"

	mock "github.com/stretchr/testify/mock"
)

// AccountsUtils is an autogenerated mock type for the AccountsUtils type
type AccountsUtils struct {
	mock.Mock
}

// GetPrivateKey provides a mock function with given fields: address, password, keystorePath
func (_m *AccountsUtils) GetPrivateKey(address string, password string, keystorePath string) *ecdsa.PrivateKey {
	ret := _m.Called(address, password, keystorePath)

	var r0 *ecdsa.PrivateKey
	if rf, ok := ret.Get(0).(func(string, string, string) *ecdsa.PrivateKey); ok {
		r0 = rf(address, password, keystorePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecdsa.PrivateKey)
		}
	}

	return r0
}
