package domain

type ProdukFoto struct {
	ID        string
	ProdukID  string
	Path      string
	IsUtama   string
	CreatedAt string
	UpdatedAt string
}

type ProdukFotoService interface {
	Create(req *ProdukFoto) *Response
	Delete(produkID string, produkFotoID string) *Response
}

type ProdukFotoRepository interface {
	Create(req *ProdukFoto) (*ProdukFoto, error)
	Delete(produkID string, produkFotoID string) error
}
