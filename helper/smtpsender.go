package helper

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func SendGmail(gmail string) {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	emailTo := []string{gmail}
	fmt.Println("isi emailto :", emailTo)
	data := struct {
		ReceiverName string
		SenderName   string
	}{
		ReceiverName: "David Gilmour",
		SenderName:   "Invoicein",
	}

	status, err := SendEmailSMTP(emailTo, data, "sample_template.html")
	if err != nil {
		log.Println(err)
	}
	if status {
		log.Println("Email sent successfully using SMTP")
	}
}
