package domain

type Kategori struct {
	ID        int
	ParentID  int
	Nama      string
	Slug      string
	Icon      string
	Cover     string
	Thumbnail string
	RGT       int
	LFT       int
	CreatedAt string
	UpdatedAt string
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
	Create(req *Kategori) *Kategori
	Read() []Kategori
	Update(req *Kategori) *Kategori
	Delete(id int) *Kategori
}
