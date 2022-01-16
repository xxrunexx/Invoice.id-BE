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

func toBillIssuerDetailRecord(bid billissuerdetail.BillIssuerDetailCore) BillIssuerDetail {
	return BillIssuerDetail{
		Model: gorm.Model{
			ID:        bid.ID,
			CreatedAt: bid.CreatedAt,
			UpdatedAt: bid.UpdatedAt,
		},
		BillIssuerID:   bid.BillIssuerID,
		CompanyName:    bid.CompanyName,
		CompanyAddress: bid.CompanyAddress,
		CompanyPhone:   bid.CompanyPhone,
		CompanySite:    bid.CompanySite,
		PaymentTerms:   bid.PaymentTerms,
	}
}

func toBillIssuerDetailCore(bid BillIssuerDetail) billissuerdetail.BillIssuerDetailCore {
	return billissuerdetail.BillIssuerDetailCore{
		ID:             bid.ID,
		BillIssuerID:   bid.BillIssuerID,
		CompanyName:    bid.CompanyName,
		CompanyAddress: bid.CompanyAddress,
		CompanyPhone:   bid.CompanyPhone,
		CompanySite:    bid.CompanySite,
		PaymentTerms:   bid.PaymentTerms,
	}
}
