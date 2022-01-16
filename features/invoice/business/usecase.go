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
	invoices, err := inBusiness.GetInvoiceByStatus(status)

	if err != nil {
		return nil, err
	}
	return invoices, nil
}
