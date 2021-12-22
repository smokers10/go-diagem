package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
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

func (cs *checkoutServiceImpl) Ongkir(req *etc.RajaOngkirReqBody) *domain.Response {
	res := domain.Response{}
	url := "https://api.rajaongkir.com/starter/cost"
	payload := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s", req.Origin, req.Destination, req.Weight, req.Curier)
	reader := strings.NewReader(payload)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		res.Message = "error saat membuat request"
		res.Status = 500
		return &res
	}

	request.Header.Add("key", "your-api-key")
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		res.Message = "error saat do response"
		res.Status = 500
		return &res
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		res.Message = "error saat membaca response body"
		res.Status = 500
		return &res
	}

	res.Data = string(body)
	res.Message = "perhitungan ongkir selesai"
	res.Status = 200
	return &res
}
