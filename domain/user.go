package domain

type User struct {
	ID         int    `json:"id,omitempty" form:"id"`
	Nama       string `json:"nama,omitempty" form:"nama"`
	HP         string `json:"hp,omitempty" form:"hp"`
	Email      string `json:"email,omitempty" form:"email"`
	Password   string `json:"password,omitempty" form:"password"`
	IsVerified bool   `json:"is_verified,omitempty" form:"is_verified"`
	Tahun      int    `json:"tahun,omitempty" form:"tahun"`
	Bulan      int    `json:"bulan,omitempty" form:"bulan"`
	Tanggal    int    `json:"tanggal,omitempty" form:"tanggal"`
	CreatedAt  string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty" form:"updated_at"`
}

type UserBasicData struct {
	ID       int    `json:"id,omitempty" form:"id"`
	Nama     string `json:"nama,omitempty" form:"nama"`
	HP       string `json:"hp,omitempty" form:"hp"`
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
	Tahun    int    `json:"tahun,omitempty" form:"tahun"`
	Bulan    int    `json:"bulan,omitempty" form:"bulan"`
	Tanggal  int    `json:"tanggal,omitempty" form:"tanggal"`
}

type UserCredential struct {
	Email    string `json:"email,omitempty" form:"email"`
	Password string `json:"password,omitempty" form:"password"`
}

type UserService interface {
	Read() *Response
	Login(cred *UserCredential) *Response
	Registrasi(req *UserBasicData) *Response
	Update(req *UserBasicData) *Response
	GetProfile(userID int) *Response
}

type UserRepository interface {
	ReadAll() ([]User, error)
	ByEmail(email string) (*User, error)
	ByID(id int) (*User, error)
	Create(req *UserBasicData) (*User, error)
	Update(req *UserBasicData) (*UserBasicData, error)
	UpdatePassword(new_password string, id int) error
	UpdateVerificationStatus(id int) error
}
