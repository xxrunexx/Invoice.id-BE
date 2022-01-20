package presentation

import (
	"invoice-api/features/billissuerdetail"
	"invoice-api/features/billissuerdetail/presentation/request"
	"invoice-api/features/billissuerdetail/presentation/response"
	"invoice-api/helper"
	"net/http"
	"strconv"

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

	resp, err := bidHandler.billissuerdetailbusiness.CreateBillIssuerDetail(newBillIssuerDetail.ToBillIssuerDetailCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}

	return helper.SuccessResponse(e, response.ToBillIssuerDetailResponse(resp))
}

func (bidHandler *BillIssuerDetailHandler) GetBillIssuerDetailById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	data, err := bidHandler.billissuerdetailbusiness.GetBillIssuerDetailById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToBillIssuerDetailResponse(data))
}
