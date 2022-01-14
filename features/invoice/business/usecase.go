package business

import "invoice-api/features/invoice"

type InvoiceBusiness struct {
	invoiceData invoice.Data
}

func NewBusinessInvoice(inData invoice.Data) invoice.Business {
	return &InvoiceBusiness{inData}
}

func (inBusiness *InvoiceBusiness) CreateInvoice(data invoice.InvoiceCore) error {
	if err := inBusiness.invoiceData.CreateInvoice(data); err != nil {
		return err
	}
	return nil
}

func (inBusiness *InvoiceBusiness) GetAllInvoice(data invoice.InvoiceCore) ([]invoice.InvoiceCore, error) {
	invoices, err := inBusiness.invoiceData.GetAllInvoice(data)

	if err != nil {
		return nil, err
	}
	return invoices, nil
}
