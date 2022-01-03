package billissuer

type BillIssuerCore struct {
	ID       uint
	Username string
	Password string
	Email    string
}

// Untuk layer data
type Data interface {
	CreateBillIssuer(biData BillIssuerCore) (err error)
}

// Untuk layer business
type Business interface {
	CreateBillIssuer(biData BillIssuerCore) (err error)
}
