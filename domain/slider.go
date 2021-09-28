package domain

type Slider struct {
	ID          int
	Name        string
	Description string
	Img         string
	IsActive    string
	URL         string
	Type        int
	CreatedAt   string
	UpdatedAt   string
}

type SliderService interface {
	//Untuk Admin
	Create(req *Slider) *Response
	Read() *Response
	Update(req *Slider) *Response
	Deletete(req *Slider) *Response
}

type SliderRepository interface {
	//Untuk Admin
	Create(req *Slider) *Slider
	Read() []Slider
	Update(req *Slider) *Slider
	Deletete(req *Slider) *Slider
}
