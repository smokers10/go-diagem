package services

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/smokers10/go-diagem.git/domain"
	"github.com/smokers10/go-diagem.git/infrastructure/config"
	"github.com/smokers10/go-diagem.git/infrastructure/etc"
)

type checkoutServiceImpl struct {
	cartRepo domain.CartRepository
}

func CheckoutService(cart *domain.CartRepository) domain.CheckoutService {
	return &checkoutServiceImpl{
		cartRepo: *cart,
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
	payload := fmt.Sprintf("origin=%s&destination=%s&weight=%d&courier=%s", req.Origin, req.Destination, req.Weight, req.Courier)
	reader := strings.NewReader(payload)
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		res.Message = "error saat membuat request"
		res.Status = 500
		return &res
	}

	request.Header.Add("key", config.ReadConfig().Rajaongkir_API_Key)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(response)
		res.Message = "error saat do response, perikasi koneksi internet"
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
