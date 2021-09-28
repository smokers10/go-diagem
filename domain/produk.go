package domain

type Produk struct {
	ID          string
	Nama        string
	Slug        string
	Deskripsi   string
	Spesifikasi string
	KategoriID  int
	HasVariasi  bool
	Var1Nama    string
	Var2Nama    string
	Var1Value   string
	Var2Value   string
	Berat       float32
	Beratsatuan float32
	Lebar       int
	Panjang     int
	Tinggi      int
	Dilihat     int64
	CreatedAt   string
	UpdatedAt   string
}

type ProdukFilter struct {
	Nama       string
	KategoriID int
}

type ProdukService interface {
	//Untuk Admin
	Create(req *Produk) *Response
	Read() *Response
	Detail() *Response
	Update(req *Produk) *Response
	Delete(req *Produk) *Response

	//Untuk Umum
	Catalogue(filter *ProdukFilter) *Response
}

type ProdukRepository interface {
	Create(req *Produk) *Produk
	Read() []Produk
	ByID(id string) *Produk
	Update(req *Produk) *Produk
	Delete(req *Produk) *Produk
}
