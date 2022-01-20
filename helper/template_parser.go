package helper

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"invoice-api/features/invoice"
	"path/filepath"
	"strconv"
	"strings"
)

func parseTemplate(templateFileName string, data interface{}, inData invoice.InvoiceCore) (string, error) {
	templatePath, err := filepath.Abs(fmt.Sprintf("helper/email_templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	body := buf.String()
	r := strings.NewReplacer("#ClientName#", inData.ClientName, "#ClientAddress#", inData.ClientAddress, "#CreatedAt#", inData.CreatedAt.String(), "#Total#", strconv.Itoa(inData.Total), "#PaymentDue#", inData.PaymentDue.String())
	bodyReplace := r.Replace(body)
	// bodyReplace := strings.Replace(body, "#ClientName#", inData.ClientName, -1)
	// bodyReplace = strings.Replace(body, "#ClientAddress#", inData.ClientAddress, -1)
	// bodyReplace = strings.Replace(body, "#CreatedAt#", inData.CreatedAt.String(), -1)
	// bodyReplace = strings.Replace(body, "#Total#", string(rune(inData.Total)), -1)
	// bodyReplace = strings.Replace(body, "#PaymentDue#", inData.PaymentDue.String(), -1)

	return bodyReplace, nil
}
