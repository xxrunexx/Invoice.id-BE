package business

import (
	"invoice-api/features/paymentmethod"
	"invoice-api/features/paymentmethod/mocks"
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
}

func TestCreatePaymentMethod(t *testing.T) {
	t.Run("validate create payment method", func(t *testing.T) {
		mockData.On("CreatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(nil).Once()
		err := paymentmethodBusiness.CreatePaymentMethod(paymentmethod.PaymentMethodCore{})
		assert.NotNil(t, err)
	})

	// t.Run("error create payment method", func(t *testing.T) {
	// 	mockData.On("CreatePaymentMethod", mock.AnythingOfType("paymentmethod.PaymentMethodCore")).Return(errors.New("error")).Once()
	// 	err := paymentmethodBusiness.CreatePaymentMethod(paymentmethod.PaymentMethodCore{})
	// 	assert.NotNil(t, err)
	// })
	t.Run("Create payment method invalid name", func(t *testing.T) {
		err := paymentmethodBusiness.CreatePaymentMethod(paymentmethod.PaymentMethodCore{
			Name: "",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "bad request")
	})
}
