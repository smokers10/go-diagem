package domain

type ProdukFoto struct {
	ID        string
	ProdukID  string
	Path      string
	IsUtama   string
	CreatedAt string
	UpdatedAt string
}

type ProdukFotoService interface {
}

type ProdukFotoRepository interface {
}
