package domain

type Mitra struct {
	ID     string
	Nama   string
	Kontak string
	Alamat string
}

type MitraService interface {
	//Untuk Admin
	Create(req *Mitra) *Response
	Update(req *Mitra) *Response
	Delete(id string) *Response

	//Untuk Umum
	Read() *Response
	GetOne(id string) *Response
}

type MitraRepository interface {
	Create(req *Mitra) *Mitra
	Read() []Mitra
	ByID(id string) *Mitra
	Update(req *Mitra) *Mitra
	Delete(id string) *Mitra
}
