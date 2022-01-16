package data

import (
	"errors"
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

func (bidData *BillIssuerDetailData) GetBillIssuerDetailById(id int) (billissuerdetail.BillIssuerDetailCore, error) {
	var singleData BillIssuerDetail

	err := bidData.DB.First(&singleData, id).Error

	if singleData.BillIssuerID == 0 {
		return billissuerdetail.BillIssuerDetailCore{}, errors.New("data not found")
	}

	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return toBillIssuerDetailCore(singleData), nil
}
