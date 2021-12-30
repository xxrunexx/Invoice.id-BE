package request

import "invoice-api/features/billissuer"

type ReqBillIssuer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (reqData *ReqBillIssuer) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		Username: reqData.Username,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
