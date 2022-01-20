package helper

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"invoice-api/features/invoice"
	"path/filepath"
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
	bodyReplace := strings.Replace(body, "#Name#", inData.ClientName, -1)

	return bodyReplace, nil
}
