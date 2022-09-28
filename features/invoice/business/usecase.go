package business

import (
	"errors"
	"fmt"
	"invoice-api/features/invoice"
	"invoice-api/helper"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/midtrans/midtrans-go/snap"
)

type InvoiceBusiness struct {
	invoiceData        invoice.Data
	midtransClient     snap.Client
	midtransCoreClient coreapi.Client
}

func NewBusinessInvoice(inData invoice.Data, midtransClient snap.Client, midtransCoreClient coreapi.Client) invoice.Business {
	return &InvoiceBusiness{inData, midtransClient, midtransCoreClient}
}

func (inBusiness *InvoiceBusiness) CreateInvoice(data invoice.InvoiceCore) error {
	t := time.Now()
	if data.PaymentTerms == 7 {
		data.PaymentDue = t.Add(time.Hour * 24 * 7)
	} else if data.PaymentTerms == 10 {
		data.PaymentDue = t.Add(time.Hour * 24 * 10)
	} else if data.PaymentTerms == 30 {
		data.PaymentDue = t.Add(time.Hour * 24 * 30)
	}

	err := inBusiness.invoiceData.CreateInvoice(data)
	if err != nil {
		return err
	}
	return nil
}

func (inBusiness *InvoiceBusiness) SendInvoice(id int) (invoice.InvoiceCore, error) {
	inData, err := inBusiness.invoiceData.GetInvoiceById(id)

	if err != nil {
		return invoice.InvoiceCore{}, err
	}
	fmt.Println("Isi Client name : ", inData.ClientName)
	fmt.Println("Isi Total : ", inData.Total)
	fmt.Println("Isi CreatedAt : ", inData.CreatedAt)
	fmt.Println("Isi Payment terms: ", inData.PaymentTerms)
	fmt.Println("Isi Payment due: ", inData.PaymentDue)
	helper.SendGmail(inData)

	return inData, nil
}

func (inBusiness *InvoiceBusiness) GetAllInvoice(data invoice.InvoiceCore) ([]invoice.InvoiceCore, error) {
	invoices, err := inBusiness.invoiceData.GetAllInvoice(data)

	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (inBusiness *InvoiceBusiness) DeleteInvoice(id int) error {
	if err := inBusiness.invoiceData.DeleteInvoice(id); err != nil {
		return err
	}
	return nil
}

func (inBusiness *InvoiceBusiness) GetInvoiceById(id int) (invoice.InvoiceCore, error) {
	inData, err := inBusiness.invoiceData.GetInvoiceById(id)

	if err != nil {
		return invoice.InvoiceCore{}, err
	}
	return inData, nil
}

func (inBusiness *InvoiceBusiness) GetInvoiceByStatus(status string) ([]invoice.InvoiceCore, error) {
	if helper.IsEmpty(status) || !helper.ValidateStatus(status) {
		return []invoice.InvoiceCore{}, errors.New("bad request")
	}
	invoices, err := inBusiness.invoiceData.GetInvoiceByStatus(status)

	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (inBusiness *InvoiceBusiness) UpdateInvoice(data invoice.InvoiceCore) error {
	if helper.IsEmpty(data.PaymentStatus) {
		return errors.New("invalid data")
	}

	t := time.Now()
	if data.PaymentTerms == 7 {
		data.PaymentDue = t.Add(time.Hour * 24 * 7)
	} else if data.PaymentTerms == 10 {
		data.PaymentDue = t.Add(time.Hour * 24 * 10)
	} else if data.PaymentTerms == 30 {
		data.PaymentDue = t.Add(time.Hour * 24 * 30)
	}

	if data.PaymentStatus == "unpaid" {
		resp, errMidtrans := inBusiness.midtransClient.CreateTransaction(&snap.Request{
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  fmt.Sprintf("%d", data.ID),
				GrossAmt: int64(data.Total),
			},
			Expiry: &snap.ExpiryDetails{
				Unit:     "day",
				Duration: int64(data.PaymentTerms),
			},
			CreditCard: &snap.CreditCardDetails{
				Secure: true,
			},
		})

		if errMidtrans != nil {
			return errMidtrans
		}
		data.PaymentLink = resp.RedirectURL
	}

	err := inBusiness.invoiceData.UpdateInvoice(data)
	if err != nil {
		return err
	}
	inBusiness.SendInvoice(int(data.ID))
	return nil
}

func (inBusiness *InvoiceBusiness) UpdateTransactionStatus(transactionID int64) error {
	trans, err := inBusiness.midtransCoreClient.CheckTransaction(fmt.Sprintf("%d", transactionID))
	if err != nil {
		return err
	}

	if trans.TransactionStatus == "capture" || trans.TransactionStatus == "settlement" {
		PaymentStatus := "paid"
		errUpd := inBusiness.invoiceData.UpdateTransactionStatus(transactionID, PaymentStatus)
		if errUpd != nil {
			return errUpd
		}
		inBusiness.SendInvoice(int(transactionID))
	}

	return nil
}

func (inBusiness *InvoiceBusiness) GetInvoiceByNik(nik int) ([]invoice.InvoiceCore, error) {
	invoices, err := inBusiness.invoiceData.GetInvoiceByNik(nik)

	if err != nil {
		return []invoice.InvoiceCore{}, err
	}
	return invoices, nil
}

func (inBusiness *InvoiceBusiness) GetInvoiceByName(name string) ([]invoice.InvoiceCore, error) {
	if helper.IsEmpty(name) {
		return []invoice.InvoiceCore{}, errors.New("bad request")
	}
	invoices, err := inBusiness.invoiceData.GetInvoiceByName(name)

	if err != nil {
		return []invoice.InvoiceCore{}, err
	}
	return invoices, nil
}

func (inBusiness *InvoiceBusiness) CheckInvoice(data invoice.InvoiceCore) ([]invoice.InvoiceCore, error) {
	invoices, err := inBusiness.invoiceData.GetAllInvoice(data)
	if err != nil {
		return nil, err
	}
	list := []invoice.InvoiceCore{}
	today := time.Now()

	for _, invoice := range invoices {
		resetPaymentStatus := invoice.PaymentDue.Add(time.Duration(-24) * time.Hour)
		// Check if due < today
		if invoice.PaymentStatus == "unpaid" && invoice.PaymentDue.After(today) {
			// Change to .Before!
			fmt.Println("Isi list : ", list)
			list = append(list, invoice)
		} else if invoice.PaymentStatus == "paid" && resetPaymentStatus.After(today) {
			invoice.PaymentStatus = "unpaid"
			inBusiness.UpdateInvoice(invoice)
		}
	}
	for _, send := range list {
		inData, err := inBusiness.invoiceData.GetInvoiceById(int(send.ID))
		if err != nil {
			return []invoice.InvoiceCore{}, nil
		}
		helper.SendGmail(inData)
	}
	return list, nil
}
