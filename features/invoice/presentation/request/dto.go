package request

import (
	"invoice-api/features/invoice"
	"time"
)

type ReqInvoice struct {
	ClientID        int       `json:"client_id"`
	Total           int       `json:"total"`
	BillIssuerID    int       `json:"bill_issuer_id"`
	PaymentMethodID int       `json:"payment_method_id"`
	PaymentDue      time.Time `json:"payment_due"`
}

func (reqdata *ReqInvoice) ToInvoiceCore() invoice.InvoiceCore {
	return invoice.InvoiceCore{
		ClientID:        reqdata.ClientID,
		Total:           reqdata.Total,
		BillIssuerID:    reqdata.BillIssuerID,
		PaymentMethodID: reqdata.PaymentMethodID,
		PaymentDue:      reqdata.PaymentDue,
	}
}
