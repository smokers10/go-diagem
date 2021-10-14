package services

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkServiceImpl struct {
	produkRepository domain.ProdukRepository
}

func ProdukService(produk *domain.ProdukRepository) domain.ProdukService {
	return &produkServiceImpl{produkRepository: *produk}
}

//Untuk Admin
func (p *produkServiceImpl) Create(req *domain.Produk) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// buat uuidv4
	uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat membuat ID produk"
		res.Status = http.StatusInternalServerError
		return &res
	}
	req.ID = uuid.String()

	// buat slug
	req.Slug = slug.Make(req.Nama)

	// panggil repository terkait
	produk, err := p.produkRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil disimpan"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) Update(req *domain.Produk) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// buat slug
	req.Slug = slug.Make(req.Nama)

	// panggil repository terkait
	produk, err := p.produkRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat update produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) Delete(id string) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	if err := p.produkRepository.Delete(id); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "produk berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

//Untuk Umum
func (p *produkServiceImpl) Read(filter *domain.ProdukFilter) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	produk, err := p.produkRepository.Read(filter)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) Detail(id string) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	produk, err := p.produkRepository.ByID(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *produkServiceImpl) DetailBySlug(slug string) *domain.Response {
	// dekalarasi var
	res := domain.Response{}

	// panggil repository terkait
	produk, err := p.produkRepository.BySlugs(slug)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = produk
	res.Message = "produk berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
