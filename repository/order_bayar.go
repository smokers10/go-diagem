package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type orderBayarRepositoryImpl struct {
	db *sql.DB
}

var c = context.Background()

func OrderBayarRepository(database *sql.DB) domain.OrderBayarRepository {
	return &orderBayarRepositoryImpl{db: database}
}

func (ob *orderBayarRepositoryImpl) Create(req *domain.OrderBayar) (*domain.OrderBayar, error) {
	tx, err := ob.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO order_bayar (order_id, rekening_id, method, layanan, jumlah, status, tgl_bayar) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	inserted, err := stmt.ExecContext(c, req.OrderID, req.RekeningID, req.Method, req.Layanan, req.Jumlah, req.Status, req.Tglbayar)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := inserted.LastInsertId()

	req.ID = int(id)

	tx.Commit()

	return req, nil
}

func (ob *orderBayarRepositoryImpl) UpdateStatus(orderID string, status string) error {
	tx, err := ob.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE order_bayar SET status = ? WHERE order_id = ?")
	if err != nil {
		return err
	}

	if _, err := stmt.ExecContext(c, status, orderID); err != nil {
		tx.Rollback()
	}

	tx.Commit()

	return nil
}
