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

// Untuk Admin
func (s *sliderRepositoryImpl) Create(req *domain.Slider) (*domain.Slider, error) {
	statement, err := s.db.Prepare(`INSERT INTO slider(id, name, description, image, is_publish, url, type) 
	VALUES(?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.ID, req.Name, req.Description, req.Image, req.IsPublish, req.URL, req.Type); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *sliderRepositoryImpl) Read() ([]domain.Slider, error) {
	result := []domain.Slider{}
	statement, err := s.db.Prepare("SELECT id, name, description, image, is_publish, url, type FROM slider")
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
		rows.Scan(&row.ID, &row.Name, &row.Description, &row.Image, &row.IsPublish, &row.URL, &row.Type)
		result = append(result, row)
	}

	return result, nil
}

func (s *sliderRepositoryImpl) Update(req *domain.Slider) (*domain.Slider, error) {
	statement, err := s.db.Prepare(`UPDATE slider SET name = ?, description = ?, is_publish = ?, url = ?, type = ? WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	defer statement.Close()

	if _, err := statement.ExecContext(context.Background(), req.Name, req.Description, req.IsPublish, req.URL, req.Type, req.ID); err != nil {
		return nil, err
	}

	return req, nil
}

func (s *sliderRepositoryImpl) Delete(id string) error {
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

// Untuk Umum
func (s *sliderRepositoryImpl) ByID(id string) (*domain.Slider, error) {
	row := domain.Slider{}
	statement, err := s.db.Prepare("SELECT id, name, description, image, is_publish, url, type FROM slider WHERE ID = ?")
	if err != nil {
		return nil, err
	}
	defer statement.Close()

	statement.QueryRow(id).Scan(&row.ID, &row.Name, &row.Description, &row.Image, &row.IsPublish, &row.URL, &row.Type)

	return &row, nil
}

func (s *sliderRepositoryImpl) ReadOnlyPublished() ([]domain.Slider, error) {
	result := []domain.Slider{}
	statement, err := s.db.Prepare("SELECT id, name, description, image, url, type FROM slider WHERE is_publish = true")
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
		rows.Scan(&row.ID, &row.Name, &row.Description, &row.Image, &row.URL, &row.Type)
		result = append(result, row)
	}

	return result, nil
}
