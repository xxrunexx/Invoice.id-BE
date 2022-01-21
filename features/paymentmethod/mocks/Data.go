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
func (_m *Data) CreatePaymentMethod(data paymentmethod.PaymentMethodCore) (paymentmethod.PaymentMethodCore, error) {
	ret := _m.Called(data)

	var r0 paymentmethod.PaymentMethodCore
	if rf, ok := ret.Get(0).(func(paymentmethod.PaymentMethodCore) paymentmethod.PaymentMethodCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(paymentmethod.PaymentMethodCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(paymentmethod.PaymentMethodCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllPaymentMethod provides a mock function with given fields: _a0
func (_m *Data) GetAllPaymentMethod(_a0 paymentmethod.PaymentMethodCore) ([]paymentmethod.PaymentMethodCore, error) {
	ret := _m.Called(_a0)

	var r0 []paymentmethod.PaymentMethodCore
	if rf, ok := ret.Get(0).(func(paymentmethod.PaymentMethodCore) []paymentmethod.PaymentMethodCore); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentmethod.PaymentMethodCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(paymentmethod.PaymentMethodCore) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentMethodById provides a mock function with given fields: id
func (_m *Data) GetPaymentMethodById(id int) (paymentmethod.PaymentMethodCore, error) {
	ret := _m.Called(id)

	var r0 paymentmethod.PaymentMethodCore
	if rf, ok := ret.Get(0).(func(int) paymentmethod.PaymentMethodCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(paymentmethod.PaymentMethodCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentMethodByIsActive provides a mock function with given fields: isactive
func (_m *Data) GetPaymentMethodByIsActive(isactive bool) ([]paymentmethod.PaymentMethodCore, error) {
	ret := _m.Called(isactive)

	var r0 []paymentmethod.PaymentMethodCore
	if rf, ok := ret.Get(0).(func(bool) []paymentmethod.PaymentMethodCore); ok {
		r0 = rf(isactive)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]paymentmethod.PaymentMethodCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(isactive)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePaymentMethod provides a mock function with given fields: data
func (_m *Data) UpdatePaymentMethod(data paymentmethod.PaymentMethodCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(paymentmethod.PaymentMethodCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
