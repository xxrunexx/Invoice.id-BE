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
	CreateClient(data ClientCore) (client ClientCore, err error)
	GetAllClient(ClientCore) (clients []ClientCore, err error)
	GetClientById(id int) (client ClientCore, err error)
	UpdateClient(data ClientCore) error
}

// Untuk layer data / repository
type Data interface {
	CreateClient(data ClientCore) (client ClientCore, err error)
	GetAllClient(ClientCore) (clients []ClientCore, err error)
	GetClientById(id int) (client ClientCore, err error)
	UpdateClient(data ClientCore) error
	GetClientByNik(nik int) (bool, error)
}
