package domain

//kusus Umum
type Alamat struct {
	ID           int    `json:"id,omitempty" form:"id"`
	UserID       int    `json:"user_id,omitempty" form:"user_id"`
	Alamat       string `json:"alamat,omitempty" form:"alamat"`
	Nama         string `json:"nama,omitempty" form:"nama"`
	Penerima     string `json:"penerima,omitempty" form:"penerima"`
	Phone        string `json:"phone,omitempty" form:"phone"`
	KDPos        string `json:"kd_pos,omitempty" form:"kd_pos"`
	ProvinsiID   string `json:"provinsi_id,omitempty" form:"provinsi"`
	KotaID       string `json:"kota_id,omitempty" form:"kota_id"`
	NamaProvinsi string `json:"nama_provinsi,omitempty" form:"nama_provinsi"`
	NamaKota     string `json:"nama_kota,omitempty" form:"nama_kota"`
	IsUtama      bool   `json:"is_utama,omitempty" form:"is_utama"`
	CreatedAt    string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt    string `json:"updated_at,omitempty" form:"updated_at"`
}

type AlamatService interface {
	Create(req *Alamat) *Response
	Read(userID int) *Response
	Update(req *Alamat) *Response
	Delete(id int, user_id int) *Response
	MakeUtama(id int, user_id int) *Response
	GetProvinsi() *Response
	GetKota(provinsiID string) *Response
	GetAllKota() *Response
}

type AlamatRepository interface {
	Create(req *Alamat) (*Alamat, error)
	Read(userID int) ([]Alamat, error)
	Update(req *Alamat) (*Alamat, error)
	Delete(id int, user_id int) error
	MakeUtama(id int, user_id int) error
	MakeUtamaFalse(id int, user_id int) error
	ByID(id int) (*Alamat, error)
}
