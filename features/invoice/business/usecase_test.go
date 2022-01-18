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
		PaymentDue:      time.Now().AddDate(0, 0, 14),
		PaymentStatus:   "unpaid",
		PaymentTerms:    14,
	}
	os.Exit(m.Run())
}

func TestCreateInvoice(t *testing.T) {
	t.Run("Create invoice - success", func(t *testing.T) {
		mockData.On("CreateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(nil).Once()
		err := invoiceBusiness.CreateInvoice(invoiceData)
		assert.Nil(t, err)
	})

	t.Run("create invoice - error insert data", func(t *testing.T) {
		mockData.On("CreateInvoice", mock.AnythingOfType("invoice.InvoiceCore")).Return(errors.New("Error")).Once()
		err := invoiceBusiness.CreateInvoice(invoice.InvoiceCore{})
		assert.NotNil(t, err)
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

}
