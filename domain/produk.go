package domain

type Produk struct {
	ID          string `json:"id,omitempty" form:"id"`
	Nama        string `json:"nama,omitempty" form:"nama"`
	Slug        string `json:"slug,omitempty" form:"slug"`
	Deskripsi   string `json:"deskripsi,omitempty" form:"deskripsi"`
	Spesifikasi string `json:"spesifikasi,omitempty" form:"spesifikasi"`
	KategoriID  int    `json:"kategori_id,omitempty" form:"kategori_id"`
	Dilihat     int64  `json:"dilihat" form:"dilihat"`
	CreatedAt   string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty" form:"updated_at"`
}

type ProdukDetailed struct {
	ID          string              `json:"id,omitempty" form:"id"`
	Nama        string              `json:"nama,omitempty" form:"nama"`
	Slug        string              `json:"slug,omitempty" form:"slug"`
	Deskripsi   string              `json:"deskripsi,omitempty" form:"deskripsi"`
	Spesifikasi []ProdukSpesifikasi `json:"spesifikasi,omitempty" form:"spesifikasi"`
	Kategori    Kategori            `json:"kategori,omitempty" form:"kategori"`
	Variasi     []ProdukVariasi     `json:"variasi,omitempty" form:"variasi"`
	Dilihat     int64               `json:"dilihat" form:"dilihat"`
	CreatedAt   string              `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt   string              `json:"updated_at,omitempty" form:"updated_at"`
}

type ProdukDetailedTemp struct {
	ID          string   `json:"id,omitempty" form:"id"`
	Nama        string   `json:"nama,omitempty" form:"nama"`
	Slug        string   `json:"slug,omitempty" form:"slug"`
	Deskripsi   string   `json:"deskripsi,omitempty" form:"deskripsi"`
	Spesifikasi string   `json:"spesifikasi,omitempty" form:"spesifikasi"`
	Kategori    Kategori `json:"kategori,omitempty" form:"kategori"`
	Dilihat     int64    `json:"dilihat" form:"dilihat"`
	CreatedAt   string   `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt   string   `json:"updated_at,omitempty" form:"updated_at"`
}

type ProdukSpesifikasi struct {
	Nama  string `json:"nama" form:"nama"`
	Value string `json:"value" form:"value"`
}

type ProdukFilter struct {
	Nama       string `json:"nama" form:"nama"`
	KategoriID string `json:"kategori_id" form:"kategori_id"`
}

type ProdukService interface {
	//Untuk Admin
	Create(req *Produk) *Response
	Update(req *Produk) *Response
	Delete(id string) *Response

	//Untuk Umum
	Read(filter *ProdukFilter) *Response
	Detail(id string) *Response
	DetailBySlug(slug string) *Response
}

type ProdukRepository interface {
	Create(req *Produk) (*ProdukDetailed, error)
	Read(filter *ProdukFilter) ([]ProdukDetailed, error)
	ByID(id string) (*ProdukDetailed, error)
	BySlugs(slug string) (*ProdukDetailed, error)
	Update(req *Produk) (*Produk, error)
	Delete(id string) error
}
