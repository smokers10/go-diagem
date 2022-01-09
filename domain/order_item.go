package domain

import "database/sql"

type OrderItem struct {
	ID              int    `json:"id,omitempty" form:"id"`
	OrderCheckoutID string `json:"order_id,omitempty" form:"order_id"`
	ProdukID        string `json:"produk_id,omitempty" form:"produk_id"`
	VariasiID       string `json:"variasi_id,omitempty" form:"variasi_id"`
	Quantity        int    `json:"quantity,omitempty" form:"quantity"`
}

type OrderItemDetail struct {
	ID              int            `json:"id,omitempty"`
	OrderCheckoutID string         `json:"order_id,omitempty"`
	ProdukID        string         `json:"produk_id,omitempty"`
	VariasiID       string         `json:"variasi_id,omitempty"`
	Quantity        int            `json:"quantity,omitempty"`
	Produk          ProdukDetailed `json:"produk,omitempty"`
	Variasi         ProdukVariasi  `json:"variasi,omitempty"`
}

type OrderItemRepository interface {
	Create(req *OrderItem, tx *sql.Tx) error
	ByOrderID(orderID string) ([]OrderItemDetail, error)
}
