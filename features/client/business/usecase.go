package business

import "invoice-api/features/client"

type ClientBusiness struct {
	clienData client.Data
}

func NewBusinessClient(clData client.Data) client.Business {
	return &ClientBusiness{clData}
}

func (clBusiness *ClientBusiness) CreateClient(data client.ClientCore) error {
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
