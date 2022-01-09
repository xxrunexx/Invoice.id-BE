package billissuer

type BillIssuerCore struct {
	ID       uint
	Username string
	Password string
	Email    string
	Token    string
}

// Untuk layer business
type Business interface {
	// CreateBillIssuer(biData BillIssuerCore) (err error)
	CreateBillIssuer(data BillIssuerCore) (err error)
	LoginBillIssuer(BillIssuerCore) (billissuer BillIssuerCore, err error)
	GetBillIssuerById(id int) (billissuer BillIssuerCore, err error)
}

// Untuk layer data / repository
type Data interface {
	CreateBillIssuer(data BillIssuerCore) (id int, err error)
	LoginBillIssuer(BillIssuerCore) (billissuer BillIssuerCore, err error)
	GetBillIssuerById(id int) (billisser BillIssuerCore, err error)
	GetBillIssuerByEmail(email string) (bool, error)
}
