package etc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	server_key = "SB-Mid-server-72w4BpkX9z-gC8HCHmbu84dZ"
	password   = "Suckitjiwa666"
)

var enablePayments = []string{"bca_klikpay", "bca_va", "bni_va", "bri_va"}

type Midtrans struct {
	TransactionDetail TransactionDetail `json:"transaction_details"`
	CustomerDetail    CustomerDetail    `json:"customer_details"`
	ItemDetail        []ItemDetail      `json:"item_details"`
	EnabledPayments   []string          `json:"enabled_payments"`
	BCA               BCA               `json:"bca_va"`
	BNI               BNI               `json:"bni_va"`
	BRI               BRI               `json:"bri_va"`
}

type BCA struct {
	VANumber         string `json:"va_number"`
	SubCompanyNumber string `json:"sub_company_number"`
}

type BNI struct {
	VANumber string `json:"va_number"`
}

type BRI struct {
	VANumber string `json:"va_number"`
}

type CustomerDetail struct {
	Email           string          `json:"email"`
	Phone           string          `json:"phone"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

type ShippingAddress struct {
	Phone      string `json:"phone"`
	FirstName  string `json:"first_name"`
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
}

type TransactionDetail struct {
	OrderID     string  `json:"order_id"`
	GrossAmount float32 `json:"gross_amount"`
}

type ItemDetail struct {
	ID       string `json:"id"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
	Name     string `json:"name"`
}

type midtransResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

func MidtransSnap() *Midtrans {
	return &Midtrans{}
}

func (m *Midtrans) Domain() string {
	production := os.Getenv("PRODUCTION_MODE")
	var domain string
	if production == "" || production == "development" || production == "local" {
		domain = "app.sandbox.midtrans.com"
	} else {
		domain = "app.midtrans.com"
	}
	return domain
}

func (m *Midtrans) Transaction() (*midtransResponse, error) {
	httpRes := midtransResponse{}
	endpoint := fmt.Sprintf("https://%s/snap/v1/transactions", m.Domain())
	m.EnabledPayments = enablePayments
	payload, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(server_key, "")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &httpRes)

	return &httpRes, nil
}
