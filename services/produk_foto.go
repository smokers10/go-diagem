package services

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
)

type produkFotoServiceImpl struct {
	produkFotoRepository domain.ProdukFotoRepository
}

func ProdukFoto(produkFoto *domain.ProdukFotoRepository) domain.ProdukFotoService {
	return &produkFotoServiceImpl{produkFotoRepository: *produkFoto}
}

func (p *produkFotoServiceImpl) Create(req *domain.FotoProdukFile) *domain.Response {
	res := domain.Response{}
	uidCitra, _ := uuid.NewRandom()
	fotoID, _ := uuid.NewRandom()

	// Menyimpan foto
	path := fmt.Sprintf("public/uploads/produk/%s", req.ProdukID)
	if err := os.MkdirAll(path, 0755); err != nil {
		panic(err)
	}

	dec, err := base64.StdEncoding.DecodeString(req.Base64)
	if err != nil {
		panic(err)
	}

	img := fmt.Sprintf("foto-produk-%s.%s", uidCitra.String(), req.Format)
	f, err := os.Create(filepath.Join(path, filepath.Base(img)))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write(dec); err != nil {
		panic(err)
	}

	if err := f.Sync(); err != nil {
		panic(err)
	}

	foto := domain.ProdukFoto{
		ID:       fotoID.String(),
		ProdukID: req.ProdukID,
		Path:     fmt.Sprintf("/uploads/produk/%s/foto-produk-%s.%s", req.ProdukID, uidCitra.String(), req.Format),
		IsUtama:  false,
	}

	// simpan ke database
	p.produkFotoRepository.Create(&foto)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menyimpan foto"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "foto berhasil dihapus"
	res.Status = http.StatusOK
	res.Data = foto
	res.Success = true
	return &res
}

func (p *produkFotoServiceImpl) Delete(id string, path string) *domain.Response {
	res := domain.Response{}

	// hapus file di direktori
	completePath := fmt.Sprintf("public%s", path)
	if err := os.Remove(completePath); err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menghapus file foto"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// hapus data citra di database
	err := p.produkFotoRepository.Delete(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat menghapus data foto"
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
