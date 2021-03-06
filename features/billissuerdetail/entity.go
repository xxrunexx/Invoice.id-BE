package billissuerdetail

import (
	"time"
)

type BillIssuerDetailCore struct {
	ID              uint
	BillIssuerID    uint
	BillIssuerName  string
	BillIssuerEmail string
	CompanyName     string
	CompanyAddress  string
	CompanyPhone    string
	CompanySite     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// Untuk layer business
type Business interface {
	CreateBillIssuerDetail(data BillIssuerDetailCore) (billissuerdetail BillIssuerDetailCore, err error)
	GetBillIssuerDetailById(id int) (billissuerdetail BillIssuerDetailCore, err error)
	UpdateBillIssuerDetail(data BillIssuerDetailCore) error
}

// Untuk layer data / repository
type Data interface {
	CreateBillIssuerDetail(data BillIssuerDetailCore) (billissuerdetail BillIssuerDetailCore, err error)
	GetBillIssuerDetailById(id int) (billissuerdetail BillIssuerDetailCore, err error)
	GetBillIssuerDetailByBillIssuerId(id int) (billissuerdetail BillIssuerDetailCore, err error)
	UpdateBillIssuerDetail(data BillIssuerDetailCore) error
}
