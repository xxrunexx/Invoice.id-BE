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

func toClientCore(cl Client) client.ClientCore {
	return client.ClientCore{
		ID:      cl.ID,
		NIK:     cl.NIK,
		Name:    cl.Name,
		Phone:   cl.Phone,
		Address: cl.Address,
		Email:   cl.Email,
	}
}

func toClientCoreList(clList []Client) []client.ClientCore {
	convCl := []client.ClientCore{}

	for _, client := range clList {
		convCl = append(convCl, toClientCore(client))
	}
	return convCl
}
