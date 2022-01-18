package data

import (
	"invoice-api/features/invoice"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ClientID        uint
	Total           int
	BillIssuerID    uint
	PaymentMethodID uint
	PaymentDue      time.Time
	PaymentStatus   string `gorm:"default:draft"`
	PaymentTerms    int
}

func toInvoiceRecord(in invoice.InvoiceCore) Invoice {
	return Invoice{
		Model: gorm.Model{
			ID:        in.ID,
			CreatedAt: in.CreatedAt,
			UpdatedAt: in.UpdatedAt,
		},
		ClientID:        in.ClientID,
		Total:           in.Total,
		BillIssuerID:    in.BillIssuerID,
		PaymentMethodID: in.PaymentMethodID,
		PaymentDue:      in.PaymentDue,
		PaymentStatus:   in.PaymentStatus,
		PaymentTerms:    in.PaymentTerms,
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
		PaymentTerms:    in.PaymentTerms,
		CreatedAt:       in.CreatedAt,
		UpdatedAt:       in.UpdatedAt,
	}
}

func toInvoiceCoreList(inList []Invoice) []invoice.InvoiceCore {
	convIn := []invoice.InvoiceCore{}

	for _, invoice := range inList {
		convIn = append(convIn, toInvoiceCore(invoice))
	}
	return convIn
}
