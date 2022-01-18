package request

import (
	"invoice-api/features/invoice"
	"time"
)

type ReqInvoice struct {
	ClientID        uint      `json:"client_id"`
	Total           int       `json:"total"`
	BillIssuerID    uint      `json:"bill_issuer_id"`
	PaymentMethodID uint      `json:"payment_method_id"`
	PaymentDue      time.Time `json:"payment_due"`
	PaymentTerms    int       `json:"payment_terms"`
}

type ReqInvoiceUpdate struct {
	ID              uint      `json:"id"`
	ClientID        uint      `json:"client_id"`
	Total           int       `json:"total"`
	BillIssuerID    uint      `json:"bill_issuer_id"`
	PaymentMethodID uint      `json:"payment_method_id"`
	PaymentDue      time.Time `json:"payment_due"`
	PaymentTerms    int       `json:"payment_terms"`
	PaymentStatus   string    `json:"payment_status"`
}

func (reqdata *ReqInvoice) ToInvoiceCore() invoice.InvoiceCore {
	return invoice.InvoiceCore{
		ClientID:        reqdata.ClientID,
		Total:           reqdata.Total,
		BillIssuerID:    reqdata.BillIssuerID,
		PaymentMethodID: reqdata.PaymentMethodID,
		PaymentDue:      reqdata.PaymentDue,
		PaymentTerms:    reqdata.PaymentTerms,
	}
}

func (reqdata *ReqInvoiceUpdate) ToInvoiceCore() invoice.InvoiceCore {
	return invoice.InvoiceCore{
		ID:              reqdata.ID,
		ClientID:        reqdata.ClientID,
		Total:           reqdata.Total,
		BillIssuerID:    reqdata.BillIssuerID,
		PaymentMethodID: reqdata.PaymentMethodID,
		PaymentDue:      reqdata.PaymentDue,
		PaymentTerms:    reqdata.PaymentTerms,
		PaymentStatus:   reqdata.PaymentStatus,
	}
}
