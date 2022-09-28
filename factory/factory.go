package factory

import (
	"invoice-api/config"
	"invoice-api/driver"
	"log"

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

	// BillIssuerDetail Domain
	bidbus "invoice-api/features/billissuerdetail/business"
	biddata "invoice-api/features/billissuerdetail/data"
	bidpres "invoice-api/features/billissuerdetail/presentation"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type presenter struct {
	BillissuerPresentation       bipres.BillIssuerHandler
	ClientPresentation           clpres.ClientHandler
	InvoicePresentation          inpres.InvoiceHandler
	BillissuerdetailPresentation bidpres.BillIssuerDetailHandler
}

func Init() presenter {
	//Initiate client for Midtrans Snap
	var s snap.Client
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	s.New(config.MidtransApi, midtrans.Sandbox)

	var c coreapi.Client
	c.New(config.MidtransApi, midtrans.Sandbox)

	// Bill Issuer
	billissuerData := bidata.NewMySqlBillIssuer(driver.DB)
	billissuerBusiness := bibus.NewBusinessBillIssuer(billissuerData)

	// Client
	clientData := cldata.NewMySqlClient(driver.DB)
	clientBusiness := clbus.NewBusinessClient(clientData)

	// Invoice
	invoiceData := indata.NewMySqlInvoice(driver.DB)
	invoiceBusiness := inbus.NewBusinessInvoice(invoiceData, s, c)

	// Bill Issuer Detail
	billissuerdetailData := biddata.NewMySqlBillIssuerDetail(driver.DB)
	billissuerdetailBusiness := bidbus.NewBusinessBillIssuerDetail(billissuerdetailData, billissuerData)

	return presenter{
		BillissuerPresentation:       *bipres.NewHandlerBillIssuer(billissuerBusiness),
		ClientPresentation:           *clpres.NewHandlerClient(clientBusiness),
		InvoicePresentation:          *inpres.NewHandlerInvoice(invoiceBusiness),
		BillissuerdetailPresentation: *bidpres.NewHandlerBillIssuerDetail(billissuerdetailBusiness),
	}
}
