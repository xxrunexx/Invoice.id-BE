package data

import (
	"fmt"
	"invoice-api/features/invoice"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ClientID         uint
	Client           Client `gorm:"foreignKey:ID;references:ClientID"`
	ClientPhone      string
	ClientAddress    string
	ClientEmail      string
	Total            int
	BillIssuerID     uint
	BillIssuerDetail BillIssuerDetail `gorm:"foreignKey:ID;references:BillIssuerID"`
	PaymentMethodID  uint
	PaymentMethod    PaymentMethod `gorm:"foreignKey:ID;references:PaymentMethodID"`
	PaymentDue       time.Time
	PaymentStatus    string `gorm:"default:draft"`
	PaymentTerms     int
}

type BillIssuerDetail struct {
	ID              uint
	BillIssuerID    uint
	BillIssuer      BillIssuer `gorm:"foreignKey:ID;references:BillIssuerID"`
	BillIssuerEmail string
	CompanyName     string
	CompanyAddress  string
	CompanyPhone    string
	CompanySite     string
}

type BillIssuer struct {
	ID   uint
	Name string
}
type Client struct {
	ID      uint
	NIK     int
	Name    string
	Phone   string
	Address string
	Email   string
}

type PaymentMethod struct {
	ID   uint
	Name string
}

func toInvoiceRecord(in invoice.InvoiceCore) Invoice {
	fmt.Println("Isi payment due di record : ", in.PaymentDue)
	return Invoice{
		Model: gorm.Model{
			ID:        in.ID,
			CreatedAt: in.CreatedAt,
			UpdatedAt: in.UpdatedAt,
		},
		ClientID:        in.ClientID,
		Total:           in.Total,
		BillIssuerID:    in.BillIssuerID,
		PaymentMethodID: in.PaymentMethodID,
		PaymentDue:      in.PaymentDue,
		PaymentStatus:   in.PaymentStatus,
		PaymentTerms:    in.PaymentTerms,
	}
}

func toInvoiceCore(in Invoice) invoice.InvoiceCore {
	return invoice.InvoiceCore{
		ID:                in.ID,
		ClientID:          in.ClientID,
		ClientName:        in.Client.Name,
		ClientPhone:       in.Client.Phone,
		ClientAddress:     in.Client.Address,
		ClientEmail:       in.Client.Email,
		Total:             in.Total,
		BillIssuerID:      in.BillIssuerID,
		BillIssuerName:    in.BillIssuerDetail.BillIssuer.Name,
		PaymentMethodID:   in.PaymentMethodID,
		PaymentMethodName: in.PaymentMethod.Name,
		PaymentDue:        in.PaymentDue,
		PaymentStatus:     in.PaymentStatus,
		PaymentTerms:      in.PaymentTerms,
		CreatedAt:         in.CreatedAt,
		UpdatedAt:         in.UpdatedAt,
	}
}

func toInvoiceCoreList(inList []Invoice) []invoice.InvoiceCore {
	convIn := []invoice.InvoiceCore{}

	for _, invoice := range inList {
		convIn = append(convIn, toInvoiceCore(invoice))
	}
	return convIn
}
