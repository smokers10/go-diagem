package services

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkVariasiServiceImpl struct {
	produkVariasi domain.ProdukVariasiRepository
}

func ProdukVariasiService(produkVariasiRepo *domain.ProdukVariasiRepository) domain.ProdukVariasiService {
	return &produkVariasiServiceImpl{produkVariasi: *produkVariasiRepo}
}

func (p *produkVariasiServiceImpl) Create(req *domain.ProdukVariasi) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// set id
	uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat generate ID variasi"
		res.Status = http.StatusInternalServerError
		return &res
	}
	req.ID = uuid.String()

	// panggil repository
	variasi, err := p.produkVariasi.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan variasi produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = variasi
	res.Message = "variasi produk berhasil disimpan"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (p *produkVariasiServiceImpl) Update(req *domain.ProdukVariasi) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repossitory
	variasi, err := p.produkVariasi.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengupdate variasi produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = variasi
	res.Message = "variasi produk berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (p *produkVariasiServiceImpl) Delete(produkID string, produkVariasiID string) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repossitory
	err := p.produkVariasi.Delete(produkID, produkVariasiID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus variasi produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "variasi produk berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}
