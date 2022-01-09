package presentation

import (
	"invoice-api/features/billissuer"
	"invoice-api/features/billissuer/presentation/request"
	"invoice-api/features/billissuer/presentation/response"
	"invoice-api/helper"
	"net/http"

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

	return helper.SuccessResponse(e, nil)
}

func (biHandler *BillIssuerHandler) LoginBillIssuerHandler(e echo.Context) error {
	billissuerAuth := request.ReqBIllIssuerAuth{}
	e.Bind(&billissuerAuth)

	data, err := biHandler.billissuerBusiness.LoginBillIssuer(billissuerAuth.ToAccountCore())
	if err != nil {
		return e.JSON(http.StatusForbidden, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "successful operator",
		"data":    response.ToBillIssuerLoginResponse(data),
	})
}
