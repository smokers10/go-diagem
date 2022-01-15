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

func (ur *userRepositoryImpl) ReadAll() ([]domain.User, error) {
	result := []domain.User{}
	stmt, err := ur.db.Prepare("SELECT id, nama, hp, email, is_verified FROM users")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.User{}
		rows.Scan(&row.ID, &row.Nama, &row.HP, &row.Email, &row.IsVerified)
		result = append(result, row)
	}

	return result, nil
}

func (ur *userRepositoryImpl) Detail(id int) (user *domain.User, err error) {
	result := domain.User{}
	alamatRepository := AlamatRepository(ur.db)
	// get user data
	stmt, err := ur.db.Prepare("SELECT id, nama, hp, email, is_verified, tgl_lahir FROM users WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRow(id).Scan(&result.ID, &result.Nama, &result.HP, &result.Email, &result.IsVerified, &result.TglLahir)

	// get alamat
	alamat, err := alamatRepository.Read(id)
	if err != nil {
		return nil, err
	}

	result.Alamat = alamat

	return &result, nil
}

func (ur *userRepositoryImpl) ByEmail(email string) (user *domain.User, err error) {
	result := domain.User{}
	stmt, err := ur.db.Prepare("SELECT id, nama, hp, email, password, is_verified FROM users WHERE email = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRowContext(context.Background(), email).Scan(&result.ID, &result.Nama, &result.HP, &result.Email, &result.Password, &result.IsVerified)

	return &result, nil
}

func (ur *userRepositoryImpl) ByID(id int) (user *domain.User, err error) {
	result := domain.User{}
	stmt, err := ur.db.Prepare("SELECT nama, hp, email, tgl_lahir FROM users WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRowContext(context.Background(), id).Scan(&result.Nama, &result.HP, &result.Email, &result.TglLahir)

	return &result, nil
}

func (ur *userRepositoryImpl) Create(req *domain.UserBasicData) (user *domain.User, err error) {
	statement, err := ur.db.Prepare("INSERT INTO users(nama, hp, email, tgl_lahir, password) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	result, err := statement.ExecContext(context.Background(), req.Nama, req.HP, req.Email, req.TglLahir, req.Password)
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

func (ur *userRepositoryImpl) Update(req *domain.UserBasicData) (user *domain.UserBasicData, err error) {
	statement, err := ur.db.Prepare("UPDATE users SET nama = ?, hp = ?, tgl_lahir = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.ExecContext(context.Background(), req.Nama, req.HP, req.TglLahir, req.ID)

	return req, nil
}

func (ur *userRepositoryImpl) UpdatePassword(new_password string, id int) error {
	statment, err := ur.db.Prepare("UPDATE users SET password = ? WHERE id = ?")
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
