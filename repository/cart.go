package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type cartRepositoryImpl struct {
	db *sql.DB
}

func CartRepository(database *sql.DB) domain.CartRepository {
	return &cartRepositoryImpl{db: database}
}

func (c *cartRepositoryImpl) Create(req *domain.Cart) (*domain.Cart, error) {
	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO carts (user_id, produk_id, variasi_id, quantity) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	inserted, err := stmt.ExecContext(context.Background(), req.UserID, req.ProdukID, req.VariasiID, req.Quantity)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := inserted.LastInsertId()

	req.ID = int(id)

	tx.Commit()

	return req, nil
}

// func (c *cartRepositoryImpl) Read(userID int) ([]domain.Cart, error)

func (c *cartRepositoryImpl) UpdateQuantity(req *domain.CartData) (*domain.CartData, error) {
	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE carts SET quantity = ? WHERE user_id = ? AND produk_id = ? AND variasi_id = ?")
	if err != nil {
		return nil, err
	}

	if _, err := stmt.ExecContext(context.Background(), req.Quantity, req.UserID, req.ProdukID, req.VariasiID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (c *cartRepositoryImpl) Delete(req *domain.CartData) error {
	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM carts WHERE user_id = ? AND produk_id = ? AND variasi_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(context.Background(), req.UserID, req.ProdukID, req.VariasiID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
