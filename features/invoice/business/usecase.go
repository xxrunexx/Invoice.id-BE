package business

import (
	"errors"
	"fmt"
	"invoice-api/features/invoice"
	"invoice-api/helper"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type InvoiceBusiness struct {
	invoiceData    invoice.Data
	midtransClient snap.Client
}

func NewBusinessInvoice(inData invoice.Data, midtransClient snap.Client) invoice.Business {
	return &InvoiceBusiness{inData, midtransClient}
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
	fmt.Println("Isi payment due : ", data.PaymentDue)
	id, err := inBusiness.invoiceData.CreateInvoice(data)
	if err != nil {
		return err
	}
	fmt.Println("idnya : ", id)
	resp, errMidtrans := inBusiness.midtransClient.CreateTransaction(&snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("%d", id),
			GrossAmt: int64(data.Total),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
	})
	if errMidtrans != nil {
		fmt.Println("Isi error : ", errMidtrans)
		return err
	}
	fmt.Println("Isi response : ", resp)
	data.PaymentLink = resp.RedirectURL
	data.ID = id
	err1 := inBusiness.invoiceData.UpdateInvoice(data)
	if err1 != nil {
		return err1
	}
	fmt.Println("Setelah update : ")
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
	err := inBusiness.invoiceData.UpdateInvoice(data)
	if err != nil {
		return err
	}
	// if data.PaymentStatus == "unpaid" || data.PaymentStatus == "draft" || data.PaymentStatus == "processed" {
	inBusiness.SendInvoice(int(data.ID))
	// }
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

// func (inBusiness *InvoiceBusiness) CheckCSV(datas []invoice.InvoiceCore) error {
// 	if err := inBusiness.invoiceData.InsertCSV(datas); err != nil {
// 		return err
// 	}
// 	return nil
// }

func (inBusiness *InvoiceBusiness) CheckInvoice(data invoice.InvoiceCore) ([]invoice.InvoiceCore, error) {
	invoices, err := inBusiness.invoiceData.GetAllInvoice(data)
	if err != nil {
		return nil, err
	}
	list := []invoice.InvoiceCore{}
	today := time.Now()
	// due := data.PaymentDue

	for _, invoice := range invoices {
		// Check if due < today
		if invoice.PaymentStatus == "processed" && invoice.PaymentDue.After(today) {
			// Change to .Before!
			fmt.Println("Isi list : ", list)
			list = append(list, invoice)
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
