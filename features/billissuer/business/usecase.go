package business

import "invoice-api/features/billissuer"

type BillIssuerBusiness struct {
	biData billissuer.Data
}

func NewBusinessBillIssuer(biData billissuer.Data) billissuer.Business {
	return &BillIssuerBusiness{biData}
}

func (biBusiness *BillIssuerBusiness) CreateBillIssuer(biData billissuer.BillIssuerCore) error {
	if err := biBusiness.biData.CreateBillIssuer(biData); err != nil {
		return err
	}
	return nil
}
