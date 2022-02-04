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

func (or *orderRepository) Create(req *domain.Order, tx *sql.Tx) error {
	stmt, err := tx.Prepare("INSERT INTO order_checkout (id, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no) VALUE(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err := stmt.Exec(req.ID, req.UserID, req.AlamatID, req.Kurir, req.PaketKurir, req.Ongkir, req.InvoiceNo); err != nil {
		return err
	}

	return nil
}

func (or *orderRepository) UpdateSR(orderID string, status string, resi string) error {
	c := context.Background()
	tx, err := or.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE order_checkout SET status = ?, no_resi = ? WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

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
	paymentRepo := OrderBayarRepository(or.db)

	stmt, err := or.db.Prepare("SELECT id, status, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no, tgl_order FROM order_checkout WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(orderID).Scan(&result.ID, &result.Status, &result.UserID, &result.AlamatID, &result.Kurir, &result.PaketKurir, &result.Ongkir, &result.InvoiceNo, &result.TglOrder)

	// ambil alamat
	alamat, err := alamatRepo.ByID(result.AlamatID)
	if err != nil {
		return nil, err
	}

	// ambil user
	user, err := userRepo.ByID(result.UserID)
	if err != nil {
		return nil, err
	}

	// ambil item pesanan
	items, err := orderItemRepo.ByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	// payment repo
	payment, err := paymentRepo.ByOrderID(orderID)
	if err != nil {
		return nil, err
	}

	// assign
	result.Alamat = *alamat
	result.User = *user
	result.OrderItem = items
	result.OrderBayar = *payment

	return &result, nil
}

func (or *orderRepository) GetByUserID(userID int) ([]domain.OrderDetail, error) {
	result := []domain.OrderDetail{}
	orderPaymentRepo := OrderBayarRepository(or.db)

	stmt, err := or.db.Prepare("SELECT id, status, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no, tgl_order FROM order_checkout WHERE user_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(userID)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for rows.Next() {
		row := domain.OrderDetail{}
		rows.Scan(&row.ID, &row.Status, &row.UserID, &row.AlamatID, &row.Kurir, &row.PaketKurir, &row.Ongkir, &row.InvoiceNo, &row.TglOrder)
		if err != nil {
			return nil, err
		}

		// get order payment
		payment, err := orderPaymentRepo.ByOrderID(row.ID)
		if err != nil {
			return nil, err
		}

		row.OrderBayar = *payment

		result = append(result, row)
	}

	return result, nil
}

func (or *orderRepository) Read() ([]domain.OrderDetail, error) {
	result := []domain.OrderDetail{}
	userRepo := UserRepository(or.db)
	orderPaymentRepo := OrderBayarRepository(or.db)

	stmt, err := or.db.Prepare("SELECT id, status, user_id, alamat_id, kurir, paket_kurir, ongkir, invoice_no, tgl_order FROM order_checkout")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.OrderDetail{}
		rows.Scan(&row.ID, &row.Status, &row.UserID, &row.AlamatID, &row.Kurir, &row.PaketKurir, &row.Ongkir, &row.InvoiceNo, &row.TglOrder)

		// get order payment
		payment, err := orderPaymentRepo.ByOrderID(row.ID)
		if err != nil {
			return nil, err
		}

		// get user
		user, err := userRepo.ByID(row.UserID)
		if err != nil {
			return nil, err
		}

		row.User = *user
		row.OrderBayar = *payment

		result = append(result, row)
	}

	return result, nil
}

func (or *orderRepository) GetSQLInstance() *sql.DB {
	return or.db
}
