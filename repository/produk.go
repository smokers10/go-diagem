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
	return &produkRepositoryImpl{
		db: database,
	}
}

func (p *produkRepositoryImpl) Create(req *domain.Produk) (*domain.ProdukDetailed, error) {
	spesifikasi := []domain.ProdukSpesifikasi{}
	produkDetailed := domain.ProdukDetailed{}
	statement, err := p.db.Prepare("INSERT INTO produk (id, nama, slug, deskripsi, spesifikasi, kategori_id, berat, satuan_berat, lebar, panjang, tinggi, kode, harga, stok, is_has_variant, discount) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.ID, req.Nama, req.Slug, req.Deskripsi, req.Spesifikasi, req.KategoriID, req.Berat, req.SatuanBerat, req.Lebar, req.Panjang, req.Tinggi, req.Kode, req.Harga, req.Stok, req.IsHasVariant, req.Discount); err != nil {
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
	produkFoto := ProdukFotoRepository(p.db)
	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, produk.stok, 
	produk.harga, produk.dilihat, produk.created_at, produk.updated_at,
	kategori.nama, kategori.id, kategori.slug
	FROM produk JOIN kategori ON kategori.id = produk.kategori_id 
	WHERE produk.nama LIKE CONCAT('%', ?, '%') 
	AND produk.kategori_id LIKE CONCAT('%', ?, '%') 
	AND produk.deleted = false ORDER BY ` + filter.ClarifyOrder.TableName + " " + filter.ClarifyOrder.OrderMethod

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

		// temporary
		spesifikasi := []domain.ProdukSpesifikasi{}

		// scan
		rows.Scan(&row.ID, &row.Nama, &row.Slug, &row.Deskripsi,
			&row.SpesifikasiTemp, &row.Stok, &row.Harga, &row.Dilihat, &row.CreatedAt, &row.UpdatedAt,
			&row.Kategori.Nama, &row.Kategori.ID, &row.Kategori.Slug,
		)

		// unmarshal spesifikasi
		json.Unmarshal([]byte(row.SpesifikasiTemp), &spesifikasi)
		row.Spesifikasi = spesifikasi
		row.SpesifikasiTemp = ""

		// get foto utama
		fotoUtama, err := produkFoto.GetUtamaOnly(row.ID)
		if err != nil {
			return nil, err
		}

		row.ProdukSingleFoto = *fotoUtama

		// append hasil scan
		result = append(result, row)
	}

	return result, nil
}

func (p *produkRepositoryImpl) ByID(id string) (*domain.ProdukDetailed, error) {
	result := domain.ProdukDetailed{}
	spesifikasi := []domain.ProdukSpesifikasi{}
	produkFoto := ProdukFotoRepository(p.db)
	produkVariasi := ProdukVariasiRepository(p.db)

	// get single produk
	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, 
	produk.stok, produk.harga, produk.dilihat, produk.created_at, produk.updated_at, produk.kode,
	produk.berat, produk.satuan_berat, produk.lebar, produk.panjang, produk.tinggi, 
	produk.is_has_variant, kategori.nama, kategori.id, kategori.slug, produk.discount
	FROM produk JOIN kategori ON kategori.id = produk.kategori_id 
	WHERE produk.id = ? AND produk.deleted = false LIMIT 1`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), id)

	// scan row ke result
	row.Scan(&result.ID, &result.Nama, &result.Slug, &result.Deskripsi, &result.SpesifikasiTemp,
		&result.Stok, &result.Harga, &result.Dilihat, &result.CreatedAt, &result.UpdatedAt, &result.Kode,
		&result.Berat, &result.SatuanBerat, &result.Lebar, &result.Panjang, &result.Tinggi,
		&result.IsHasVariant, &result.Kategori.Nama, &result.Kategori.ID, &result.Kategori.Slug, &result.Discount)

	// unmarshall spesifikasi
	json.Unmarshal([]byte(result.SpesifikasiTemp), &spesifikasi)
	result.Spesifikasi = spesifikasi
	result.SpesifikasiTemp = ""

	// ambil semua foto produk
	foto, err := produkFoto.ReadByProdukID(result.ID)
	if err != nil {
		return nil, err
	}

	// ambil semua variasi
	variasi, err := produkVariasi.ReadByProdukID(result.ID)
	if err != nil {
		return nil, err
	}

	// assign foto produk dan variasi ke variable result
	result.ProdukFoto = foto
	result.Variasi = variasi

	return &result, nil
}

