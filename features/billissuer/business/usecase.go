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

func (biBusiness *BillIssuerBusiness) CreateBillIssuer(data billissuer.BillIssuerCore) error {
	if !helper.ValidateEmail(data.Email) || !helper.ValidatePassword(data.Password) || len(data.Username) == 0 {
		return errors.New("bad request")
	}

	isExist, err := biBusiness.billissuerData.GetBillIssuerByEmail(data.Email)
	if err != nil {
		return err
	}
	if isExist {
		setMessage := fmt.Sprintf("email %v already in use!", data.Email)
		fmt.Println("Isi setMessage : ", setMessage)
		return errors.New(setMessage)
	}

	err = biBusiness.billissuerData.CreateBillIssuer(data)
	if err != nil {
		return err
	}
	// if err := biBusiness.biData.CreateBillIssuer(data); err != nil {
	// 	return err
	// }

	return nil
}

func (biBussiness *BillIssuerBusiness) LoginBillIssuer(data billissuer.BillIssuerCore) (billissuer.BillIssuerCore, error) {
	biData, err := biBussiness.billissuerData.LoginBillIssuer(data)

	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}
	biData.Token, err = middleware.CreateToken(data.ID, data.Username)
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

func (biBussiness *BillIssuerBusiness) UpdateBillIssuer(id int) (billissuer.BillIssuerCore, error) {
	// biData, err := biBussiness.billissuerData.GetBillIssuerById(id)

	// if err != nil {
	// 	return billissuer.BillIssuerCore{}, err
	// }

	// if helper.IsEmpty(biData.Username) || helper.IsEmpty(biData.Password) || helper.IsEmpty(biData.Email) {
	// 	return billissuer.BillIssuerCore{}, errors.New("bad request")
	// }

	// _, err := biBussiness.billissuerData.GetBillIssuerById(int(data.ID))
	// if err != nil {
	// 	return err
	// }

	biData, err := biBussiness.billissuerData.UpdateBillIssuer(id)
	if err != nil {
		return billissuer.BillIssuerCore{}, err
	}

	return biData, nil
}
