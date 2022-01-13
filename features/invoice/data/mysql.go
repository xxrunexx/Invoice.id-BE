package data

import (
	"invoice-api/features/invoice"

	"gorm.io/gorm"
)

type InvoiceData struct {
	DB *gorm.DB
}

func NewMySqlInvoice(DB *gorm.DB) invoice.Data {
	return &InvoiceData{DB}
}

func (inData *InvoiceData) CreateInvoice(data invoice.InvoiceCore) error {
	convData := toInvoiceRecord(data)

	if err := inData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
