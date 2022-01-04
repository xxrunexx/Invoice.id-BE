package billissuer

type BillIssuerCore struct {
	ID       uint
	Username string
	Password string
	Email    string
}

// Untuk layer data
type Data interface {
	CreateBillIssuer(data BillIssuerCore) (id int, err error)
	LoginBillIssuer(BillIssuerCore) (billissuer BillIssuerCore, err error)
}

// Untuk layer business
type Business interface {
	// CreateBillIssuer(biData BillIssuerCore) (err error)
	CreateBillIssuer(data BillIssuerCore) (err error)
	LoginBillIssuer(BillIssuerCore) (billissuer BillIssuerCore, err error)
}
