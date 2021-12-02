package domain

type Kategori struct {
	ID         string        `json:"id,omitempty" form:"id"`
	Nama       string        `json:"nama,omitempty" form:"nama"`
	Slug       string        `json:"slug,omitempty" form:"slug"`
	Cover      string        `json:"cover,omitempty" form:"cover"`
	CoverImage kategoriCover `json:"cover_image,omitempty" form:"cover_image"`
	CreatedAt  string        `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  string        `json:"updated_at,omitempty" form:"updated_at"`
}

type kategoriCover struct {
	Base64 string `json:"base64" form:"base64"`
	Format string `json:"format" form:"format"`
}

type KategoriService interface {
	//Khusus Admin
	Create(req *Kategori) *Response
	Update(req *Kategori) *Response
	Delete(id string) *Response
	UpdateCover(req *Kategori) *Response
	//Untuk Umum
	Read() *Response
	Detail(id string) *Response
}

type KategoriRepository interface {
	// Untuk Admin
	Create(req *Kategori) (*Kategori, error)
	Update(req *Kategori) (*Kategori, error)
	Delete(id string) error

	// Untuk Umum
	Read() ([]Kategori, error)
	Detail(id string) (*Kategori, error)
}
