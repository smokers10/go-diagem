package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type bankRepositoryImpl struct {
	db *sql.DB
}

func BankRepository(database *sql.DB) domain.BankRepository {
	return &bankRepositoryImpl{db: database}
}

func (br *bankRepositoryImpl) Update(req *domain.Bank) error {
	c := context.Background()
	tx, err := br.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE bank SET no_va = ?, status = ? WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.Exec(req.NoVa, req.Status, req.ID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (br *bankRepositoryImpl) GetActive() ([]domain.Bank, error) {
	result := []domain.Bank{}
	stmt, err := br.db.Prepare("SELECT id, no_va, vendor, status FROM bank WHERE status = true")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Bank{}
		rows.Scan(&row.ID, &row.NoVa, &row.Vendor, &row.Status)
		result = append(result, row)
	}

	return result, nil
}

func (br *bankRepositoryImpl) Read() ([]domain.Bank, error) {
	result := []domain.Bank{}
	stmt, err := br.db.Prepare("SELECT id, no_va, vendor, status FROM bank")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Bank{}
		rows.Scan(&row.ID, &row.NoVa, &row.Vendor, &row.Status)
		result = append(result, row)
	}

	return result, nil
}

func (br *bankRepositoryImpl) Segmented() (*domain.BankSegmented, error) {
	segmented := domain.BankSegmented{}
	stmt, err := br.db.Prepare("SELECT id, no_va, vendor, status FROM bank")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Bank{}
		rows.Scan(&row.ID, &row.NoVa, &row.Vendor, &row.Status)
		if row.Vendor == "bni_va" {
			segmented.BNI.VANumber = row.NoVa
		}

		if row.Vendor == "bri_va" {
			segmented.BRI.VANumber = row.NoVa
		}

		if row.Vendor == "bca_va" {
			segmented.BCA.SubCompanyNumber = "n/a"
			segmented.BCA.VANumber = row.NoVa
		}
	}

	return &segmented, nil
}
