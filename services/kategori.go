package services

import (
	"fmt"
	"net/http"

	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
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

func (k *kategoriServiceImpl) Delete(id int) *domain.Response {
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

//Untuk Umum
func (k *kategoriServiceImpl) Read() *domain.Response {
	res := domain.Response{}

	kategori, err := k.kategoriRepository.Read()
	if err != nil {
		fmt.Println(err.Error())
		res.Message = "error saat menghapus data kategori"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "data kategori berhasil dihapus"
	res.Data = kategori
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
