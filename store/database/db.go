// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

//Ref: https://github.com/drone/drone/blob/master/store/shared/db/db.go

package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Driver defines the database driver.
type Driver int

// Database driver enums.
const (
	Sqlite = iota + 1
)

type (
	// A Scanner represents an object that can be scanned
	// for values.
	Scanner interface {
		Scan(dest ...interface{}) error
	}

	// A Locker represents an object that can be locked and unlocked.
	Locker interface {
		Lock()
		Unlock()
		RLock()
		RUnlock()
	}

	// Binder interface defines database field bindings.
	Binder interface {
		BindNamed(query string, arg interface{}) (string, []interface{}, error)
	}

	// Queryer interface defines a set of methods for
	// querying the database.
	Queryer interface {
		Query(query string, args ...interface{}) (*sql.Rows, error)
		QueryRow(query string, args ...interface{}) *sql.Row
	}

	// Execer interface defines a set of methods for executing
	// read and write commands against the database.
	Execer interface {
		Queryer
		Exec(query string, args ...interface{}) (sql.Result, error)
	}

	// DB is a pool of zero or more underlying connections to
	// the drone database.
	DB struct {
		conn   *sqlx.DB
		lock   Locker
		driver Driver
	}
)

// View executes a function within the context of a managed read-only
// transaction. Any error that is returned from the function is returned
// from the View() method.
func (db *DB) View(fn func(Queryer, Binder) error) error {
	db.lock.RLock()
	err := fn(db.conn, db.conn)
	db.lock.RUnlock()
	return err
}

// Lock obtains a write lock to the database (sqlite only) and executes
// a function. Any error that is returned from the function is returned
// from the Lock() method.
func (db *DB) Lock(fn func(Execer, Binder) error) error {
	db.lock.Lock()
	err := fn(db.conn, db.conn)
	db.lock.Unlock()
	return err
}

// Update executes a function within the context of a read-write managed
// transaction. If no error is returned from the function then the
// transaction is committed. If an error is returned then the entire
// transaction is rolled back. Any error that is returned from the function
// or returned from the commit is returned from the Update() method.
func (db *DB) Update(fn func(Execer, Binder) error) error {
	db.lock.Lock()
	defer db.lock.Unlock()

	tx, err := db.conn.Begin()
	if err != nil {
		return err
	}

	err = fn(tx, db.conn)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// Driver returns the name of the SQL driver.
func (db *DB) Driver() Driver {
	return db.driver
}

// Close cloes the database connection.
func (db *DB) Close() error {
	return db.conn.Close()
}
