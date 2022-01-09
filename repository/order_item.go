package repository

import (
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type orderItemRepository struct {
	db *sql.DB
}

func OrderItemRepository(database *sql.DB) domain.OrderItemRepository {
	return &orderItemRepository{db: database}
}

func (oir *orderItemRepository) Create(req *domain.OrderItem, tx *sql.Tx) error {
	stmt, err := tx.Prepare(`INSERT INTO order_item (order_checkout_id, produk_id, variasi_id, quantity) VALUES(?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(req.OrderCheckoutID, req.ProdukID, req.VariasiID, req.Quantity); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (oir *orderItemRepository) ByOrderID(orderID string) ([]domain.OrderItemDetail, error) {
	result := []domain.OrderItemDetail{}
	produkRepository := ProdukRepository(oir.db)
	produkVariasiRepository := ProdukVariasiRepository(oir.db)
	stmt, err := oir.db.Prepare(`SELECT id, order_checkout_id, produk_id, variasi_id, quantity FROM order_item WHERE order_checkout_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(orderID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.OrderItemDetail{}

		rows.Scan(&row.ID, &row.OrderCheckoutID, &row.ProdukID, &row.VariasiID, &row.Quantity)

		// ambil data produk
		produk, err := produkRepository.ByID(row.ProdukID)
		if err != nil {
			return nil, err
		}
		row.Produk = *produk

		// ambil data variasi (jika ada)
		if produk.IsHasVariant {
			varian, err := produkVariasiRepository.ByID(row.VariasiID)
			if err != nil {
				return nil, err
			}
			row.Variasi = *varian
		}

		result = append(result, row)
	}

	return result, nil
}
