package repository

import (
	"context"
	"database/sql"

	"github.com/smokers10/go-diagem.git/domain"
)

type sliderRepositoryImpl struct {
	db *sql.DB
}

func SliderRepository(database *sql.DB) domain.SliderRepository {
	return &sliderRepositoryImpl{db: database}
}

func (s *sliderRepositoryImpl) Create(req *domain.Slider) (*domain.Slider, error) {
	statement, err := s.db.Prepare("INSERT INTO slider (name, description, img, is_active, url, type) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	inserted, err := statement.ExecContext(context.Background(), req.Name, req.Description, req.Img, req.IsActive, req.URL, req.Type)
	if err != nil {
		return nil, err
	}

	id, _ := inserted.LastInsertId()

	req.ID = int(id)

	return req, nil
}

func (s *sliderRepositoryImpl) Read() ([]domain.Slider, error) {
	result := []domain.Slider{}
	statement, err := s.db.Prepare("SELECT id, name, description, img, is_active, url, type FROM slider")
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	rows, err := statement.QueryContext(context.Background())
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		row := domain.Slider{}
		rows.Scan(&row.ID, &row.Name, &row.Description, &row.Img, &row.IsActive, &row.URL, &row.Type)
		result = append(result, row)
	}

	return result, nil
}

func (s *sliderRepositoryImpl) Update(req *domain.Slider) (*domain.Slider, error) {
	statement, err := s.db.Prepare(`UPDATE slider SET name = ?, description = ?, img = ?, 
	is_active = ?, url = ?, type = ? 
	WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Name, req.Description, req.Img, req.IsActive, req.URL, req.Type, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *sliderRepositoryImpl) Delete(id int) error {
	statement, err := s.db.Prepare("DELETE FROM slider WHERE id = ?")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), id); err != nil {
		return err
	}

	return nil
}
