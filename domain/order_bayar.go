package domain

type OrderBayar struct {
	ID              int     `json:"id,omitempty" form:"id"`
	OrderCheckoutID string  `json:"order_checkout_id,omitempty" form:"order_checkout_id"`
	Jumlah          float32 `json:"jumlah,omitempty" form:"jumlah"`
	Status          string  `json:"status,omitempty" form:"status"`
	Token           string  `json:"token,omitempty" form:"token"`
	RedirectURL     string  `json:"redirect_url,omitempty" form:"redirect_url"`
	TGLBayar        string  `json:"tgl_bayar,omitempty" form:"tgl_bayar"`
}

type OrderBayarRepository interface {
	Create(req *OrderBayar) error
	UpdateStatus(token string, status string) error
	ByToken(token string) (*OrderBayar, error)
	ByOrderID(OrderID string) (*OrderBayar, error)
}
