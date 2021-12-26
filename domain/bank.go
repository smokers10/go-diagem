package domain

type Bank struct {
	ID     int    `json:"id,omitempty" form:"id"`
	NoVa   string `json:"no_va,omitempty" form:"no_va"`
	Status bool   `json:"status" form:"status"`
	Vendor string `json:"vendor" form:"vendor"`
}

type BankService interface {
	Update(req *Bank) *Response
	Read() *Response
}

type BankRepository interface {
	Update(req *Bank) error
	GetActive() ([]Bank, error)
	Read() ([]Bank, error)
}
