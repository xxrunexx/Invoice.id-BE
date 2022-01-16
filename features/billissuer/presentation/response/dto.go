package response

import "invoice-api/features/billissuer"

type RespBillIssuer struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RespBillIssuerLogin struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func ToBillIssuerResponse(bi billissuer.BillIssuerCore) RespBillIssuer {
	return RespBillIssuer{
		Id:       bi.ID,
		Username: bi.Username,
		Password: bi.Password,
		Email:    bi.Email,
	}
}

func ToBillIssuerLoginResponse(bi billissuer.BillIssuerCore) RespBillIssuerLogin {
	return RespBillIssuerLogin{
		Id:       bi.ID,
		Username: bi.Username,
		Token:    bi.Token,
	}
}
