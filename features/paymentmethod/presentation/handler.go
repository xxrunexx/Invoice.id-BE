package presentation

import (
	"invoice-api/features/paymentmethod"
	"invoice-api/features/paymentmethod/presentation/request"
	"invoice-api/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler struct {
	paymentmethodBusiness paymentmethod.Business
}

func NewHandlerPaymentMethod(paymentmethodBusiness paymentmethod.Business) *PaymentMethodHandler {
	return &PaymentMethodHandler{paymentmethodBusiness}
}

func (pmHandler *PaymentMethodHandler) CreatePaymentMethod(e echo.Context) error {
	newPaymentMethod := request.ReqPaymentMethod{}

	if err := e.Bind(&newPaymentMethod); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := pmHandler.paymentmethodBusiness.CreatePaymentMethod(newPaymentMethod.ToPaymentMethodCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, newPaymentMethod)
}
