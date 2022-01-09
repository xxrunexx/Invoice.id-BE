package request

import "invoice-api/features/client"

type ReqClient struct {
	NIK     int    `json:"nik"`
	Name    string `json:"name"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func (reqData *ReqClient) ToClientCore() client.ClientCore {
	return client.ClientCore{
		NIK:     reqData.NIK,
		Name:    reqData.Name,
		Phone:   reqData.Phone,
		Address: reqData.Address,
		Email:   reqData.Email,
	}
}
