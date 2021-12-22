package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type alamatRepository struct {
	db *sql.DB
}

func AlamatRepository(databse *sql.DB) domain.AlamatRepository {
	return &alamatRepository{db: databse}
}

func (a *alamatRepository) Create(req *domain.Alamat) (*domain.Alamat, error) {
	c := context.Background()
	tx, err := a.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`INSERT INTO alamat 
	(user_id, alamat, nama, penerima, phone, kd_pos, is_utama, provinsi_id, kota_id, nama_provinsi, nama_kota) 
	VALUES(?, ? ,?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	inserted, err := stmt.ExecContext(c, req.UserID, req.Alamat, req.Nama, req.Penerima, req.Phone, req.KDPos, req.IsUtama, req.ProvinsiID, req.KotaID, req.NamaProvinsi, req.NamaKota)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	id, _ := inserted.LastInsertId()

	req.ID = int(id)

	return req, nil
}

func (a *alamatRepository) Read(userID int) ([]domain.Alamat, error) {
	result := []domain.Alamat{}
	stmt, err := a.db.Prepare(`
		SELECT id, alamat, nama, penerima, phone, provinsi_id, kota_id, nama_provinsi, nama_kota, kd_pos, is_utama
		FROM alamat 
		WHERE user_id = ?
	`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.QueryContext(context.Background(), userID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Alamat{}
		rows.Scan(&row.ID, &row.Alamat, &row.Nama, &row.Penerima, &row.Phone, &row.ProvinsiID,
			&row.KotaID, &row.NamaProvinsi, &row.NamaKota, &row.KDPos, &row.IsUtama)
		result = append(result, row)
	}

	return result, nil
}

func (a *alamatRepository) Update(req *domain.Alamat) (*domain.Alamat, error) {
	c := context.Background()
	tx, err := a.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`UPDATE alamat SET 
	alamat = ?, nama = ?, penerima = ?, phone = ?, kd_pos = ?, provinsi_id = ?, kota_id = ?, nama_provinsi = ?, nama_kota = ?
	WHERE id = ? AND user_id = ?`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.Alamat, req.Nama, req.Penerima, req.Phone, req.KDPos, req.ProvinsiID, req.KotaID, req.NamaProvinsi, req.NamaKota, req.ID, req.UserID); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return req, nil
}

func (a *alamatRepository) Delete(id int, user_id int) error {
	c := context.Background()
	tx, err := a.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM alamat WHERE id = ? AND user_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, id, user_id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (a *alamatRepository) MakeUtama(id int, user_id int) error {
	c := context.Background()
	tx, err := a.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE alamat SET is_utama = true WHERE id = ? AND user_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(c, id, user_id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (a *alamatRepository) MakeUtamaFalse(id int, user_id int) error {
	c := context.Background()
	tx, err := a.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE alamat SET is_utama = false WHERE id <> ? AND user_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(c, id, user_id); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
