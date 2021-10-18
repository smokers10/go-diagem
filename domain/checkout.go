package domain

type Checkout struct {
	UserID   int `json:"user_id" form:"user_id"`
	AlamatID int `json:"alamat_id" form:"alamat_id"`
}

type CheckoutService interface {
	Checkout(req *Checkout) *Response
}
