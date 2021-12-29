package repository

import (
	"database/sql"
	"time"

	"github.com/smokers10/go-diagem.git/domain"
)

type orderBayarRepository struct {
	db *sql.DB
}

func OrderBayarRepository(database *sql.DB) domain.OrderBayarRepository {
	return &orderBayarRepository{db: database}
}

func (ob *orderBayarRepository) Create(req *domain.OrderBayar) error {
	stmt, _ := ob.db.Prepare("INSERT INTO order_bayar(order_checkout_id, jumlah, token, redirect_url) VALUES (?, ?, ?, ?)")
	if _, err := stmt.Exec(req.OrderCheckoutID, req.Jumlah, req.Token, req.RedirectURL); err != nil {
		return err
	}
	return nil
}

func (ob *orderBayarRepository) UpdateStatus(token string, status string) error {
	now := time.Now().Local()
	stmt, _ := ob.db.Prepare("UPDATE order_bayar SET status = ?, tgl_bayar = ? WHERE token = ?")
	if _, err := stmt.Exec(status, now, token); err != nil {
		return err
	}
	return nil
}

func (ob *orderBayarRepository) ByToken(token string) (*domain.OrderBayar, error) {
	result := domain.OrderBayar{}
	stmt, _ := ob.db.Prepare("SELECT id, order_checkout_id, jumlah, status, token, redirect_url, tgl_bayar WHERE token = ?")
	stmt.QueryRow(token).Scan(&result.ID, &result.OrderCheckoutID, &result.Jumlah, &result.Status, &result.Token, &result.RedirectURL, &result.TGLBayar)
	return &result, nil
}

func (ob *orderBayarRepository) ByOrderID(OrderID string) (*domain.OrderBayar, error) {
	result := domain.OrderBayar{}
	stmt, _ := ob.db.Prepare("SELECT id, order_checkout_id, jumlah, status, token, redirect_url, tgl_bayar WHERE token = ?")
	stmt.QueryRow(OrderID).Scan(&result.ID, &result.OrderCheckoutID, &result.Jumlah, &result.Status, &result.Token, &result.RedirectURL, &result.TGLBayar)
	return &result, nil
}
