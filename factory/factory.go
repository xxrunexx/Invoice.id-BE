package factory

import (
	// Billissuer Domain
	"invoice-api/driver"
	bibus "invoice-api/features/billissuer/business"
	bidata "invoice-api/features/billissuer/data"
	bipres "invoice-api/features/billissuer/presentation"

	// Client Domain
	clbus "invoice-api/features/client/business"
	cldata "invoice-api/features/client/data"
	clpres "invoice-api/features/client/presentation"
	// Reserved for other domains
)

type presenter struct {
	BillissuerPresentation bipres.BillIssuerHandler
	ClientPresentation     clpres.ClientHandler
}

func Init() presenter {
	// Bill Issuer
	billissuerData := bidata.NewMySqlBillIssuer(driver.DB)
	billissuerBusiness := bibus.NewBusinessBillIssuer(billissuerData)

	// Client
	clientData := cldata.NewMySqlClient(driver.DB)
	clientBusiness := clbus.NewBusinessClient(clientData)
	return presenter{
		BillissuerPresentation: *bipres.NewHandlerBillIssuer(billissuerBusiness),
		ClientPresentation:     *clpres.NewHandlerClient(clientBusiness),
	}
}
