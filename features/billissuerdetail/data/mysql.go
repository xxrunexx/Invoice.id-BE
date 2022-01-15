package data

import (
	"invoice-api/features/billissuerdetail"

	"gorm.io/gorm"
)

type BillIssuerDetailData struct {
	DB *gorm.DB
}

func NewMySqlBillIssuerDetail(DB *gorm.DB) billissuerdetail.Data {
	return &BillIssuerDetailData{DB}
}

func (bidData *BillIssuerDetailData) CreateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) error {
	convData := toBillIssuerDetailRecord(data)

	if err := bidData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
