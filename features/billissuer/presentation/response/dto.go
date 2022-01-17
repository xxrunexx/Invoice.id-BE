package response

import "invoice-api/features/billissuer"

type RespBillIssuer struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type RespBillIssuerLogin struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
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
		Id:    bi.ID,
		Email: bi.Email,
		Token: bi.Token,
	}
}
