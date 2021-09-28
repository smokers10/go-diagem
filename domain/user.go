package domain

type User struct {
	ID            string
	Nama          string
	HP            string
	Email         string
	Password      string
	APIToken      string
	DeviceToken   string
	RememberToken string
	CreatedAt     string
	UpdatedAt     string
}

type UserBasicData struct {
	Nama     string
	HP       string
	Email    string
	Password string
}

type UserCredential struct {
	Email    string
	Password string
}

type UserResetPassword struct {
	Token           string
	InputCode       string
	NewPassword     string
	ConfirmPassword string
}

type UserVerifikasi struct {
	Email string
	Code  string
}

type UserService interface {
	Login(cred *UserCredential) *Response
	ForgotPassword(email string) *Response
	ResetPassword(UserResetPassword) *Response
	Verifikasi(req *UserVerifikasi) *Response
	Registrasi(req *UserBasicData) *Response
}

type UserRepository interface {
	ByEmail(email string) *User
	ByID(id string) *User
	Create(req *UserBasicData) *User
	Update(req *UserBasicData) *User
}
