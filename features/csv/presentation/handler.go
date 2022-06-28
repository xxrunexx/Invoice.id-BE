package presentation

// import (
// 	"invoice-api/features/csv"
// 	"invoice-api/features/csv/presentation/request"
// 	"invoice-api/helper"
// 	"io"
// 	"net/http"
// 	"os"

// 	"github.com/gocarina/gocsv"
// 	"github.com/labstack/echo/v4"
// )

// type CsvHandler struct {
// 	csvBusiness csv.Business
// }

// func NewHandlerCsv(csvBusiness csv.Business) *CsvHandler {
// 	return &CsvHandler{csvBusiness}
// }

// func (csvHandler *CsvHandler) CreateBulkInvoice(e echo.Context) error {
// 	file, err := e.FormFile("file")
// 	// fmt.Println("Isi file", file)
// 	if err != nil {
// 		return err
// 	}

// 	// data, err := inHandler.invoiceBusiness.CheckInvoice(file)

// 	src, err := file.Open()
// 	if err != nil {
// 		return err
// 	}
// 	defer src.Close()

// 	dst, err := os.Create(file.Filename)
// 	if err != nil {
// 		return err
// 	}
// 	defer dst.Close()

// 	if _, err = io.Copy(dst, src); err != nil {
// 		return err
// 	}
// 	// fmt.Printf("<p>File %s uploaded successfully", file.Filename)

// 	invoicesFile, err := os.OpenFile(file.Filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer invoicesFile.Close()

// 	invoices := []request.InvoiceCSV{}
// 	if err := e.Bind(&invoices); err != nil {
// 		return helper.ErrorResponse(e, http.StatusBadRequest, "bad request", err)
// 	}

// 	if err := gocsv.UnmarshalFile(invoicesFile, &invoices); err != nil { // Load clients from file
// 		panic(err)
// 	}
// 	csvHandler.csvBusiness.CreateBulkInvoice(request.ToInvoiceCoreList(invoices))
// 	// for _, invoice := range invoices {
// 	// 	if _,err := csvHandler.csvBusiness.CreateBulkInvoice(invoice.ToInvoiceCoreLis); err != nil {
// 	// 		return helper.ErrorResponse(e, http.StatusInternalServerError, "internal server error", err)
// 	// 	}
// 	// }
// 	// // csvContent, err := gocsv.MarshalString(&invoices) // Get all clients as CSV string
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
// 	// fmt.Println(csvContent) // Display all clients as CSV string

// 	return nil
// }
