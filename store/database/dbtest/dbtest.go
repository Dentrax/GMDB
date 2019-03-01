package dbtest

import (
	"gmdb/store/database"

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
		return nil
	})
}

// Disconnect closes the database connection.
func Disconnect(d *database.DB) error {
	return d.Close()
}
