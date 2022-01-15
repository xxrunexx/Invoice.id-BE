package request

import "invoice-api/features/billissuerdetail"

type ReqBillIssuerDetail struct {
	BillIssuerID   uint   `json:"bill_issuer_id"`
	CompanyName    string `json:"company_name"`
	CompanyAddress string `json:"company_address"`
	CompanyPhone   string `json:"company_phone"`
	CompanySite    string `json:"company_site"`
	PaymentTerms   int    `json:"payment_terms"`
}

func (data *ReqBillIssuerDetail) ToBillIssuerDetailCore() billissuerdetail.BillIssuerDetailCore {
	return billissuerdetail.BillIssuerDetailCore{
		BillIssuerID:   data.BillIssuerID,
		CompanyName:    data.CompanyName,
		CompanyAddress: data.CompanyAddress,
		CompanyPhone:   data.CompanyPhone,
		CompanySite:    data.CompanySite,
		PaymentTerms:   data.PaymentTerms,
	}
}
