package presentation

import (
	"invoice-api/features/paymentmethod"
	"invoice-api/features/paymentmethod/presentation/request"
	"invoice-api/features/paymentmethod/presentation/response"
	"invoice-api/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler struct {
	paymentmethodBusiness paymentmethod.Business
}

func NewHandlerPaymentMethod(paymentmethodBusiness paymentmethod.Business) *PaymentMethodHandler {
	return &PaymentMethodHandler{paymentmethodBusiness}
}

func (pmHandler *PaymentMethodHandler) CreatePaymentMethodHandler(e echo.Context) error {
	newPaymentMethod := request.ReqPaymentMethod{}

	if err := e.Bind(&newPaymentMethod); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	resp, err := pmHandler.paymentmethodBusiness.CreatePaymentMethod(newPaymentMethod.ToPaymentMethodCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, response.ToPaymentMethodResponse(resp))
}

func (pmHandler *PaymentMethodHandler) GetPaymentMethodByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := pmHandler.paymentmethodBusiness.GetPaymentMethodById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToPaymentMethodResponse(data))
}
