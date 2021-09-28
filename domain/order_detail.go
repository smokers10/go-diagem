package domain

type OrderDetail struct {
	ID        int
	OrderID   string
	BisnisID  string
	ProdukID  string
	VariasiID string
	Harga     float32
	Qty       int
	SubTotal  float32
	CouponID  string
	CreatedAt string
	UpdatedAt string
}

type OrderDetailService interface{}

type OrderDetailRepository interface{}
