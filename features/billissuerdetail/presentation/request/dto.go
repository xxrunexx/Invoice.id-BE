package request

import "invoice-api/features/billissuerdetail"

type ReqBillIssuerDetail struct {
	BillIssuerID   uint
	CompanyName    string
	CompanyAddress string
	CompanyPhone   string
	CompanySite    string
	PaymentTerms   int
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
