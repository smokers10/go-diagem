package domain

type ProdukFoto struct {
	ID        string `json:"id,omitempty" form:"id"`
	ProdukID  string `json:"produk_id,omitempty" form:"produk_id"`
	Path      string `json:"path,omitempty" form:"path"`
	IsUtama   bool   `json:"is_utama,omitempty" form:"is_utama"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" form:"updated_at"`
}

type FotoProdukFile struct {
	Base64   string `form:"base64" json:"base64"`
	Format   string `form:"format" json:"format"`
	ProdukID string `json:"produk_id,omitempty" form:"produk_id"`
}

type ProdukFotoService interface {
	Create(req *FotoProdukFile) *Response
	Delete(id string, path string) *Response
	MakeUtama(id string, produkID string) *Response
}

type ProdukFotoRepository interface {
	ReadByProdukID(produkID string) ([]ProdukFoto, error)
	GetUtamaOnly(produkID string) (*ProdukFoto, error)
	Create(req *ProdukFoto) error
	Delete(id string) error
	UpdateIsUtamaToTrue(id string) error
	UpdateRestIsUtamaToFalse(id string, produkID string) error
}
