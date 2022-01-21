package data

import (
	"invoice-api/features/billissuerdetail"

	"gorm.io/gorm"
)

type BillIssuerDetail struct {
	gorm.Model
	BillIssuerID    uint
	BillIssuer      BillIssuer `gorm:"foreignKey:ID;references:BillIssuerID"`
	BillIssuerEmail string
	CompanyName     string
	CompanyAddress  string
	CompanyPhone    string
	CompanySite     string
}

type BillIssuer struct {
	ID       uint
	Name     string
	Password string
	Email    string
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
	}
}

func toBillIssuerDetailCore(bid BillIssuerDetail) billissuerdetail.BillIssuerDetailCore {
	return billissuerdetail.BillIssuerDetailCore{
		ID:              bid.ID,
		BillIssuerID:    bid.BillIssuerID,
		BillIssuerName:  bid.BillIssuer.Name,
		BillIssuerEmail: bid.BillIssuer.Email,
		CompanyName:     bid.CompanyName,
		CompanyAddress:  bid.CompanyAddress,
		CompanyPhone:    bid.CompanyPhone,
		CompanySite:     bid.CompanySite,
		CreatedAt:       bid.CreatedAt,
		UpdatedAt:       bid.UpdatedAt,
	}
}
