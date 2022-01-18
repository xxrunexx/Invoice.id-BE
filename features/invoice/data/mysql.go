package data

import (
	"errors"
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

func (inData *InvoiceData) GetAllInvoice(data invoice.InvoiceCore) ([]invoice.InvoiceCore, error) {
	var invoices []Invoice

	err := inData.DB.Find(&invoices).Error

	if err != nil {
		return nil, err
	}

	return toInvoiceCoreList(invoices), nil
}

func (inData *InvoiceData) DeleteInvoice(id int) error {
	var singleInvoice Invoice

	err := inData.DB.Where("id = ?", id).Delete(&singleInvoice).Error
	if err != nil {
		return err
	}
	return nil
}

func (inData *InvoiceData) GetInvoiceById(id int) (invoice.InvoiceCore, error) {
	var singleData Invoice

	err := inData.DB.First(&singleData, id).Error

	if singleData.ID == 0 {
		return invoice.InvoiceCore{}, errors.New("data not found")
	}

	if err != nil {
		return invoice.InvoiceCore{}, err
	}

	return toInvoiceCore(singleData), nil
}

func (inData *InvoiceData) GetInvoiceByStatus(status string) ([]invoice.InvoiceCore, error) {
	var invoices []Invoice

	err := inData.DB.Where("payment_status = ?", status).Find(&invoices).Error
	if err != nil {
		return nil, err
	}

	return toInvoiceCoreList(invoices), nil
}

func (inData *InvoiceData) UpdateInvoice(data invoice.InvoiceCore) error {
	var singleData Invoice
	convData := toInvoiceRecord(data)
	err := inData.DB.Model(&singleData).Where("id = ?", data.ID).Updates(&convData).Error

	if err != nil {
		return err
	}
	return nil
}
