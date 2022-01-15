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

func (pmBusiness *PaymentMethodBusiness) CreatePaymentMethod(data paymentmethod.PaymentMethodCore) error {
	if !helper.IsEmpty(data.Name) {
		return errors.New("bad request")
	}

	err := pmBusiness.paymentmethodData.CreatePaymentMethod(data)
	if err != nil {
		return err
	}
	return nil
}
