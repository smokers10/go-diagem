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
	statement, err := ar.db.Prepare("SELECT id, nama, username, email, password FROM admins WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), email)
	row.Scan(&result.ID, &result.Nama, &result.Username, &result.Email, &result.Password)

	return &result, nil
}
