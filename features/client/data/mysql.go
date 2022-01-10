package data

import (
	"errors"
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

func (clData *ClientData) GetAllCient(data client.ClientCore) ([]client.ClientCore, error) {
	var clients []Client

	err := clData.DB.Find(&clients).Error

	if err != nil {
		return nil, err
	}
	return toClientCoreList(clients), nil
}

func (clData *ClientData) GetClientById(id int) (client.ClientCore, error) {
	var singleData Client

	err := clData.DB.First(&singleData, id).Error

	if singleData.ID == 0 && singleData.Name == "" {
		return client.ClientCore{}, errors.New("data not found")
	}
	if err != nil {
		return client.ClientCore{}, err
	}
	return toClientCore(singleData), nil
}
