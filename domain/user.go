package domain

type User struct {
	ID            int    `json:"id,omitempty" form:"id"`
	Nama          string `json:"nama,omitempty" form:"nama"`
	HP            string `json:"hp,omitempty" form:"hp"`
	Email         string `json:"email,omitempty" form:"email"`
	Password      string `json:"password,omitempty" form:"password"`
	APIToken      string `json:"api_token,omitempty" form:"api_token"`
	DeviceToken   string `json:"device_token,omitempty" form:"device_token"`
	RememberToken string `json:"remember_token,omitempty" form:"remember_token"`
	IsVerified    bool   `json:"is_verified,omitempty" form:"is_verified"`
	CreatedAt     string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt     string `json:"updated_at,omitempty" form:"updated_at"`
}

type UserBasicData struct {
	ID       int    `json:"id,omitempty" form:"id"`
	Nama     string `json:"nama,omitempty" form:"nama"`
	HP       string `json:"hp,omitempty" form:"hp"`
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
}

type UserCredential struct {
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
}

type UserService interface {
	Login(cred *UserCredential) *Response
	Registrasi(req *UserBasicData) *Response
}

type UserRepository interface {
	ByEmail(email string) (*User, error)
	ByID(id int) (*User, error)
	Create(req *UserBasicData) (*User, error)
	Update(req *UserBasicData) (*User, error)
	UpdatePassword(new_password string, id int) error
	UpdateVerificationStatus(id int) error
}
