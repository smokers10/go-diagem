package services

import (
	"fmt"

	"github.com/smokers10/go-diagem.git/domain"
)

type bankServiceImpl struct {
	bankRepository domain.BankRepository
}

func BankService(bank *domain.BankRepository) domain.BankService {
	return &bankServiceImpl{bankRepository: *bank}
}

func (bs *bankServiceImpl) Update(req *domain.Bank) *domain.Response {
	res := domain.Response{}

	if err := bs.bankRepository.Update(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat pada update bank"
		res.Status = 500
		return &res
	}

	res.Message = "bank berhasil diupdate"
	res.Status = 200
	res.Success = true
	return &res
}

func (bs *bankServiceImpl) Read() *domain.Response {
	res := domain.Response{}

	banks, err := bs.bankRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat pada mengambil data bank"
		res.Status = 500
		return &res
	}

	res.Data = banks
	res.Message = "bank berhasil diambil"
	res.Status = 200
	res.Success = true
	return &res
}
