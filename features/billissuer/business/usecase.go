package business

type BillIssuerBusiness struct {
	biData billissuer.Data
}

func NewBusinessBillIssuer(biData billissuer.Data) billissuer.Business {
	return &BillIssuerBusiness{biData}
}
