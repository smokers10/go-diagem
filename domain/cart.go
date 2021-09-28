package domain

//khusus umum
type Cart struct {
	ID        int
	UserID    int
	ProdukID  string
	VariasiID string
	Harga     int
	Quantity  int
	CreatedAt string
	UpdatedAt string
}

type CartUpdateQuantityData struct {
	ID       int
	ProdukID int
	UserID   int
	Quantity int
}

//khusus untuk umum
type CartService interface {
	Read(userID int) *Response
	AddtoCart(req *Cart) *Response
	UpdateQuantity(req *CartUpdateQuantityData) *Response
	Delete(id int) *Response
}

type CartRepository interface {
	Create(req *Cart) *Cart
	Read(userID int) []Cart
	UpdateQuantity(req *CartUpdateQuantityData) *Response
	Delete(id int) *Cart
}
