package presentation

import (
	"invoice-api/features/invoice"
	"invoice-api/features/invoice/presentation/request"
	"invoice-api/helper"
	"net/http"

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

	if err := e.Bind(&newInvoice); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := inHandler.invoiceBusiness.CreateInvoice(newInvoice.ToInvoiceCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, newInvoice)
}
