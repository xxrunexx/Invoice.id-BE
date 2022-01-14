package request

import "invoice-api/features/billissuer"

type ReqBillIssuer struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqBillIssuerUpdate struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqBIllIssuerAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (data *ReqBIllIssuerAuth) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		Username: data.Username,
		Password: data.Password,
	}
}

func (reqData *ReqBillIssuer) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		Username: reqData.Username,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
func (reqData *ReqBillIssuerUpdate) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		ID:       reqData.ID,
		Username: reqData.Username,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
