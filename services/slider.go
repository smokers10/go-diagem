package services

import (
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type sliderServiceImpl struct {
	sliderRepository domain.SliderRepository
}

func SliderService(slider *domain.SliderRepository) domain.SliderService {
	return &sliderServiceImpl{sliderRepository: *slider}
}

// Untuk Admin
func (s *sliderServiceImpl) Create(req *domain.Slider) *domain.Response {
	res := domain.Response{}

	// buat id
	id, err := uuid.NewRandom()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat membuat ID promo"
		res.Status = http.StatusInternalServerError
		return &res
	}
	req.ID = id.String()

	// simpan data citra sampul
	path := "public/uploads/slider/"                                  // file path
	filename := fmt.Sprintf("cover-%s.%s", req.ID, req.Sampul.Format) // image/file name
	fullpath, err := etc.UploadFile(path, filename, req.Sampul.Base64)
	if err != nil {
		panic(err)
	}

	// simpan data ke database
	req.Image = fullpath
	slider, err := s.sliderRepository.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "slider baru berhasil disimpan"
	res.Data = slider
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (s *sliderServiceImpl) Read() *domain.Response {
	res := domain.Response{}

	slider, err := s.sliderRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "slider berhasil diambil"
	res.Data = slider
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (s *sliderServiceImpl) Update(req *domain.Slider) *domain.Response {
	res := domain.Response{}

	slider, err := s.sliderRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat update slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "slider berhasil diupdate"
	res.Data = slider
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (s *sliderServiceImpl) Delete(id string) *domain.Response {
	res := domain.Response{}

	if err := s.sliderRepository.Delete(id); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	os.Remove(fmt.Sprintf("public/uploads/slider/%s.jpeg", id))

	res.Message = "slider berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (s *sliderServiceImpl) UpdateCover(req *domain.Slider) *domain.Response {
	res := domain.Response{}

	path := "public/uploads/slider/"                                  // file path
	filename := fmt.Sprintf("cover-%s.%s", req.ID, req.Sampul.Format) // image/file name
	if _, err := etc.UploadFile(path, filename, req.Sampul.Base64); err != nil {
		fmt.Println(err)
		res.Message = "cover slider gagal diupdate"
		res.Status = http.StatusOK
		res.Success = true
	}

	res.Message = "cover slider berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

// Untuk Umum
func (s *sliderServiceImpl) ByID(id string) *domain.Response {
	res := domain.Response{}

	slider, err := s.sliderRepository.ByID(id)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "slider berhasil diambil"
	res.Data = slider
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (s *sliderServiceImpl) ReadPublished() *domain.Response {
	res := domain.Response{}

	slider, err := s.sliderRepository.ReadOnlyPublished()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "slider berhasil diambil"
	res.Data = slider
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
