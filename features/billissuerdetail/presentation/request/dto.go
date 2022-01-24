package request

import "invoice-api/features/billissuerdetail"

type ReqBillIssuerDetail struct {
	BillIssuerID   uint   `json:"bill_issuer_id"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	CompanyPhone   string `json:"company_phone"`
	CompanySite    string `json:"company_site"`
}

type ReqBillIssuerDetailUpdate struct {
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	CompanyPhone   string `json:"company_phone"`
	CompanySite    string `json:"company_site"`
}

func (reqData *ReqBillIssuerDetail) ToBillIssuerDetailCore() billissuerdetail.BillIssuerDetailCore {
	return billissuerdetail.BillIssuerDetailCore{
		BillIssuerID:   reqData.BillIssuerID,
		CompanyName:    reqData.CompanyName,
		CompanyAddress: reqData.CompanyAddress,
		CompanyPhone:   reqData.CompanyPhone,
		CompanySite:    reqData.CompanySite,
	}
}
func (reqData *ReqBillIssuerDetailUpdate) ToBillIssuerDetailCore() billissuerdetail.BillIssuerDetailCore {
	return billissuerdetail.BillIssuerDetailCore{
		CompanyName:    reqData.CompanyName,
		CompanyAddress: reqData.CompanyAddress,
		CompanyPhone:   reqData.CompanyPhone,
		CompanySite:    reqData.CompanySite,
	}
}
