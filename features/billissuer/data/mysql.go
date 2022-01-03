package data

import (
	"invoice-api/features/billissuer"

	"gorm.io/gorm"
)

type BillIssuerData struct {
	DB *gorm.DB
}

func NewMySqlBillIssuer(DB *gorm.DB) billissuer.Data {
	return &BillIssuerData{DB}
}

func (biData *BillIssuerData) CreateBillIssuer(billissuer billissuer.BillIssuerCore) error {
	convData := toBillIssuerRecord(billissuer)

	if err := biData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
