package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type cartServiceImpl struct {
	cartRepo domain.CartRepository
}

func CartService(cart *domain.CartRepository) domain.CartService {
	return &cartServiceImpl{cartRepo: *cart}
}

func (cs *cartServiceImpl) Read(userID int) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	cart, err := cs.cartRepo.Read(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = cart
	res.Message = "cart berhasil diambil"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (cs *cartServiceImpl) AddtoCart(req *domain.Cart) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	cart, err := cs.cartRepo.Create(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan cart baru"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = cart
	res.Message = "cart bertambah"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (cs *cartServiceImpl) UpdateQuantity(req *domain.CartData) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	cart, err := cs.cartRepo.UpdateQuantity(req)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengupdate quantitas item cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = cart
	res.Message = "quantitas item cart berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (cs *cartServiceImpl) Delete(req *domain.CartData) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// panggil repository
	if err := cs.cartRepo.Delete(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus quantitas item cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "quantitas item cart berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}
