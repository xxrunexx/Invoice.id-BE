package response

import (
	"invoice-api/features/invoice"
	"time"
)

type RespInvoice struct {
	ID                uint      `json:"id"`
	ClientID          uint      `json:"client_id"`
	ClientName        string    `json:"client_name"`
	ClientPhone       string    `json:"client_phone"`
	ClientAddress     string    `json:"client_address"`
	ClientEmail       string    `json:"client_email"`
	Total             int       `json:"total"`
	BillIssuerID      uint      `json:"bill_issuer_id"`
	BillIssuerName    string    `json:"bill_issuer_name"`
	PaymentMethodID   uint      `json:"payment_method_id"`
	PaymentMethodName string    `json:"payment_method_name"`
	PaymentDue        time.Time `json:"payment_due"`
	PaymentStatus     string    `json:"payment_status"`
	PaymentTerms      int       `json:"payment_terms"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func ToInvoiceResponse(in invoice.InvoiceCore) RespInvoice {
	return RespInvoice{
		ID:                in.ID,
		ClientID:          in.ClientID,
		ClientName:        in.ClientName,
		ClientPhone:       in.ClientPhone,
		ClientAddress:     in.ClientAddress,
		ClientEmail:       in.ClientEmail,
		Total:             in.Total,
		BillIssuerID:      in.BillIssuerID,
		BillIssuerName:    in.BillIssuerName,
		PaymentMethodID:   in.PaymentMethodID,
		PaymentMethodName: in.PaymentMethodName,
		PaymentDue:        in.PaymentDue,
		PaymentStatus:     in.PaymentStatus,
		PaymentTerms:      in.PaymentTerms,
		CreatedAt:         in.CreatedAt,
		UpdatedAt:         in.UpdatedAt,
	}
}

func ToInvoiceResponseList(inList []invoice.InvoiceCore) []RespInvoice {
	convIn := []RespInvoice{}

	for _, invoice := range inList {
		convIn = append(convIn, ToInvoiceResponse(invoice))
	}
	return convIn
}
