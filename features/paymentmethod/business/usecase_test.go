package business

import (
	"errors"
	"invoice-api/features/paymentmethod"
	"invoice-api/features/paymentmethod/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData              mocks.Data
	paymentmethodBusiness paymentmethod.Business
	paymentmethodData     paymentmethod.PaymentMethodCore
)

func TestMain(m *testing.M) {
	paymentmethodBusiness = NewBusinessPaymentMethod(&mockData)

	paymentmethodData = paymentmethod.PaymentMethodCore{
		Name:     "BRI",
		IsActive: true,
	}
	os.Exit(m.Run())
}

func TestCreatePaymentMethod(t *testing.T) {
	t.Run("validate create payment method", func(t *testing.T) {
		mockData.On("CreatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(paymentmethod.PaymentMethodCore{}, nil).Once()
		resp, err := paymentmethodBusiness.CreatePaymentMethod(paymentmethod.PaymentMethodCore{})
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("create payment method - success", func(t *testing.T) {
		mockData.On("CreatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(paymentmethod.PaymentMethodCore{}, nil).Once()
		resp, err := paymentmethodBusiness.CreatePaymentMethod(paymentmethodData)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	// t.Run("error create payment method", func(t *testing.T) {
	// 	mockData.On("CreatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(errors.New("error")).Once()
	// 	err := paymentmethodBusiness.CreatePaymentMethod(paymentmethod.PaymentMethodCore{})
	// 	assert.NotNil(t, err)
	// })
	t.Run("Create payment method invalid name", func(t *testing.T) {
		resp, err := paymentmethodBusiness.CreatePaymentMethod(paymentmethod.PaymentMethodCore{
			Name: "",
		})
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, err.Error(), "bad request")
	})
}

func TestGetPaymentMethodById(t *testing.T) {
	t.Run("validate get payment method by id", func(t *testing.T) {
		mockData.On("GetPaymentMethodById", mock.AnythingOfType("int")).Return(paymentmethod.PaymentMethodCore{}, nil).Once()
		resp, err := paymentmethodBusiness.GetPaymentMethodById(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("error get payment method by id", func(t *testing.T) {
		mockData.On("GetPaymentMethodById", mock.AnythingOfType("int")).Return(paymentmethod.PaymentMethodCore{}, errors.New("error")).Once()
		resp, err := paymentmethodBusiness.GetPaymentMethodById(3)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestUpdatePaymentMethod(t *testing.T) {
	t.Run("validate update payment method", func(t *testing.T) {
		mockData.On("UpdatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(nil).Once()
		err := paymentmethodBusiness.UpdatePaymentMethod(paymentmethod.PaymentMethodCore{})
		assert.NotNil(t, err)
	})

	t.Run("Update payment method - error insert data", func(t *testing.T) {
		mockData.On("UpdatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(errors.New("error")).Once()
		err := paymentmethodBusiness.UpdatePaymentMethod(paymentmethodData)
		assert.Nil(t, err)
	})
}
