package business

// import (
// 	"invoice-api/features/client"
// 	"invoice-api/features/csv"
// 	"invoice-api/features/invoice"
// )

// type CsvBusiness struct {
// 	invoiceBusiness invoice.Business
// 	clientBusiness  client.Business
// }

// func NewBusinessCsv(inBus invoice.Business, clBus client.Business) csv.Business {
// 	return &CsvBusiness{
// 		invoiceBusiness: inBus,
// 		clientBusiness:  clBus,
// 	}
// }

// func (csvBusiness *CsvBusiness) CreateBulkInvoice(in []csv.InvoiceCore, cl []csv.ClientCore) error {
// 	for _, client := range csv {
// 		for _, newData := range client {
// 			csvBusiness.clientBusiness.CreateClient()
// 		}
// 	}
// 	return nil
// }
