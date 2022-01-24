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
			Name:     "Raviy Setiaji",
			Password: "admonadmonadmon",
			Email:    "rasyiajaja@gmail.com",
		},
		{
			ID:       2,
			Name:     "Dyah Ayu",
			Password: "newpassword",
			Email:    "dyah.ayu@gmail.com",
		},
	}

	billissuerData = billissuer.BillIssuerCore{
		Name:     "Raviy Setiaji",
		Password: "admonadmonadmon",
		Email:    "rasyiajaja@gmail.com",
	}

	billissuerlogin = billissuer.BillIssuerCore{
		Email:    "rasyiajaja@gmail.com",
		Password: "admonadmonadmon",
	}

	os.Exit(m.Run())
}

func TestCreateBillIssuer(t *testing.T) {
	t.Run("Create bill issuer - success", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		mockData.On("CreateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuer.BillIssuerCore{}, nil).Once()
		resp, err := billissuerBusiness.CreateBillIssuer(billissuerData)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Create bill issuer error - invalid email", func(t *testing.T) {
		resp, err := billissuerBusiness.CreateBillIssuer(billissuer.BillIssuerCore{
			Email: "pastierror",
		})
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, err.Error(), "bad request")
	})

	t.Run("Create bill issuer error - GetBillIssuerByEmail", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		resp, err := billissuerBusiness.CreateBillIssuer(billissuerData)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Create bill issuer error - email exist", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(true, nil).Once()
		resp, err := billissuerBusiness.CreateBillIssuer(billissuerData)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("Create bill issuer - error insert data", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		mockData.On("CreateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuer.BillIssuerCore{}, errors.New("error")).Once()
		resp, err := billissuerBusiness.CreateBillIssuer(billissuerData)
		assert.NotNil(t, err)
		assert.NotNil(t, resp)
	})
}

func TestLoginBillIssuer(t *testing.T) {
	t.Run("login bill issuer success", func(t *testing.T) {
		mockData.On("LoginBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuerData, nil).Once()
		resp, err := billissuerBusiness.LoginBillIssuer(billissuerlogin)
		assert.Equal(t, billissuerData.Email, resp.Email)
		assert.Nil(t, err)
	})

	t.Run("error login bill issuer", func(t *testing.T) {
		mockData.On("LoginBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuer.BillIssuerCore{}, errors.New("error")).Once()
		resp, err := billissuerBusiness.LoginBillIssuer(billissuer.BillIssuerCore{
			Email:    "fakeemail",
			Password: "capstonealterra",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "error")
		assert.Empty(t, resp.Email)
	})

	t.Run("Login error check bill issuer", func(t *testing.T) {
		mockData.On("LoginBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuer.BillIssuerCore{}, errors.New("error check data")).Once()
		resp, err := billissuerBusiness.LoginBillIssuer(billissuerlogin)
		assert.Equal(t, "error check data", err.Error())
		assert.NotNil(t, err)
		assert.Empty(t, resp.ID)
	})
}

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
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		mockData.On("UpdateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(nil).Once()
		err := billissuerBusiness.UpdateBillIssuer(billissuer.BillIssuerCore{})
		assert.NotNil(t, err)
	})

	t.Run("Update bill issuer error - invalid email", func(t *testing.T) {
		err := billissuerBusiness.UpdateBillIssuer(billissuer.BillIssuerCore{
			Email: "pastierror",
		})
		assert.NotNil(t, err)
		assert.Equal(t, err.Error(), "invalid data")
	})

	t.Run("Update bill issuer error - GetBillIssuerByEmail", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
		err := billissuerBusiness.UpdateBillIssuer(billissuerData)
		assert.Nil(t, err)
	})

	t.Run("Update bill issuer error - email exist", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(true, nil).Once()
		err := billissuerBusiness.UpdateBillIssuer(billissuerData)
		assert.NotNil(t, err)
	})

	t.Run("Update bill issuer - error insert data", func(t *testing.T) {
		mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
		mockData.On("UpdateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(errors.New("error")).Once()
		err := billissuerBusiness.UpdateBillIssuer(billissuerData)
		assert.NotNil(t, err)
	})
}

func TestGetAllBillIssuer(t *testing.T) {
	t.Run("validate get bill issuers", func(t *testing.T) {
		mockData.On("GetAllBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(billissuerDatas, nil).Once()
		resp, err := billissuerBusiness.GetAllBillIssuer(billissuer.BillIssuerCore{})
		assert.Nil(t, err)
		assert.Equal(t, len(resp), 2)
	})

	t.Run("error get bill issuers", func(t *testing.T) {
		mockData.On("GetAllBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(nil, errors.New("error"))
		resp, err := billissuerBusiness.GetAllBillIssuer(billissuer.BillIssuerCore{})
		assert.NotNil(t, err)
		assert.Nil(t, resp)
	})
}
