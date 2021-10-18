package domain

type OrderItem struct {
	ID        int     `json:"id" form:"id"`
	OrderID   string  `json:"order_id" form:"order_id"`
	ProdukID  string  `json:"produk_id" form:"produk_id"`
	VariasiID string  `json:"variasi_id" form:"variasi_id"`
	Harga     float32 `json:"harga" form:"harga"`
	Quantity  int     `json:"quantity" form:"quantity"`
	SubTotal  float32 `json:"sub_total" form:"sub_total"`
	CreatedAt string  `json:"created_at" form:"created_at"`
	UpdatedAt string  `json:"updated_at" form:"updated_at"`
}

type OrderItemRepository interface {
	Create(req *OrderItem) (*OrderItem, error)
}
