package store

import (
	"gmdb/models"

	"golang.org/x/net/context"
)

type Store interface {
	CreateMovie(*models.Movie) error
	GetMovie(int64) (*models.Movie, error)

	Ping() error
}

func CreateMovie(c context.Context, movie *models.Movie) error {
	return FromContext(c).CreateMovie(movie)
}

func GetMovie(c context.Context, id int64) (*models.Movie, error) {
	return FromContext(c).GetMovie(id)
}
