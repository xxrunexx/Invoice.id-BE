package data

import (
	"invoice-api/features/billissuerdetail"

	"gorm.io/gorm"
)

type BillIssuerDetail struct {
	gorm.Model
	BillIssuerID   uint
	CompanyName    string
	CompanyAddress string
	CompanyPhone   string
	CompanySite    string
	PaymentTerms   int
}

func toBillIssuerDetailRecord(billissuerdetail billissuerdetail.BillIssuerDetailCore) BillIssuerDetail {
	return BillIssuerDetail{
		Model: gorm.Model{
			ID:        billissuerdetail.ID,
			CreatedAt: billissuerdetail.CreatedAt,
			UpdatedAt: billissuerdetail.UpdatedAt,
		},
		BillIssuerID:   billissuerdetail.BillIssuerID,
		CompanyName:    billissuerdetail.CompanyName,
		CompanyAddress: billissuerdetail.CompanyAddress,
		CompanyPhone:   billissuerdetail.CompanyPhone,
		CompanySite:    billissuerdetail.CompanySite,
		PaymentTerms:   billissuerdetail.PaymentTerms,
	}
}
