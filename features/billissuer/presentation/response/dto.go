package response

import "invoice-api/features/billissuer"

type RespBillIssuer struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func ToBillIssuerResponse(billissuer billissuer.BillIssuerCore) RespBillIssuer {
	return RespBillIssuer{
		Id:       billissuer.ID,
		Username: billissuer.Username,
		Password: billissuer.Password,
		Email:    billissuer.Email,
	}
}
