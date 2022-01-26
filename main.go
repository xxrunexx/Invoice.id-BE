package main

import (
	"fmt"
	"invoice-api/driver"
	"invoice-api/factory"
	_middleware "invoice-api/middleware"
	"invoice-api/routes"
	"os"
	"time"

	"github.com/robfig/cron"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	params := os.Args
	paramsLength := len(params)
	if paramsLength < 2 {
		log.Println("Please add SERVER or CRONJOB along with go run main.go command")
		log.Println("Eg: go run main.go SERVER or go run main.go CRONJOB")
		os.Exit(1)
	}

	inputMethod := os.Args[1]
	valid := IsValidInputMethod(inputMethod)

	driver.InitDB()
	e := routes.New()

	// Log Middleware
	_middleware.LogMiddlewareInit(e)
	if valid {
		if inputMethod == "SERVER" {
			// Starting The Server
			e.Start(":8000")
		}

		if inputMethod == "CRONJOB" {
			presenter := factory.Init()
			// Cron Job
			log.Info("Create new cron")
			c := cron.New()
			c.AddFunc("@every 3s", func() {
				fmt.Println("Every minute")
				presenter.InvoicePresentation.CheckInvoiceHandler()
			})

			log.Info("Start cron")
			c.Start()
			time.Sleep(10080 * time.Hour)
		}
	}
}

func IsValidInputMethod(method string) bool {
	switch method {
	case
		"SERVER",
		"CRONJOB":
		return true
	}
	return false
}
