package response

import "invoice-api/features/billissuer"

type RespBillIssuer struct {
	Id       uint   `json:"id"`
	Name     string `json:"Name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RespBillIssuerLogin struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToBillIssuerResponse(bi billissuer.BillIssuerCore) RespBillIssuer {
	return RespBillIssuer{
		Id:       bi.ID,
		Name:     bi.Name,
		Password: bi.Password,
		Email:    bi.Email,
	}
}

func ToBillIssuerLoginResponse(bi billissuer.BillIssuerCore) RespBillIssuerLogin {
	return RespBillIssuerLogin{
		Id:    bi.ID,
		Name:  bi.Name,
		Email: bi.Email,
		Token: bi.Token,
	}
}
