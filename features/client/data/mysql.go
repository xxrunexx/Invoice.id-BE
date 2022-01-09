package data

import (
	"invoice-api/features/client"

	"gorm.io/gorm"
)

type ClientData struct {
	DB *gorm.DB
}

func NewMySqlClient(DB *gorm.DB) client.Data {
	return &ClientData{DB}
}

func (clData *ClientData) CreateClient(data client.ClientCore) error {
	convData := toClientRecord(data)

	if err := clData.DB.Create(&convData).Error; err != nil {
		return err
	}
	return nil
}
