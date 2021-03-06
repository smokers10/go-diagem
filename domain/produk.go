package domain

type Produk struct {
	ID              string           `json:"id,omitempty" form:"id"`
	Nama            string           `json:"nama,omitempty" form:"nama"`
	Slug            string           `json:"slug,omitempty" form:"slug"`
	Deskripsi       string           `json:"deskripsi,omitempty" form:"deskripsi"`
	Spesifikasi     string           `json:"spesifikasi,omitempty" form:"spesifikasi"`
	KategoriID      string           `json:"kategori_id,omitempty" form:"kategori_id"`
	Dilihat         int64            `json:"dilihat" form:"dilihat"`
	Berat           int              `json:"berat,omitempty" form:"berat"`
	SatuanBerat     string           `json:"satuan_berat,omitempty" form:"satuan_berat"`
	Lebar           float32          `json:"lebar,omitempty" form:"lebar"`
	Panjang         float32          `json:"panjang,omitempty" form:"panjang"`
	Tinggi          float32          `json:"tinggi,omitempty" form:"tinggi"`
	Kode            string           `json:"kode,omitempty" form:"kode"`
	Harga           int              `json:"harga,omitempty" form:"harga"`
	Stok            int              `json:"stok,omitempty" form:"stok"`
	Variasi         []ProdukVariasi  `json:"variasi,omitempty" form:"variasi"`
	ProdukFoto      []FotoProdukFile `json:"produk_foto,omitempty" form:"produk_foto"`
	JumlahPenjualan int              `json:"jumlah_penjualan,omitempty" form:"jumlah_penjualan"`
	IsHasVariant    bool             `json:"is_has_variant,omitempty" form:"is_has_varian"`
	Discount        int              `json:"discount,omitempty" form:"discount"`
	CreatedAt       string           `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt       string           `json:"updated_at,omitempty" form:"updated_at"`
}

type ProdukDetailed struct {
	ID               string              `json:"id,omitempty" form:"id"`
	Nama             string              `json:"nama,omitempty" form:"nama"`
	Slug             string              `json:"slug,omitempty" form:"slug"`
	Deskripsi        string              `json:"deskripsi,omitempty" form:"deskripsi"`
	Spesifikasi      []ProdukSpesifikasi `json:"spesifikasi,omitempty" form:"spesifikasi"`
	SpesifikasiTemp  string              `json:"spesifikasi_temp,omitempty" form:"spesifikasi_temp"` // untuk menyimpan sementara data string spesifikasi dari DB sebelum di Unmarhsal.
	Kategori         Kategori            `json:"kategori,omitempty" form:"kategori"`
	Variasi          []ProdukVariasi     `json:"variasi,omitempty" form:"variasi"`
	ProdukFoto       []ProdukFoto        `json:"produk_foto,omitempty" form:"produk_foto"`               // menyimpan beberapa foto produk yang tersedia
	ProdukSingleFoto ProdukFoto          `json:"produk_single_foto,omitempty" form:"produk_single_foto"` // menyimpan satu foto utama saja
	Dilihat          int64               `json:"dilihat" form:"dilihat"`
	Berat            int                 `json:"berat,omitempty" form:"berat"`
	SatuanBerat      string              `json:"satuan_berat,omitempty" form:"satuan_berat"`
	Lebar            float32             `json:"lebar,omitempty" form:"lebar"`
	Panjang          float32             `json:"panjang,omitempty" form:"panjang"`
	Tinggi           float32             `json:"tinggi,omitempty" form:"tinggi"`
	Kode             string              `json:"kode,omitempty" form:"kode"`
	Harga            int                 `json:"harga,omitempty" form:"harga"`
	Stok             int                 `json:"stok,omitempty" form:"stok"`
	JumlahPenjualan  int                 `json:"jumlah_penjualan,omitempty" form:"jumlah_penjualan"`
	IsHasVariant     bool                `json:"is_has_variant,omitempty" form:"is_has_varian"`
	Discount         int                 `json:"discount,omitempty" form:"discount"`
	KategoriID       string              `json:"kategori_id,omitempty" form:"kategori_id"`
	CreatedAt        string              `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt        string              `json:"updated_at,omitempty" form:"updated_at"`
}

type ProdukSpesifikasi struct {
	Nama  string `json:"nama" form:"nama"`
	Value string `json:"value" form:"value"`
}

type ProdukFilter struct {
	Nama         string       `json:"nama" form:"nama"`
	KategoriID   string       `json:"kategori_id" form:"kategori_id"`
	ShortBy      string       `json:"short_by,omitempty" form:"short_by"`
	Pagination   pagination   `json:"pagination,omitempty" form:"pagination"`
	ClarifyOrder clarifyOrder `json:"order_by,omitempty" form:"order_by"`
}

type pagination struct {
	From  int `json:"from,omitempty" form:"from"`
	Until int `json:"until,omitempty" form:"until"`
}

type clarifyOrder struct {
	TableName   string
	OrderMethod string
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
	GetPopular() *Response
}

type ProdukRepository interface {
	Create(req *Produk) (*ProdukDetailed, error)
	Read(filter *ProdukFilter) ([]ProdukDetailed, error)
	ByID(id string) (*ProdukDetailed, error)
	ByIDSimplified(id string) (*ProdukDetailed, error)
	BySlugs(slug string) (*ProdukDetailed, error)
	Update(req *Produk) (*Produk, error)
	UpdateStok(produkID string, changeValue int, operationType string) error
	UpdateSellingPoint(produkID string, changeValue int) error
	Delete(id string) error
	GetPopular() ([]ProdukDetailed, error)
}
