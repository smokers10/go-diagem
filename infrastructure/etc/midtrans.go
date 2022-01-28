package etc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/smokers10/go-diagem.git/infrastructure/config"
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
	OrderID     string `json:"order_id"`
	GrossAmount int    `json:"gross_amount"`
}

type ItemDetail struct {
	ID       string `json:"id"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Name     string `json:"name"`
}

type midtransResponse struct {
	Token       string `json:"token"`
	RedirectURL string `json:"redirect_url"`
}

type midtransErrorResponse struct {
	ErrorMessages []string `json:"error_messages,omitempty"`
}

type midtransStatusResponse struct {
	TransactionStatus string `json:"transaction_status"`
	TransactionTime   string `json:"transaction_time"`
	StatusMessage     string `json:"status_message"`
	StatusCode        string `json:"status_code"`
	OrderID           string `json:"order_id"`
}

func MidtransSnap() *Midtrans {
	return &Midtrans{}
}

func (m *Midtrans) Domain() string {
	var domain string
	if mode := config.ReadConfig().Application.APP_Production_Mode; mode == "local" || mode == "development" {
		domain = "app.sandbox.midtrans.com"
	} else {
		domain = "app.midtrans.com"
	}
	return domain
}

func (m *Midtrans) APIDomain() string {
	var domain string
	if mode := config.ReadConfig().Application.APP_Production_Mode; mode == "local" || mode == "development" {
		domain = "api.sandbox.midtrans.com"
	} else {
		domain = "api.midtrans.com"
	}
	return domain
}

func (m *Midtrans) Transaction() (*midtransResponse, error) {
	httpRes := midtransResponse{}
	httpErrorRes := midtransErrorResponse{}
	endpoint := fmt.Sprintf("https://%s/snap/v1/transactions", m.Domain())
	m.EnabledPayments = enablePayments
	payload, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(config.ReadConfig().ETC.Midtrans_Server_key, "")

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if status := response.Status; status == "400 Bad Request" {
		json.Unmarshal(body, &httpErrorRes)
		fmt.Println(httpErrorRes)
		fmt.Println()
		fmt.Println(m.ItemDetail)
		fmt.Println()
		fmt.Println(m.TransactionDetail.GrossAmount)
	}

	json.Unmarshal(body, &httpRes)

	return &httpRes, nil
}

func (m *Midtrans) CheckStatus(orderID string) (*midtransStatusResponse, error) {
	// deklarasi variable penting
	var httpResponse midtransStatusResponse
	endpoint := fmt.Sprintf("https://%s/v2/%s/status", m.APIDomain(), orderID)

	// make request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	// set request header
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(config.ReadConfig().ETC.Midtrans_Server_key, "")

	// kirim request and handle response dari http request
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// read response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// parse response ke struct
	json.Unmarshal(body, &httpResponse)

	return &httpResponse, nil
}
