package data

import (
	"invoice-api/features/invoice"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ClientID        int
	Total           int
	BillIssuerID    int
	PaymentMethodID int
	PaymentDue      time.Time
	PaymentStatus   string
}

func toInvoiceRecord(invoice invoice.InvoiceCore) Invoice {
	return Invoice{
		Model: gorm.Model{
			ID:        invoice.ID,
			CreatedAt: invoice.CreatedAt,
			UpdatedAt: invoice.UpdatedAt,
		},
		ClientID:        invoice.ClientID,
		Total:           invoice.Total,
		BillIssuerID:    invoice.BillIssuerID,
		PaymentMethodID: invoice.PaymentMethodID,
		PaymentDue:      invoice.PaymentDue,
		PaymentStatus:   invoice.PaymentStatus,
	}
}

func toInvoiceCore(in Invoice) invoice.InvoiceCore {
	return invoice.InvoiceCore{
		ID:              in.ID,
		ClientID:        in.ClientID,
		Total:           in.Total,
		BillIssuerID:    in.BillIssuerID,
		PaymentMethodID: in.PaymentMethodID,
		PaymentDue:      in.PaymentDue,
		PaymentStatus:   in.PaymentStatus,
	}
}
