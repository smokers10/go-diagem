package domain

type PasswordResets struct {
	Email        string `json:"email,omitempty" form:"email"`
	Token        string `json:"token,omitempty" form:"token"`
	Kode         string `json:"kode,omitempty" form:"kode"`
	NewPassword  string `json:"new_password,omitempty" form:"new_password"`
	ConfPassword string `json:"conf_password,omitempty" form:"conf_password"`
	CreatedAt    string `json:"created_at,omitempty" form:"created_at"`
}

type PasswordResetsService interface {
	Create(email string) *Response
	Reset(req *PasswordResets) *Response
}

type PasswordResetsRepository interface {
	Create(req *PasswordResets) error
	Update(req *PasswordResets) error
	Delete(email string) error
	ByEmail(email string) (*PasswordResets, error)
	ByToken(token string) (*PasswordResets, error)
}
