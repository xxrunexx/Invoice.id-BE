package data

import (
	"invoice-api/features/paymentmethod"

	"gorm.io/gorm"
)

type PaymentMethodData struct {
	DB *gorm.DB
}

func NewMySqlPaymentMethod(DB *gorm.DB) paymentmethod.Data {
	return &PaymentMethodData{DB}
}

func (pmData *PaymentMethodData) CreatePaymentMethod(data paymentmethod.PaymentMethodCore) error {
	convData := toPaymentMethodRecord(data)

	if err := pmData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
