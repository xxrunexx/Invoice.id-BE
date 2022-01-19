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

	// err := bidData.DB.First(&singleData, id).Joins("left join bill_issuers on bill_issuers.id = bill_issuer_details.bill_issuer_id").Error
	err := bidData.DB.Where("bill_issuer_details.bill_issuer_id = ?", id).Joins("BillIssuer").Find(&singleData).Error
	// err := bidData.DB.Model(&singleData).Where("id = ?", id).Joins("left join bill_issuers on bill_issuers.id = bill_issuer_details.bill_issuer_id").Error
	// err := bidData.DB.Joins("BillIssuers").First(&singleData, id).Error
	// err := bidData.DB.Model(&singleData).Where("id = ?", id).Joins("left join bill_issuers on bill_issuer_details.bill_issuer_id = bill_issuers.id").Error

	if singleData.BillIssuerID == 0 {
		return billissuerdetail.BillIssuerDetailCore{}, errors.New("data not found")
	}

	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return toBillIssuerDetailCore(singleData), nil
}
