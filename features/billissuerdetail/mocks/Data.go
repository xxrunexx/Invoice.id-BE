// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	billissuerdetail "invoice-api/features/billissuerdetail"

	mock "github.com/stretchr/testify/mock"
)

// Data is an autogenerated mock type for the Data type
type Data struct {
	mock.Mock
}

// CreateBillIssuerDetail provides a mock function with given fields: data
func (_m *Data) CreateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) (billissuerdetail.BillIssuerDetailCore, error) {
	ret := _m.Called(data)

	var r0 billissuerdetail.BillIssuerDetailCore
	if rf, ok := ret.Get(0).(func(billissuerdetail.BillIssuerDetailCore) billissuerdetail.BillIssuerDetailCore); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Get(0).(billissuerdetail.BillIssuerDetailCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(billissuerdetail.BillIssuerDetailCore) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBillIssuerDetailByBillIssuerId provides a mock function with given fields: id
func (_m *Data) GetBillIssuerDetailByBillIssuerId(id int) (billissuerdetail.BillIssuerDetailCore, error) {
	ret := _m.Called(id)

	var r0 billissuerdetail.BillIssuerDetailCore
	if rf, ok := ret.Get(0).(func(int) billissuerdetail.BillIssuerDetailCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(billissuerdetail.BillIssuerDetailCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBillIssuerDetailById provides a mock function with given fields: id
func (_m *Data) GetBillIssuerDetailById(id int) (billissuerdetail.BillIssuerDetailCore, error) {
	ret := _m.Called(id)

	var r0 billissuerdetail.BillIssuerDetailCore
	if rf, ok := ret.Get(0).(func(int) billissuerdetail.BillIssuerDetailCore); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(billissuerdetail.BillIssuerDetailCore)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateBillIssuerDetail provides a mock function with given fields: data
func (_m *Data) UpdateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) error {
	ret := _m.Called(data)

	var r0 error
	if rf, ok := ret.Get(0).(func(billissuerdetail.BillIssuerDetailCore) error); ok {
		r0 = rf(data)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
