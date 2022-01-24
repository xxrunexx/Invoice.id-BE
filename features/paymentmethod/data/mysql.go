package data

import (
	"errors"
	"invoice-api/features/paymentmethod"

	"gorm.io/gorm"
)

type PaymentMethodData struct {
	DB *gorm.DB
}

func NewMySqlPaymentMethod(DB *gorm.DB) paymentmethod.Data {
	return &PaymentMethodData{DB}
}

func (pmData *PaymentMethodData) CreatePaymentMethod(data paymentmethod.PaymentMethodCore) (paymentmethod.PaymentMethodCore, error) {
	convData := toPaymentMethodRecord(data)

	if err := pmData.DB.Create(&convData).Error; err != nil {
		return paymentmethod.PaymentMethodCore{}, err
	}

	record, err := pmData.GetPaymentMethodById(int(convData.ID))
	if err != nil {
		return paymentmethod.PaymentMethodCore{}, err
	}
	return record, nil
}

func (pmData *PaymentMethodData) GetPaymentMethodById(id int) (paymentmethod.PaymentMethodCore, error) {
	var singleData PaymentMethod

	err := pmData.DB.First(&singleData, id).Error

	if singleData.Name == "" && singleData.ID == 0 {
		return paymentmethod.PaymentMethodCore{}, errors.New("data not found")
	}

	if err != nil {
		return paymentmethod.PaymentMethodCore{}, err
	}
	return toPaymentMethodCore(singleData), nil
}

func (pmData *PaymentMethodData) GetAllPaymentMethod(data paymentmethod.PaymentMethodCore) ([]paymentmethod.PaymentMethodCore, error) {
	var paymentmethods []PaymentMethod

	err := pmData.DB.Find(&paymentmethods).Error

	if err != nil {
		return nil, err
	}
	return toPaymentMethodCoreList(paymentmethods), nil
}

func (pmData *PaymentMethodData) UpdatePaymentMethod(data paymentmethod.PaymentMethodCore) error {
	var singleData PaymentMethod

	// convData := toPaymentMethodRecord(data)
	err := pmData.DB.Model(&singleData).Where("id = ?", data.ID).Updates(PaymentMethod{Name: data.Name, IsActive: data.IsActive}).Error
	if err != nil {
		return err
	}
	return nil
}

func (pmData *PaymentMethodData) GetPaymentMethodByIsActive(isactive bool) ([]paymentmethod.PaymentMethodCore, error) {
	var paymentmethods []PaymentMethod

	err := pmData.DB.Where("payment_methods.is_active = ?", isactive).Find(&paymentmethods).Error
	if err != nil {
		return nil, err
	}
	return toPaymentMethodCoreList(paymentmethods), nil
}
