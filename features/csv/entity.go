package csv

import (
	"time"
)

// type InvoiceCsvCore struct {
// 	ID            uint
// 	ClientID      uint
// 	ClientNIK     int
// 	ClientName    string
// 	ClientPhone   string
// 	ClientAddress string
// 	ClientEmail   string
// 	Item          string
// 	Total         int
// 	BillIssuerID  uint
// 	// PaymentMethodID uint   `csv:"payment_method_id"`
// 	PaymentDue    time.Time
// 	PaymentStatus string
// 	PaymentTerms  int
// 	ClientCore    client.ClientCore
// 	CreatedAt     time.Time
// 	UpdatedAt     time.Time
// }

type InvoiceCore struct {
	ID                uint
	ClientID          uint
	Item              string
	Total             int
	BillIssuerID      uint
	PaymentMethodID   uint
	PaymentMethodName string
	PaymentDue        time.Time
	PaymentStatus     string
	PaymentTerms      int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type ClientCore struct {
	ID      uint
	NIK     int
	Name    string
	Phone   string
	Address string
	Email   string
}

type Business interface {
	CreateBulkInvoice(csv []InvoiceCore) error
}

// type CSVInvoice struct {
// 	ClientID     uint   `csv:"client_id"`
// 	Item         string `csv:"item"`
// 	Total        int    `csv:"total"`
// 	BillIssuerID uint   `csv:"bill_issuer_id"`
// 	// PaymentMethodID uint   `csv:"payment_method_id"`
// 	PaymentTerms int `csv:"payment_terms"`
// }
// type Client struct {
// 	ClientNIK     int    `csv:"client_nik"`
// 	ClientName    string `csv:"client_name"`
// 	ClientPhone   string `csv:"client_phone"`
// 	ClientAddress string `csv:"client_address"`
// 	ClientEmail   string `csv:"client_email"`
// }
