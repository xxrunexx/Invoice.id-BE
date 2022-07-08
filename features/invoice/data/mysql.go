package data

import (
	"errors"
	"fmt"
	"invoice-api/features/invoice"
	"strings"

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
	err := inData.DB.Joins("Client").Joins("BillIssuer").Joins("PaymentMethod").Find(&invoices).Error

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

	err := inData.DB.Where("invoices.id = ?", id).Joins("Client").Joins("BillIssuer").Joins("PaymentMethod").Find(&singleData).Error
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

	// err := inData.DB.Where("invoices.payment_status = ?", status).Joins("Client").Joins("BillIssuerDetail").Joins("PaymentMethod").Find(&invoices).Error
	err := inData.DB.Where("payment_status = ?", status).Joins("Client").Joins("BillIssuerDetail").Joins("PaymentMethod").Find(&invoices).Error
	if err != nil {
		return nil, err
	}

	return toInvoiceCoreList(invoices), nil
}

func (inData *InvoiceData) UpdateInvoice(data invoice.InvoiceCore) error {
	fmt.Println("Isi data di data : ", data)
	// var singleData Invoice
	convData := toInvoiceRecord(data)
	fmt.Println("Isi convData di data : ", convData)
	err := inData.DB.Debug().Model(&Invoice{}).Where("id = ?", data.ID).Updates(&convData).Error
	if err != nil {
		return err
	}
	err = inData.UpdatePaymentLink(convData.PaymentLink, convData.ID)
	if err != nil {
		return err
	}
	return nil
}

func (inData *InvoiceData) UpdateTransactionStatus(transactionID int64, PaymentStatus string) error {
	err := inData.DB.Model(&Invoice{}).Where("id = ?", transactionID).Update("payment_status", PaymentStatus).Error
	if err != nil {
		return err
	}
	return nil
}

func (inData *InvoiceData) UpdatePaymentLink(url string, id uint) error {
	err := inData.DB.Debug().Model(&Invoice{}).Where("id = ?", id).Update("payment_link", url).Error
	if err != nil {
		return err
	}
	return nil
}

func (inData *InvoiceData) GetInvoiceByNik(nik int) ([]invoice.InvoiceCore, error) {
	var invoices []Invoice
	// var clients []Client

	err := inData.DB.Joins("JOIN clients ON clients.id = invoices.client_id AND clients.nik = ?", nik).Joins("Client").Joins("BillIssuer").Joins("PaymentMethod").Find(&invoices).Error
	// db.Joins("JOIN clients ON clients.id = invoices.client_id AND clients.nik = ?", nik
	if err != nil {
		return nil, err
	}

	return toInvoiceCoreList(invoices), nil
}

func (inData *InvoiceData) GetInvoiceByName(name string) ([]invoice.InvoiceCore, error) {
	var invoices []Invoice

	myslice := []string{"%", name, "%"}
	result := strings.Join(myslice, "")
	fmt.Println(result)

	// err := inData.DB.Where("clients.name = ?", name).Joins("Invoice").Joins("Client").Joins("BillIssuerDetail").Joins("PaymentMethod").Find(&invoices).Error
	err := inData.DB.Joins("JOIN clients ON clients.id = invoices.client_id AND clients.name LIKE ?", result).Joins("Client").Joins("BillIssuer").Joins("PaymentMethod").First(&invoices).Error
	if err != nil {
		return nil, err
	}

	return toInvoiceCoreList(invoices), nil
}

func (inData *InvoiceData) InsertCSV(datas []invoice.InvoiceCore) error {
	// convData := toInvoiceRecordList(datas)
	var invoices []Invoice
	if err := inData.DB.Create(&invoices).Error; err != nil {
		return err
	}
	return nil
}
