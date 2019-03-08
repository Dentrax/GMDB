// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package store

import (
	"github.com/Dentrax/GMDB/models"
)

// New returns a new Store service.
func New(
	movies models.MovieStore,
) Stores {
	return Stores{
		Movies: movies,
	}
}

// Stores represents a list of store models.
type Stores struct {
	Movies models.MovieStore
}
