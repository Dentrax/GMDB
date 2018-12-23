package database

import (
	"gmdb/models"
)

func (db *Database) CreateMovie(model *models.Movie) error {
	return nil
}

func (db *Database) GetMovie(id int64) (*models.Movie, error) {
	example := models.Movie{
		Info: models.MovieInfo{
			Title: "Test",
			Year:  "2020",
		},
	}
	return &example, nil
}
