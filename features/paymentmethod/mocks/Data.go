// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	paymentmethod "invoice-api/features/paymentmethod"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreatePaymentMethod provides a mock function with given fields: data
func (_m *Data) CreatePaymentMethod(data paymentmethod.PaymentMethodCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(paymentmethod.PaymentMethodCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}