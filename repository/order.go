package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type orderRepositoryImpl struct {
	db *sql.DB
}

func OrderRepository(database *sql.DB) domain.OrderRepository {
	return &orderRepositoryImpl{db: database}
}

// // khusus Admin
// func (or *orderRepositoryImpl) Read() ([]domain.Order, error)

// // khusus umum
// func (or *orderRepositoryImpl) GetByID(id int) (*domain.Order, error)

// func (or *orderRepositoryImpl) GetByUserID(UserID int) ([]domain.Order, error)

func (or *orderRepositoryImpl) Create(req *domain.Order) (*domain.Order, error) {
	c := context.Background()
	tx, err := or.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO order(id, user_id, alamat_id, sub_total, biaya_kirim, invoice_no, tgl_transaksi, final_total) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	if _, err := stmt.ExecContext(c, req.ID, req.UserID, req.AlamatID, req.SubTotal, req.BiayaKirim, req.InvoiceNo, req.TglTransaksi, req.FinalTotal); err != nil {
		return nil, err
	}

	return req, nil
}
