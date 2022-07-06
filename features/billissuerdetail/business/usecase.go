package business

import (
	"errors"
	"invoice-api/features/billissuer"
	"invoice-api/features/billissuerdetail"
	"invoice-api/helper"
)

type BillIssuerDetailBusiness struct {
	billissuerdetailData billissuerdetail.Data
	billissuerBusiness   billissuer.Business
}

func NewBusinessBillIssuerDetail(bidData billissuerdetail.Data, biBus billissuer.Business) billissuerdetail.Business {
	return &BillIssuerDetailBusiness{
		billissuerdetailData: bidData,
		billissuerBusiness:   biBus,
	}
}

func (bidBusiness *BillIssuerDetailBusiness) CreateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) (billissuerdetail.BillIssuerDetailCore, error) {
	if helper.IsEmpty(data.CompanyName) || helper.IsEmpty(data.CompanyAddress) || helper.IsEmpty(data.CompanyPhone) || helper.IsEmpty(data.CompanySite) {
		return billissuerdetail.BillIssuerDetailCore{}, errors.New("bad request")
	}

	_, err := bidBusiness.billissuerBusiness.GetBillIssuerById(int(data.BillIssuerID))
	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, errors.New("user not found")
	}

	result, err := bidBusiness.billissuerdetailData.CreateBillIssuerDetail(data)
	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return result, nil
}

func (bidBusiness *BillIssuerDetailBusiness) GetBillIssuerDetailById(id int) (billissuerdetail.BillIssuerDetailCore, error) {
	bidData, err := bidBusiness.billissuerdetailData.GetBillIssuerDetailByBillIssuerId(id)

	if err != nil {
		return billissuerdetail.BillIssuerDetailCore{}, err
	}
	return bidData, nil
}

func (bidBusiness *BillIssuerDetailBusiness) UpdateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) error {
	err := bidBusiness.billissuerdetailData.UpdateBillIssuerDetail(data)

	if err != nil {
		return err
	}
	return nil
}
