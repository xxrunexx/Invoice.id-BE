// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	billissuer "invoice-api/features/billissuer"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreateBillIssuer provides a mock function with given fields: data
func (_m *Data) CreateBillIssuer(data billissuer.BillIssuerCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(billissuer.BillIssuerCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllBillIssuer provides a mock function with given fields: _a0
func (_m *Data) GetAllBillIssuer(_a0 billissuer.BillIssuerCore) ([]billissuer.BillIssuerCore, error) {
	ret := _m.Called(_a0)

	var r0 []billissuer.BillIssuerCore
	if rf, ok := ret.Get(0).(func(billissuer.BillIssuerCore) []billissuer.BillIssuerCore); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]billissuer.BillIssuerCore)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(billissuer.BillIssuerCore) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBillIssuerByEmail provides a mock function with given fields: email
func (_m *Data) GetBillIssuerByEmail(email string) (bool, error) {
	ret := _m.Called(email)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(email)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBillIssuerById provides a mock function with given fields: id
func (_m *Data) GetBillIssuerById(id int) (billissuer.BillIssuerCore, error) {
	ret := _m.Called(id)

	var r0 billissuer.BillIssuerCore
	if rf, ok := ret.Get(0).(func(int) billissuer.BillIssuerCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(billissuer.BillIssuerCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LoginBillIssuer provides a mock function with given fields: data
func (_m *Data) LoginBillIssuer(data billissuer.BillIssuerCore) (billissuer.BillIssuerCore, error) {
	ret := _m.Called(data)

	var r0 billissuer.BillIssuerCore
	if rf, ok := ret.Get(0).(func(billissuer.BillIssuerCore) billissuer.BillIssuerCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(billissuer.BillIssuerCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(billissuer.BillIssuerCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBillIssuer provides a mock function with given fields: data
func (_m *Data) UpdateBillIssuer(data billissuer.BillIssuerCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(billissuer.BillIssuerCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
