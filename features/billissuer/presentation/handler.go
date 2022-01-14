package presentation

import (
	"invoice-api/features/billissuer"
	"invoice-api/features/billissuer/presentation/request"
	"invoice-api/features/billissuer/presentation/response"
	"invoice-api/helper"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BillIssuerHandler struct {
	billissuerBusiness billissuer.Business
}

func NewHandlerBillIssuer(billissuerBusiness billissuer.Business) *BillIssuerHandler {
	return &BillIssuerHandler{billissuerBusiness}
}

func (biHandler *BillIssuerHandler) CreateBillIssuerHandler(e echo.Context) error {
	newBillIssuer := request.ReqBillIssuer{}

	if err := e.Bind(&newBillIssuer); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := biHandler.billissuerBusiness.CreateBillIssuer(newBillIssuer.ToBillIssuerCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, newBillIssuer)
}

func (biHandler *BillIssuerHandler) LoginBillIssuerHandler(e echo.Context) error {
	billissuerAuth := request.ReqBIllIssuerAuth{}

	if err := e.Bind(&billissuerAuth); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := biHandler.billissuerBusiness.LoginBillIssuer(billissuerAuth.ToBillIssuerCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusForbidden, "Mismatch Data", err)
	}

	return helper.SuccessResponse(e, response.ToBillIssuerLoginResponse(data))
}

func (biHandler BillIssuerHandler) GetBillIssuerByIdHandler(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := biHandler.billissuerBusiness.GetBillIssuerById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "successful operator",
		"data":    response.ToBillIssuerResponse(data),
	})
}

func (biHandler BillIssuerHandler) UpdateBillIssuerHandler(e echo.Context) error {
	updateData := request.ReqBillIssuerUpdate{}

	if err := e.Bind(&updateData); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	// claims := middleware.ExtractClaim(e)
	if err := biHandler.billissuerBusiness.UpdateBillIssuer(updateData.ToBillIssuerCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, updateData)
}
