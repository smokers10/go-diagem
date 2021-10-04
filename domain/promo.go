package domain

type Promo struct {
	ID         string `json:"id,omitempty" form:"id"`
	Judul      string `json:"judul,omitempty" form:"judul"`
	Slug       string `json:"slug,omitempty" form:"slug"`
	Deskripsi  string `json:"deskripsi,omitempty" form:"deskripsi"`
	Image      string `json:"image,omitempty" form:"image"`
	IsActive   bool   `json:"is_active,omitempty" form:"is_active"`
	IsFeatured bool   `json:"is_featured,omitempty" form:"is_featured"`
	TglMulai   string `json:"tgl_mulai,omitempty" form:"tgl_mulai"`
	TglSelesai string `json:"tgl_selesai,omitempty" form:"tgl_selesai"`
	CreatedAt  string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty" form:"updated_at"`
}

type PromoService interface {
	//Untuk Admin
	Create(req *Promo) *Response
	Update(req *Promo) *Response
	Delete(id string) *Response
	ReadOnlyByAdmin() *Response

	//untuk umum
	Read() *Response
	Detail(id string) *Response
}

type PromoRepository interface {
	//Untuk Admin
	Create(req *Promo) (*Promo, error)
	Update(req *Promo) (*Promo, error)
	Delete(id string) error
	ReadOnlyByAdmin() ([]Promo, error)

	// khusus umum
	Read() ([]Promo, error)
	ByID(id string) (*Promo, error)
	BySlug(slug string) (*Promo, error)
}
