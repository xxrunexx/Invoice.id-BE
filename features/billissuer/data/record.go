package data

import "gorm.io/gorm"

type BillIssuer struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func toBillIssuerRecord(billissuer billissuer.BillIssuerCore) BillIssuer {
	return BillIssuer{
		Model: gorm.Model{
			ID: billissuer.ID
		},
		Username: billissuer.Username,
		Password: billissuer.Password,
		Email: billissuer.Email
	}
}

func toBillIssuerCore(bi BillIssuer) billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore {
		ID: bi.ID,
		Username: bi.Username,
		Password: bi.Password,
		Email: bi.Email
	}
}