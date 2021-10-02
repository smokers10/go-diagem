package domain

type Mitra struct {
	ID       int    `json:"id,omitempty" form:"id"`
	Nama     string `json:"nama,omitempty" form:"nama"`
	Kontak   string `json:"kontak,omitempty" form:"kontak"`
	Alamat   string `json:"alamat,omitempty" form:"alamat"`
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Username string `json:"username,omitempty" form:"username"`
}

type MitraService interface {
	//Untuk Admin
	Create(req *Mitra) *Response
	Update(req *Mitra) *Response
	Delete(id int) *Response

	//Untuk Umum
	Read() *Response
	GetOne(id int) *Response

	//untuk Mitra
}

type MitraRepository interface {
	Create(req *Mitra) (*Mitra, error)
	Read() ([]Mitra, error)
	ByID(id int) (*Mitra, error)
	ByEmail(email string) (*Mitra, error)
	ByUsername(username string) (*Mitra, error)
	Update(req *Mitra) (*Mitra, error)
	UpdateEmail(id int, email string) error
	UpdatePassword(id int, password string) error
	Delete(id int) error
}
