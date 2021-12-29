package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type orderRepository struct {
	db *sql.DB
}

func OrderRepository(database *sql.DB) domain.OrderRepository {
	return &orderRepository{db: database}
}

func (or *orderRepository) Create(req *domain.Order) error {
	c := context.Background()
	tx, err := or.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO order_checkout (id, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no) VALUE(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(req.ID, req.UserID, req.AlamatID, req.Kurir, req.PaketKurir, req.Ongkir, req.InvoiceNo); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (or *orderRepository) UpdateStatus(orderID string, status string) error {
	c := context.Background()
	tx, err := or.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE order_checkout SET status = ? WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(status, orderID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (or *orderRepository) GetByID(orderID string) (*domain.OrderDetail, error) {
	result := domain.OrderDetail{}
	alamatRepo := AlamatRepository(or.db)
	userRepo := UserRepository(or.db)
	orderItemRepo := OrderItemRepository(or.db)

	stmt, err := or.db.Prepare("SELECT id, status, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no, tgl_order FROM order_checkout WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	stmt.QueryRow(orderID).Scan(&result.ID, &result.Status, &result.UserID, &result.AlamatID, &result.Kurir, &result.PaketKurir, &result.Ongkir, &result.InvoiceNo, &result.TglOrder)

	alamat, err := alamatRepo.ByID(result.AlamatID)
	if err != nil {
		return nil, err
	}

	user, err := userRepo.ByID(result.UserID)
	if err != nil {
		return nil, err
	}

	orderItem, err := orderItemRepo.ByOrderID(result.ID)
	if err != nil {
		return nil, err
	}

	result.Alamat = *alamat
	result.User = *user
	result.OrderItem = orderItem

	return &result, nil
}

func (or *orderRepository) GetByUserID(userID int) ([]domain.OrderDetail, error) {
	orderBayarRepo := OrderBayarRepository(or.db)
	result := []domain.OrderDetail{}

	stmt, err := or.db.Prepare("SELECT id, status, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no, tgl_order FROM order_checkout WHERE user_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.OrderDetail{}
		rows.Scan(&row.ID, &row.Status, &row.UserID, &row.AlamatID, &row.Kurir, &row.PaketKurir, &row.Ongkir, &row.InvoiceNo, &row.TglOrder)
		ob, err := orderBayarRepo.ByOrderID(row.ID)
		if err != nil {
			return nil, err
		}
		row.OrderBayar = *ob
		result = append(result, row)
	}

	return result, nil
}

func (or *orderRepository) Read() ([]domain.OrderDetail, error) {
	result := []domain.OrderDetail{}
	userRepo := UserRepository(or.db)

	stmt, err := or.db.Prepare("SELECT id, status, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no, tgl_order FROM order_checkout")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.OrderDetail{}
		rows.Scan(&row.ID, &row.Status, &row.UserID, &row.AlamatID, &row.Kurir, &row.PaketKurir, &row.Ongkir, &row.InvoiceNo, &row.TglOrder)

		// get user
		user, err := userRepo.ByID(row.UserID)
		if err != nil {
			return nil, err
		}

		row.User = *user

		result = append(result, row)
	}

	return result, nil
}