func (p *produkRepositoryImpl) ByIDSimplified(id string) (*domain.ProdukDetailed, error) {
	result := domain.ProdukDetailed{}

	// get single produk
	query := `SELECT id, nama, harga, slug, stok, discount, is_has_variant, berat FROM produk WHERE id = ? AND deleted = false`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	statement.QueryRowContext(context.Background(), id).Scan(&result.ID, &result.Nama, &result.Harga, &result.Slug, &result.Stok, &result.Discount, &result.IsHasVariant, &result.Berat)

	return &result, nil
}

func (p *produkRepositoryImpl) BySlugs(slug string) (*domain.ProdukDetailed, error) {
	result := domain.ProdukDetailed{}
	spesifikasi := []domain.ProdukSpesifikasi{}
	produkFoto := ProdukFotoRepository(p.db)
	produkVariasi := ProdukVariasiRepository(p.db)

	// get single produk
	query := `SELECT produk.id, produk.nama, produk.slug, produk.deskripsi, produk.spesifikasi, 
	produk.stok, produk.harga, produk.dilihat, produk.created_at, produk.updated_at, produk.kode,
	produk.berat, produk.satuan_berat, produk.lebar, produk.panjang, produk.tinggi, 
	produk.is_has_variant, kategori.nama, kategori.id, kategori.slug, produk.discount
	FROM produk JOIN kategori ON kategori.id = produk.kategori_id 
	WHERE produk.slug = ? AND produk.deleted = false LIMIT 1`

	statement, err := p.db.Prepare(query)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	row := statement.QueryRowContext(context.Background(), slug)

	// scan row ke result
	row.Scan(&result.ID, &result.Nama, &result.Slug, &result.Deskripsi, &result.SpesifikasiTemp,
		&result.Stok, &result.Harga, &result.Dilihat, &result.CreatedAt, &result.UpdatedAt, &result.Kode,
		&result.Berat, &result.SatuanBerat, &result.Lebar, &result.Panjang, &result.Tinggi,
		&result.IsHasVariant, &result.Kategori.Nama, &result.Kategori.ID, &result.Kategori.Slug, &result.Discount)

	// unmarshall spesifikasi
	json.Unmarshal([]byte(result.SpesifikasiTemp), &spesifikasi)
	result.Spesifikasi = spesifikasi
	result.SpesifikasiTemp = ""

	// ambil semua foto produk
	foto, err := produkFoto.ReadByProdukID(result.ID)
	if err != nil {
		return nil, err
	}

	// ambil semua variasi
	variasi, err := produkVariasi.ReadByProdukID(result.ID)
	if err != nil {
		return nil, err
	}

	// assign foto produk dan variasi ke variable result
	result.ProdukFoto = foto
	result.Variasi = variasi

	return &result, nil
}

func (p *produkRepositoryImpl) Update(req *domain.Produk) (*domain.Produk, error) {
	statement, err := p.db.Prepare(`
		UPDATE produk SET
		nama = ?, kategori_id = ?, deskripsi = ?, slug = ?, spesifikasi = ?, berat = ?, 
		lebar = ?, panjang = ?, tinggi = ?, harga = ?, kode = ?, stok = ?, discount = ?
		WHERE id = ?
	`)

	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Nama, req.KategoriID, req.Deskripsi, req.Slug, req.Spesifikasi,
		req.Berat, req.Lebar, req.Panjang, req.Tinggi, req.Harga, req.Kode, req.Stok, req.Discount, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (p *produkRepositoryImpl) Delete(id string) error {
	statement, err := p.db.Prepare("UPDATE produk SET deleted = true WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}

func (p *produkRepositoryImpl) UpdateStok(produkID string, stok int) error {
	statement, err := p.db.Prepare(`UPDATE produk SET stok = ? WHERE id = ?`)
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), stok, produkID); err != nil {
		return err
	}

	return nil
}
