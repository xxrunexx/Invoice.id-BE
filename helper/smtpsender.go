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
		ReceiverName string
		SenderName   string
		Total        int
		CreatedAt    time.Time
		PaymentTerms int
		PaymentDue   time.Time
	}{
		ReceiverName: inData.ClientName,
		SenderName:   "Invoicein",
		Total:        inData.Total,
		CreatedAt:    inData.CreatedAt,
		PaymentTerms: inData.PaymentTerms,
		PaymentDue:   inData.PaymentDue,
	}

	status, err := SendEmailSMTP(emailTo, data, "sample_template.html", inData)
	if err != nil {
		log.Println(err)
	}
	if status {
		log.Println("Email sent successfully using SMTP")
	}
}
