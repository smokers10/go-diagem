package domain

import "github.com/smokers10/go-diagem.git/infrastructure/etc"

type Checkout struct {
	UserID   int `json:"user_id" form:"user_id"`
	AlamatID int `json:"alamat_id" form:"alamat_id"`
}

type CheckoutService interface {
	Checkout(req *Checkout) *Response
	Ongkir(req *etc.RajaOngkirReqBody) *Response
}
