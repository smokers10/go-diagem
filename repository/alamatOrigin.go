package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type alamatOriginRepositoryImpl struct {
	db *sql.DB
}

func AlamatOriginRepository(database *sql.DB) domain.AlamatOriginRepository {
	return &alamatOriginRepositoryImpl{db: database}
}

func (aor *alamatOriginRepositoryImpl) Read() (*domain.AlamatOrigin, error) {
	c := context.Background()
	result := domain.AlamatOrigin{}
	stmt, err := aor.db.Prepare(`SELECT id, kontak_admin, alamat_lengkap, provinsi_id, kota_id, nama_provinsi, nama_kota, kd_pos FROM alamat_origin LIMIT 1`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	stmt.QueryRowContext(c).Scan(&result.ID, &result.KontakAdmin, &result.AlamatLengkap, &result.ProvinsiID, &result.KotaID, &result.NamaProvinsi, &result.NamaKota, &result.KDPos)

	return &result, nil
}

func (aor *alamatOriginRepositoryImpl) Update(req *domain.AlamatOrigin) (*domain.AlamatOrigin, error) {
	c := context.Background()
	tx, err := aor.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
		UPDATE alamat_origin
		SET kontak_admin = ?, alamat_lengkap = ?, provinsi_id = ?, 
		kota_id = ?, nama_provinsi = ?, nama_kota = ?, kd_pos = ?
		WHERE id = ?
	`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if _, err := stmt.ExecContext(c, req.KontakAdmin, req.AlamatLengkap, req.ProvinsiID,
		req.KotaID, req.NamaProvinsi, req.NamaKota, req.KDPos, req.ID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}
