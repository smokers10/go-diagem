package domain

type Slider struct {
	ID          string `json:"id,omitempty" form:"id"`
	Name        string `json:"name,omitempty" form:"name"`
	Description string `json:"description,omitempty" form:"description"`
	Image       string `json:"image,omitempty" form:"image"`
	IsPublish   bool   `json:"is_publish,omitempty" form:"is_publish"`
	URL         string `json:"url,omitempty" form:"url"`
	Type        int    `json:"type,omitempty" form:"type"`
	CreatedAt   string `json:"created_at,omitempty" form:"created_at"`
	UpdatedAt   string `json:"updated_at,omitempty" form:"updated_at"`
	Sampul      *sliderCover
}

type sliderCover struct {
	Base64 string `json:"base64" form:"base64"`
	Format string `json:"format" form:"format"`
}

type SliderService interface {
	// Untuk Admin
	Create(req *Slider) *Response
	Read() *Response
	Update(req *Slider) *Response
	Delete(id string) *Response
	UpdateCover(req *Slider) *Response

	// Untuk Umum
	ByID(id string) *Response
	ReadPublished() *Response
}

type SliderRepository interface {
	// Untuk Admin
	Create(req *Slider) (*Slider, error)
	Read() ([]Slider, error)
	Update(req *Slider) (*Slider, error)
	Delete(id string) error

	// Untuk Umum
	ByID(id string) (*Slider, error)
	ReadOnlyPublished() ([]Slider, error)
}
