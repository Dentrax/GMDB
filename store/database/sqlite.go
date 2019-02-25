package database

import (
	"database/sql"
	"log"
	"time"

	"gmdb/store"
	"gmdb/store/database/ddl"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB

	Driver string
	Config string
}

func New(driver string, config string) store.Store {
	return &Database{
		DB:     open(driver, config),
		Driver: driver,
		Config: config,
	}
}

func open(driver, config string) *sql.DB {
	db, err := sql.Open(driver, config)

	if err != nil {
		log.Fatalln(err)
	}

	if err := pingDatabase(db); err != nil {
		log.Fatalln(err)
	}

	if err := setupDatabase(db); err != nil {
		log.Println(err)
		log.Fatalln("migration failed")
	}

	return db
}

func pingDatabase(db *sql.DB) (err error) {
	for i := 0; i < 3; i++ {
		err = db.Ping()
		if err == nil {
			return
		}
		log.Println("database ping failed. retry in 1s")
		time.Sleep(time.Second)
	}
	return
}

func setupDatabase(db *sql.DB) error {
	return ddl.Migrate(db)
}
