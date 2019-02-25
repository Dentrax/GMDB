package ddl

import (
	"database/sql"
)

//Ref: https://github.com/drone/drone/blob/master/store/datastore/ddl/sqlite/ddl_gen.go
var migrations = []struct {
	name string
	stmt string
}{
	{
		name: "create-table-movies",
		stmt: createTableMovies,
	},
	{
		name: "create-table-movie-library",
		stmt: createTableMovieLibrary,
	},
	{
		name: "create-table-parents-guide",
		stmt: createTableParentsGuide,
	},
	{
		name: "create-table-history",
		stmt: createTableHistory,
	},
}

// Migrate performs the database migration. If the migration fails
// and error is returned.
func Migrate(db *sql.DB) error {
	if err := createTable(db); err != nil {
		return err
	}
	completed, err := selectCompleted(db)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, migration := range migrations {
		if _, ok := completed[migration.name]; ok {

			continue
		}

		if _, err := db.Exec(migration.stmt); err != nil {
			return err
		}
		if err := insertMigration(db, migration.name); err != nil {
			return err
		}

	}
	return nil
}

func createTable(db *sql.DB) error {
	_, err := db.Exec(migrationTableCreate)
	return err
}

func insertMigration(db *sql.DB, name string) error {
	_, err := db.Exec(migrationInsert, name)
	return err
}

func selectCompleted(db *sql.DB) (map[string]struct{}, error) {
	migrations := map[string]struct{}{}
	rows, err := db.Query(migrationSelect)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		migrations[name] = struct{}{}
	}
	return migrations, nil
}

//
// migration table ddl and sql
//

var migrationTableCreate = `
CREATE TABLE IF NOT EXISTS migrations (
 name VARCHAR(255)
,UNIQUE(name)
)
`

var migrationInsert = `
INSERT INTO migrations (name) VALUES (?)
`

var migrationSelect = `
SELECT name FROM migrations
`

var createTableMovies = `
CREATE TABLE IF NOT EXISTS movies (
 movie_id      INTEGER PRIMARY KEY AUTOINCREMENT
,movie_title   TEXT
,movie_year    INTEGER
,movie_rating  INTEGER
,movie_tv_show BOOLEAN
,movie_genre   TEXT
,movie_active  BOOLEAN
,movie_like    BOOLEAN
,UNIQUE(movie_title)
);
`

var createTableMovieLibrary = `
CREATE TABLE IF NOT EXISTS movie_library (
 mlib_id          INTEGER PRIMARY KEY AUTOINCREMENT
,mlib_movie_id    INTEGER
,mlib_watched     BOOLEAN
,UNIQUE(mlib_movie_id)
);
`

var createTableParentsGuide = `
CREATE TABLE IF NOT EXISTS parents_guide (
 pg_id          INTEGER PRIMARY KEY AUTOINCREMENT
,pg_movie_id    INTEGER
,pg_nudity      INTEGER
,pg_violence    INTEGER
,pg_profanity   INTEGER
,pg_alcohol     INTEGER
,pg_frightening INTEGER
,UNIQUE(pg_movie_id)
);
`

var createTableHistory = `
CREATE TABLE IF NOT EXISTS histories (
 history_id      INTEGER PRIMARY KEY AUTOINCREMENT
,history_name    TEXT
,history_created INTEGER
,UNIQUE(history_name)
);
`
