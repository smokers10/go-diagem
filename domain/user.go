package domain

type User struct {
	ID         int      `json:"id,omitempty" form:"id"`
	Nama       string   `json:"nama,omitempty" form:"nama"`
	HP         string   `json:"hp,omitempty" form:"hp"`
	Email      string   `json:"email,omitempty" form:"email"`
	Password   string   `json:"password,omitempty" form:"password"`
	TglLahir   string   `json:"tgl_lahir,omitempty" form:"tgl_lahir"`
	IsVerified bool     `json:"is_verified,omitempty" form:"is_verified"`
	CreatedAt  string   `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  string   `json:"updated_at,omitempty" form:"updated_at"`
	Alamat     []Alamat `json:"alamat,omitempty" form:"alamata"`
}

type UserBasicData struct {
	ID       int    `json:"id,omitempty" form:"id"`
	Nama     string `json:"nama,omitempty" form:"nama"`
	HP       string `json:"hp,omitempty" form:"hp"`
	Email    string `json:"email,omitempty" form:"email"`
	TglLahir string `json:"tgl_lahir,omitempty" form:"tgl_lahir"`
	Password string `json:"password,omitempty" form:"password"`
}

type UserCredential struct {
	Email            string `json:"email,omitempty" form:"email"`
	Password         string `json:"password,omitempty" form:"password"`
	RechaptaResponse string `json:"g-recaptcha-response,omitempty" form:"g-recaptcha-response"`
}

type UserService interface {
	Read() *Response
	Detail(id int) *Response
	Login(cred *UserCredential) *Response
	Registrasi(req *UserBasicData) *Response
	Update(req *UserBasicData) *Response
	GetProfile(userID int) *Response
}

type UserRepository interface {
	ReadAll() ([]User, error)
	Detail(id int) (*User, error)
	ByEmail(email string) (*User, error)
	ByID(id int) (*User, error)
	Create(req *UserBasicData) (*User, error)
	Update(req *UserBasicData) (*UserBasicData, error)
	UpdatePassword(new_password string, id int) error
	UpdateVerificationStatus(id int) error
}
