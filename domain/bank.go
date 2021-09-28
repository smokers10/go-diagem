package domain

type Bank struct {
	ID        int
	Name      string
	Code      string
	CreatedAt string
	UpdatedAt string
}

type BankService interface{}

type BankRepository interface{}
