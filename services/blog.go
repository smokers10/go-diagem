package services

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type blogServiceImpl struct {
	blogRepository domain.BlogRepository
}

func BlogService(blog *domain.BlogRepository) domain.BlogService {
	return &blogServiceImpl{blogRepository: *blog}
}

// khusus admin
func (bs *blogServiceImpl) Create(req *domain.Blog) *domain.Response {
	// start transaction
	c := context.Background()
	tx, err := bs.blogRepository.GetSqlInstance().BeginTx(c, nil)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat check keberadaan blog",
			Status:  500,
		}
	}

	// check judul apakah ada atau tidak
	isExist, err := bs.IsTitleExists(req.Judul)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat check keberadaan blog",
			Status:  500,
		}
	}

	if isExist {
		tx.Rollback()
		return &domain.Response{
			Message: "error saat check keberadaan blog",
			Status:  500,
		}
	}

	// buat id
	id, _ := uuid.NewRandom()
	req.ID = id.String()

	// buat slug
	req.Slug = slug.Make(req.Judul)

	// upload thumbnails blog
	path := "public/uploads/blog/"
	filename := fmt.Sprintf("%s.%s", req.ID, req.Thumbnail.Format)
	fpath, err := etc.UploadFile(path, filename, req.Thumbnail.Base64)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat upload thumbnail blog",
			Status:  500,
		}
	}
	req.Image = fpath

	// jika judul belum terpakai call repository
	if err := bs.blogRepository.Create(req, tx); err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat membuat blog",
			Status:  500,
		}
	}

	// commit changes
	tx.Commit()

	// return response
	return &domain.Response{
		Status:  200,
		Success: true,
		Message: "blog berhasil dibuat",
	}
}

func (bs *blogServiceImpl) Update(req *domain.Blog) *domain.Response {
	// buat slug
	req.Slug = slug.Make(req.Judul)

	// call repository
	if err := bs.blogRepository.Update(req); err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat update blog",
			Status:  500,
		}
	}

	return &domain.Response{
		Status:  200,
		Success: true,
		Message: "blog berhasil diupdate",
	}
}

func (bs *blogServiceImpl) Delete(id string) *domain.Response {
	// call repository
	image, err := bs.blogRepository.Delete(id)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat menghapus blog",
			Status:  500,
		}
	}

	// hapus citra promo
	if err := os.Remove(image); err != nil {
		fmt.Println(err)
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat menghapus thumbnail blog",
			Status:  500,
		}
	}

	return &domain.Response{
		Status:  200,
		Success: true,
		Message: "blog berhasil dihapus",
	}
}

func (bs *blogServiceImpl) ReadAll() *domain.Response {
	// call repository
	response, err := bs.blogRepository.ReadAll()
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil blog",
			Status:  500,
		}
	}

	return &domain.Response{
		Data:    response,
		Status:  200,
		Success: true,
		Message: "blog berhasil diambil",
	}
}

func (bs *blogServiceImpl) UpdateThumbnail(req *domain.Blog) *domain.Response {
	// upload thumbnails blog
	path := "public/uploads/blog/"
	filename := fmt.Sprintf("%s.%s", req.ID, req.Thumbnail.Format)
	fpath, err := etc.UploadFile(path, filename, req.Thumbnail.Base64)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat upload thumbnail blog",
			Status:  500,
		}
	}
	req.Image = fpath

	return &domain.Response{
		Status:  200,
		Success: true,
		Message: "thumbnail blog berhasil diupdate",
	}
}

// khusus umum
func (bs *blogServiceImpl) PublishedOnly(judul string) *domain.Response {
	// call repository
	response, err := bs.blogRepository.PublishedOnly(judul)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil blog",
			Status:  500,
		}
	}

	return &domain.Response{
		Data:    response,
		Status:  200,
		Success: true,
		Message: "blog berhasil diambil",
	}
}

func (bs *blogServiceImpl) Detail(slug string) *domain.Response {
	// call repository
	response, err := bs.blogRepository.BySlug(slug)
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil blog",
			Status:  500,
		}
	}

	return &domain.Response{
		Data:    response,
		Status:  200,
		Success: true,
		Message: "blog berhasil diambil",
	}
}

func (bs *blogServiceImpl) Latest() *domain.Response {
	response, err := bs.blogRepository.Latest()
	if err != nil {
		fmt.Println(err)
		return &domain.Response{
			Message: "error saat mengambil blog",
			Status:  500,
		}
	}

	return &domain.Response{
		Data:    response,
		Status:  200,
		Success: true,
		Message: "blog berhasil diambil",
	}
}

// out of abtract (check judul)
func (bs *blogServiceImpl) IsTitleExists(judul string) (bool, error) {
	blog, err := bs.blogRepository.CheckJudul(judul)
	if err != nil {
		return false, err
	}

	if blog.Judul != "" {
		return true, nil
	}

	return false, nil
}
