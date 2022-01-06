package factory

import (
	// Billissuer Domain
	"invoice-api/driver"
	bibus "invoice-api/features/billissuer/business"
	bidata "invoice-api/features/billissuer/data"
	bipres "invoice-api/features/billissuer/presentation"
	// Reserved for other domains
)

type presenter struct {
	BillissuerPresentation bipres.BillIssuerHandler
}

func Init() presenter {
	// Bill Issuer
	billissuerData := bidata.NewMySqlBillIssuer(driver.DB)
	billissuerBusiness := bibus.NewBusinessBillIssuer(billissuerData)

	return presenter{
		BillissuerPresentation: *bipres.NewHandlerBillIssuer(billissuerBusiness),
	}
}
