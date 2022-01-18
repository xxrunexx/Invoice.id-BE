package business

import (
	bi_m "invoice-api/features/billissuer/mocks"
	"invoice-api/features/billissuerdetail"
	bid_m "invoice-api/features/billissuerdetail/mocks"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockbidData              bid_m.Data
	mockbiData               bi_m.Data
	billissuerdetailBusiness billissuerdetail.Business
	billissuerdetailData     billissuerdetail.BillIssuerDetailCore
)

func TestMain(m *testing.M) {
	billissuerdetailBusiness = NewBusinessBillIssuerDetail(&mockbidData, &mockbiData)

	billissuerdetailData = billissuerdetail.BillIssuerDetailCore{
		BillIssuerID:   1,
		CompanyName:    "Alterra",
		CompanyAddress: "Malang",
		CompanyPhone:   "08128080977",
		CompanySite:    "www.alta.id",
	}

	os.Exit(m.Run())
}

func TestCreateBillIssuerDetail(t *testing.T) {
	t.Run("create bill issuer detail - success", func(t *testing.T) {
		mockbiData.On("GetBillIssuerById", mock.AnythingOfType("int")).Return(false, nil).Once()
		mockbidData.On("CreateBillIssuerDetail", mock.AnythingOfType("billissuerdetail.BillIssuerDetailCore")).Return(nil).Once()
		err := billissuerdetailBusiness.CreateBillIssuerDetail(billissuerdetailData)
		assert.Nil(t, err)
	})

	// t.Run("Create bill issuer error - invalid email", func(t *testing.T) {
	// 	err := billissuerBusiness.CreateBillIssuer(billissuer.BillIssuerCore{
	// 		Email: "pastierror",
	// 	})
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, err.Error(), "bad request")
	// })

	// t.Run("Create bill issuer error - GetBillIssuerByEmail", func(t *testing.T) {
	// 	mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, errors.New("error")).Once()
	// 	err := billissuerBusiness.CreateBillIssuer(billissuerData)
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Create bill issuer error - email exist", func(t *testing.T) {
	// 	mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(true, nil).Once()
	// 	err := billissuerBusiness.CreateBillIssuer(billissuerData)
	// 	assert.NotNil(t, err)
	// })

	// t.Run("Create bill issuer - error insert data", func(t *testing.T) {
	// 	mockData.On("GetBillIssuerByEmail", mock.AnythingOfType("string")).Return(false, nil).Once()
	// 	mockData.On("CreateBillIssuer", mock.AnythingOfType("billissuer.BillIssuerCore")).Return(errors.New("error")).Once()
	// 	err := billissuerBusiness.CreateBillIssuer(billissuerData)
	// 	assert.NotNil(t, err)
	// })
}

func TestGetBillIssuerDetailById(t *testing.T) {

}
