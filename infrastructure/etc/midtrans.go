package etc

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	clientMidtransKey = "Mid-client-qSHcGy_zZfUArqlP"
	serverMidtransKey = "Mid-server-kaHeYZv7SmquzLF1slqQm859"
)

var enablePayments = []string{"bca_klikpay", "bca_va", "bni_va", "bri_va"}

type Midtrans struct {
	TransactionDetail TransactionDetail `json:"transaction_detail"`
	CustomerDetail    CustomerDetail    `json:"customer_detail"`
	ItemDetail        []ItemDetail      `json:"item_detail"`
	EnabledPayments   []string          `json:"enable_payments"`
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
	Price    string `json:"price"`
	Quantity string `json:"Quantity"`
	Name     string `json:"name"`
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

func (m *Midtrans) AuthString() string {
	preformatedKey := fmt.Sprintf("%s:", serverMidtransKey)
	return base64.StdEncoding.EncodeToString([]byte(preformatedKey))
}

func (m *Midtrans) Transaction() (string, error) {
	endpoint := fmt.Sprintf("https://%s/snap/v1/transactions", m.Domain())
	m.EnabledPayments = enablePayments
	payload, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", m.AuthString())

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
