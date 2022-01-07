package repository

import (
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type blogRepositoryImpl struct {
	db *sql.DB
}

func BlogRepository(database *sql.DB) domain.BlogRepository {
	return &blogRepositoryImpl{db: database}
}

//khusus admin
func (br *blogRepositoryImpl) Create(req *domain.Blog, tx *sql.Tx) error {
	stmt, _ := tx.Prepare(`
		INSERT INTO blogs(id, judul, slug, isi, image, is_publish, seo_keyword, seo_tags, seo_deskripsi) 
		VALUE(?, ?, ?, ?, ?, ?, ?, ?, ?)
	`)

	if _, err := stmt.Exec(req.ID, req.Judul, req.Slug, req.Isi, req.Image, req.IsPublish, req.SEOKeyword, req.SEOTags, req.SEODeskripsi); err != nil {
		return err
	}

	return nil
}

func (br *blogRepositoryImpl) Update(req *domain.Blog) error {
	stmt, _ := br.db.Prepare(`
		UPDATE blogs 
		SET judul = ?, slug = ?, isi = ?, is_publish = ?, seo_keyword = ?, seo_tags = ?, seo_deskripsi = ?, tgl_update = now()
		WHERE id = ? 
	`)

	if _, err := stmt.Exec(req.Judul, req.Slug, req.Isi, req.IsPublish, req.SEOKeyword, req.SEOTags, req.SEODeskripsi, req.ID); err != nil {
		return err
	}

	return nil
}

func (br *blogRepositoryImpl) Delete(id string) error {
	stmt, _ := br.db.Prepare(`DELETE FROM blogs WHERE id = ?`)

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}

func (br *blogRepositoryImpl) CheckJudul(judul string) (*domain.Blog, error) {
	result := domain.Blog{}

	stmt, _ := br.db.Prepare(`
		SELECT judul
		FROM blogs 
		WHERE judul LIKE CONCAT('%', ?, '%') LIMIT 1
	`)

	stmt.QueryRow().Scan(&result.Judul)

	return &result, nil
}

func (br *blogRepositoryImpl) ReadAll() ([]domain.Blog, error) {
	result := []domain.Blog{}

	stmt, _ := br.db.Prepare(`
		SELECT id, judul, slug, isi, image, is_publish, seo_keyword, seo_tags, seo_deskripsi, tgl_publish, tgl_update 
		FROM blogs
	`)

	rows, _ := stmt.Query()
	defer rows.Close()

	for rows.Next() {
		row := domain.Blog{}

		rows.Scan(&row.ID, &row.Judul, &row.Slug, &row.Isi,
			&row.Image, &row.IsPublish, &row.SEOKeyword, &row.SEOTags, &row.SEODeskripsi,
			&row.TGLPublish, &row.TGLUpdate)

		result = append(result, row)
	}

	return result, nil
}

//khusus umum
func (br *blogRepositoryImpl) PublishedOnly(judul string) ([]domain.Blog, error) {
	result := []domain.Blog{}

	stmt, _ := br.db.Prepare(`
		SELECT id, judul, slug, isi, image, is_publish, seo_keyword, seo_tags, seo_deskripsi, tgl_publish, tgl_update 
		FROM blogs 
		WHERE 
		judul LIKE CONCAT('%', ?, '%')
		AND is_publish = true
	`)

	rows, _ := stmt.Query(judul)
	defer rows.Close()

	for rows.Next() {
		row := domain.Blog{}

		rows.Scan(&row.ID, &row.Judul, &row.Slug, &row.Isi,
			&row.Image, &row.IsPublish, &row.SEOKeyword, &row.SEOTags, &row.SEODeskripsi,
			&row.TGLPublish, &row.TGLUpdate)

		result = append(result, row)
	}

	return result, nil
}

func (br *blogRepositoryImpl) BySlug(slug string) (*domain.Blog, error) {
	result := domain.Blog{}
	stmt, _ := br.db.Prepare(`
		SELECT id, judul, slug, isi, image, is_publish, seo_keyword, seo_tags, seo_deskripsi, tgl_publish, tgl_update 
		FROM blogs 
		WHERE slug = ? LIMIT 1
	`)

	stmt.QueryRow(slug).Scan(&result.ID, &result.Judul, &result.Slug, &result.Isi,
		&result.Image, &result.IsPublish, &result.SEOKeyword, &result.SEOTags, &result.SEODeskripsi,
		&result.TGLPublish, &result.TGLUpdate)

	return &result, nil
}

func (br *blogRepositoryImpl) Latest() ([]domain.Blog, error) {
	result := []domain.Blog{}

	stmt, _ := br.db.Prepare(`
		SELECT id, judul, slug, isi, image, is_publish, seo_keyword, seo_tags, seo_deskripsi, tgl_publish, tgl_update 
		FROM blogs 
		ORDER BY tgl_publish DESC 
		LIMIT 10
	`)

	rows, _ := stmt.Query()

	for rows.Next() {
		row := domain.Blog{}
		rows.Scan(&row.ID, &row.Judul, &row.Slug, &row.Isi,
			&row.Image, &row.IsPublish, &row.SEOKeyword, &row.SEOTags, &row.SEODeskripsi,
			&row.TGLPublish, &row.TGLUpdate)
		result = append(result, row)
	}

	return result, nil
}

//for transaction
func (br *blogRepositoryImpl) GetSqlInstance() *sql.DB {
	return br.db
}
