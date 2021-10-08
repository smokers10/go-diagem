package domain

type Verifikasi struct {
	UserID int    `json:"user_id,omitempty" form:"user_id"`
	Kode   string `json:"kode,omitempty" form:"kode"`
}

type VerifikasiService interface {
	Verificate(req *Verifikasi) *Response
	Create(req *Verifikasi) *Response
}

type VerifikasiRepository interface {
	Create(req *Verifikasi) error
	ByUserID(userID int) (*Verifikasi, error)
	Update(req *Verifikasi) error
	Delete(userID int) error
}
