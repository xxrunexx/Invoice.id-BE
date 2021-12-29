package main

func main() {
	driver.InitDB()
	e := routes.New()

	// Log Middleware
	_middleware.LogMiddlewareInit(e)

	// Starting The Server
	e.Start(":8000")
}
