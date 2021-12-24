package domain

type Order struct {
	ID         string `json:"id,omitempty" form:"id"`
	Status     string `json:"status,omitempty" form:"status"`
	UserID     int    `json:"user_id,omitempty" form:"user_id"`
	AlamatID   int    `json:"alamat_id,omitempty" form:"alamat_id"`
	Kurir      string `json:"kurir,omitempty" form:"kurir"`
	PaketKurir string `json:"paket_kurir,omitempty" form:"paket_kurir"`
	Ongkir     int    `json:"ongkir,omitempty" form:"ongkir"`
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
	OrderItem  []OrderItemDetail `json:"order_item,omitempty" form:"order_item"`
	Alamat     Alamat            `json:"alamat,omitempty" form:"alamat"`
	User       User              `json:"user,omitempty" form:"user"`
}

type OrderService interface {
	Create(req *Order) *Response
	ReadAll() *Response
	Detail(orderID string) *Response
	UpdateStatus(orderID string, status string) *Response
	ReadByUser(userID int) *Response
}

type OrderRepository interface {
	Create(req *Order) error
	UpdateStatus(orderID string, status string) error
	GetByID(orderID string) (*OrderDetail, error)
	GetByUserID(userID int) ([]Order, error)
	Read() ([]OrderDetail, error)
}
