package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type sliderServiceImpl struct {
	sliderRepository domain.SliderRepository
}

func SliderService(slider *domain.SliderRepository) domain.SliderService {
	return &sliderServiceImpl{sliderRepository: *slider}
}

func (s *sliderServiceImpl) Create(req *domain.Slider) *domain.Response {
	res := domain.Response{}

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

func (s *sliderServiceImpl) Delete(id int) *domain.Response {
	res := domain.Response{}

	if err := s.sliderRepository.Delete(id); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus slider"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "slider berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
