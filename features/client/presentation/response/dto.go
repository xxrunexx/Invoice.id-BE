package response

import "invoice-api/features/client"

type RespClient struct {
	ID      uint   `json:"id"`
	NIK     int    `json:"nik"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

func ToClientResponse(client client.ClientCore) RespClient {
	return RespClient{
		ID:      client.ID,
		NIK:     client.NIK,
		Name:    client.Name,
		Phone:   client.Phone,
		Address: client.Address,
		Email:   client.Email,
	}
}

func ToClientResponseList(clList []client.ClientCore) []RespClient {
	convCl := []RespClient{}

	for _, client := range clList {
		convCl = append(convCl, ToClientResponse(client))
	}
	return convCl
}
