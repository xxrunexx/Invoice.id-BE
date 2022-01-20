package data

import (
	"invoice-api/features/paymentmethod"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name     string
	IsActive bool
}

func toPaymentMethodRecord(pm paymentmethod.PaymentMethodCore) PaymentMethod {
	return PaymentMethod{
		Model: gorm.Model{
			ID:        pm.ID,
			CreatedAt: pm.CreatedAt,
			UpdatedAt: pm.UpdatedAt,
		},
		Name:     pm.Name,
		IsActive: pm.IsActive,
	}
}

func toPaymentMethodCore(pm PaymentMethod) paymentmethod.PaymentMethodCore {
	return paymentmethod.PaymentMethodCore{
		ID:       pm.ID,
		Name:     pm.Name,
		IsActive: pm.IsActive,
	}
}
