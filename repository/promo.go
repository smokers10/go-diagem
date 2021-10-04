package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/smokers10/go-diagem.git/domain"
)

type promoRepositoryImpl struct {
	db *sql.DB
}

func PromoRepository(database *sql.DB) domain.PromoRepository {
	return &promoRepositoryImpl{db: database}
}

func (p *promoRepositoryImpl) Create(req *domain.Promo) (*domain.Promo, error) {
	statement, err := p.db.Prepare("INSERT INTO promo (id, judul, slug, deskripsi, image, is_active, is_featured, tgl_mulai, tgl_selesai) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.ID, req.Judul, req.Slug, req.Deskripsi, req.Image, req.IsActive, req.IsFeatured, req.TglMulai, req.TglSelesai); err != nil {
		return nil, err
	}

	return req, nil
}

func (p *promoRepositoryImpl) Read() ([]domain.Promo, error) {
	result := []domain.Promo{}
	statement, err := p.db.Prepare("SELECT * FROM promo WHERE is_active = true")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Promo{}
		rows.Scan(&row.ID, &row.Judul, &row.Slug, &row.Deskripsi, &row.IsActive, &row.Image, &row.IsFeatured, &row.TglMulai, &row.TglSelesai, &row.CreatedAt, &row.UpdatedAt)
		result = append(result, row)
		fmt.Println(row)
	}

	return result, nil
}

func (p *promoRepositoryImpl) ReadOnlyByAdmin() ([]domain.Promo, error) {
	result := []domain.Promo{}
	statement, err := p.db.Prepare("SELECT id, judul, slug, deskripsi, image, is_active, is_featured, tgl_mulai, tgl_selesai FROM promo")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Promo{}
		rows.Scan(&row.ID, &row.Judul, &row.Slug, &row.Deskripsi, &row.Image, &row.IsActive, &row.IsFeatured, &row.TglMulai, &row.TglSelesai)
		result = append(result, row)
	}

	return result, nil
}

func (p *promoRepositoryImpl) Update(req *domain.Promo) (*domain.Promo, error) {
	statement, err := p.db.Prepare("UPDATE promo SET judul = ?, slug = ?, deskripsi = ?, image = ? ,is_active = ?, is_featured = ?, tgl_mulai = ?, tgl_selesai = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Judul, req.Slug, req.Deskripsi, req.Image, req.IsActive, req.IsFeatured, req.TglMulai, req.TglSelesai, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (p *promoRepositoryImpl) Delete(id string) error {
	statement, err := p.db.Prepare("DELETE FROM promo WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}

func (p *promoRepositoryImpl) ByID(id string) (*domain.Promo, error) {
	row := domain.Promo{}
	statement, err := p.db.Prepare("SELECT * FROM promo WHERE id = ? dLIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), id).Scan(&row.ID, &row.Judul, &row.Slug, &row.Deskripsi, &row.IsActive, &row.IsFeatured, &row.TglMulai, &row.TglSelesai)

	return &row, nil
}

func (p *promoRepositoryImpl) BySlug(slug string) (*domain.Promo, error) {
	row := domain.Promo{}
	statement, err := p.db.Prepare("SELECT * FROM promo WHERE id = ? LIMIT 1")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), slug).Scan(slug)

	return &row, nil
}
