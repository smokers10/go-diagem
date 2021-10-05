package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type rekeningServiceImpl struct {
	rekeningRepository domain.RekeningRepository
}

func RekeningService(rekening *domain.RekeningRepository) domain.RekeningService {
	return &rekeningServiceImpl{rekeningRepository: *rekening}
}

//untuk Admin
func (r *rekeningServiceImpl) Create(req *domain.Rekening) *domain.Response {
	res := domain.Response{}

	rekening, err := r.rekeningRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengimpan data rekening"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "rekengin baru berhasil disimpan"
	res.Data = rekening
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (r *rekeningServiceImpl) Update(req *domain.Rekening) *domain.Response {
	res := domain.Response{}

	rekening, err := r.rekeningRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat update data rekening"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "rekening berhasil diupdate"
	res.Data = rekening
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (r *rekeningServiceImpl) Delete(id int) *domain.Response {
	res := domain.Response{}

	if err := r.rekeningRepository.Delete(id); err != nil {
		fmt.Println(err)
		res.Message = "error saat hapus data rekening"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "rekening berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

//untuk umum
func (r *rekeningServiceImpl) Read() *domain.Response {
	res := domain.Response{}

	rekening, err := r.rekeningRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data rekening"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "rekening berhasil diambil"
	res.Data = rekening
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
