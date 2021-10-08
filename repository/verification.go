package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type verificationRepositoryImpl struct {
	db *sql.DB
}

func VerificationRepository(database *sql.DB) domain.VerifikasiRepository {
	return &verificationRepositoryImpl{db: database}
}

func (v *verificationRepositoryImpl) Create(req *domain.Verifikasi) error {
	statement, err := v.db.Prepare("INSERT INTO verifikasi (user_id, kode) VALUES(?, ?)")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.UserID, req.Kode); err != nil {
		return err
	}

	return nil
}

func (v *verificationRepositoryImpl) ByUserID(userID int) (*domain.Verifikasi, error) {
	result := domain.Verifikasi{}
	statement, err := v.db.Prepare("SELECT * FROM verifikasi WHERE user_id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), userID).Scan(&result.UserID, &result.Kode)

	return &result, nil
}

func (v *verificationRepositoryImpl) Update(req *domain.Verifikasi) error {
	statement, err := v.db.Prepare("UPDATE verifikasi SET kode = ? WHERE user_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Kode, req.UserID); err != nil {
		return err
	}

	return nil
}

func (v *verificationRepositoryImpl) Delete(userID int) error {
	statement, err := v.db.Prepare("DELETE FROM verifikasi WHERE user_id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), userID); err != nil {
		return err
	}

	return nil
}
