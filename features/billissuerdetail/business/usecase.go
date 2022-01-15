package business

import (
	"errors"
	"invoice-api/features/billissuerdetail"
	"invoice-api/helper"
)

type BillIssuerDetailBusiness struct {
	billissuerdetailData billissuerdetail.Data
}

func NewBusinessBillIssuerDetail(bidData billissuerdetail.Data) billissuerdetail.Business {
	return &BillIssuerDetailBusiness{bidData}
}

func (bidBusiness *BillIssuerDetailBusiness) CreateBillIssuerDetail(data billissuerdetail.BillIssuerDetailCore) error {
	if helper.IsEmpty(data.CompanyName) || helper.IsEmpty(data.CompanyAddress) || helper.IsEmpty(data.CompanyPhone) || helper.IsEmpty(data.CompanySite) {
		return errors.New("bad request")
	}

	err := bidBusiness.billissuerdetailData.CreateBillIssuerDetail(data)
	if err != nil {
		return err
	}
	return nil
}
