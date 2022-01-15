package data

import (
	"invoice-api/features/paymentmethod"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Name     string
	IsActive bool `gorm:"default:true"`
}

func toPaymentMethodRecord(paymentmethod paymentmethod.PaymentMethodCore) PaymentMethod {
	return PaymentMethod{
		Model: gorm.Model{
			ID:        paymentmethod.ID,
			CreatedAt: paymentmethod.CreatedAt,
			UpdatedAt: paymentmethod.UpdatedAt,
		},
		Name:     paymentmethod.Name,
		IsActive: paymentmethod.IsActive,
	}
}
