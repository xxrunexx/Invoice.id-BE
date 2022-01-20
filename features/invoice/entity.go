package invoice

import "time"

type InvoiceCore struct {
	ID                uint
	ClientID          uint
	ClientName        string
	ClientPhone       string
	ClientAddress     string
	ClientEmail       string
	Total             int
	BillIssuerID      uint
	BillIssuerName    string
	PaymentMethodID   uint
	PaymentMethodName string
	PaymentDue        time.Time
	PaymentStatus     string
	PaymentTerms      int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Business interface {
	CreateInvoice(data InvoiceCore) (err error)
	GetAllInvoice(InvoiceCore) (invoices []InvoiceCore, err error)
	GetInvoiceById(id int) (invoice InvoiceCore, err error)
	DeleteInvoice(id int) error
	GetInvoiceByStatus(status string) (invoices []InvoiceCore, err error)
	UpdateInvoice(data InvoiceCore) error
	SendInvoice(id int) (invoice InvoiceCore, err error)
}

type Data interface {
	CreateInvoice(data InvoiceCore) (err error)
	GetAllInvoice(InvoiceCore) (invoices []InvoiceCore, err error)
	GetInvoiceById(id int) (invoice InvoiceCore, err error)
	DeleteInvoice(id int) error
	GetInvoiceByStatus(status string) (invoices []InvoiceCore, err error)
	UpdateInvoice(data InvoiceCore) error
}
