// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	ecdsa "crypto/ecdsa"
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	mock "github.com/stretchr/testify/mock"
)

// BindUtils is an autogenerated mock type for the BindUtils type
type BindUtils struct {
	mock.Mock
}

// NewKeyedTransactorWithChainID provides a mock function with given fields: key, chainID
func (_m *BindUtils) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	ret := _m.Called(key, chainID)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(*ecdsa.PrivateKey, *big.Int) *bind.TransactOpts); ok {
		r0 = rf(key, chainID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ecdsa.PrivateKey, *big.Int) error); ok {
		r1 = rf(key, chainID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
