package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/smokers10/go-diagem.git/domain"
)

type orderService struct {
	orderRepository         domain.OrderRepository
	produkRepository        domain.ProdukRepository
	produkVariasiRepository domain.ProdukRepository
	cartRepository          domain.CartRepository
	orderItemRepository     domain.OrderItemRepository
}

func OrderService(order *domain.OrderRepository, produk *domain.ProdukRepository, variasi *domain.ProdukVariasiRepository, cart *domain.CartRepository, item *domain.OrderItemRepository) domain.OrderService {
	return &orderService{
		orderRepository:         *order,
		produkRepository:        *produk,
		produkVariasiRepository: *produk,
		cartRepository:          *cart,
		orderItemRepository:     *item,
	}
}

func (osc *orderService) Create(req *domain.Order) *domain.Response {
	res := domain.Response{}

	// ambil data cart
	carts, err := osc.cartRepository.Read(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data order"
		res.Status = 500
		return &res
	}

	// check stok pada tiap cart item
	for i := 0; i < len(carts); i++ {
		quantity := carts[i].Quantity
		if carts[i].Produk.IsHasVariant {
			variantStok := carts[i].Variasi.Stok

			if osc.IsStokSufficient(variantStok, quantity) {
				res.Message = fmt.Sprintf("Produk %s Variasi %s Sudah Habis", carts[i].Produk.Nama, carts[i].Variasi.Variant)
				res.Status = 200
				return &res
			}
		} else {
			produkStok := carts[i].Produk.Stok

			if osc.IsStokSufficient(produkStok, quantity) {
				res.Message = fmt.Sprintf("Produk %s Sudah Habis", carts[i].Produk.Nama)
				res.Status = 200
				return &res
			}
		}
	}

	// generate id
	id, _ := uuid.NewRandom()
	req.ID = id.String()

	// generate invoice
	year, month, day := time.Now().Date()
	unixTime := time.Now().Unix()
	invoiceNo := fmt.Sprintf("DGM-%d%d%d%d", unixTime, day, int(month), year)
	req.InvoiceNo = invoiceNo

	// buat order
	if err := osc.orderRepository.Create(req); err != nil {
		fmt.Println(err)
		res.Message = "error saat menyimpan data order"
		res.Status = 500
		return &res
	}

	// pindahkan data cart ke order item
	for _, cart := range carts {
		orderItem := domain.OrderItem{
			OrderCheckoutID: id.String(),
			VariasiID:       cart.VariasiID,
			ProdukID:        cart.ProdukID,
			Quantity:        cart.Quantity,
		}

		if err := osc.orderItemRepository.Create(&orderItem); err != nil {
			fmt.Println(err)
			res.Message = "error saat menyimpan data order item"
			res.Status = 500
			return &res
		}

		osc.cartRepository.Delete(&domain.Cart{ID: cart.ID, UserID: cart.UserID})
	}

	// write response sukses
	res.Message = "checkout berhasil"
	res.Status = 200
	res.Success = true
	return &res
}

func (osc *orderService) ReadAll() *domain.Response {
	res := domain.Response{}

	order, err := osc.orderRepository.Read()
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data order"
		res.Status = 500
		return &res
	}

	res.Data = order
	res.Message = "data order berhasil diambil"
	res.Status = 200
	return &res
}

func (osc *orderService) Detail(orderID string) *domain.Response {
	res := domain.Response{}

	order, err := osc.orderRepository.GetByID(orderID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil detail order"
		res.Status = 500
		return &res
	}

	res.Data = order
	res.Message = "detail order berhasil diambil"
	res.Status = 200
	return &res
}

func (osc *orderService) UpdateStatus(orderID string, status string) *domain.Response {
	res := domain.Response{}

	if err := osc.orderRepository.UpdateStatus(orderID, status); err != nil {
		fmt.Println(err)
		res.Message = "error saat update status order"
		res.Status = 500
		return &res
	}

	res.Message = "status order berhasil diubah"
	res.Status = 200
	return &res
}

func (osc *orderService) ReadByUser(userID int) *domain.Response {
	res := domain.Response{}

	order, err := osc.orderRepository.GetByUserID(userID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data order"
		res.Status = 500
		return &res
	}

	res.Data = order
	res.Message = "data order berhasil diambil"
	res.Status = 200
	return &res
}

// OUT OF ORDER SERVICE SCOPE
func (osc *orderService) IsStokSufficient(stok int, quantity int) bool {
	check := stok - quantity
	return check < 0
}
