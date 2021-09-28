package domain

type OrderBayar struct {
	ID         int
	OrderID    string
	Method     string
	RekeningID int
	Layanan    string
	Jumlah     float32
	Status     string
	Tglbayar   string
	CreatedAt  string
	UpdatedAt  string
}

type OrderBayarService interface {
}

type OrderBayarRepository interface {
}
