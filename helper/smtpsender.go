package helper

import (
	"fmt"
	"invoice-api/features/invoice"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func SendGmail(inData invoice.InvoiceCore) {
	err := godotenv.Load("app.env")
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	emailTo := []string{inData.ClientEmail}
	fmt.Println("isi emailto :", emailTo)
	data := struct {
		ReceiverName  string
		SenderName    string
		Total         int
		CreatedAt     time.Time
		PaymentTerms  int
		PaymentDue    time.Time
		PaymentStatus string
	}{
		ReceiverName:  inData.ClientName,
		SenderName:    "Invoicein",
		Total:         inData.Total,
		CreatedAt:     inData.CreatedAt,
		PaymentTerms:  inData.PaymentTerms,
		PaymentDue:    inData.PaymentDue,
		PaymentStatus: inData.PaymentStatus,
	}

	if data.PaymentStatus == "paid" {
		status, err := SendEmailSMTP(emailTo, data, "success_invoice.html", inData)

		if err != nil {
			log.Println(err)
		}
		if status {
			log.Println("Email sent successfully using SMTP")
		}
	} else if data.PaymentStatus == "processed" {
		status, err := SendEmailSMTP(emailTo, data, "new_invoice.html", inData)

		if err != nil {
			log.Println(err)
		}
		if status {
			log.Println("Email sent successfully using SMTP")
		}
	}

}
