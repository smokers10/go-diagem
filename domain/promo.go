package domain

type Promo struct {
	ID         string
	Judul      string
	Slug       string
	Deskripsi  string
	Image      string
	IsActive   bool
	Isfeatured bool
	TglMulai   string
	TglSelesai string
	CreatedAt  string
	UpdatedAt  string
}

type PromoService interface {
	//Untuk Admin
	Create(req *Promo) *Response
	Read() *Response
	Update(req *Promo) *Response
	Delete(id string) *Response

	//untuk umum
	Detail(id string) *Response
}

type PromoRepository interface {
	//Untuk Admin
	Create(req *Promo) *Promo
	Read() []Promo
	Update(req *Promo) *Promo
	Delete(id string) *Promo
	ByID(id string) *Promo
	BySlug(slug string) *Promo
}
