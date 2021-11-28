package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/smokers10/go-diagem.git/domain"
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

	// buat directory jika tidak ada dan set permission
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}

	dec, err := base64.StdEncoding.DecodeString(req.Sampul.Base64)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filepath.Join(path, filepath.Base(filename)))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err := file.Write(dec); err != nil {
		panic(err)
	}

	if err := file.Sync(); err != nil {
		panic(err)
	}

	// set path(web) data citra
	req.Image = fmt.Sprintf("/uploads/promo/sampul-%s.%s", req.ID, req.Sampul.Format)

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
	//deklarasi var
	res := domain.Response{}

	err := p.promoRepository.Delete(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus promo"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "promo baru berhasil dihapus"
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
