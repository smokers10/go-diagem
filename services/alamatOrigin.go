package services

import (
	"fmt"

	"github.com/smokers10/go-diagem.git/domain"
)

type alamatOriginServiceImpl struct {
	alamatOriginRepository domain.AlamatOriginRepository
}

func AlamatOriginService(alamatOrigin *domain.AlamatOriginRepository) domain.AlamatOriginService {
	return &alamatOriginServiceImpl{
		alamatOriginRepository: *alamatOrigin,
	}
}

func (ao *alamatOriginServiceImpl) Read() *domain.Response {
	res := domain.Response{}
	origin, err := ao.alamatOriginRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat mengambil alamat"
		res.Status = 500
		return &res
	}

	res.Data = origin
	res.Message = "data alamat berhasil diambil"
	res.Status = 200
	return &res
}

func (ao *alamatOriginServiceImpl) Update(req *domain.AlamatOrigin) *domain.Response {
	res := domain.Response{}
	origin, err := ao.alamatOriginRepository.Update(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error terjadi saat mengupdate alamat"
		res.Status = 500
		return &res
	}
	res.Data = origin
	res.Message = "alamat berhasil diupdate"
	res.Status = 200
	return &res
}
