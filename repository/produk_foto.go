package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type produkFotoRepositoryImpl struct {
	db *sql.DB
}

func ProdukFotoRepository(database *sql.DB) domain.ProdukFotoRepository {
	return &produkFotoRepositoryImpl{db: database}
}

func (p *produkFotoRepositoryImpl) ReadByProdukID(produkID string) ([]domain.ProdukFoto, error) {
	result := []domain.ProdukFoto{}
	stmt, err := p.db.Prepare("SELECT id, produk_id, path, is_utama FROM produk_foto WHERE produk_id = ?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(c, produkID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.ProdukFoto{}
		rows.Scan(&row.ID, &row.ProdukID, &row.Path, &row.IsUtama)
		result = append(result, row)
	}

	return result, nil
}

func (p *produkFotoRepositoryImpl) GetUtamaOnly(produkID string) (*domain.ProdukFoto, error) {
	result := domain.ProdukFoto{}
	stmt, err := p.db.Prepare("SELECT id, produk_id, path, is_utama FROM produk_foto WHERE produk_id = ? and is_utama = true LIMIT 1")
	if err != nil {
		return nil, err
	}

	stmt.QueryRowContext(c, produkID).Scan(&result.ID, &result.ProdukID, &result.Path, &result.IsUtama)

	return &result, nil
}

func (p *produkFotoRepositoryImpl) Create(req *domain.ProdukFoto) error {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO produk_foto (id, produk_id, path, is_utama) VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(c, req.ID, req.ProdukID, req.Path, req.IsUtama); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return err
}

func (p *produkFotoRepositoryImpl) Delete(id string) error {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM produk_foto WHERE id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(c, id); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (p *produkFotoRepositoryImpl) UpdateIsUtamaToTrue(id string) error {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE foto_produk WHERE id = ? AND produk_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(c, id); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (p *produkFotoRepositoryImpl) UpdateRestIsUtamaToFalse(id string, produkID string) error {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE foto_produk SET is_utama = false WHERE produk_id = ? AND id <> ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	if _, err := stmt.ExecContext(c, produkID, id); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
