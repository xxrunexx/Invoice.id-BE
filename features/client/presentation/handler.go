package handler

import (
	"invoice-api/features/client"
	"invoice-api/features/client/presentation/request"
	"invoice-api/features/client/presentation/response"
	"invoice-api/helper"
	"net/http"
	"strconv"

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

func (clHandler *ClientHandler) GetAllClientHandler(e echo.Context) error {
	data, err := clHandler.clientBusiness.GetAllCient(client.ClientCore{})

	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToClientResponseList(data))
}

func (clHandler *ClientHandler) GetClientById(e echo.Context) error {
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}
	data, err := clHandler.clientBusiness.GetClientById(id)
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToClientResponse(data))
}
