package domain

type Rekening struct {
	ID         int
	Bank       string
	Kode       string
	Icon       string
	BankID     int
	RekeningNo string
	Nama       string
	CreatedAt  string
	UpdatedAt  string
}

type RekeningInterface interface {
	//untuk Admin
	Create(req *Rekening) *Response
	Read() *Response
	Update(req *Rekening) *Response
	Delete(id int) *Response
}

type RekeningRepository interface {
	Create(req *Rekening) *Rekening
	Read() []Rekening
	Update(req *Rekening) *Rekening
	Delete(id int) *Rekening
}
