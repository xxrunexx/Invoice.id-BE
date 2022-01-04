package business

import (
	"errors"
	"invoice-api/features/billissuer"
	"invoice-api/helper"
)

type BillIssuerBusiness struct {
	biData billissuer.Data
}

func NewBusinessBillIssuer(biData billissuer.Data) billissuer.Business {
	return &BillIssuerBusiness{biData}
}

func (biBusiness *BillIssuerBusiness) CreateBillIssuer(data billissuer.BillIssuerCore) error {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) || len(data.Username) == 0 {
		return errors.New("bad request")
	}

	_, err := biBusiness.biData.CreateBillIssuer(data)
	if err != nil {
		return err
	}

	// if err := biBusiness.biData.CreateBillIssuer(data); err != nil {
	// 	return err
	// }

	return nil
}
