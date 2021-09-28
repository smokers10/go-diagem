package domain

type Order struct {
	ID           int
	UserID       int
	Status       string
	SubTotal     float32
	AlamatKirim  string
	BiayaKirim   string
	InvoiceNo    string
	TglTransaksi string
	FinalTotal   string
	CreatedAt    string
	UpdatedAt    string
}

type OrderService interface {
	//Khusu Admin
	Read() *Response

	//Khusus Umum
	MyOrder(userID int) *Response
}

type OrderRepository interface {
	Read() []Order
	MyOrder(UserID int) []Order
}
