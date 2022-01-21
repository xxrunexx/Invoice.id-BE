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

func ToPaymentMethodResponseList(pmList []paymentmethod.PaymentMethodCore) []RespPaymentMethod {
	convPm := []RespPaymentMethod{}

	for _, paymentmethod := range pmList {
		convPm = append(convPm, ToPaymentMethodResponse(paymentmethod))
	}
	return convPm
}
