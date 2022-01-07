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
