// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package database

import (
	"database/sql"
	"sync"
	"time"

	"github.com/Dentrax/GMDB/store/database/ddl"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func Connect(driver, datasource string) (*DB, error) {
	db, err := sql.Open(driver, datasource)

	if err != nil {
		return nil, err
	}

	if err := pingDatabase(db); err != nil {
		return nil, err
	}

	if err := setupDatabase(db); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(1)

	var engine Driver = Sqlite
	var locker Locker = &sync.RWMutex{}

	return &DB{
		conn:   sqlx.NewDb(db, driver),
		lock:   locker,
		driver: engine,
	}, nil
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		time.Sleep(time.Second)
	}
	return
}

func setupDatabase(db *sql.DB) error {
	return ddl.Migrate(db)
}
