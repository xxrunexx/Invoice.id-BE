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
	jwt := e.Group("")
	jwt.Use(middleware.JWT([]byte(config.JWT_KEY)))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// Bill Issuer
	e.POST("/billissuer/register", presenter.BillissuerPresentation.CreateBillIssuerHandler)
	return e
}
