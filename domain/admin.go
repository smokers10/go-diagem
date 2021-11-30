package domain

type Admin struct {
	ID              int    `json:"id,omitempty" form:"id"`
	Nama            string `json:"nama,omitempty" form:"nama"`
	Username        string `json:"username,omitempty" form:"username"`
	Email           string `json:"email,omitempty" form:"email"`
	Password        string `json:"password,omitempty" form:"password"`
	Roles           string `json:"roles,omitempty" form:"roles"`
	CreatedAt       string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt       string `json:"updated_at,omitempty" form:"updated_at"`
	EmailVerifiedAt string `json:"email_Verified_at,omitempty" form:"email_Verified_at"`
}

type AdminCredentials struct {
	Email    string
	Password string
}

type AdminService interface {
	Login(cred *AdminCredentials) *Response
	Create(req *Admin) *Response
	Update(req *Admin, logged_id int, logged_role string) *Response
	UpdatePassword(req *Admin, logged_id int, logged_role string) *Response
	Delete(id int) *Response
	Read(logged_id int) *Response
	Detail(id int) *Response
}

type AdminRepository interface {
	Create(req *Admin) (*Admin, error)
	Update(req *Admin) (*Admin, error)
	UpdatePassword(req *Admin) (*Admin, error)
	Delete(id int) error
	GetByEmail(email string) (*Admin, error)
	GetByID(id int) (*Admin, error)
	Read(logged_id int) ([]Admin, error)
}
