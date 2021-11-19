package domain

type ProdukVariasi struct {
	ID        string  `json:"id,omitempty" form:"id"`
	Variant   string  `json:"variant,omitempty" form:"variant"`
	SKU       string  `json:"sku,omitempty" form:"sku"`
	Harga     float32 `json:"harga,omitempty" form:"harga"`
	ProdukID  string  `json:"produk_id,omitempty" form:"produk_id"`
	Stok      int     `json:"stok,omitempty" form:"stok"`
	CreatedAt string  `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt string  `json:"updated_at,omitempty" form:"updated_at"`
}

type ProdukVariasiService interface {
	Create(req *ProdukVariasi) *Response
	Update(req *ProdukVariasi) *Response
	Delete(produkID string, produkVariasiID string) *Response
}

type ProdukVariasiRepository interface {
	ReadByProdukID(produkID string) ([]ProdukVariasi, error)
	Create(req *ProdukVariasi) (*ProdukVariasi, error)
	Update(req *ProdukVariasi) (*ProdukVariasi, error)
	Delete(produkID string, produkVariasiID string) error
}
