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

	resp, err := clHandler.clientBusiness.CreateClient(newClient.ToClientCore())
	if err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, response.ToClientResponse(resp))
}

func (clHandler *ClientHandler) GetAllClientHandler(e echo.Context) error {
	data, err := clHandler.clientBusiness.GetAllClient(client.ClientCore{})

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

func (clHandler *ClientHandler) UpdateClient(e echo.Context) error {
	updateData := request.ReqClientUpdate{}

	if err := e.Bind(&updateData); err != nil {
		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
	}
	if err := clHandler.clientBusiness.UpdateClient(updateData.ToClientCore()); err != nil {
		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
	}
	return helper.SuccessResponse(e, updateData)
}
