package domain

type Checkout struct{}

type CheckoutService interface {
	Create() *Response
}

type CheckoutRepository interface{}
