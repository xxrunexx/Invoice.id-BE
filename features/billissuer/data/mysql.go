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

func (biData *BillIssuerData) CreateBillIssuer(data billissuer.BillIssuerCore) (int, error) {
	convData := toBillIssuerRecord(data)

	if err := biData.DB.Create(&convData).Error; err != nil {
		return 0, err
	}
	return int(convData.ID), nil
}
