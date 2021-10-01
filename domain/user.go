package domain

type User struct {
	ID            int
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
	ID       int
	Nama     string
	HP       string
	Email    string
	Password string
}

type UserCredential struct {
	Email    string
	Password string
}

type UserService interface {
	Login(cred *UserCredential) *Response
	Registrasi(req *UserBasicData) *Response
}

type UserRepository interface {
	ByEmail(email string) (*User, error)
	ByID(id string) (*User, error)
	Create(req *UserBasicData) (*User, error)
	Update(req *UserBasicData) (*User, error)
	UpdatePassword(new_password string, id int) error
}
