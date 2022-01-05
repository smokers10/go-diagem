package domain

import "database/sql"

type Blog struct {
	ID           string         `json:"id,omitempty" form:"id"`
	Judul        string         `json:"judul,omitempty" form:"judul"`
	Slug         string         `json:"slug,omitempty" form:"slug"`
	Isi          string         `json:"isi,omitempty" form:"isi"`
	Image        string         `json:"image,omitempty" form:"image"`
	IsPublish    bool           `json:"is_publish" form:"is_publish"`
	SEOKeyword   string         `json:"seo_keyword,omitempty" form:"seo_keyword"`
	SEOTags      string         `json:"seo_tags,omitempty" form:"seo_tags"`
	SEODeskripsi string         `json:"seo_deskripsi,omitempty" form:"seo_deskripsi"`
	TGLPublish   string         `json:"tgl_publish,omitempty" form:"tgl_publish"`
	TGLUpdate    string         `json:"tgl_update,omitempty" form:"tgl_update"`
	Thumbnail    *blogThumbnail `json:"thumbnail,omitempty" form:"thumbnail"`
}

type blogThumbnail struct {
	Base64 string `json:"base64" form:"base64"`
	Format string `json:"format" form:"format"`
}

type BlogService interface {
	//khusus admin
	Create(req *Blog) *Response
	Update(req *Blog) *Response
	Delete(id string) *Response
	ReadAll() *Response
	UpdateThumbnail(req *Blog) *Response

	//khusus umum
	PublishedOnly() *Response
	Detail(slug string) *Response
}

type BlogRepository interface {
	//khusus admin
	Create(req *Blog, tx *sql.Tx) error
	Update(req *Blog) error
	Delete(blogID string) error
	CheckJudul(judul string) (*Blog, error)
	ReadAll() ([]Blog, error)

	//khusus umum
	PublishedOnly() ([]Blog, error)
	BySlug(slug string) (*Blog, error)

	//for transaction
	GetSqlInstance() *sql.DB
}
