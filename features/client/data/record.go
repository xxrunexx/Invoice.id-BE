package data

import (
	"invoice-api/features/client"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	NIK     int
	Name    string
	Phone   string
	Address string
	Email   string
}

func toClientRecord(client client.ClientCore) Client {
	return Client{
		Model: gorm.Model{
			ID: client.ID,
		},
		NIK:     client.NIK,
		Name:    client.Name,
		Phone:   client.Phone,
		Address: client.Address,
		Email:   client.Email,
	}
}
