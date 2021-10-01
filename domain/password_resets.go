package domain

type PasswordResets struct {
	Email     string `json:"email,omitempty" form:"email"`
	Token     string `json:"token,omitempty" form:"token"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
}

type PasswordResetsCreate struct {
	Email    string `json:"email,omitempty" form:"email"`
	UserType string `json:"user_type,omitempty" form:"user_type"`
}

type PasswordResetsService interface {
	Create(req *PasswordResetsCreate) *Response
	ForgotPassword(email string) *Response
	ResetPassword(token string) *Response
}

type PasswordResetsRepository interface {
	Create(req *PasswordResets) error
	Delete(email string) error
	ByEmail(email string) (*PasswordResets, error)
}
