package presentation

import (
	"invoice-api/features/billissuerdetail"
	"invoice-api/features/billissuerdetail/presentation/request"
	"invoice-api/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BillIssuerDetailHandler struct {
	billissuerdetailbusiness billissuerdetail.Business
}

func NewHandlerBillIssuerDetail(billissuerdetailbusiness billissuerdetail.Business) *BillIssuerDetailHandler {
	return &BillIssuerDetailHandler{billissuerdetailbusiness}
}

func (bidHandler *BillIssuerDetailHandler) CreateBillIssuerDetailHandler(e echo.Context) error {
	newBillIssuerDetail := request.ReqBillIssuerDetail{}

	if err := e.Bind(&newBillIssuerDetail); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := bidHandler.billissuerdetailbusiness.CreateBillIssuerDetail(newBillIssuerDetail.ToBillIssuerDetailCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, newBillIssuerDetail)
}
