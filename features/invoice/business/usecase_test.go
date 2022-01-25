package business

import (
	"errors"
	"invoice-api/features/invoice"
	"invoice-api/features/invoice/mocks"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData        mocks.Data
	invoiceBusiness invoice.Business
	invoiceDatas    []invoice.InvoiceCore
	invoiceData     invoice.InvoiceCore
	invoiceData2    invoice.InvoiceCore
	invoiceData3    invoice.InvoiceCore
)

func TestMain(m *testing.M) {
	invoiceBusiness = NewBusinessInvoice(&mockData)

	invoiceDatas = []invoice.InvoiceCore{
		{
			ID:              1,
			ClientID:        1,
			Total:           1000000,
			BillIssuerID:    1,
			PaymentMethodID: 1,
			PaymentDue:      time.Now().AddDate(0, 0, 14),
			PaymentStatus:   "unpaid",
			PaymentTerms:    14,
		},
	}

	invoiceData = invoice.InvoiceCore{
		ClientID:        1,
		Total:           1000000,
		BillIssuerID:    1,
		PaymentMethodID: 1,
		PaymentStatus:   "unpaid",
		PaymentTerms:    7,
	}

	invoiceData2 = invoice.InvoiceCore{
		ClientID:        1,
		Total:           1000000,
		BillIssuerID:    1,
		PaymentMethodID: 1,
		PaymentStatus:   "unpaid",
		PaymentTerms:    10,
	}
	invoiceData3 = invoice.InvoiceCore{
		ClientID:        1,
		Total:           1000000,
		BillIssuerID:    1,
		PaymentMethodID: 1,
		PaymentStatus:   "unpaid",
		PaymentTerms:    30,
	}
	os.Exit(m.Run())
}

