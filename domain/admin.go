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

type AdminService interface {
	Login(cred *AdminCredentials) *Response
}

type AdminRepository interface {
	GetByEmail(email string) (*Admin, error)
}
