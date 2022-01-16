package response

import (
	"invoice-api/features/billissuerdetail"
	"time"
)

type RespBillIssuerDetail struct {
	ID             uint      `json:"id"`
	BillIssuerID   uint      `json:"bill_issuer_id"`
	CompanyName    string    `json:"company_name"`
	CompanyAddress string    `json:"company_address"`
	CompanyPhone   string    `json:"company_phone"`
	CompanySite    string    `json:"company_site"`
	PaymentTerms   int       `json:"payment_terms"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func ToBillIssuerDetailResponse(bid billissuerdetail.BillIssuerDetailCore) RespBillIssuerDetail {
	return RespBillIssuerDetail{
		ID:             bid.ID,
		BillIssuerID:   bid.BillIssuerID,
		CompanyName:    bid.CompanyName,
		CompanyAddress: bid.CompanyAddress,
		CompanyPhone:   bid.CompanyPhone,
		CompanySite:    bid.CompanySite,
		PaymentTerms:   bid.PaymentTerms,
		CreatedAt:      bid.CreatedAt,
		UpdatedAt:      bid.UpdatedAt,
	}
}