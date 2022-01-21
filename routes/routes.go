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

	// Client
	e.POST("/client/add", presenter.ClientPresentation.CreateClientHandler)
	e.GET("/client", presenter.ClientPresentation.GetAllClientHandler)
	e.GET("/client/:id", presenter.ClientPresentation.GetClientById)
	e.PUT("/client", presenter.ClientPresentation.UpdateClient)

	// Invoice
	e.POST("/invoice/add", presenter.InvoicePresentation.CreateInvoiceHandler)
	e.GET("/invoice", presenter.InvoicePresentation.GetAllInvoiceHandler)
	e.GET("/invoice/:id", presenter.InvoicePresentation.GetInvoiceByIdHandler)
	e.GET("/invoice/status/:status", presenter.InvoicePresentation.GetInvoiceByStatus)
	e.DELETE("/invoice/:id", presenter.InvoicePresentation.DeleteInvoiceHandler)
	e.PUT("/invoice/update", presenter.InvoicePresentation.UpdateInvoiceHandler)
	e.POST("/invoice/send/:id", presenter.InvoicePresentation.SendInvoiceHandler)

	// Bill Issuer Detail
	e.POST("/billissuerdetail/add", presenter.BillissuerdetailPresentation.CreateBillIssuerDetailHandler)
	e.GET("/billissuerdetail/:id", presenter.BillissuerdetailPresentation.GetBillIssuerDetailById)

	// Payment Method
	e.POST("/paymentmethod/add", presenter.PaymentmethodPresentation.CreatePaymentMethodHandler)
	e.GET("/paymentmethod/:id", presenter.PaymentmethodPresentation.GetPaymentMethodByIdHandler)
	e.PUT("/paymentmethod/update", presenter.PaymentmethodPresentation.UpdatePaymentMethodHandler)

	return e
}
