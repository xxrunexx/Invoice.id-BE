package billissuer

type BillIssuerCore struct {
	ID       uint
	Name     string
	Password string
	Email    string
	Token    string
}

// Untuk layer business
type Business interface {
	CreateBillIssuer(data BillIssuerCore) (err error)
	LoginBillIssuer(data BillIssuerCore) (billissuer BillIssuerCore, err error)
	GetBillIssuerById(id int) (billissuer BillIssuerCore, err error)
	UpdateBillIssuer(data BillIssuerCore) error
	GetAllBillIssuer(BillIssuerCore) (billissuers []BillIssuerCore, err error)
}

// Untuk layer data / repository
type Data interface {
	CreateBillIssuer(data BillIssuerCore) (err error)
	LoginBillIssuer(data BillIssuerCore) (billissuer BillIssuerCore, err error)
	GetBillIssuerById(id int) (billisser BillIssuerCore, err error)
	GetBillIssuerByEmail(email string) (bool, error)
	UpdateBillIssuer(data BillIssuerCore) error
	GetAllBillIssuer(BillIssuerCore) (billissuers []BillIssuerCore, err error)
}
