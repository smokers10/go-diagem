package domain

type Kategori struct {
	ID        int    `json:"id,omitempty" form:"id"`
	Nama      string `json:"nama,omitempty" form:"nama"`
	Slug      string `json:"slug,omitempty" form:"slug"`
	Icon      string `json:"icon,omitempty" form:"icon"`
	Cover     string `json:"cover,omitempty" form:"cover"`
	Thumbnail string `json:"thumbnail,omitempty" form:"thumbnail"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" form:"updated_at"`
}

type KategoriService interface {
	//Khusus Admin
	Create(req *Kategori) *Response
	Update(req *Kategori) *Response
	Delete(id int) *Response

	//Untuk Umum
	Read() *Response
}

type KategoriRepository interface {
	Create(req *Kategori) (*Kategori, error)
	Read() ([]Kategori, error)
	Update(req *Kategori) (*Kategori, error)
	Delete(id int) error
}
