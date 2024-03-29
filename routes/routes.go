package routes

import (
	"invoice-api/config"
	"invoice-api/factory"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	presenter := factory.Init()

	// Initiate Echo & JWT
	e := echo.New()
	e.Use(middleware.CORS())
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWTsecret)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// Bill Issuer
	e.POST("/billissuer/register", presenter.BillissuerPresentation.CreateBillIssuerHandler)
	e.POST("/billissuer/login", presenter.BillissuerPresentation.LoginBillIssuerHandler)
	e.GET("/billissuer/:id", presenter.BillissuerPresentation.GetBillIssuerByIdHandler)
	e.GET("/billissuer", presenter.BillissuerPresentation.GetAllBillIssuerHandler)
	e.PUT("/billissuer", presenter.BillissuerPresentation.UpdateBillIssuerHandler)

	// Bill Issuer Detail
	e.POST("/billissuerdetail/add", presenter.BillissuerdetailPresentation.CreateBillIssuerDetailHandler)
	e.GET("/billissuerdetail/:id", presenter.BillissuerdetailPresentation.GetBillIssuerDetailByIdHandler)
	e.PUT("/billissuerdetail", presenter.BillissuerdetailPresentation.UpdateBillIssuerDetailHandler)

	// Client
	e.POST("/client/add", presenter.ClientPresentation.CreateClientHandler)
	e.GET("/client", presenter.ClientPresentation.GetAllClientHandler)
	e.GET("/client/:id", presenter.ClientPresentation.GetClientById)
	e.PUT("/client", presenter.ClientPresentation.UpdateClient)

	// Invoice
	e.POST("/invoice/add", presenter.InvoicePresentation.CreateInvoiceHandler)
	e.GET("/invoice", presenter.InvoicePresentation.GetAllInvoiceHandler)
	e.GET("/invoice/:id", presenter.InvoicePresentation.GetInvoiceByIdHandler)
	e.GET("/invoice/status/:status", presenter.InvoicePresentation.GetInvoiceByStatusHandler)
	e.GET("/invoice/nik/:nik", presenter.InvoicePresentation.GetInvoiceByNikHandler)
	e.GET("/invoice/name/:name", presenter.InvoicePresentation.GetInvoiceByNameHandler)
	e.POST("/transactions/callback", presenter.InvoicePresentation.CallbackHandler)
	e.DELETE("/invoice/:id", presenter.InvoicePresentation.DeleteInvoiceHandler)
	e.PUT("/invoice/update", presenter.InvoicePresentation.UpdateInvoiceHandler)
	e.POST("/invoice/send/:id", presenter.InvoicePresentation.SendInvoiceHandler)
	return e
}
