package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type userRepositoryImpl struct {
	db *sql.DB
}

func UserRepository(database *sql.DB) domain.UserRepository {
	return &userRepositoryImpl{db: database}
}

func (ur *userRepositoryImpl) ByEmail(email string) (user *domain.User, err error) {
	result := domain.User{}
	statement, err := ur.db.Prepare("SELECT id, nama, hp, email, password, is_verified FROM users WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), email).Scan(&result.ID, &result.Nama, &result.HP, &result.Email, &result.Password, &result.IsVerified)

	return &result, nil
}

func (ur *userRepositoryImpl) ByID(id int) (user *domain.User, err error) {
	result := domain.User{}
	statement, err := ur.db.Prepare("SELECT * FROM users WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), id).Scan(&result.ID, &result.Nama, &result.HP, &result.Email, &result.Password, &result.IsVerified)

	return &result, nil
}

func (ur *userRepositoryImpl) Create(req *domain.UserBasicData) (user *domain.User, err error) {
	statement, err := ur.db.Prepare("INSERT INTO users (nama, hp, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.ExecContext(context.Background(), req.Nama, req.HP, req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &domain.User{
		ID:    int(id),
		Nama:  req.Nama,
		Email: req.Email,
	}, nil
}

func (ur *userRepositoryImpl) Update(req *domain.UserBasicData) (user *domain.User, err error) {
	statement, err := ur.db.Prepare("UPDATE users SET nama = ?, hp = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.ExecContext(context.Background(), req.Nama, req.HP, req.ID)

	return &domain.User{
		ID:    req.ID,
		Nama:  req.Nama,
		Email: req.Email,
	}, nil
}

func (ur *userRepositoryImpl) UpdatePassword(new_password string, id int) error {
	statment, err := ur.db.Prepare("UPDATE users password = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer statment.Close()

	if _, err := statment.ExecContext(context.Background(), new_password, id); err != nil {
		return err
	}

	return nil
}

func (ur *userRepositoryImpl) UpdateVerificationStatus(id int) error {
	statment, err := ur.db.Prepare("UPDATE users SET is_verified = true WHERE id = ?")
	if err != nil {
		return err
	}

	defer statment.Close()

	if _, err := statment.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}
