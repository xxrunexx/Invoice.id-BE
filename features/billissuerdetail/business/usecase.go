package business

import (
	"errors"
	"invoice-api/features/billissuer"
	"invoice-api/features/billissuerdetail"
	"invoice-api/helper"
)

type BillIssuerDetailBusiness struct {
	billissuerdetailData billissuerdetail.Data
	billissuerData       billissuer.Data
}

func NewBusinessBillIssuerDetail(bidData billissuerdetail.Data, biData billissuer.Data) billissuerdetail.Business {
	return &BillIssuerDetailBusiness{
		billissuerdetailData: bidData,
		billissuerData:       biData,
	}
}

func (bidBusiness *BillIssuerDetailBusiness) CreateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) error {
	if helper.IsEmpty(data.CompanyName) || helper.IsEmpty(data.CompanyAddress) || helper.IsEmpty(data.CompanyPhone) || helper.IsEmpty(data.CompanySite) {
		return errors.New("bad request")
	}

	_, err := bidBusiness.billissuerData.GetBillIssuerById(int(data.BillIssuerID))
	if err != nil {
		return errors.New("user not found")
	}

	err = bidBusiness.billissuerdetailData.CreateBillIssuerDetail(data)
	if err != nil {
		return err
	}
	return nil
}

func (bidBusiness *BillIssuerDetailBusiness) GetBillIssuerDetailById(id int) (billissuerdetail.BillIssuerDetailCore, error) {
	bidData, err := bidBusiness.billissuerdetailData.GetBillIssuerDetailById(id)

	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return bidData, nil
}
