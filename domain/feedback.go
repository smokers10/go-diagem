package domain

type Feedback struct {
	ID       int    `json:"id,omitempty" form:"id"`
	Comment  string `json:"comment,omitempty" form:"comment"`
	Rating   string `json:"rating,omitempty" form:"rating"`
	ProdukID string `json:"produk_id,omitempty" form:"produk_id"`
	UserID   int    `json:"user_id,omitempty" form:"user_id"`
}

type FeedbackDetailed struct {
	ID        int    `json:"id,omitempty" form:"id"`
	Comment   string `json:"comment,omitempty" form:"comment"`
	Rating    string `json:"rating,omitempty" form:"rating"`
	ProdukID  string `json:"produk_id,omitempty" form:"produk_id"`
	UserID    int    `json:"user_id,omitempty" form:"user_id"`
	CreatedAt string `json:"created_at,omitempty" form:"created_at"`
	User      User   `json:"user,omitempty"`
}

type FeedbackRateCount struct {
	FiveStar  int `json:"five_star"`
	FourStar  int `json:"four_star"`
	ThreeStar int `json:"three_star"`
	TwoStar   int `json:"two_star"`
	OneStar   int `json:"one_star"`
}

type FeedbackData struct {
	FeedbackList []FeedbackDetailed `json:"feedback_list"`
	ByRate       FeedbackRateCount  `json:"by_rate"`
}

type FeedbackRepository interface {
	Create(req *Feedback) error
	Update(req *Feedback) error
	Read(ProdukID string) (*FeedbackData, error)
	Delete(req *Feedback) error
	OnProduct(userid int, productID string) (*FeedbackDetailed, error)
	ByRating(productID string, rating int) (*FeedbackData, error)
}

type FeedbackService interface {
	Create(req *Feedback) *Response
	Update(req *Feedback) *Response
	Read(produkID string) *Response
	Delete(req *Feedback) *Response
	ByRating(productID string, rating int) *Response
}
