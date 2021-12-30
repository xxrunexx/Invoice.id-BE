package main

import (
	"invoice-api/driver"
	_middleware "invoice-api/middleware"
	"invoice-api/routes"
)

func main() {
	driver.InitDB()
	e := routes.New()

	// Log Middleware
	_middleware.LogMiddlewareInit(e)
	// _middleware.LogMiddlewareInit(e)

	// Starting The Server
	e.Start(":8000")
}
