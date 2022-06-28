package request

// import (
// 	"invoice-api/features/csv"
// )

// type InvoiceCSV struct {
// 	ClientID      uint   `csv:"client_id"`
// 	ClientNIK     int    `csv:"client_nik"`
// 	ClientName    string `csv:"client_name"`
// 	ClientPhone   string `csv:"client_phone"`
// 	ClientAddress string `csv:"client_address"`
// 	ClientEmail   string `csv:"client_email"`
// 	Item          string `csv:"item"`
// 	Total         int    `csv:"total"`
// 	BillIssuerID  uint   `csv:"bill_issuer_id"`
// 	// PaymentMethodID uint   `csv:"payment_method_id"`
// 	PaymentTerms int `csv:"payment_terms"`
// }

// func (reqData *InvoiceCSV) toInvoiceCore() csv.InvoiceCsvCore {
// 	return csv.InvoiceCsvCore{
// 		ClientID:      reqData.ClientID,
// 		ClientNIK:     reqData.ClientNIK,
// 		ClientName:    reqData.ClientName,
// 		ClientPhone:   reqData.ClientPhone,
// 		ClientAddress: reqData.ClientAddress,
// 		ClientEmail:   reqData.ClientEmail,
// 		Item:          reqData.Item,
// 		Total:         reqData.Total,
// 		BillIssuerID:  reqData.BillIssuerID,
// 		PaymentTerms:  reqData.PaymentTerms,
// 	}
// }

// func ToInvoiceCoreList(icsv []InvoiceCSV) []csv.InvoiceCsvCore {
// 	arrData := []csv.InvoiceCsvCore{}
// 	for _, invoice := range icsv {
// 		arrData = append(arrData, invoice.toInvoiceCore())
// 	}
// 	return arrData
// }
