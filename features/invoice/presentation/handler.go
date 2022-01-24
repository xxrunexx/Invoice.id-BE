package presentation

import (
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

func (inHandler *InvoiceHandler) SendInvoiceHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	_, err = inHandler.invoiceBusiness.GetInvoiceById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	_, err = inHandler.invoiceBusiness.SendInvoice(id)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "failed send data",
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "successfully sending email",
	})
}

func (inHandler *InvoiceHandler) GetInvoiceByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := inHandler.invoiceBusiness.GetInvoiceById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, response.ToInvoiceResponse(data))
}

func (inHandler *InvoiceHandler) DeleteInvoiceHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}
	err = inHandler.invoiceBusiness.DeleteInvoice(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, map[string]interface{}{
		"message": "data successfully deleted",
	})
}

func (inHandler *InvoiceHandler) GetInvoiceByStatusHandler(e echo.Context) error {
	status := e.Param("status")

	data, err := inHandler.invoiceBusiness.GetInvoiceByStatus(status)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToInvoiceResponseList(data))
}

func (inHandler *InvoiceHandler) UpdateInvoiceHandler(e echo.Context) error {
	updateData := request.ReqInvoiceUpdate{}

	if err := e.Bind(&updateData); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := inHandler.invoiceBusiness.UpdateInvoice(updateData.ToInvoiceCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, updateData)
}

func (inHandler *InvoiceHandler) GetInvoiceByNikHandler(e echo.Context) error {
	nik, err := strconv.Atoi(e.Param("nik"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}
	data, err := inHandler.invoiceBusiness.GetInvoiceByNik(nik)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToInvoiceResponseList(data))
}
