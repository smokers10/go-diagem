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

func (c *cartRepositoryImpl) Read(userID int) ([]domain.CartDetail, error) {
	result := []domain.CartDetail{}
	stmt, err := c.db.Prepare(`SELECT carts.produk_id, carts.variasi_id, carts.quantity, produk.nama, produk.id, produk_variasi.variant, produk_variasi.harga, 
	carts.quantity * produk_variasi.harga sub_total FROM carts
	JOIN produk ON carts.produk_id = produk.id
	JOIN produk_variasi ON carts.variasi_id = produk_variasi.id
	WHERE carts.user_id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.CartDetail{}
		rows.Scan(&row.ProdukID, &row.VariasiID, &row.Quantity, &row.Produk.Nama, &row.Produk.ID, &row.Variasi.Variant, &row.Variasi.Harga, &row.SubTotal)
		result = append(result, row)
	}

	return result, nil
}

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
