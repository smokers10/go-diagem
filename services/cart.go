package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type cartServiceImpl struct {
	cartRepo                domain.CartRepository
	produkRepository        domain.ProdukRepository
	produkVariasiRepository domain.ProdukVariasiRepository
}

func CartService(cart *domain.CartRepository, produk *domain.ProdukRepository, variasi *domain.ProdukVariasiRepository) domain.CartService {
	return &cartServiceImpl{
		cartRepo:                *cart,
		produkRepository:        *produk,
		produkVariasiRepository: *variasi,
	}
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

	// check apakah stok tersedia atau tidak
	produk, err := cs.produkRepository.ByID(req.ProdukID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat check stok produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if cs.isStockInsuficient(produk, req) {
		res.Message = "stok tidak memadai"
		res.Status = http.StatusOK
		return &res
	}

	if _, err := cs.cartRepo.Create(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan cart baru dengan variasi"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "barang masuk ke keranjang"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

func (cs *cartServiceImpl) UpdateQuantity(req *domain.Cart) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// ambil data cart
	cart, err := cs.cartRepo.ByID(req.ID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if cart.VariasiID != "" {
		response := cs.UpdateQuantityVariant(cart.VariasiID, req)
		return response
	} else {
		response := cs.UpdateQuantityNonVariant(cart.ProdukID, req)
		return response
	}
}

func (cs *cartServiceImpl) Delete(req *domain.Cart) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	if err := cs.cartRepo.Delete(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat menghapus data cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// panggil repository
	res.Message = "quantitas item cart berhasil dihapus"
	res.Status = http.StatusOK
	res.Success = true

	return &res
}

func (cs *cartServiceImpl) Count(userID int) *domain.Response {
	// deklarasi variable
	res := domain.Response{}

	// pemanggilan repository
	count, err := cs.cartRepo.Count(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat menghitung jumlah carts"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Data = count
	res.Message = "carts Anda berhasil dihitung"
	res.Status = 200
	return &res
}

// DILUAR ABTRAKSI DOMAIN SERVICE
func (cs *cartServiceImpl) UpdateQuantityVariant(variantID string, req *domain.Cart) *domain.Response {
	res := domain.Response{}

	variasi, err := cs.produkVariasiRepository.ByID(variantID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data variasi"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if variasi.Stok == 0 {
		res.Message = "stok barang sudah habis"
		res.Status = http.StatusOK
		return &res
	}

	if variasi.Stok == req.Quantity {
		res.Message = "operasi update tidak dilakukan"
		res.Status = http.StatusOK
		return &res
	}

	if isNegative := variasi.Stok - req.Quantity; isNegative < 0 {
		res.Message = "stok barang tidak memenuhi jumlah permintaan Anda"
		res.Status = http.StatusOK
		return &res
	}

	if _, err := cs.cartRepo.UpdateQuantityVariant(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat mengupdate quantitas cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "Quantitas berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

// DILUAR ABTRAKSI DOMAIN SERVICE
func (cs *cartServiceImpl) UpdateQuantityNonVariant(produkID string, req *domain.Cart) *domain.Response {
	res := domain.Response{}

	produk, err := cs.produkRepository.ByID(produkID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data produk"
		res.Status = http.StatusInternalServerError
		return &res
	}

	if produk.Stok == 0 {
		res.Message = "stok barang sudah habis"
		res.Status = http.StatusOK
		return &res
	}

	if produk.Stok == req.Quantity {
		res.Message = "operasi update tidak dilakukan"
		res.Status = http.StatusOK
		return &res
	}

	if isSufficient := produk.Stok - req.Quantity; isSufficient < 0 {
		res.Message = "stok barang tidak memenuhi jumlah permintaan Anda"
		res.Status = http.StatusOK
		return &res
	}

	if _, err := cs.cartRepo.UpdateQuantityVariant(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat mengupdate quantitas cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	res.Message = "Quantitas berhasil diupdate"
	res.Status = http.StatusOK
	res.Success = true
	return &res
}

// DILUAR ABTRAKSI DOMAIN SERVICE
func (cs *cartServiceImpl) isStockInsuficient(produk *domain.ProdukDetailed, req *domain.Cart) bool {
	if produk.IsHasVariant {
		// check stok jika punya varian
		variant := produk.Variasi
		for i := 0; i < len(variant); i++ {
			if variant[i].ID == req.VariasiID {
				if variant[i].Stok < req.Quantity {
					fmt.Println("stock is insufficient")
					return true
				}
			}
		}
	} else {
		if produk.Stok < req.Quantity {
			fmt.Println("stock is insufficient")
			return true
		}
	}

	return false
}
