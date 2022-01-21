package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type feedbackRepositoryImpl struct {
	db *sql.DB
}

func FeedbackRepository(database *sql.DB) domain.FeedbackRepository {
	return &feedbackRepositoryImpl{db: database}
}

func (fr *feedbackRepositoryImpl) Create(req *domain.Feedback) error {
	c := context.Background()
	tx, _ := fr.db.BeginTx(c, nil)

	stmt, _ := tx.Prepare(`INSERT INTO feedback (comment, rating, produk_id, user_id) VALUES(?, ?, ?, ?)`)
	defer stmt.Close()

	_, err := stmt.Exec(req.Comment, req.Rating, req.ProdukID, req.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (fr *feedbackRepositoryImpl) Update(req *domain.Feedback) error {
	c := context.Background()
	tx, _ := fr.db.BeginTx(c, nil)

	stmt, _ := tx.Prepare(`UPDATE feedback SET comment = ?, rating = ? WHERE id = ? AND user_id = ?`)
	defer stmt.Close()

	_, err := stmt.Exec(req.Comment, req.Rating, req.ID, req.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (fr *feedbackRepositoryImpl) Read(ProdukID string) (*domain.FeedbackData, error) {
	userRepository := UserRepository(fr.db)
	result := domain.FeedbackData{}

	stmt, _ := fr.db.Prepare(`SELECT id, comment, rating, user_id, created_at FROM feedback WHERE produk_id = ?`)
	defer stmt.Close()

	rows, err := stmt.Query(ProdukID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.FeedbackDetailed{}
		rows.Scan(&row.ID, &row.Comment, &row.Rating, &row.UserID, &row.CreatedAt)

		// get user
		user, err := userRepository.ByID(row.UserID)
		if err != nil {
			return nil, err
		}
		row.User = *user

		result.FeedbackList = append(result.FeedbackList, row)
	}

	countRating, err := fr.CountByRating(ProdukID)
	if err != nil {
		return nil, err
	}

	result.ByRate = *countRating

	return &result, nil
}

func (fr *feedbackRepositoryImpl) OnProduct(userid int, productID string) (*domain.FeedbackDetailed, error) {
	result := domain.FeedbackDetailed{}
	stmt, err := fr.db.Prepare(`SELECT id, comment, rating, user_id FROM feedback WHERE user_id = ? AND produk_id = ?`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	stmt.QueryRow(userid, productID).Scan(&result.ID, &result.Comment, &result.Rating, &result.UserID)

	return &result, nil
}

func (fr *feedbackRepositoryImpl) Delete(req *domain.Feedback) error {
	stmt, _ := fr.db.Prepare(`DELETE FROM feedback WHERE id = ? AND user_id = ?`)
	defer stmt.Close()

	_, err := stmt.Exec(req.ID, req.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (fr *feedbackRepositoryImpl) ByRating(productID string, rating int) (*domain.FeedbackData, error) {
	result := domain.FeedbackData{}
	userRepository := UserRepository(fr.db)

	stmt, err := fr.db.Prepare(`SELECT id, comment, rating, user_id, created_at FROM feedback WHERE produk_id = ? AND rating = ? `)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, _ := stmt.Query(productID, rating)

	for rows.Next() {
		row := domain.FeedbackDetailed{}
		rows.Scan(&row.ID, &row.Comment, &row.Rating, &row.UserID, &row.CreatedAt)

		// get user detail
		user, err := userRepository.ByID(row.UserID)
		if err != nil {
			return nil, err
		}
		row.User = *user

		result.FeedbackList = append(result.FeedbackList, row)
	}

	return &result, nil
}

// Diluar Scope Repsitory
func (fr *feedbackRepositoryImpl) CountByRating(productID string) (*domain.FeedbackRateCount, error) {
	result := domain.FeedbackRateCount{}

	// ambil jumlah bintang 1
	stmt1, _ := fr.db.Prepare("SELECT COUNT(*) as four_start FROM feedback WHERE rating = 1 AND produk_id = ?;")
	defer stmt1.Close()
	stmt1.QueryRow(productID).Scan(&result.OneStar)

	// ambil jumlah bintang 2
	stmt2, _ := fr.db.Prepare("SELECT COUNT(*) as four_start FROM feedback WHERE rating = 2 AND produk_id = ?;")
	defer stmt2.Close()
	stmt2.QueryRow(productID).Scan(&result.TwoStar)

	// ambil jumlah bintang 3
	stmt3, _ := fr.db.Prepare("SELECT COUNT(*) as four_start FROM feedback WHERE rating = 3 AND produk_id = ?;")
	defer stmt3.Close()
	stmt3.QueryRow(productID).Scan(&result.ThreeStar)

	// ambil jumlah bintang 4
	stmt4, _ := fr.db.Prepare("SELECT COUNT(*) as four_start FROM feedback WHERE rating = 4 AND produk_id = ?;")
	defer stmt4.Close()
	stmt4.QueryRow(productID).Scan(&result.FourStar)

	// ambil jumlah bintang 5
	stmt5, _ := fr.db.Prepare("SELECT COUNT(*) as five_start FROM feedback WHERE rating = 5 AND produk_id = ?;")
	defer stmt5.Close()
	stmt5.QueryRow(productID).Scan(&result.FiveStar)

	return &result, nil
}
