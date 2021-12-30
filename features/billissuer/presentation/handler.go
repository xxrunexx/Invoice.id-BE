package presentation

import (
	"invoice-api/features/billissuer"
	"invoice-api/features/billissuer/presentation/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BillIssuerHandler struct {
	billissuerBusiness billissuer.Business
}

func NewHandlerBillIssuer(billissuerBusiness billissuer.Business) *BillIssuerHandler {
	return &BillIssuerHandler{billissuerBusiness}
}

func (biHandler *BillIssuerHandler) CreateBillIssuerHandler(e echo.Context) (err error) {
	newBillIssuer := request.ReqBillIssuer{}

	if err := e.Bind(&newBillIssuer); err != nil {
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	if err := biHandler.billissuerBusiness.CreateAccount(newBillIssuer.ToBillIssuerCore()); err != nil {
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"data":    newBillIssuer,
	})
}
