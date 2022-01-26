package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"invoice-api/features/invoice"
	"strconv"
	"strings"
)

func parseTemplate(templateFileName string, data interface{}, inData invoice.InvoiceCore) (string, error) {
	templatePath := fmt.Sprintf("helper/email_templates/%s", templateFileName)

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	r := strings.NewReplacer("#ClientName#", inData.ClientName,
		"#ClientAddress#", inData.ClientAddress,
		"#CreatedAt#", inData.CreatedAt.String(),
		"#Total#", strconv.Itoa(inData.Total),
		"#PaymentDue#", inData.PaymentDue.Format("2006-01-02 15:04:05"),
		"#ClientEmail#", inData.ClientEmail,
		"#ClientPhone#", inData.ClientPhone,
		"#InvoiceDate#", inData.CreatedAt.Format("2006-01-02 15:04:05"),
		"#PaymentMethod#", inData.PaymentMethodName,
		"#Item#", inData.Item, "#InvoiceCode#", strconv.Itoa(int(inData.ID)),
		"#BillIssuerCompany#", inData.BillIssuerName)
	bodyReplace := r.Replace(body)

	return bodyReplace, nil
}
