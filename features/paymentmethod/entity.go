package paymentmethod

import "time"

type PaymentMethodCore struct {
	ID        uint
	Name      string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Untuk layer business
type Business interface {
	CreatePaymentMethod(data PaymentMethodCore) (err error)
}

// Untuk layer data / repository
type Data interface {
	CreatePaymentMethod(data PaymentMethodCore) (err error)
}
