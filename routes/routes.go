package routes

import (
	"invoice-api/config"
	"invoice-api/factory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	presenter := factory.Init()

	// Initiate Echo & JWT
	e := echo.New()
	e.Use(middleware.CORS())
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// Bill Issuer
	e.POST("/billissuer/register", presenter.BillissuerPresentation.CreateBillIssuerHandler)
	e.POST("/billissuer/login", presenter.BillissuerPresentation.LoginBillIssuerHandler)
	e.GET("/billissuer/:id", presenter.BillissuerPresentation.GetBillIssuerByIdHandler)
	e.PUT("/billissuer", presenter.BillissuerPresentation.UpdateBillIssuerHandler)

	// Client
	e.POST("/client/add", presenter.ClientPresentation.CreateClientHandler)
	e.GET("/client", presenter.ClientPresentation.GetAllClientHandler)
	e.GET("/client/:id", presenter.ClientPresentation.GetClientById)

	// Invoice
	e.POST("/invoice/add", presenter.InvoicePresentation.CreateInvoiceHandler)
	e.GET("/invoice", presenter.InvoicePresentation.GetAllInvoiceHandler)
	e.GET("/invoice/:id", presenter.InvoicePresentation.GetInvoiceByIdHandler)

	return e
}
