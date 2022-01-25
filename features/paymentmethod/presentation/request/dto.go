package request

import "invoice-api/features/paymentmethod"

type ReqPaymentMethod struct {
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

type ReqPaymentMethodUpdate struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func (data *ReqPaymentMethod) ToPaymentMethodCore() paymentmethod.PaymentMethodCore {
	return paymentmethod.PaymentMethodCore{
		Name:     data.Name,
		IsActive: data.IsActive,
	}
}
func (data *ReqPaymentMethodUpdate) ToPaymentMethodCore() paymentmethod.PaymentMethodCore {
	return paymentmethod.PaymentMethodCore{
		ID:       data.ID,
		Name:     data.Name,
		IsActive: data.IsActive,
	}
}