func TestCreateInvoice(t *testing.T) {
	t.Run("Create invoice - success add payment terms", func(t *testing.T) {
		mockData.On("CreateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(nil).Once()
		err := invoiceBusiness.CreateInvoice(invoiceData)
		assert.Nil(t, err)
	})

	t.Run("Create invoice - success add different payment terms", func(t *testing.T) {
		mockData.On("CreateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(nil).Once()
		err := invoiceBusiness.CreateInvoice(invoiceData2)
		assert.Nil(t, err)
	})
	t.Run("Create invoice - success add different payment terms", func(t *testing.T) {
		mockData.On("CreateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(nil).Once()
		err := invoiceBusiness.CreateInvoice(invoiceData3)
		assert.Nil(t, err)
	})

	t.Run("create invoice - error insert data", func(t *testing.T) {
		mockData.On("CreateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(errors.New("Error")).Once()
		err := invoiceBusiness.CreateInvoice(invoice.InvoiceCore{})
		assert.NotNil(t, err)
	})
}

func TestSendInvoice(t *testing.T) {
	t.Run("Send invoice - success", func(t *testing.T) {
		mockData.On("GetInvoiceById", mock.AnythingOfType("int")).Return(invoice.InvoiceCore{}, nil).Once()
		resp, err := invoiceBusiness.GetInvoiceById(int(invoiceData.ID))
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Send Invoice  - error", func(t *testing.T) {
		mockData.On("GetInvoiceById", mock.AnythingOfType("int")).Return(invoice.InvoiceCore{}, errors.New("error")).Once()
		resp, err := invoiceBusiness.GetInvoiceById(int(invoiceData.ID))
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}
func TestGetInvoiceById(t *testing.T) {
	t.Run("Get invoice by id- success", func(t *testing.T) {
		mockData.On("GetInvoiceById", mock.AnythingOfType("int")).Return(invoice.InvoiceCore{}, nil).Once()
		resp, err := invoiceBusiness.GetInvoiceById(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("error get invoice by id", func(t *testing.T) {
		mockData.On("GetInvoiceById", mock.AnythingOfType("int")).Return(invoice.InvoiceCore{}, errors.New("error")).Once()
		resp, err := invoiceBusiness.GetInvoiceById(3)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestGetInvoiceByStatus(t *testing.T) {
	t.Run("Get invoice by id - success", func(t *testing.T) {
		mockData.On("GetInvoiceByStatus", mock.AnythingOfType("string")).Return([]invoice.InvoiceCore{}, nil).Once()
		resp, err := invoiceBusiness.GetInvoiceByStatus("draft")
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Create bill issuer error - invalid status", func(t *testing.T) {
		_, err := invoiceBusiness.GetInvoiceByStatus("")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "bad request")
	})

	t.Run("error get invoice by status", func(t *testing.T) {
		mockData.On("GetInvoiceByStatus", mock.AnythingOfType("string")).Return([]invoice.InvoiceCore{}, errors.New("error")).Once()
		_, err := invoiceBusiness.GetInvoiceByStatus("draft")
		assert.NotNil(t, err)
	})
}

func TestGetAllInvoice(t *testing.T) {
	t.Run("validate get invoices", func(t *testing.T) {
		mockData.On("GetAllInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(invoiceDatas, nil).Once()
		resp, err := invoiceBusiness.GetAllInvoice(invoice.InvoiceCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get invoices", func(t *testing.T) {
		mockData.On("GetAllInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(nil, errors.New("error"))
		resp, err := invoiceBusiness.GetAllInvoice(invoice.InvoiceCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func TestUpdateInvoice(t *testing.T) {
	t.Run("Update invoice - error insert data", func(t *testing.T) {
		mockData.On("UpdateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(errors.New("error")).Once()
		err := invoiceBusiness.UpdateInvoice(invoiceData)
		assert.NotNil(t, err)
	})

	t.Run("Update invoice error - invalid payment status", func(t *testing.T) {
		err := invoiceBusiness.UpdateInvoice(invoice.InvoiceCore{
			PaymentStatus: "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid data")
	})
}
func TestDeleteInvoice(t *testing.T) {
	t.Run("Delete invoice - error insert data", func(t *testing.T) {
		mockData.On("DeleteInvoice", mock.AnythingOfType("int")).Return(errors.New("error")).Once()
		err := invoiceBusiness.DeleteInvoice(int(invoiceData.ID))
		assert.NotNil(t, err)
	})
	t.Run("Delete invoice - error insert data", func(t *testing.T) {
		mockData.On("DeleteInvoice", mock.AnythingOfType("int")).Return(nil).Once()
		err := invoiceBusiness.DeleteInvoice(int(invoiceData.ID))
		assert.Nil(t, err)
	})
}

func TestGetInvoiceByName(t *testing.T) {
	t.Run("error create invoice", func(t *testing.T) {
		mockData.On("GetInvoiceByName", mock.AnythingOfType("string")).Return(invoiceDatas, errors.New("bad request")).Once()
		resp, err := invoiceBusiness.GetInvoiceByName("")
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "bad request")
		assert.NotNil(t, resp)
	})

	t.Run("Get invoice by name - success", func(t *testing.T) {
		mockData.On("GetInvoiceByName", mock.AnythingOfType("string")).Return(invoiceDatas, nil).Once()
		resp, err := invoiceBusiness.GetInvoiceByName("Harun")
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error get invoice by name", func(t *testing.T) {
		mockData.On("GetInvoiceByName", mock.AnythingOfType("string")).Return(invoice.InvoiceCore{}, errors.New("error")).Once()
		resp, err := invoiceBusiness.GetInvoiceByName("Harun")
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestGetInvoiceByNik(t *testing.T) {
	t.Run("Get invoice by name - success", func(t *testing.T) {
		mockData.On("GetInvoiceByNik", mock.AnythingOfType("int")).Return(invoiceDatas, nil).Once()
		resp, err := invoiceBusiness.GetInvoiceByNik(3174051805000009)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("error get invoice by name", func(t *testing.T) {
		mockData.On("GetInvoiceByNik", mock.AnythingOfType("int")).Return([]invoice.InvoiceCore{}, errors.New("error")).Once()
		resp, err := invoiceBusiness.GetInvoiceByNik(3174051805000009)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, err.Error(), "error")
	})
}
