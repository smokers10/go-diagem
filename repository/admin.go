package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type adminRepositoryImpl struct {
	db *sql.DB
}

func AdminRepository(database *sql.DB) domain.AdminRepository {
	return &adminRepositoryImpl{db: database}
}

func (ar *adminRepositoryImpl) GetByEmail(email string) (admin *domain.Admin, err error) {
	result := domain.Admin{}
	statement, err := ar.db.Prepare("SELECT id, nama, username, email, password, roles FROM admins WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), email)
	row.Scan(&result.ID, &result.Nama, &result.Username, &result.Email, &result.Password, &result.Roles)

	return &result, nil
}

func (ar *adminRepositoryImpl) GetByID(id int) (admin *domain.Admin, err error) {
	result := domain.Admin{}
	statement, err := ar.db.Prepare("SELECT id, nama, username, email, password, roles FROM admins WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), id)
	row.Scan(&result.ID, &result.Nama, &result.Username, &result.Email, &result.Password, &result.Roles)

	return &result, nil
}

func (ar *adminRepositoryImpl) Create(req *domain.Admin) (*domain.Admin, error) {
	c := context.Background()
	tx, err := ar.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO admins (nama, username, email, password, roles) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	inserted, err := stmt.ExecContext(c, req.Nama, req.Username, req.Email, req.Password, req.Roles)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	id, _ := inserted.LastInsertId()
	req.ID = int(id)

	tx.Commit()

	return req, nil
}

func (ar *adminRepositoryImpl) Update(req *domain.Admin) (*domain.Admin, error) {
	c := context.Background()
	tx, err := ar.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE admins SET nama = ?, username = ?, email = ?, roles = ?, updated_at = now() WHERE ID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.Nama, req.Username, req.Email, req.Roles, req.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (ar *adminRepositoryImpl) UpdatePassword(req *domain.Admin) (*domain.Admin, error) {
	c := context.Background()
	tx, err := ar.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE admins SET password = ? WHERE ID = ?")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.Password, req.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (ar *adminRepositoryImpl) Delete(id int) error {
	c := context.Background()
	tx, err := ar.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM admins WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (ar *adminRepositoryImpl) Read(logged_id int) ([]domain.Admin, error) {
	result := []domain.Admin{}
	c := context.Background()

	stmt, err := ar.db.Prepare("SELECT id, nama, username, email, roles, created_at, updated_at FROM admins WHERE id <> ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(c, logged_id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Admin{}
		rows.Scan(&row.ID, &row.Nama, &row.Username, &row.Email, &row.Roles, &row.CreatedAt, &row.UpdatedAt)
		result = append(result, row)
	}

	return result, nil
}
