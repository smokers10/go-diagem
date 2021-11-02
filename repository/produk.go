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
	statement, err := p.db.Prepare("INSERT INTO produk (id, nama, slug, deskripsi, spesifikasi, kategori_id, berat, satuan_berat, lebar, panjang, tinggi, kode, harga, is_has_variant) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.ID, req.Nama, req.Slug, req.Deskripsi, req.Spesifikasi, req.KategoriID, req.Berat, req.SatuanBerat, req.Lebar, req.Panjang, req.Tinggi, req.Kode, req.Harga, req.IsHasVariant); err != nil {
		return nil, err
	}

	json.Unmarshal([]byte(req.Spesifikasi), &spesifikasi)

	produkDetailed.ID = req.ID
	produkDetailed.Nama = req.Nama
	produkDetailed.Slug = req.Slug
	produkDetailed.Deskripsi = req.Deskripsi
	produkDetailed.Spesifikasi = spesifikasi
	produkDetailed.Berat = req.Berat
	produkDetailed.SatuanBerat = req.SatuanBerat
	produkDetailed.Panjang = req.Panjang
	produkDetailed.Lebar = req.Lebar
	produkDetailed.Tinggi = req.Tinggi
	produkDetailed.Kode = req.Kode
	produkDetailed.Harga = req.Harga
	produkDetailed.IsHasVariant = req.IsHasVariant

	return &produkDetailed, nil
}

func (p *produkRepositoryImpl) Read(filter *domain.ProdukFilter) ([]domain.ProdukDetailed, error) {
	result := []domain.ProdukDetailed{}
	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, produk.dilihat, produk.created_at, produk.updated_at, kategori.nama, kategori.id, kategori.slug 
	FROM produk JOIN kategori ON kategori.id = produk.kategori_id 
	WHERE produk.nama LIKE CONCAT('%', ?, '%') AND produk.kategori_id LIKE CONCAT('%', ?, '%')`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background(), filter.Nama, filter.KategoriID)
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
	variasi := []domain.ProdukVariasi{}

	// get single produk
	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, produk.dilihat, produk.harga, produk.kode, produk.is_has_variant
	 	  	  produk.created_at, produk.updated_at, kategori.nama, kategori.id, kategori.slug FROM produk 
			  JOIN kategori ON kategori.id = produk.kategori_id WHERE produk.id = ? LIMIT 1`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), id)

	// scan row ke result temporary
	row.Scan(&resultTemp.ID, &resultTemp.Nama, &resultTemp.Slug, &resultTemp.Deskripsi,
		&resultTemp.Spesifikasi, &resultTemp.Dilihat, &resultTemp.Harga, &resultTemp.Kode, &resultTemp.IsHasVariant, &resultTemp.CreatedAt, &resultTemp.UpdatedAt,
		&resultTemp.Kategori.Nama, &resultTemp.Kategori.ID, &resultTemp.Kategori.Slug)

	// unmarshall spesifikasi dari result temporary ke variable spesifikasi
	json.Unmarshal([]byte(resultTemp.Spesifikasi), &spesifikasi)

	// get variasi
	statement2, err := p.db.Prepare("SELECT id, variant, harga, stok FROM produk_variasi WHERE produk_id = ?")
	if err != nil {
		return nil, err
	}

	defer statement2.Close()

	variasiRows, err := statement2.QueryContext(context.Background(), resultTemp.ID)
	if err != nil {
		return nil, err
	}

	for variasiRows.Next() {
		variasiRow := domain.ProdukVariasi{}
		variasiRows.Scan(&variasiRow.ID, &variasiRow.Variant, &variasiRow.Harga, &variasiRow.Stok)
		variasi = append(variasi, variasiRow)
	}

	// assign result temporary ke result
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
	result.Variasi = variasi
	result.Kode = resultTemp.Kode
	result.Harga = resultTemp.Harga
	result.IsHasVariant = resultTemp.IsHasVariant

	return &result, nil
}

func (p *produkRepositoryImpl) BySlugs(slug string) (*domain.ProdukDetailed, error) {
	result := domain.ProdukDetailed{}
	resultTemp := domain.ProdukDetailedTemp{}
	spesifikasi := []domain.ProdukSpesifikasi{}
	variasi := []domain.ProdukVariasi{}

	// get single produk
	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, produk.harga, produk.kode, produk.is_has_variant, produk.dilihat,
	 	  	  produk.created_at, produk.updated_at, kategori.nama, kategori.id, kategori.slug FROM produk 
			  JOIN kategori ON kategori.id = produk.kategori_id WHERE produk.slug = ? LIMIT 1`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), slug)

	// scan row ke result temporary
	row.Scan(&resultTemp.ID, &resultTemp.Nama, &resultTemp.Slug, &resultTemp.Deskripsi,
		&resultTemp.Spesifikasi, &resultTemp.Dilihat, &resultTemp.CreatedAt, &resultTemp.UpdatedAt,
		&resultTemp.Kategori.Nama, &resultTemp.Kategori.ID, &resultTemp.Kategori.Slug)

	// unmarshall spesifikasi dari result temporary ke variable spesifikasi
	json.Unmarshal([]byte(resultTemp.Spesifikasi), &spesifikasi)

	// get variasi
	statement2, err := p.db.Prepare("SELECT id, variant, harga, stok FROM produk_variasi WHERE produk_id = ?")
	if err != nil {
		return nil, err
	}

	defer statement2.Close()

	variasiRows, err := statement2.QueryContext(context.Background(), resultTemp.ID)
	if err != nil {
		return nil, err
	}

	for variasiRows.Next() {
		variasiRow := domain.ProdukVariasi{}
		variasiRows.Scan(&variasiRow.ID, &variasiRow.Variant, &variasiRow.Harga, &variasiRow.Stok)
		variasi = append(variasi, variasiRow)
	}

	// assign result temporary ke result
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
	result.Variasi = variasi
	result.Kode = resultTemp.Kode
	result.Harga = resultTemp.Harga
	result.IsHasVariant = resultTemp.IsHasVariant

	return &result, nil
}

func (p *produkRepositoryImpl) Update(req *domain.Produk) (*domain.Produk, error) {
	statement, err := p.db.Prepare("UPDATE produk SET nama = ?, slug = ?, deskripsi = ?, spesifikasi = ?, kategori_id = ?, harga = ?, kode = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Nama, req.Slug, req.Deskripsi, req.Spesifikasi, req.KategoriID, req.Harga, req.Kode, req.ID); err != nil {
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
