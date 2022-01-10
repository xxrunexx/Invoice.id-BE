package handler

import (
	"invoice-api/features/client"
	"invoice-api/features/client/presentation/request"
	"invoice-api/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ClientHandler struct {
	clientBusiness client.Business
}

func NewHandlerClient(clientBusiness client.Business) *ClientHandler {
	return &ClientHandler{clientBusiness}
}

func (clHandler *ClientHandler) CreateClientHandler(e echo.Context) error {
	newClient := request.ReqClient{}

	if err := e.Bind(&newClient); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}

	if err := clHandler.clientBusiness.CreateClient(newClient.ToClientCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, newClient)
}
