package business

import (
	"errors"
	"fmt"
	"invoice-api/features/client"
	"invoice-api/helper"
)

type ClientBusiness struct {
	clienData client.Data
}

func NewBusinessClient(clData client.Data) client.Business {
	return &ClientBusiness{clData}
}

func (clBusiness *ClientBusiness) CreateClient(data client.ClientCore) error {
	if data.NIK == 0 || helper.IsEmpty(data.Phone) || helper.IsEmpty(data.Address) || helper.IsEmpty(data.Email) {
		return errors.New("invalid data")
	}

	isExist, err := clBusiness.clienData.GetClientByNik(data.NIK)
	if err != nil {
		return err
	}
	if isExist {
		setMessage := fmt.Sprintf("nik %v already in use!", data.NIK)
		return errors.New(setMessage)
	}

	if err := clBusiness.clienData.CreateClient(data); err != nil {
		return err
	}
	return nil
}

func (clBussiness *ClientBusiness) GetAllCient(data client.ClientCore) ([]client.ClientCore, error) {
	clients, err := clBussiness.clienData.GetAllCient(data)

	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (clBussiness *ClientBusiness) GetClientById(id int) (client.ClientCore, error) {
	clData, err := clBussiness.clienData.GetClientById(id)

	if err != nil {
		return client.ClientCore{}, err
	}
	return clData, nil
}

func (clBusiness *ClientBusiness) UpdateClient(data client.ClientCore) error {
	if data.NIK == 0 || helper.IsEmpty(data.Phone) || helper.IsEmpty(data.Address) || helper.IsEmpty(data.Email) {
		return errors.New("invalid data")
	}

	isExist, err := clBusiness.clienData.GetClientByNik(data.NIK)
	if err != nil {
		return err
	}
	if isExist {
		setMessage := fmt.Sprintf("nik %v already in use!", data.NIK)
		return errors.New(setMessage)
	}

	err = clBusiness.clienData.UpdateClient(data)
	if err != nil {
		return err
	}
	return nil
}
