package domain

type Slider struct {
	ID          int    `json:"id,omitempty" form:"id"`
	Name        string `json:"name,omitempty" form:"name"`
	Description string `json:"description,omitempty" form:"description"`
	Img         string `json:"img,omitempty" form:"img"`
	IsActive    bool   `json:"is_active,omitempty" form:"is_active"`
	URL         string `json:"url,omitempty" form:"url"`
	Type        int    `json:"type,omitempty" form:"type"`
	CreatedAt   string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty" form:"updated_at"`
}

type SliderService interface {
	//Untuk Admin
	Create(req *Slider) *Response
	Read() *Response
	Update(req *Slider) *Response
	Delete(id int) *Response
}

type SliderRepository interface {
	//Untuk Admin
	Create(req *Slider) (*Slider, error)
	Read() ([]Slider, error)
	Update(req *Slider) (*Slider, error)
	Delete(id int) error
}
