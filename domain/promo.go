package domain

type Promo struct {
	ID           string      `json:"id,omitempty" form:"id"`
	Judul        string      `json:"judul,omitempty" form:"judul"`
	Slug         string      `json:"slug,omitempty" form:"slug"`
	Deskripsi    string      `json:"deskripsi,omitempty" form:"deskripsi"`
	Image        string      `json:"image,omitempty" form:"image"`
	IsPublish    bool        `json:"is_publish,omitempty" form:"is_publish"`
	IsFeatured   bool        `json:"is_featured,omitempty" form:"is_featured"`
	SEOKeyword   string      `json:"seo_keyword,omitempty" form:"seo_keyword"`
	SEOTags      string      `json:"seo_tags,omitempty" form:"seo_tags"`
	SEODeskripsi string      `json:"seo_deskripsi,omitempty" form:"seo_deskripsi"`
	TglMulai     string      `json:"tgl_mulai,omitempty" form:"tgl_mulai"`
	TglSelesai   string      `json:"tgl_selesai,omitempty" form:"tgl_selesai"`
	CreatedAt    string      `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt    string      `json:"updated_at,omitempty" form:"updated_at"`
	Sampul       promoSampul `json:"sampul,omitempty" form:"sampul"`
}

type promoSampul struct {
	Base64 string `json:"base64" form:"base64"`
	Format string `form:"format" json:"format"`
}

type PromoService interface {
	// Untuk Admin
	Create(req *Promo) *Response
	Update(req *Promo) *Response
	Delete(id string) *Response
	UpdateCover(req *Promo) *Response

	// untuk umum
	Read() *Response
	Detail(id string) *Response
}

type PromoRepository interface {
	// Untuk Admin
	Create(req *Promo) (*Promo, error)
	Update(req *Promo) (*Promo, error)
	Delete(id string) error

	// khusus umum
	Read() ([]Promo, error)
	ByID(id string) (*Promo, error)
	BySlug(slug string) (*Promo, error)
}
