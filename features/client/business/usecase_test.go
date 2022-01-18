package business

import (
	"errors"
	"invoice-api/features/client"
	"invoice-api/features/client/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData       mocks.Data
	clientBusiness client.Business
	clientDatas    []client.ClientCore
	clientData     client.ClientCore
)

func TestMain(m *testing.M) {
	clientBusiness = NewBusinessClient(&mockData)

	clientDatas = []client.ClientCore{
		{
			ID:      1,
			NIK:     3714051805000009,
			Name:    "Harun Ar Rasyid",
			Phone:   "08128080977",
			Address: "Depok",
			Email:   "rasyid.id3@gmail.com",
		},
	}

	clientData = client.ClientCore{
		NIK:     3714051805000009,
		Name:    "Harun Ar Rasyid",
		Phone:   "08128080977",
		Address: "Depok, Tangerang Selatan",
		Email:   "rasyid.id3@gmail.com",
	}

	os.Exit(m.Run())
}

func TestCreateClient(t *testing.T) {
	t.Run("create client - success", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(false, nil).Once()
		mockData.On("CreateClient", mock.AnythingOfType("client.ClientCore")).Return(nil).Once()
		// mockData.On("CreateClient", mock.AnythingOfType("client.ClientCore")).Return(nil).Once()
		err := clientBusiness.CreateClient(clientData)
		assert.Nil(t, err)
	})

	t.Run("Create client error - invalid NIK", func(t *testing.T) {
		err := clientBusiness.CreateClient(client.ClientCore{
			NIK: 0,
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid data")
	})

	t.Run("Create client error - GetBillIssuerByNik", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(false, errors.New("error")).Once()
		err := clientBusiness.CreateClient(clientData)
		assert.NotNil(t, err)
	})

	t.Run("Create client error - NIK exist", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(true, nil).Once()
		err := clientBusiness.CreateClient(clientData)
		assert.NotNil(t, err)
	})
}

func TestGetClientById(t *testing.T) {
	t.Run("Get client by id - success", func(t *testing.T) {
		mockData.On("GetClientById", mock.AnythingOfType("int")).Return(client.ClientCore{}, nil).Once()
		resp, err := clientBusiness.GetClientById(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("Get client by id - invalid data", func(t *testing.T) {
		mockData.On("GetClientById", mock.AnythingOfType("int")).Return(client.ClientCore{}, errors.New("error")).Once()
		resp, err := clientBusiness.GetClientById(3)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestGetAllClient(t *testing.T) {
	t.Run("validate get clients", func(t *testing.T) {
		mockData.On("GetAllClient", mock.AnythingOfType("client.ClientCore")).Return(clientDatas, nil).Once()
		resp, err := clientBusiness.GetAllClient(client.ClientCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 1)
	})

	t.Run("error get clients", func(t *testing.T) {
		mockData.On("GetAllClient", mock.AnythingOfType("client.ClientCore")).Return(nil, errors.New("error"))
		resp, err := clientBusiness.GetAllClient(client.ClientCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}

func Test(t *testing.T) {
	t.Run("Update client - success", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(false, nil).Once()
		mockData.On("UpdateClient", mock.AnythingOfType("client.ClientCore")).Return(nil).Once()
		err := clientBusiness.UpdateClient(client.ClientCore{})
		assert.NotNil(t, err)
	})

	t.Run("Update client error - invalid NIK", func(t *testing.T) {
		err := clientBusiness.UpdateClient(client.ClientCore{
			NIK: 0,
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid data")
	})

	t.Run("Update client error - GetClientByNik", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(false, errors.New("error")).Once()
		err := clientBusiness.UpdateClient(clientData)
		assert.Nil(t, err)
	})

	t.Run("Update client error - email exist", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(true, nil).Once()
		err := clientBusiness.UpdateClient(clientData)
		assert.NotNil(t, err)
	})

	t.Run("Update client - error insert data", func(t *testing.T) {
		mockData.On("GetClientByNik", mock.AnythingOfType("int")).Return(false, nil).Once()
		mockData.On("UpdateClient", mock.AnythingOfType("client.ClientCore")).Return(errors.New("error")).Once()
		err := clientBusiness.UpdateClient(clientData)
		assert.NotNil(t, err)
	})
}
