package domain

type OrderBayar struct {
	ID         int     `json:"id" form:"id"`
	OrderID    string  `json:"order_id" form:"order_id"`
	RekeningID int     `json:"rekening_id" form:"rekening_id"`
	Method     string  `json:"method" form:"method"`
	Layanan    string  `json:"layanan" form:"layanan"`
	Jumlah     float32 `json:"jumlah" form:"jumlah"`
	Status     string  `json:"status" form:"status"`
	Tglbayar   string  `json:"tgl_bayar" form:"tgl_bayar"`
	CreatedAt  string  `json:"created_at" form:"created_at"`
	UpdatedAt  string  `json:"updated_at" form:"updated_at"`
}

type OrderBayarService interface {
	Create(req *OrderBayar) *Response
}

type OrderBayarRepository interface {
	Create(req *OrderBayar) (*OrderBayar, error)
	UpdateStatus(orderID string, status string) error
}
