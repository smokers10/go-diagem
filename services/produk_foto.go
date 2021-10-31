package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type produkFotoServiceImpl struct {
	produkFotoRepository domain.ProdukFotoRepository
}

func ProdukFoto(produkFoto *domain.ProdukFotoRepository) domain.ProdukFotoService {
	return &produkFotoServiceImpl{produkFotoRepository: *produkFoto}
}

func (p *produkFotoServiceImpl) Create(req *domain.ProdukFoto) *domain.Response {
	res := domain.Response{}

	err := p.produkFotoRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menyimpan foto"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "foto berhasil disimpan"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (p *produkFotoServiceImpl) Delete(id string) *domain.Response {
	res := domain.Response{}

	err := p.produkFotoRepository.Delete(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menyimpan foto"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "foto berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (p *produkFotoServiceImpl) MakeUtama(id string, produkID string) *domain.Response {
	res := domain.Response{}

	if err := p.produkFotoRepository.UpdateIsUtamaToTrue(id); err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat membuat foto utama"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if err := p.produkFotoRepository.UpdateRestIsUtamaToFalse(id, produkID); err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat membuat foto utama (vice versa)"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "foto berhasil membuat foto utama"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}
