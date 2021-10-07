package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/smokers10/go-diagem.git/domain"
)

type produkRepositoryImpl struct {
	db *sql.DB
}

func ProdukRepository(database *sql.DB) domain.ProdukRepository {
	return &produkRepositoryImpl{db: database}
}

func (p *produkRepositoryImpl) Create(req *domain.Produk) (*domain.ProdukDetailed, error) {
	spesifikasi := []domain.ProdukSpesifikasi{}
	produkDetailed := domain.ProdukDetailed{}

	statement, err := p.db.Prepare("INSERT INTO produk (id, nama, slug, deskripsi, spesifikasi, kategori_id) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.ID, req.Nama, req.Slug, req.Deskripsi, req.Spesifikasi, req.KategoriID); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(req.Spesifikasi), &spesifikasi)
	produkDetailed.ID = req.ID
	produkDetailed.Nama = req.Nama
	produkDetailed.Slug = req.Slug
	produkDetailed.Deskripsi = req.Deskripsi
	produkDetailed.Spesifikasi = spesifikasi

	return &produkDetailed, nil
}

func (p *produkRepositoryImpl) Read(filter *domain.ProdukFilter) ([]domain.ProdukDetailed, error) {
	result := []domain.ProdukDetailed{}
	statement, err := p.db.Prepare(`SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, 
	produk.spesifikasi, produk.dilihat, produk.created_at, produk.updated_at, kategori.nama, 
	kategori.id, kategori.slug 
	FROM produk 
	JOIN kategori ON kategori.id = produk.kategori_id 
	WHERE produk.nama LIKE %?%`)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background(), filter.Nama)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.ProdukDetailed{}
		tempRow := domain.ProdukDetailedTemp{}
		spesifikasi := []domain.ProdukSpesifikasi{}

		rows.Scan(&tempRow.ID, &tempRow.Nama, &tempRow.Slug, &tempRow.Deskripsi,
			&tempRow.Spesifikasi, &tempRow.Dilihat, &tempRow.CreatedAt, &tempRow.UpdatedAt,
			&tempRow.Kategori.Nama, &tempRow.Kategori.ID, &tempRow.Kategori.Slug)

		json.Unmarshal([]byte(tempRow.Spesifikasi), &spesifikasi)

		row.ID = tempRow.ID
		row.Nama = tempRow.Nama
		row.Slug = tempRow.Slug
		row.Deskripsi = tempRow.Deskripsi
		row.Spesifikasi = spesifikasi
		row.Dilihat = tempRow.Dilihat
		row.CreatedAt = tempRow.CreatedAt
		row.UpdatedAt = tempRow.UpdatedAt
		row.Kategori.ID = tempRow.Kategori.ID
		row.Kategori.Nama = tempRow.Kategori.Nama
		row.Kategori.Slug = tempRow.Kategori.Slug

		result = append(result, row)
	}

	return result, nil
}

func (p *produkRepositoryImpl) ByID(id string) (*domain.ProdukDetailed, error) {
	result := domain.ProdukDetailed{}
	resultTemp := domain.ProdukDetailedTemp{}
	spesifikasi := []domain.ProdukSpesifikasi{}

	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, produk.dilihat,
	 	  	  produk.created_at, produk.updated_at, kategori.nama, kategori.id, kategori.slug FROM produk 
			  JOIN kategori ON kategori.id = produk.kategori_id WHERE produk.id = ? LIMIT 1`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), id)
	row.Scan(&resultTemp.ID, &resultTemp.Nama, &resultTemp.Slug, &resultTemp.Deskripsi,
		&resultTemp.Spesifikasi, &resultTemp.Dilihat, &resultTemp.CreatedAt, &resultTemp.UpdatedAt,
		&resultTemp.Kategori.Nama, &resultTemp.Kategori.ID, &resultTemp.Kategori.Slug)

	//assign ke result asli
	json.Unmarshal([]byte(resultTemp.Spesifikasi), &spesifikasi)

	result.ID = resultTemp.ID
	result.Nama = resultTemp.Nama
	result.Slug = resultTemp.Slug
	result.Deskripsi = resultTemp.Deskripsi
	result.Dilihat = resultTemp.Dilihat
	result.Spesifikasi = spesifikasi
	result.Kategori.ID = resultTemp.Kategori.ID
	result.Kategori.Nama = resultTemp.Kategori.Nama
	result.Kategori.Slug = resultTemp.Kategori.Slug
	result.CreatedAt = resultTemp.CreatedAt
	result.UpdatedAt = resultTemp.UpdatedAt

	return &result, nil
}

func (p *produkRepositoryImpl) Update(req *domain.Produk) (*domain.Produk, error) {
	statement, err := p.db.Prepare("UPDATE produk SET nama = ?, slug = ?, deskripsi = ?, spesifikasi = ?, kategori_id = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Nama, req.Slug, req.Deskripsi, req.Spesifikasi, req.KategoriID, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (p *produkRepositoryImpl) Delete(id string) error {
	statement, err := p.db.Prepare("DELETE FROM produk WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}
