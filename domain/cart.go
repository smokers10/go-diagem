package domain

// khusus umum
type Cart struct {
	ID        int    `json:"id,omitempty" form:"id"`
	UserID    int    `json:"user_id,omitempty" form:"user_id"`
	ProdukID  string `json:"produk_id,omitempty" form:"produk_id"`
	VariasiID string `json:"variasi_id,omitempty" form:"variasi_id"`
	Quantity  int    `json:"quantity,omitempty" form:"quantity"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" form:"updated_at"`
}

type CartDetail struct {
	ID               int            `json:"id,omitempty" form:"id"`
	UserID           int            `json:"user_id,omitempty" form:"user_id"`
	ProdukID         string         `json:"produk_id,omitempty" form:"produk_id"`
	VariasiID        string         `json:"variasi_id,omitempty" form:"variasi_id"`
	Produk           ProdukDetailed `json:"produk,omitempty" form:"produk"`
	Variasi          ProdukVariasi  `json:"variasi,omitempty" form:"variasi"`
	ProdukSingleFoto ProdukFoto     `json:"produk_single_foto,omitempty" form:"produk_single_foto"`
	SubTotal         int            `json:"subtotal,omitempty" form:"subtotal"`
	Quantity         int            `json:"quantity,omitempty" form:"quantity"`
}

type CartCount struct {
	TotalCarts int `json:"total_carts,omitempty" form:"total_carts"`
}

// khusus untuk umum
type CartService interface {
	Read(userID int) *Response
	AddtoCart(req *Cart) *Response
	UpdateQuantity(req *Cart) *Response
	Delete(req *Cart) *Response
	Count(userID int) *Response
}

type CartRepository interface {
	Create(req *Cart) (*Cart, error)
	ByID(cartID int) (*Cart, error)
	Read(userID int) ([]CartDetail, error)
	UpdateQuantityVariant(req *Cart) (*Cart, error)
	Delete(req *Cart) error
	Count(userID int) (*CartCount, error)
}
