package presentation

import (
	"fmt"
	"invoice-api/features/invoice"
	"invoice-api/features/invoice/presentation/request"
	"invoice-api/features/invoice/presentation/response"
	"invoice-api/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InvoiceHandler struct {
	invoiceBusiness invoice.Business
}

func NewHandlerInvoice(invoiceBusiness invoice.Business) *InvoiceHandler {
	return &InvoiceHandler{invoiceBusiness}
}

func (inHandler *InvoiceHandler) CreateInvoiceHandler(e echo.Context) error {
	newInvoice := request.ReqInvoice{}
	fmt.Println("isi new invoice", newInvoice)

	if err := e.Bind(&newInvoice); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := inHandler.invoiceBusiness.CreateInvoice(newInvoice.ToInvoiceCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, newInvoice)
}

func (inHandler *InvoiceHandler) GetAllInvoiceHandler(e echo.Context) error {
	data, err := inHandler.invoiceBusiness.GetAllInvoice(invoice.InvoiceCore{})

	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToInvoiceResponseList(data))
}

func (inHandler *InvoiceHandler) DeleteInvoiceHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.QueryParam("id"))
	// fmt.Println("Isi id : ", id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}
	err = inHandler.invoiceBusiness.DeleteInvoice(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, map[string]interface{}{
		"message": "data deleted",
	})
}
