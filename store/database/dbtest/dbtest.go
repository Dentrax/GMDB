// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package dbtest

import (
	"github.com/Dentrax/GMDB/store/database"

	_ "github.com/mattn/go-sqlite3"
)

// Connect opens a new test database connection.
func Connect() (*database.DB, error) {
	var (
		driver = "sqlite3"
		config = ":memory:?_foreign_keys=1"
	)
	return database.Connect(driver, config)
}

// Reset resets the database state.
func Reset(d *database.DB) error {
	return d.Lock(func(tx database.Execer, _ database.Binder) error {
		tx.Exec("DELETE FROM movies")
		tx.Exec("DELETE FROM movie_watchlaters")
		tx.Exec("DELETE FROM movie_searches")
		tx.Exec("DELETE FROM movie_notes")
		tx.Exec("DELETE FROM movie_learns")
		tx.Exec("DELETE FROM movie_parentsguides")
		return nil
	})
}

// Disconnect closes the database connection.
func Disconnect(d *database.DB) error {
	return d.Close()
}
