package store

import (
	"gmdb/models"
)

func New(
	movies models.MovieStore,
) Stores {
	return Stores{
		Movies: movies,
	}
}

type Stores struct {
	Movies models.MovieStore
}
