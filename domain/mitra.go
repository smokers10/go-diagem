package domain

type Mitra struct {
	ID       int    `json:"id,omitempty" form:"id"`
	Nama     string `json:"nama,omitempty" form:"nama"`
	SellerID string `json:"seller_id,omitempty" form:"seller_id"`
	Kontak   string `json:"kontak,omitempty" form:"kontak"`
	Alamat   string `json:"alamat,omitempty" form:"alamat"`
	Email    string `json:"email,omitempty" form:"email"`
	Kota     string `json:"kota,omitempty" form:"kota"`
	KotaID   string `json:"kota_id,omitempty" form:"kota_id"`
}

type MitraService interface {
	//Untuk Admin
	Create(req *Mitra) *Response
	Update(req *Mitra) *Response
	Delete(id int) *Response

	//Untuk Umum
	Read() *Response
	ReadWithFilter(req *Mitra) *Response
	GetOne(id int) *Response
}

type MitraRepository interface {
	Create(req *Mitra) (*Mitra, error)
	Read() ([]Mitra, error)
	ByID(id int) (*Mitra, error)
	BySellerID(sellerId string) (*Mitra, error)
	ByEmail(email string) (*Mitra, error)
	Update(req *Mitra) (*Mitra, error)
	Delete(id int) error
	ReadWithFilter(req *Mitra) ([]Mitra, error)
}
