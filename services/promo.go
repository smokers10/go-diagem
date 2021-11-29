package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type promoServiceImpl struct {
	promoRepository domain.PromoRepository
}

func PromoService(promo *domain.PromoRepository) domain.PromoService {
	return &promoServiceImpl{promoRepository: *promo}
}

// Untuk Admin
func (p *promoServiceImpl) Create(req *domain.Promo) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	// buat ID
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat membuat ID promo"
		res.Status = http.StatusInternalServerError
		return &res
	}
	req.ID = id.String()

	// buat slug
	req.Slug = slug.Make(req.Judul)

	// simpan data citra sampul
	path := "public/uploads/promo/"                                    // file path
	filename := fmt.Sprintf("sampul-%s.%s", req.ID, req.Sampul.Format) // image/file name
	fullpath, err := etc.UploadFile(path, filename, req.Sampul.Base64)
	if err != nil {
		panic(err)
	}

	// set path(web) data citra
	req.Image = fullpath

	// menyimpan data promo
	promo, err := p.promoRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan promo"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = promo
	res.Message = "promo baru berhasil disimpan"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *promoServiceImpl) Update(req *domain.Promo) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	// buat slug
	req.Slug = slug.Make(req.Judul)

	promo, err := p.promoRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengupdate promo"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = promo
	res.Message = "promo baru berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *promoServiceImpl) Delete(id string) *domain.Response {
	// deklarasi var
	res := domain.Response{}

	// hapus promo di database
	err := p.promoRepository.Delete(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus promo"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// hapus citra promo
	if err := os.Remove(fmt.Sprintf("public/uploads/promo/sampul-%s.jpeg", id)); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus citra sampul"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "promo baru berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *promoServiceImpl) UpdateCover(req *domain.Promo) *domain.Response {
	// deklarasi var
	res := domain.Response{}

	// simpan data citra sampul
	path := "public/uploads/promo/"                                    // file path
	filename := fmt.Sprintf("sampul-%s.%s", req.ID, req.Sampul.Format) // image/file name
	if _, err := etc.UploadFile(path, filename, req.Sampul.Base64); err != nil {
		panic(err)
	}

	res.Message = "sampul promo berhasil duupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

// untuk umum
func (p *promoServiceImpl) Read() *domain.Response {
	//deklarasi var
	res := domain.Response{}

	promo, err := p.promoRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil promo"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = promo
	res.Message = "promo baru berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (p *promoServiceImpl) Detail(id string) *domain.Response {
	//deklarasi var
	res := domain.Response{}

	promo, err := p.promoRepository.ByID(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil promo"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = promo
	res.Message = "promo baru berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
