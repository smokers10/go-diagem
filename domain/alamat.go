package domain

//kusus Umum
type Alamat struct {
	ID          int
	UserID      int
	Alamat      string
	Nama        string
	Penerima    string
	Phone       string
	KelurahanID string
	KodePOS     string
	IsUtama     bool
	Lat         float32
	Lng         float32
	CreatedAt   string
	UpdatedAt   string
}

type AlamatService interface {
	Create(req *Alamat) *Response
	Read(userID int) *Response
	Update(req *Alamat) *Response
	Delete(id int) *Response
}

type AlamatRepository interface {
	Create(req *Alamat) *Alamat
	Read(userID int) []Alamat
	Update(req *Alamat) *Alamat
	Delete(id int) *Alamat
}
