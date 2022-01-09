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

func (ob *orderBayarRepository) Create(req *domain.OrderBayar, tx *sql.Tx) error {
	stmt, _ := tx.Prepare("INSERT INTO order_bayar(order_checkout_id, jumlah, token, redirect_url) VALUES (?, ?, ?, ?)")
	defer stmt.Close()
	if _, err := stmt.Exec(req.OrderCheckoutID, req.Jumlah, req.Token, req.RedirectURL); err != nil {
		return err
	}
	return nil
}

func (ob *orderBayarRepository) UpdateStatus(token string, status string) error {
	now := time.Now().Local()
	stmt, _ := ob.db.Prepare("UPDATE order_bayar SET status = ?, tgl_bayar = ? WHERE token = ?")
	defer stmt.Close()
	if _, err := stmt.Exec(status, now, token); err != nil {
		return err
	}
	return nil
}

func (ob *orderBayarRepository) ByToken(token string) (*domain.OrderBayar, error) {
	result := domain.OrderBayar{}
	stmt, _ := ob.db.Prepare("SELECT id, order_checkout_id, jumlah, status, token, redirect_url, tgl_bayar FROM order_bayar WHERE token = ? LIMIT 1")
	defer stmt.Close()

	row := stmt.QueryRow(token)
	if err := row.Scan(&result.ID, &result.OrderCheckoutID, &result.Jumlah, &result.Status, &result.Token, &result.RedirectURL, &result.TGLBayar); err != nil {
		return nil, err
	}

	return &result, nil
}

func (ob *orderBayarRepository) ByOrderID(OrderID string) (*domain.OrderBayar, error) {
	result := domain.OrderBayar{}
	stmt, _ := ob.db.Prepare("SELECT id, order_checkout_id, jumlah, status, token, redirect_url, tgl_bayar FROM order_bayar WHERE order_checkout_id = ? LIMIT 1")
	defer stmt.Close()

	row := stmt.QueryRow(OrderID)
	if err := row.Scan(&result.ID, &result.OrderCheckoutID, &result.Jumlah, &result.Status, &result.Token, &result.RedirectURL, &result.TGLBayar); err != nil {
		return nil, err
	}

	return &result, nil
}
