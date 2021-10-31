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
	Delete(id string) *Response
	MakeUtama(id string, produkID string) *Response
}

type ProdukFotoRepository interface {
	ReadByID(produkID string) ([]ProdukFoto, error)
	Create(req *ProdukFoto) error
	Delete(id string) error
	UpdateIsUtamaToTrue(id string) error
	UpdateRestIsUtamaToFalse(id string, produkID string) error
}
