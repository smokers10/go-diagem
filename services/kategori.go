package services

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type kategoriServiceImpl struct {
	kategoriRepository domain.KategoriRepository
}

func KategoriService(kategori *domain.KategoriRepository) domain.KategoriService {
	return &kategoriServiceImpl{kategoriRepository: *kategori}
}

//Khusus Admin
func (k *kategoriServiceImpl) Create(req *domain.Kategori) *domain.Response {
	res := domain.Response{}
	req.Slug = slug.Make(req.Nama)

	uuid, _ := uuid.NewRandom()
	req.ID = uuid.String()

	// simpan data citra sampul
	path := "public/uploads/kategori/"                                    // file path
	filename := fmt.Sprintf("cover-%s.%s", req.ID, req.CoverImage.Format) // image/file name
	fullpath, err := etc.UploadFile(path, filename, req.CoverImage.Base64)
	if err != nil {
		panic(err)
	}

	req.Cover = fullpath

	kategori, err := k.kategoriRepository.Create(req)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat membuat kategori"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "kategori berhasill disimpan"
	res.Data = kategori
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (k *kategoriServiceImpl) Update(req *domain.Kategori) *domain.Response {
	res := domain.Response{}
	req.Slug = slug.Make(req.Nama)

	kategori, err := k.kategoriRepository.Update(req)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat update kategori"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "kategori berhasill diupdate"
	res.Data = kategori
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (k *kategoriServiceImpl) Delete(id string) *domain.Response {
	res := domain.Response{}

	err := k.kategoriRepository.Delete(id)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat update kategori"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "kategori berhasill diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (k *kategoriServiceImpl) UpdateCover(req *domain.Kategori) *domain.Response {
	res := domain.Response{}

	path := "public/uploads/kategori/"                                    // file path
	filename := fmt.Sprintf("cover-%s.%s", req.ID, req.CoverImage.Format) // image/file name
	_, err := etc.UploadFile(path, filename, req.CoverImage.Base64)
	if err != nil {
		panic(err)
	}

	res.Message = "sampul kategori berhasill diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

//Untuk Umum
func (k *kategoriServiceImpl) Read() *domain.Response {
	res := domain.Response{}

	kategori, err := k.kategoriRepository.Read()
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat mengambil data kategori"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "data kategori berhasil diambil"
	res.Data = kategori
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (k *kategoriServiceImpl) Detail(id string) *domain.Response {
	res := domain.Response{}

	kategori, err := k.kategoriRepository.Detail(id)
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat update kategori"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = kategori
	res.Message = "kategori berhasill diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
