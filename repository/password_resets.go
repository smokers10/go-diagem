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
	statement, err := pr.db.Prepare("INSERT INTO password_resets email, token VALUES(? ,?)")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Email, req.Token); err != nil {
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

func (pr *passwordResetRepositoryImpl) ByEmail(email string) (passwordResets *domain.PasswordResets, err error) {
	statement, err := pr.db.Prepare("SELECT * FROM password_resets WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if err := statement.QueryRowContext(context.Background(), email).Scan(&passwordResets); err != nil {
		return nil, err
	}

	return passwordResets, nil
}
