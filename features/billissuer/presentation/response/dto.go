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

func ToBillIssuerResponse(billissuer billissuer.BillIssuerCore) RespBillIssuer {
	return RespBillIssuer{
		Id:       billissuer.ID,
		Username: billissuer.Username,
		Password: billissuer.Password,
		Email:    billissuer.Email,
	}
}

func ToBillIssuerLoginResponse(billissuer billissuer.BillIssuerCore) RespBillIssuerLogin {
	return RespBillIssuerLogin{
		Id:       billissuer.ID,
		Username: billissuer.Username,
		Token:    billissuer.Token,
	}
}
