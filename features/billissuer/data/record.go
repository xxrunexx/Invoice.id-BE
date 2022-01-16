package data

import (
	"invoice-api/features/billissuer"

	"gorm.io/gorm"
)

type BillIssuer struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

func toBillIssuerRecord(bi billissuer.BillIssuerCore) BillIssuer {
	return BillIssuer{
		Model: gorm.Model{
			ID: bi.ID,
		},
		Username: bi.Username,
		Password: bi.Password,
		Email:    bi.Email,
	}
}

func toBillIssuerCore(bi BillIssuer) billissuer.BillIssuerCore {
	return billissuer.BillIssuerCore{
		ID:       bi.ID,
		Username: bi.Username,
		Password: bi.Password,
		Email:    bi.Email,
	}
}

// Get All / Array
// func toBillIssuerCoreList(biList []BillIssuer) []billissuer.BillIssuerCore {
// 	convBi := []billissuer.BillIssuerCore{}

// 	for _, billissuer := range biList {
// 		convBi = append(convBi, toBillIssuerCore(billissuer))
// 	}
// 	return convBi
// }
