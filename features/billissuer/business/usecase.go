package business

import (
	"errors"
	"fmt"
	"invoice-api/features/billissuer"
	"invoice-api/helper"
	"invoice-api/middleware"
)

type BillIssuerBusiness struct {
	billissuerData billissuer.Data
}

func NewBusinessBillIssuer(biData billissuer.Data) billissuer.Business {
	return &BillIssuerBusiness{biData}
}

func (biBusiness *BillIssuerBusiness) CreateBillIssuer(data billissuer.BillIssuerCore) (billissuer.BillIssuerCore, error) {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) || helper.IsEmpty(data.Name) {
		return billissuer.BillIssuerCore{}, errors.New("bad request")
	}

	isExist, err := biBusiness.billissuerData.GetBillIssuerByEmail(data.Email)
	if err != nil {
		return billissuer.BillIssuerCore{}, errors.New("duplicate data")
	}
	if isExist {
		setMessage := fmt.Sprintf("email %v already in use!", data.Email)
		return billissuer.BillIssuerCore{}, errors.New(setMessage)
	}

	result, err := biBusiness.billissuerData.CreateBillIssuer(data)
	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}
	return result, nil
}

func (biBussiness *BillIssuerBusiness) LoginBillIssuer(data billissuer.BillIssuerCore) (billissuer.BillIssuerCore, error) {
	biData, err := biBussiness.billissuerData.LoginBillIssuer(data)

	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}
	biData.Token, err = middleware.CreateToken(biData.ID, biData.Name, biData.Email)
	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}
	return biData, nil
}

func (biBussiness *BillIssuerBusiness) GetBillIssuerById(id int) (billissuer.BillIssuerCore, error) {
	biData, err := biBussiness.billissuerData.GetBillIssuerById(id)

	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}
	return biData, nil
}

func (biBussiness *BillIssuerBusiness) UpdateBillIssuer(data billissuer.BillIssuerCore) error {
	if helper.IsEmpty(data.Name) || helper.IsEmpty(data.Password) || helper.IsEmpty(data.Email) {
		return errors.New("invalid data")
	}

	// isExist, err := biBussiness.billissuerData.GetBillIssuerByEmail(data.Email)
	// if err != nil {
	// 	return err
	// }
	// if isExist {
	// 	setMessage := fmt.Sprintf("email %v already in use!", data.Email)
	// 	return errors.New(setMessage)
	// }

	err := biBussiness.billissuerData.UpdateBillIssuer(data)
	if err != nil {
		return err
	}

	return nil
}

func (biBussiness *BillIssuerBusiness) GetAllBillIssuer(data billissuer.BillIssuerCore) ([]billissuer.BillIssuerCore, error) {
	billissuers, err := biBussiness.billissuerData.GetAllBillIssuer(data)

	if err != nil {
		return nil, err
	}
	return billissuers, nil
}
