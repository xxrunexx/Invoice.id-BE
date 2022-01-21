package business

import (
	"errors"
	"invoice-api/features/paymentmethod"
	"invoice-api/helper"
)

type PaymentMethodBusiness struct {
	paymentmethodData paymentmethod.Data
}

func NewBusinessPaymentMethod(pmData paymentmethod.Data) paymentmethod.Business {
	return &PaymentMethodBusiness{pmData}
}

func (pmBusiness *PaymentMethodBusiness) CreatePaymentMethod(data paymentmethod.PaymentMethodCore) (paymentmethod.PaymentMethodCore, error) {
	if helper.IsEmpty(data.Name) {
		return paymentmethod.PaymentMethodCore{}, errors.New("bad request")
	}

	result, err := pmBusiness.paymentmethodData.CreatePaymentMethod(data)
	if err != nil {
		return paymentmethod.PaymentMethodCore{}, err
	}
	return result, nil
}

func (pmBusiness *PaymentMethodBusiness) GetPaymentMethodById(id int) (paymentmethod.PaymentMethodCore, error) {
	pmData, err := pmBusiness.paymentmethodData.GetPaymentMethodById(id)

	if err != nil {
		return paymentmethod.PaymentMethodCore{}, err
	}
	return pmData, nil
}

func (pmBusiness *PaymentMethodBusiness) GetAllPaymentMethod(data paymentmethod.PaymentMethodCore) ([]paymentmethod.PaymentMethodCore, error) {
	pmDatas, err := pmBusiness.paymentmethodData.GetAllPaymentMethod(data)

	if err != nil {
		return nil, err
	}
	return pmDatas, nil
}

func (pmBusiness *PaymentMethodBusiness) UpdatePaymentMethod(data paymentmethod.PaymentMethodCore) error {
	if helper.IsEmpty(data.Name) {
		return errors.New("invalid data")
	}

	err := pmBusiness.paymentmethodData.UpdatePaymentMethod(data)
	if err != nil {
		return err
	}
	return nil
}
