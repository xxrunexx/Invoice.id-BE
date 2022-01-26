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
	jwt.POST("/client/add", presenter.ClientPresentation.CreateClientHandler)
	jwt.GET("/client", presenter.ClientPresentation.GetAllClientHandler)
	jwt.GET("/client/:id", presenter.ClientPresentation.GetClientById)
	jwt.PUT("/client", presenter.ClientPresentation.UpdateClient)

	// Invoice
	jwt.POST("/invoice/add", presenter.InvoicePresentation.CreateInvoiceHandler)
	jwt.GET("/invoice", presenter.InvoicePresentation.GetAllInvoiceHandler)
	jwt.GET("/invoice/:id", presenter.InvoicePresentation.GetInvoiceByIdHandler)
	jwt.GET("/invoice/status/:status", presenter.InvoicePresentation.GetInvoiceByStatusHandler)
	jwt.GET("/invoice/nik/:nik", presenter.InvoicePresentation.GetInvoiceByNikHandler)
	jwt.GET("/invoice/name/:name", presenter.InvoicePresentation.GetInvoiceByNameHandler)
	jwt.DELETE("/invoice/:id", presenter.InvoicePresentation.DeleteInvoiceHandler)
	jwt.PUT("/invoice/update", presenter.InvoicePresentation.UpdateInvoiceHandler)
	jwt.POST("/invoice/send/:id", presenter.InvoicePresentation.SendInvoiceHandler)

	// Payment Method
	jwt.POST("/paymentmethod/add", presenter.PaymentmethodPresentation.CreatePaymentMethodHandler)
	jwt.GET("/paymentmethod/:id", presenter.PaymentmethodPresentation.GetPaymentMethodByIdHandler)
	jwt.GET("/paymentmethod", presenter.PaymentmethodPresentation.GetAllPaymentMethodHandler)
	jwt.GET("/paymentmethod/status/:status", presenter.PaymentmethodPresentation.GetPaymentMethodByIsActiveHandler)
	jwt.PUT("/paymentmethod/update", presenter.PaymentmethodPresentation.UpdatePaymentMethodHandler)
	return e
}
