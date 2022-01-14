package factory

import (
	"invoice-api/driver"
	// Billissuer Domain

	bibus "invoice-api/features/billissuer/business"
	bidata "invoice-api/features/billissuer/data"
	bipres "invoice-api/features/billissuer/presentation"

	// Client Domain
	clbus "invoice-api/features/client/business"
	cldata "invoice-api/features/client/data"
	clpres "invoice-api/features/client/presentation"

	// Invoice Domain
	inbus "invoice-api/features/invoice/business"
	indata "invoice-api/features/invoice/data"
	inpres "invoice-api/features/invoice/presentation"
	// Reserved for other domains
)

type presenter struct {
	BillissuerPresentation bipres.BillIssuerHandler
	ClientPresentation     clpres.ClientHandler
	InvoicePresentation    inpres.InvoiceHandler
}

func Init() presenter {
	// Bill Issuer
	billissuerData := bidata.NewMySqlBillIssuer(driver.DB)
	billissuerBusiness := bibus.NewBusinessBillIssuer(billissuerData)

	// Client
	clientData := cldata.NewMySqlClient(driver.DB)
	clientBusiness := clbus.NewBusinessClient(clientData)

	// Invoice
	invoiceData := indata.NewMySqlInvoice(driver.DB)
	invoiceBusiness := inbus.NewBusinessInvoice(invoiceData)

	return presenter{
		BillissuerPresentation: *bipres.NewHandlerBillIssuer(billissuerBusiness),
		ClientPresentation:     *clpres.NewHandlerClient(clientBusiness),
		InvoicePresentation:    *inpres.NewHandlerInvoice(invoiceBusiness),
	}
}
