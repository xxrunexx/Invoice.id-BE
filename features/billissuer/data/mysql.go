package data

import (
	"gorm.io/gorm"
)

type BillIssuerData struct {
	DB *gorm.DB
}

func NewMySqlBillIssuer(DB *gorm.DB) billissuer.Data {
	return &BillIssuerData{DB}
}

func (biData *BillIssuerData) CreateAccount(billissuer billIssuer.BillIssuerCore) error {
	convData := toBillIssuerRecord(billissuer)

	if err := biData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
