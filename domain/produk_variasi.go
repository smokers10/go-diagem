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

type ProdukVariasiService interface{}

type ProdukVariasiRepository interface{}
