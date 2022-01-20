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

func (bidData *BillIssuerDetailData) CreateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) (billissuerdetail.BillIssuerDetailCore, error) {
	convData := toBillIssuerDetailRecord(data)

	if err := bidData.DB.Create(&convData).Error; err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}

	record, err := bidData.GetBillIssuerDetailById(int(convData.ID))
	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return record, nil
}

func (bidData *BillIssuerDetailData) GetBillIssuerDetailById(id int) (billissuerdetail.BillIssuerDetailCore, error) {
	var singleData BillIssuerDetail

	err := bidData.DB.Where("bill_issuer_details.bill_issuer_id = ?", id).Joins("BillIssuer").Find(&singleData).Error
	if singleData.BillIssuerID == 0 {
		return billissuerdetail.BillIssuerDetailCore{}, errors.New("data not found")
	}

	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return toBillIssuerDetailCore(singleData), nil
}
