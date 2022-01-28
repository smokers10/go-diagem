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

func (p *produkVariasiRepository) ReadByProdukID(produkID string) ([]domain.ProdukVariasi, error) {
	result := []domain.ProdukVariasi{}
	c := context.Background()
	stmt, err := p.db.Prepare("SELECT id, variant, sku, harga, produk_id, stok FROM produk_variasi WHERE produk_id = ? AND deleted = false")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.QueryContext(c, produkID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.ProdukVariasi{}
		rows.Scan(&row.ID, &row.Variant, &row.SKU, &row.Harga, &row.ProdukID, &row.Stok)
		result = append(result, row)
	}

	return result, nil
}

func (p *produkVariasiRepository) ByID(id string) (*domain.ProdukVariasi, error) {
	result := domain.ProdukVariasi{}
	stmt, err := p.db.Prepare("SELECT id, variant, sku, harga, produk_id, stok FROM produk_variasi WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if err := stmt.QueryRow(id).Scan(&result.ID, &result.Variant, &result.SKU, &result.Harga, &result.ProdukID, &result.Stok); err != nil {
		return nil, err
	}

	return &result, nil
}

func (p *produkVariasiRepository) Create(req *domain.ProdukVariasi) (*domain.ProdukVariasi, error) {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO produk_variasi (id, variant, sku, harga, produk_id, stok) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.ID, req.Variant, req.SKU, req.Harga, req.ProdukID, req.Stok); err != nil {
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

	stmt, err := tx.Prepare("UPDATE produk_variasi SET variant = ?, sku = ?, harga = ?, stok = ? WHERE id = ? AND produk_id = ?")
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(c, req.Variant, req.SKU, req.Harga, req.Stok, req.ID, req.ProdukID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return req, nil
}

func (p *produkVariasiRepository) UpdateStokVariant(id string, changeValue int, operationType string) error {
	variant, err := p.ByID(id)
	if err != nil {
		return err
	}

	if operationType == "addition" {
		variant.Stok = variant.Stok + changeValue
	}

	if operationType == "subtraction" {
		variant.Stok = variant.Stok - changeValue
	}

	stmt, err := p.db.Prepare("UPDATE produk_variasi SET stok = ? WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.Exec(variant.Stok, id); err != nil {
		return err
	}

	return nil
}

func (p *produkVariasiRepository) Delete(produkID string, produkVariasiID string) error {
	c := context.Background()
	tx, err := p.db.BeginTx(c, nil)
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("UPDATE produk_variasi SET deleted = true WHERE id = ? AND produk_id = ?")
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
