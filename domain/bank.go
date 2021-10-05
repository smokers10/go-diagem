package domain

type Bank struct {
	ID        int    `json:"id,omitempty" form:"id"`
	Name      string `json:"name,omitempty" form:"name"`
	Code      string `json:"code,omitempty" form:"code"`
	Icon      string `json:"icon,omitempty" form:"icon"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt string `json:"updated_at,omitempty" form:"updated_at"`
}

type BankService interface{}

type BankRepository interface{}
