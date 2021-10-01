package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type kategoriRepositoryImpl struct {
	db *sql.DB
}

func KategoriRepository(database *sql.DB) domain.KategoriRepository {
	return &kategoriRepositoryImpl{db: database}
}

func (k *kategoriRepositoryImpl) Create(req *domain.Kategori) (*domain.Kategori, error) {
	insertedItem := domain.Kategori{}
	statement, err := k.db.Prepare("INSERT INTO kategori (nama, slug) VALUES(?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	result, err := statement.ExecContext(context.Background(), req.Nama, req.Slug)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	insertedItem.ID = int(id)
	insertedItem.Nama = req.Nama
	insertedItem.Slug = req.Slug

	return &insertedItem, nil
}

func (k *kategoriRepositoryImpl) Read() ([]domain.Kategori, error) {
	result := []domain.Kategori{}
	statement, err := k.db.Prepare("SELECT id, nama, slug FROM kategori")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Kategori{}
		rows.Scan(&row.ID, &row.Nama, &row.Slug)
		result = append(result, row)
	}

	return result, nil
}

func (k *kategoriRepositoryImpl) Update(req *domain.Kategori) (*domain.Kategori, error) {
	statement, err := k.db.Prepare("UPDATE kategori SET nama = ?, slug = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Nama, req.Slug, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (k *kategoriRepositoryImpl) Delete(id int) error {
	statement, err := k.db.Prepare("DELETE FROM kategori WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}
