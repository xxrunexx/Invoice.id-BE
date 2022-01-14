package response

import (
	"invoice-api/features/invoice"
	"time"
)

type RespInvoice struct {
	ID              uint      `json:"id"`
	ClientID        uint      `json:"client_id"`
	Total           int       `json:"total"`
	BillIssuerID    uint      `json:"bill_issuer_id"`
	PaymentMethodID uint      `json:"payment_method_id"`
	PaymentDue      time.Time `json:"payment_due"`
	PaymentStatus   string    `json:"payment_status"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func ToInvoiceResponse(invoice invoice.InvoiceCore) RespInvoice {
	return RespInvoice{
		ID:              invoice.ID,
		ClientID:        invoice.ClientID,
		Total:           invoice.Total,
		BillIssuerID:    invoice.BillIssuerID,
		PaymentMethodID: invoice.PaymentMethodID,
		PaymentDue:      invoice.PaymentDue,
		PaymentStatus:   invoice.PaymentStatus,
		CreatedAt:       invoice.CreatedAt,
		UpdatedAt:       invoice.UpdatedAt,
	}
}

func ToInvoiceResponseList(inList []invoice.InvoiceCore) []RespInvoice {
	convIn := []RespInvoice{}

	for _, invoice := range inList {
		convIn = append(convIn, ToInvoiceResponse(invoice))
	}
	return convIn
}
