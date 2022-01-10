package client

type ClientCore struct {
	ID      uint
	NIK     int
	Name    string
	Phone   string
	Address string
	Email   string
}

// Untuk layer business / service
type Business interface {
	CreateClient(data ClientCore) (err error)
}

// Untuk layer data / repository
type Data interface {
	CreateClient(data ClientCore) (err error)
}
