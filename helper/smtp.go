package helper

import (
	"errors"
	"fmt"
	"invoice-api/features/invoice"
	"net/smtp"
	"os"
)

var emailAuth smtp.Auth

func SendEmailSMTP(to []string, data interface{}, template string, inData invoice.InvoiceCore) (bool, error) {
	var subject_type string
	if inData.PaymentStatus == "paid" {
		subject_type = "Payment Success"
	} else {
		subject_type = "Payment Reminder"
	}

	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPort := os.Getenv("EMAIL_PORT")

	emailAuth = smtp.PlainAuth("", emailFrom, emailPassword, emailHost)

	emailBody, err := parseTemplate(template, data, inData)
	// fmt.Println("isi template", template)
	// fmt.Println("isi data", data)
	// fmt.Println("isi intData", inData)
	if err != nil {
		return false, errors.New("unable to parse email template")
	}

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + subject_type + "\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", emailHost, emailPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}
