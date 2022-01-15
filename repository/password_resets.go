package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type passwordResetRepositoryImpl struct {
	db *sql.DB
}

func PasswordResetRepository(database *sql.DB) domain.PasswordResetsRepository {
	return &passwordResetRepositoryImpl{db: database}
}

func (pr *passwordResetRepositoryImpl) Create(req *domain.PasswordResets) error {
	stmt, err := pr.db.Prepare(`INSERT INTO password_resets (email, token, code, created_at) VALUE(?, ?, ?, now())`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(req.Email, req.Token, req.Kode); err != nil {
		return err
	}

	return nil
}

func (pr *passwordResetRepositoryImpl) Delete(email string) error {
	statement, err := pr.db.Prepare("DELETE FROM password_resets WHERE email = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), email); err != nil {
		return err
	}

	return nil
}

func (pr *passwordResetRepositoryImpl) ByEmail(email string) (*domain.PasswordResets, error) {
	result := domain.PasswordResets{}
	statement, err := pr.db.Prepare("SELECT email, token, code FROM password_resets WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), email).Scan(&result.Email, &result.Token, &result.Kode)

	return &result, nil
}

func (pr *passwordResetRepositoryImpl) ByToken(token string) (*domain.PasswordResets, error) {
	result := domain.PasswordResets{}
	statement, err := pr.db.Prepare("SELECT email, token, code FROM password_resets WHERE token = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), token).Scan(&result.Email, &result.Token, &result.Kode)

	return &result, nil
}

func (pr *passwordResetRepositoryImpl) Update(req *domain.PasswordResets) error {
	stmt, err := pr.db.Prepare(`UPDATE password_resets SET token = ?, code = ? WHERE email = ?`)
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(req.Token, req.Kode, req.Email); err != nil {
		return err
	}

	return nil
}
