package domain

type Admin struct {
	ID              int
	Nama            string
	Username        string
	Email           string
	Password        string
	RememberToken   string
	CreatedAt       string
	UpdatedAt       string
	EmailVerifiedAt string
}

type AdminCredentials struct {
	Email    string
	Password string
}

type AdminResetPassword struct {
	Token           string
	InputCode       string
	NewPassword     string
	ConfirmPassword string
}

type AdminService interface {
	ForgotPassword(email string) *Response
	ResetPassword(req *AdminResetPassword) *Response
	Login(cred *AdminCredentials) *Response
}

type AdminRepository interface {
	GetByEmail(email string) *Admin
}
