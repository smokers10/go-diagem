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
	statement, err := k.db.Prepare("INSERT INTO kategori (id, nama, slug, cover) VALUES(?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.ID, req.Nama, req.Slug, req.Cover); err != nil {
		return nil, err
	}

	insertedItem.Nama = req.Nama
	insertedItem.Slug = req.Slug

	return &insertedItem, nil
}

func (k *kategoriRepositoryImpl) Read() ([]domain.Kategori, error) {
	result := []domain.Kategori{}
	statement, err := k.db.Prepare("SELECT id, nama, slug, cover FROM kategori WHERE deleted = false")
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
		rows.Scan(&row.ID, &row.Nama, &row.Slug, row.Cover)
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

func (k *kategoriRepositoryImpl) Delete(id string) error {
	stmt, err := k.db.Prepare("UPDATE kategori SET deleted = true WHERE id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()

	if _, err := stmt.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}

func (k *kategoriRepositoryImpl) Detail(id string) (*domain.Kategori, error) {
	result := domain.Kategori{}
	stmt, err := k.db.Prepare("SELECT id, nama, cover, slug FROM kategori WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(id).Scan(&result.ID, &result.Nama, &result.Cover, &result.Slug)

	return &result, nil
}
