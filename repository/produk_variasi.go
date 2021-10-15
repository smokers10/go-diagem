package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type produkVariasiRepository struct {
	db *sql.DB
}

func ProdukVariasiRepository(database *sql.DB) domain.ProdukVariasiRepository {
	return &produkVariasiRepository{db: database}
}

func (p *produkVariasiRepository) Create(req *domain.ProdukVariasi) (*domain.ProdukVariasi, error) {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO produk_variasi (id, variant, harga, produk_id, stok) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.ID, req.Variant, req.Harga, req.ProdukID, req.Stok); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (p *produkVariasiRepository) Update(req *domain.ProdukVariasi) (*domain.ProdukVariasi, error) {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE produk_variasi SET variant = ?, harga = ?, stok = ? WHERE id = ? AND produk_id = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.Variant, req.Harga, req.Stok, req.ID, req.ProdukID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (p *produkVariasiRepository) Delete(produkID string, produkVariasiID string) error {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("DELETE FROM produk_variasi WHERE id = ? AND produk_id = ?")
	if err != nil {
		tx.Rollback()
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, produkVariasiID, produkID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
