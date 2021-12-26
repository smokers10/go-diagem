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

func (fr *feedbackRepositoryImpl) Read(ProdukID string) ([]domain.FeedbackDetailed, error) {
	c := context.Background()
	userRepository := UserRepository(fr.db)
	result := []domain.FeedbackDetailed{}
	tx, _ := fr.db.BeginTx(c, nil)

	stmt, _ := tx.Prepare(`SELECT id, comment, rating, user_id FROM feedback WHERE produk_id = ?`)
	defer stmt.Close()

	rows, err := stmt.Query(ProdukID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.FeedbackDetailed{}
		rows.Scan(&row.ID, &row.Comment, &row.Rating, &row.UserID)

		// get user
		user, err := userRepository.ByID(row.UserID)
		if err != nil {
			return nil, err
		}
		row.User = *user

		result = append(result, row)
	}

	return result, nil
}

func (fr *feedbackRepositoryImpl) ByUserId(UserID int) ([]domain.FeedbackDetailed, error) {
	c := context.Background()
	result := []domain.FeedbackDetailed{}
	tx, _ := fr.db.BeginTx(c, nil)

	stmt, _ := tx.Prepare(`SELECT id, comment, rating FROM feedback WHERE user_id = ?`)
	defer stmt.Close()

	rows, err := stmt.Query(UserID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.FeedbackDetailed{}
		rows.Scan(&row.ID, &row.Comment, &row.Rating)
		result = append(result, row)
	}

	return result, nil
}

func (fr *feedbackRepositoryImpl) Delete(req *domain.Feedback) error {
	c := context.Background()
	tx, _ := fr.db.BeginTx(c, nil)

	stmt, _ := tx.Prepare(`DELETE FROM feedback WHERE id = ? AND user_id = ?`)
	defer stmt.Close()

	_, err := stmt.Exec(req.ID, req.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
