package presentation

type BillIssuerHandler struct {
	billissuerBusiness billissuer.Business
}

func NewHandlerBillIssuer(billissuerBusiness billissuer.Business) *BillIssuerHandler {
	return &BillIssuerHandler{billissuerBusiness}
}
