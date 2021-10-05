package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type rekeningRepositoryImpl struct {
	db *sql.DB
}

func RekeningRepository(database *sql.DB) domain.RekeningRepository {
	return &rekeningRepositoryImpl{db: database}
}

func (r *rekeningRepositoryImpl) Create(req *domain.Rekening) (*domain.Rekening, error) {
	statement, err := r.db.Prepare("INSERT INTO rekening (bank_id, rekening_no, nama) VALUES (?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	inserted, err := statement.ExecContext(context.Background(), req.BankID, req.RekeningNo, req.Nama)
	if err != nil {
		return nil, err
	}

	id, _ := inserted.LastInsertId()

	req.ID = int(id)

	return req, nil
}

func (r *rekeningRepositoryImpl) Read() ([]domain.Rekening, error) {
	result := []domain.Rekening{}
	statement, err := r.db.Prepare(`SELECT rekening.id, rekening.rekening_no, rekening.nama, bank.name, bank.code, bank.icon FROM rekening
	JOIN bank ON rekening.bank_id = bank.id
	`)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Rekening{}
		rows.Scan(&row.ID, &row.RekeningNo, &row.Nama, &row.Bank.Name, &row.Bank.Code, &row.Bank.Icon)
		result = append(result, row)
	}

	return result, nil
}

func (r *rekeningRepositoryImpl) Update(req *domain.Rekening) (*domain.Rekening, error) {
	statement, err := r.db.Prepare("UPDATE rekening SET bank_id = ?, rekening_no = ?, nama = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.BankID, req.RekeningNo, req.Nama, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (r *rekeningRepositoryImpl) Delete(id int) error {
	statement, err := r.db.Prepare("DELETE FROM rekening WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}
