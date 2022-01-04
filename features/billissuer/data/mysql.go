package data

import (
	"errors"
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

func (biData *BillIssuerData) LoginBillIssuer(data billissuer.BillIssuerCore) (billissuer.BillIssuerCore, error) {
	var billissuerData BillIssuer
	err := biData.DB.Where("username = ? and password = ?", data.Username, data.Password).First(&billissuerData).Error

	// Eliminate null data
	if billissuerData.Username == "" && billissuerData.ID == 0 {
		return billissuer.BillIssuerCore{}, errors.New("user not found")
	}

	// Validate with DB
	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}

	return toBillIssuerCore(billissuerData), nil
}
