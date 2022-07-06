package data

import (
	"fmt"
	"invoice-api/features/invoice"
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	ClientID        uint
	Client          Client `gorm:"foreignKey:ID;references:ClientID"`
	ClientNIK       int
	ClientPhone     string
	ClientAddress   string
	ClientEmail     string
	Item            string
	Total           int
	BillIssuerID    uint
	BillIssuer      BillIssuer `gorm:"foreignKey:ID;references:BillIssuerID"`
	PaymentMethodID uint
	PaymentMethod   PaymentMethod `gorm:"foreignKey:ID;references:PaymentMethodID"`
	PaymentDue      time.Time
	PaymentStatus   string `gorm:"default:draft"`
	PaymentTerms    int
	PaymentLink     string
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
		Item:            in.Item,
		Total:           in.Total,
		BillIssuerID:    in.BillIssuerID,
		PaymentMethodID: in.PaymentMethodID,
		PaymentDue:      in.PaymentDue,
		PaymentStatus:   in.PaymentStatus,
		PaymentTerms:    in.PaymentTerms,
		PaymentLink:     in.PaymentLink,
	}
}

func toInvoiceCore(in Invoice) invoice.InvoiceCore {
	return invoice.InvoiceCore{
		ID:                in.ID,
		ClientID:          in.ClientID,
		ClientNIK:         in.Client.NIK,
		ClientName:        in.Client.Name,
		ClientPhone:       in.Client.Phone,
		ClientAddress:     in.Client.Address,
		ClientEmail:       in.Client.Email,
		Item:              in.Item,
		Total:             in.Total,
		BillIssuerID:      in.BillIssuerID,
		BillIssuerName:    in.BillIssuer.Name,
		PaymentMethodID:   in.PaymentMethodID,
		PaymentMethodName: in.PaymentMethod.Name,
		PaymentDue:        in.PaymentDue,
		PaymentStatus:     in.PaymentStatus,
		PaymentTerms:      in.PaymentTerms,
		CreatedAt:         in.CreatedAt,
		UpdatedAt:         in.UpdatedAt,
		PaymentLink:       in.PaymentLink,
	}
}

func toInvoiceCoreList(inList []Invoice) []invoice.InvoiceCore {
	convIn := []invoice.InvoiceCore{}

	for _, invoice := range inList {
		convIn = append(convIn, toInvoiceCore(invoice))
	}
	return convIn
}
