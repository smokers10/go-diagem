package domain

import "database/sql"

type Order struct {
	ID         string `json:"id,omitempty" form:"id"`
	Status     string `json:"status,omitempty" form:"status"`
	UserID     int    `json:"user_id,omitempty" form:"user_id"`
	AlamatID   int    `json:"alamat_id,omitempty" form:"alamat_id"`
	Kurir      string `json:"kurir,omitempty" form:"kurir"`
	PaketKurir string `json:"paket_kurir,omitempty" form:"paket_kurir"`
	Ongkir     int    `json:"ongkir,omitempty" form:"ongkir"`
	Total      int    `json:"total,omitempty" form:"total"`
	InvoiceNo  string `json:"invoice_no,omitempty" form:"invoice_no"`
	TglOrder   string `json:"tgl_order,omitempty" form:"tgl_order"`
}

type OrderDetail struct {
	ID         string            `json:"id,omitempty" form:"id"`
	Status     string            `json:"status,omitempty" form:"status"`
	UserID     int               `json:"user_id,omitempty" form:"user_id"`
	AlamatID   int               `json:"alamat_id,omitempty" form:"alamat_id"`
	Kurir      string            `json:"kurir,omitempty" form:"kurir"`
	PaketKurir string            `json:"paket_kurir,omitempty" form:"paket_kurir"`
	Ongkir     string            `json:"ongkir,omitempty" form:"ongkir"`
	InvoiceNo  string            `json:"invoice_no,omitempty" form:"invoice_no"`
	TglOrder   string            `json:"tgl_order,omitempty" form:"tgl_order"`
	OrderBayar OrderBayar        `json:"order_bayar,omitempty" form:"order_bayar"`
	OrderItem  []OrderItemDetail `json:"order_item,omitempty" form:"order_item"`
	Alamat     Alamat            `json:"alamat,omitempty" form:"alamat"`
	User       User              `json:"user,omitempty" form:"user"`
}

type OrderService interface {
	Create(req *Order) *Response
	ReadAll() *Response
	Detail(orderID string) *Response
	UpdateSR(orderID string, status string, resi string) *Response
	ReadByUser(userID int) *Response
}

type OrderRepository interface {
	Create(req *Order, tx *sql.Tx) error
	UpdateSR(orderID string, status string, resi string) error
	GetByID(orderID string) (*OrderDetail, error)
	GetByUserID(userID int) ([]OrderDetail, error)
	Read() ([]OrderDetail, error)
	GetSQLInstance() *sql.DB
}
