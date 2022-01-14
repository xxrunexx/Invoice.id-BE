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

func (biData *BillIssuerData) CreateBillIssuer(data billissuer.BillIssuerCore) error {
	convData := toBillIssuerRecord(data)

	if err := biData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
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

func (biData *BillIssuerData) GetBillIssuerById(id int) (billissuer.BillIssuerCore, error) {
	var singleData BillIssuer

	err := biData.DB.First(&singleData, id).Error

	if singleData.Username == "" && singleData.ID == 0 {
		return billissuer.BillIssuerCore{}, errors.New("data not found")
	}

	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}

	return toBillIssuerCore(singleData), nil
}

func (biData BillIssuerData) GetBillIssuerByEmail(email string) (bool, error) {
	var singleData BillIssuer
	err := biData.DB.Where("email = ?", email).Find(&singleData).Error
	if err != nil || singleData.ID == 0 {
		return false, err
	}
	return true, nil
}

func (biData BillIssuerData) UpdateBillIssuer(data billissuer.BillIssuerCore) error {
	var singleData BillIssuer
	convData := toBillIssuerRecord(data)
	err := biData.DB.Model(&singleData).Where("id = ?", data.ID).Updates(&convData).Error

	if err != nil {
		return err
	}
	return nil
}
