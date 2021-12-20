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

	defer stmt.Close()

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

func (c *cartRepositoryImpl) ByID(cartID int) (*domain.Cart, error) {
	result := domain.Cart{}
	stmt, err := c.db.Prepare(`SELECT id, user_id, produk_id, variasi_id, quantity FROM carts WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRowContext(context.Background(), cartID)
	row.Scan(&result.ID, &result.UserID, &result.ProdukID, &result.VariasiID, &result.Quantity)

	return &result, nil
}

func (c *cartRepositoryImpl) Read(userID int) ([]domain.CartDetail, error) {
	fotoProdukRepository := ProdukFotoRepository(c.db)
	produkRepository := ProdukRepository(c.db)
	variasiRepository := ProdukVariasiRepository(c.db)

	result := []domain.CartDetail{}
	stmt, err := c.db.Prepare(`SELECT carts.id, carts.user_id, carts.produk_id, carts.variasi_id, carts.quantity FROM carts WHERE carts.user_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		// scan data cart
		row := domain.CartDetail{}
		rows.Scan(&row.ID, &row.UserID, &row.ProdukID, &row.VariasiID, &row.Quantity)

		//1. ambil data foto utama produk
		fotoUtama, err := fotoProdukRepository.GetUtamaOnly(row.ProdukID)
		if err != nil {
			return nil, err
		}

		//2. ambil data produk
		produk, err := produkRepository.ByIDSimplified(row.ProdukID)
		if err != nil {
			return nil, err
		}

		//3. ambil data variasi (jika ada)
		if produk.IsHasVariant {
			variants, err := variasiRepository.ByID(row.VariasiID)
			if err != nil {
				return nil, err
			}

			row.Variasi = *variants

			row.SubTotal = int(variants.Harga) * row.Quantity
		} else {
			row.SubTotal = produk.Harga * row.Quantity
		}

		//4. assign data produk dan foto utama ke struct row
		row.Produk = *produk
		row.ProdukSingleFoto = *fotoUtama

		result = append(result, row)
	}

	return result, nil
}

func (c *cartRepositoryImpl) UpdateQuantityVariant(req *domain.Cart) (*domain.Cart, error) {
	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE carts SET quantity = ? WHERE user_id = ? AND id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(context.Background(), req.Quantity, req.UserID, req.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (c *cartRepositoryImpl) Delete(req *domain.Cart) error {
	tx, err := c.db.BeginTx(context.Background(), nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM carts WHERE id = ? AND user_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	if _, err := stmt.ExecContext(context.Background(), req.ID, req.UserID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (c *cartRepositoryImpl) Count(userID int) (*domain.CartCount, error) {
	result := domain.CartCount{}

	stmt, err := c.db.Prepare("SELECT COUNT(id) as total_carts FROM carts WHERE user_id = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRowContext(context.Background(), userID).Scan(&result.TotalCarts)

	return &result, nil
}
