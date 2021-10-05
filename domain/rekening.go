package domain

type Rekening struct {
	ID         int    `json:"id,omitempty" form:"id"`
	BankID     int    `json:"bank_id,omitempty" form:"bank_id"`
	RekeningNo string `json:"rekening_no,omitempty" form:"rekening_no"`
	Nama       string `json:"nama,omitempty" form:"nama"`
	CreatedAt  string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty" form:"updated_at"`
	Bank       Bank   `json:"bank,omitempty"`
}

type RekeningService interface {
	//untuk Admin
	Create(req *Rekening) *Response
	Update(req *Rekening) *Response
	Delete(id int) *Response

	//untuk umum
	Read() *Response
}

type RekeningRepository interface {
	Create(req *Rekening) (*Rekening, error)
	Read() ([]Rekening, error)
	Update(req *Rekening) (*Rekening, error)
	Delete(id int) error
}
