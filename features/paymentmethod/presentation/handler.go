package presentation

import (
	"fmt"
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

func (pmHandler *PaymentMethodHandler) GetAllPaymentMethodHandler(e echo.Context) error {
	data, err := pmHandler.paymentmethodBusiness.GetAllPaymentMethod(paymentmethod.PaymentMethodCore{})

	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToPaymentMethodResponseList(data))
}

func (pmHandler *PaymentMethodHandler) UpdatePaymentMethodHandler(e echo.Context) error {
	updateData := request.ReqPaymentMethodUpdate{}

	if err := e.Bind(&updateData); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := pmHandler.paymentmethodBusiness.UpdatePaymentMethod(updateData.ToPaymentMethodCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, updateData)
}

func (pmHandler *PaymentMethodHandler) GetPaymentMethodByIsActiveHandler(e echo.Context) error {
	isActive, err := strconv.ParseBool(e.Param("status"))
	fmt.Println("Isi Is Active : ", isActive)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := pmHandler.paymentmethodBusiness.GetPaymentMethodByIsActive(isActive)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToPaymentMethodResponseList(data))
}
