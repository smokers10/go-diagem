package domain

type Post struct {
	ID        int
	Slug      string
	Title     string
	Body      string
	View      string
	Status    string
	CreatedAt int
}

type PostService interface {
	//khusus admin
	Create(req *Post) *Response
	Update(req *Post) *Response
	Delete(id int) *Response

	//khusus umum
	Read() *Response
	Detail(slug string) *Response
}

type PostRepositrory interface {
	//khusus admin
	Create(req *Post) *Post
	Update(req *Post) *Post
	Delete(id int) *Post

	//khusus umum
	Read() *Post
	BySlug(slug string) *Post
}
