package domain

// khusus umum
type Cart struct {
	ID        int    `json:"id" form:"id"`
	UserID    int    `json:"user_id" form:"user_id"`
	ProdukID  string `json:"produk_id" form:"produk_id"`
	VariasiID string `json:"variasi_id" form:"variasi_id"`
	Quantity  int    `json:"quantity" form:"quantity"`
	CreatedAt string `json:"created_at" form:"created_at"`
	UpdatedAt string `json:"updated_at" form:"updated_at"`
}

type CartData struct {
	ID        int    `json:"id" form:"id"`
	ProdukID  string `json:"produk_id" form:"produk_id"`
	VariasiID string `json:"variasi_id" form:"variasi_id"`
	UserID    int    `json:"user_id" form:"user_id"`
	Quantity  int    `json:"quantity" form:"quantity"`
}

type CartDetail struct {
	ID        int           `json:"id" form:"id"`
	UserID    int           `json:"user_id" form:"user_id"`
	ProdukID  string        `json:"produk_id" form:"produk_id"`
	VariasiID string        `json:"variasi_id" form:"variasi_id"`
	Produk    Produk        `json:"produk" form:"produk"`
	Variasi   ProdukVariasi `json:"variasi" form:"variasi"`
	SubTotal  float32       `json:"sub_total" form:"sub_total"`
	Quantity  int           `json:"quantity" form:"quantity"`
}

// khusus untuk umum
type CartService interface {
	Read(userID int) *Response
	AddtoCart(req *Cart) *Response
	UpdateQuantity(req *CartData) *Response
	Delete(req *CartData) *Response
}

type CartRepository interface {
	Create(req *Cart) (*Cart, error)
	Read(userID int) ([]CartDetail, error)
	UpdateQuantity(req *CartData) (*CartData, error)
	Delete(req *CartData) error
}
