package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type orderItemRepositoryImpl struct {
	db *sql.DB
}

func OrderItemRepository(database *sql.DB) domain.OrderItemRepository {
	return &orderItemRepositoryImpl{db: database}
}

func (oi *orderItemRepositoryImpl) Create(req *domain.OrderItem) (*domain.OrderItem, error) {
	c := context.Background()
	tx, err := oi.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO order_item (order_id, produk_id, variasi_id, harga, quantity, sub_total) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	inserted, err := stmt.ExecContext(c, req.OrderID, req.ProdukID, req.VariasiID, req.Harga, req.Quantity, req.SubTotal)
	if err != nil {
		return nil, err
	}

	id, _ := inserted.LastInsertId()

	req.ID = int(id)

	return req, nil
}
