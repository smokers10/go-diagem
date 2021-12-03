package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type mitraRepositoryImpl struct {
	db *sql.DB
}

func MitraRepository(database *sql.DB) domain.MitraRepository {
	return &mitraRepositoryImpl{db: database}
}

func (m *mitraRepositoryImpl) Create(req *domain.Mitra) (*domain.Mitra, error) {
	statement, err := m.db.Prepare("INSERT INTO mitra (nama, kontak, alamat, email) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.ExecContext(context.Background(), req.Nama, req.Kontak, req.Alamat, req.Email)
	if err != nil {
		return nil, err
	}

	insertedID, _ := result.LastInsertId()
	req.ID = int(insertedID)

	return req, nil
}

func (m *mitraRepositoryImpl) Read() ([]domain.Mitra, error) {
	result := []domain.Mitra{}
	statement, err := m.db.Prepare("SELECT id, nama, kontak, email, alamat FROM mitra")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Mitra{}
		rows.Scan(&row.ID, &row.Nama, &row.Kontak, &row.Email, &row.Alamat)
		result = append(result, row)
	}

	return result, nil
}

func (m *mitraRepositoryImpl) ByID(id int) (*domain.Mitra, error) {
	result := domain.Mitra{}
	stmt, err := m.db.Prepare("SELECT id, nama, email, alamat, kontak FROM mitra WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRowContext(context.Background(), id).Scan(&result.ID, &result.Nama, &result.Email, &result.Alamat, &result.Kontak)

	return &result, nil
}

func (m *mitraRepositoryImpl) ByEmail(email string) (*domain.Mitra, error) {
	result := domain.Mitra{}
	stmt, err := m.db.Prepare("SELECT id, nama, email, alamat, kontak FROM mitra WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRowContext(context.Background(), email).Scan(&result.ID, &result.Nama, &result.Email, &result.Alamat, &result.Kontak)

	return &result, nil
}

func (m *mitraRepositoryImpl) Update(req *domain.Mitra) (*domain.Mitra, error) {
	statement, err := m.db.Prepare("UPDATE mitra SET nama = ?, kontak = ?, alamat = ?, email = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Nama, req.Kontak, req.Alamat, req.Email, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (m *mitraRepositoryImpl) Delete(id int) error {
	statement, err := m.db.Prepare("DELETE FROM mitra WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}
