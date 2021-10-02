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
	statement, err := m.db.Prepare("INSERT INTO mitra (nama, username, kontak, alamat, email, password) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.ExecContext(context.Background(), req.Nama, req.Username, req.Kontak, req.Alamat, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	insertedID, _ := result.LastInsertId()
	req.ID = int(insertedID)

	return req, nil
}

func (m *mitraRepositoryImpl) Read() ([]domain.Mitra, error) {
	result := []domain.Mitra{}
	statement, err := m.db.Prepare("SELECT * FROM mitra")
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
		rows.Scan(&row.ID, &row.Alamat, &row.Email, &row.Kontak, &row.Nama, &row.Password, &row.Username)
		result = append(result, row)
	}

	return result, nil
}

func (m *mitraRepositoryImpl) ByID(id int) (*domain.Mitra, error) {
	result := domain.Mitra{}
	statement, err := m.db.Prepare("SELECT * FROM mitra WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), id)
	row.Scan(&result.ID, &result.Alamat, &result.Email, &result.Kontak, &result.Nama, &result.Password, &result.Username)

	return &result, nil
}

func (m *mitraRepositoryImpl) ByEmail(email string) (*domain.Mitra, error) {
	result := domain.Mitra{}
	statement, err := m.db.Prepare("SELECT * FROM mitra WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), email)
	row.Scan(&result.ID, &result.Alamat, &result.Email, &result.Kontak, &result.Nama, &result.Password, &result.Username)

	return &result, nil
}

func (m *mitraRepositoryImpl) ByUsername(username string) (*domain.Mitra, error) {
	result := domain.Mitra{}
	statement, err := m.db.Prepare("SELECT * FROM mitra WHERE username = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), username)
	row.Scan(&result.ID, &result.Alamat, &result.Email, &result.Kontak, &result.Nama, &result.Password, &result.Username)

	return &result, nil
}

func (m *mitraRepositoryImpl) Update(req *domain.Mitra) (*domain.Mitra, error) {
	statement, err := m.db.Prepare("UPDATE mitra SET nama = ?, kontak = ?, alamat = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Nama, req.Kontak, req.Alamat, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (m *mitraRepositoryImpl) UpdatePassword(id int, password string) error {
	statement, err := m.db.Prepare("UPDATE mitra SET password = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), password, id); err != nil {
		return err
	}

	return nil
}

func (m *mitraRepositoryImpl) UpdateEmail(id int, email string) error {
	statement, err := m.db.Prepare("UPDATE mitra SET email = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), email, id); err != nil {
		return err
	}

	return nil
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
