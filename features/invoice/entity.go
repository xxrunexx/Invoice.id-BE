package invoice

import "time"

type InvoiceCore struct {
	ID              uint
	ClientID        uint
	Total           int
	BillIssuerID    uint
	PaymentMethodID uint
	PaymentDue      time.Time
	PaymentStatus   string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Business interface {
	CreateInvoice(data InvoiceCore) (err error)
	GetAllInvoice(InvoiceCore) (invoices []InvoiceCore, err error)
	GetInvoiceById(id int) (invoice InvoiceCore, err error)
	DeleteInvoice(id int) error
	GetInvoiceByStatus(status string) (invoices []InvoiceCore, err error)
}

type Data interface {
	CreateInvoice(data InvoiceCore) (err error)
	GetAllInvoice(InvoiceCore) (invoices []InvoiceCore, err error)
	GetInvoiceById(id int) (invoice InvoiceCore, err error)
	DeleteInvoice(id int) error
	GetInvoiceByStatus(status string) (invoices []InvoiceCore, err error)
}
