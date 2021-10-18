package domain

type Order struct {
	ID           string  `json:"id" form:"id"`
	UserID       int     `json:"user_id" form:"user_id"`
	AlamatID     int     `json:"alamat_id" form:"alamat_id"`
	Status       string  `json:"status" form:"status"`
	SubTotal     float32 `json:"sub_total" form:"sub_total"`
	BiayaKirim   string  `json:"biaya_kirim" form:"biaya_kirim"`
	InvoiceNo    string  `json:"invoice_no" form:"invoice_no"`
	TglTransaksi string  `json:"tgl_transaksi" form:"tgl_transaksi"`
	FinalTotal   string  `json:"final_total" form:"final_total"`
	CreatedAt    string  `json:"created_at" form:"created_at"`
	UpdatedAt    string  `json:"updated_at" form:"updated_at"`
}

type OrderDetail struct {
	ID           string      `json:"id" form:"id"`
	User         User        `json:"user" form:"user"`
	Alamat       Alamat      `json:"alamat" form:"alamat"`
	Item         []OrderItem `json:"item" from:"item"`
	Pembayaran   OrderItem   `json:"pembayaran" from:"pembayaran"`
	Status       string      `json:"status" form:"status"`
	SubTotal     float32     `json:"sub_total" form:"sub_total"`
	BiayaKirim   string      `json:"biaya_kirim" form:"biaya_kirim"`
	InvoiceNo    string      `json:"invoice_no" form:"invoice_no"`
	TglTransaksi string      `json:"tgl_transaksi" form:"tgl_transaksi"`
	FinalTotal   string      `json:"final_total" form:"final_total"`
	CreatedAt    string      `json:"created_at" form:"created_at"`
	UpdatedAt    string      `json:"updated_at" form:"updated_at"`
}

type OrderService interface {
	//Khusu Admin
	Read() *Response

	//Khusus Umum
	MyOrder(userID int) *Response
}

type OrderRepository interface {
	//khusus Admin
	Read() ([]Order, error)

	// khusus umum
	Create(req *Order) (*Order, error)
	GetByID(id int) (*Order, error)
	GetByUserID(UserID int) ([]Order, error)
}
