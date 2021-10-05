package domain

type ProdukVariasi struct {
	ID        string
	Variant   string
	SKU       string
	Harga     string
	ProdukID  string
	Stok      int
	CreatedAt string
	UpdatedAt string
}

type ProdukVariasiService interface {
	Create(req *ProdukVariasi) *Response
	Update(req *ProdukVariasi) *Response
	Delete(produkID string, produkVariasiID string) *Response
}

type ProdukVariasiRepository interface {
	Create(req *ProdukVariasi) (*ProdukVariasi, error)
	Update(req *ProdukVariasi) (*ProdukVariasi, error)
	Delete(produkID string, produkVariasiID string) error
}
