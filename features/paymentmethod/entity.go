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
	CreatePaymentMethod(data PaymentMethodCore) (paymentmethod PaymentMethodCore, err error)
	GetPaymentMethodById(id int) (paymentmethod PaymentMethodCore, err error)
	GetAllPaymentMethod(PaymentMethodCore) (paymentmethods []PaymentMethodCore, err error)
	UpdatePaymentMethod(data PaymentMethodCore) error
}

// Untuk layer data / repository
type Data interface {
	CreatePaymentMethod(data PaymentMethodCore) (paymentmethod PaymentMethodCore, err error)
	GetPaymentMethodById(id int) (paymentmethod PaymentMethodCore, err error)
	GetAllPaymentMethod(PaymentMethodCore) (paymentmethods []PaymentMethodCore, err error)
	UpdatePaymentMethod(data PaymentMethodCore) error
}
