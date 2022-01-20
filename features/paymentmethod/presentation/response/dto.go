package response

import "invoice-api/features/paymentmethod"

type RespPaymentMethod struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}

func ToPaymentMethodResponse(pm paymentmethod.PaymentMethodCore) RespPaymentMethod {
	return RespPaymentMethod{
		Id:       pm.ID,
		Name:     pm.Name,
		IsActive: pm.IsActive,
	}
}
