package billissuerdetail

import (
	"time"
)

type BillIssuerDetailCore struct {
	ID             uint
	BillIssuerID   uint
	CompanyName    string
	CompanyAddress string
	CompanyPhone   string
	CompanySite    string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// Untuk layer business
type Business interface {
	CreateBillIssuerDetail(data BillIssuerDetailCore) (err error)
	GetBillIssuerDetailById(id int) (billissuerdetail BillIssuerDetailCore, err error)
}

// Untuk layer data / repository
type Data interface {
	CreateBillIssuerDetail(data BillIssuerDetailCore) (err error)
	GetBillIssuerDetailById(id int) (billissuerdetail BillIssuerDetailCore, err error)
}
