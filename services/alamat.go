package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type alamatServiceImpl struct {
	alamatRepository domain.AlamatRepository
}

func AlamtService(alamt *domain.AlamatRepository) domain.AlamatService {
	return &alamatServiceImpl{alamatRepository: *alamt}
}

func (a *alamatServiceImpl) Create(req *domain.Alamat) *domain.Response {
	res := domain.Response{}

	alamat, err := a.alamatRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menyimpan alamat"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = alamat
	res.Message = "alamat berhasil disimpan"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (a *alamatServiceImpl) Read(userID int) *domain.Response {
	res := domain.Response{}

	alamat, err := a.alamatRepository.Read(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat mengambil alamat"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = alamat
	res.Message = "alamat berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (a *alamatServiceImpl) Update(req *domain.Alamat) *domain.Response {
	res := domain.Response{}

	alamat, err := a.alamatRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat update alamat"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = alamat
	res.Message = "alamat berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (a *alamatServiceImpl) Delete(id int, user_id int) *domain.Response {
	res := domain.Response{}

	if err := a.alamatRepository.Delete(id, user_id); err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menghapus alamat"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "alamat berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}
