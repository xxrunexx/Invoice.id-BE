package business

import (
	"errors"
	"invoice-api/features/billissuer"
	"invoice-api/features/billissuer/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockData           mocks.Data
	billissuerBusiness billissuer.Business
	billissuerDatas    []billissuer.BillIssuerCore
	billissuerData     billissuer.BillIssuerCore
	billissuerlogin    billissuer.BillIssuerCore
)

func TestMain(m *testing.M) {
	billissuerBusiness = NewBusinessBillIssuer(&mockData)

	billissuerDatas = []billissuer.BillIssuerCore{
		{
			ID:       1,
			Username: "Harun Ar Rasyid",
			Password: "adminadminadmin",
			Email:    "rasyid.id3@gmail.com",
		},
	}

	billissuerData = billissuer.BillIssuerCore{
		Username: "Raviy Setiaji",
		Password: "iniadalahsebuahtesting",
		Email:    "testingtest@gmail.com",
	}

	billissuerlogin = billissuer.BillIssuerCore{
		Username: "Raviy Setiaji",
		Password: "iniadalahsebuahtesting",
	}

	os.Exit(m.Run())
}

func TestCreateBillIssuer(t *testing.T) {
	t.Run("validate create bill issuer", func(t *testing.T) {
		mockData.On("CreateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(nil).Once()
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		err := billissuerBusiness.CreateBillIssuer(billissuer.BillIssuerCore{})
		assert.NotNil(t, err)
	})

	t.Run("error invalid email", func(t *testing.T) {
		err := billissuerBusiness.CreateBillIssuer(billissuer.BillIssuerCore{
			Email: "pastierror",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "bad request")
	})

	t.Run("error GetBillIssuerByEmail", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := billissuerBusiness.CreateBillIssuer(billissuerData)
		assert.NotNil(t, err)
	})

	t.Run("error create bill issuer", func(t *testing.T) {
		mockData.On("CreateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(errors.New("error")).Once()
		err := billissuerBusiness.CreateBillIssuer(billissuer.BillIssuerCore{})
		assert.NotNil(t, err)
	})
}

// func TestLoginBillIssuer(t *testing.T) {
// 	t.Run("validate login bill issuer", func(t *testing.T) {
// 		mockData.On("LoginBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuer.BillIssuerCore{}, nil).Once()
// 		resp, err := billissuerBusiness.LoginBillIssuer(billissuer.BillIssuerCore{})
// 		assert.Nil(t, err)
// 		assert.NotNil(t, resp)
// 	})
// 	t.Run("error login bill issuer", func(t *testing.T) {
// 		mockData.On("LoginBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuer.BillIssuerCore{}, errors.New("error")).Once()
// 		resp, err := billissuerBusiness.LoginBillIssuer(billissuer.BillIssuerCore{})
// 		assert.NotNil(t, err)
// 		assert.Nil(t, 0, int(resp.ID))
// 	})
// }

func TestGetBillIssuerById(t *testing.T) {
	t.Run("validate get bill issuer by id", func(t *testing.T) {
		mockData.On("GetBillIssuerById", mock.AnythingOfType("int")).Return(billissuer.BillIssuerCore{}, nil).Once()
		resp, err := billissuerBusiness.GetBillIssuerById(3)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})
	t.Run("error get bill issuer by id", func(t *testing.T) {
		mockData.On("GetBillIssuerById", mock.AnythingOfType("int")).Return(billissuer.BillIssuerCore{}, errors.New("error")).Once()
		resp, err := billissuerBusiness.GetBillIssuerById(3)
		assert.NotNil(t, err)
		assert.Equal(t, 0, int(resp.ID))
	})
}

func TestUpdateBIllIssuer(t *testing.T) {
	t.Run("validate update bill issuer", func(t *testing.T) {
		mockData.On("UpdateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(nil).Once()
		err := billissuerBusiness.UpdateBillIssuer(billissuer.BillIssuerCore{})
		assert.NotNil(t, err)
	})
	t.Run("error update bill issuer", func(t *testing.T) {
		mockData.On("UpdateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(errors.New("error")).Once()
		err := billissuerBusiness.UpdateBillIssuer(billissuer.BillIssuerCore{})
		assert.NotNil(t, err)
	})
}
