package services

import (
	"fmt"
	"net/http"

	"github.com/smokers10/go-diagem.git/domain"
)

type checkoutServiceImpl struct {
	cartRepo      domain.CartRepository
	orderRepo     domain.OrderRepository
	orderItemRepo domain.OrderItemRepository
}

func CheckoutService(cart *domain.CartRepository, order *domain.OrderRepository, orderItem *domain.OrderItemRepository) domain.CheckoutService {
	return &checkoutServiceImpl{
		cartRepo:      *cart,
		orderRepo:     *order,
		orderItemRepo: *orderItem,
	}
}

func (cs *checkoutServiceImpl) Checkout(req *domain.Checkout) *domain.Response {
	// deklarasi variable
	res := domain.Response{}
	var subtotal int

	// get cart
	carts, err := cs.cartRepo.Read(req.UserID)
	if err != nil {
		fmt.Println(err)
		res.Message = "error saat mengambil data cart"
		res.Status = http.StatusInternalServerError
		return &res
	}

	// check cart
	for i := 0; i < len(carts); i++ {
		subtotal += carts[i].SubTotal
	}

	fmt.Println(subtotal)

	res.Message = "on the test"
	res.Status = http.StatusOK
	return &res
}
