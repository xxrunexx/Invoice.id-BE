package request

import "invoice-api/features/billissuer"

type ReqBillIssuer struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqBillIssuerUpdate struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type ReqBIllIssuerAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (reqData *ReqBIllIssuerAuth) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		Email:    reqData.Email,
		Password: reqData.Password,
	}
}

func (reqData *ReqBillIssuer) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		Name:     reqData.Name,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
func (reqData *ReqBillIssuerUpdate) ToBillIssuerCore() billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		ID:       reqData.ID,
		Name:     reqData.Name,
		Password: reqData.Password,
		Email:    reqData.Email,
	}
}
