package response

import "invoice-api/features/billissuer"

type RespBillIssuer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToBillIssuerResponse(billissuer billissuer.BillIssuerCore) RespBillIssuer {
	return RespBillIssuer{
		Username: billissuer.Username,
		Password: billissuer.Password,
		Email:    billissuer.Email,
	}
}
