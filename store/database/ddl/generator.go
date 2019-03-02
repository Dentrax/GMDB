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
		name: "create-table-movie-watchlater",
		stmt: createTableMovieWatchLater,
	},
	{
		name: "create-table-movie-parents-guide",
		stmt: createTableMovieParentsGuide,
	},
	{
		name: "create-table-movie-search-history",
		stmt: createTableMovieSearchHistory,
	},
	{
		name: "create-table-movie-note",
		stmt: createTableMovieNote,
	},
	{
		name: "create-table-movie-learn",
		stmt: createTableMovieLearn,
	},
}

// Migrate performs the database migration.
// If the migration fails and error is returned.
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
,movie_title TEXT
,movie_year TEXT
,movie_rating TEXT
,movie_votes TEXT
,movie_reviews TEXT
,movie_duration TEXT
,movie_released TEXT
,movie_istvseries INTEGER
,movie_seasons TEXT
,movie_episodes TEXT
,movie_summary TEXT
,movie_metascore TEXT
,movie_review_count_user TEXT
,movie_review_count_critic TEXT
,movie_rtmeter TEXT
,movie_url_trailer_imdb TEXT
,movie_url_poster_imdb TEXT
);
`

var createTableMovieWatchLater = `
CREATE TABLE IF NOT EXISTS movie_watchlaters (
 wl_id          INTEGER PRIMARY KEY AUTOINCREMENT
,wl_movie_id    INTEGER
,wl_created     INTEGER
,wl_updated     INTEGER
,wl_watched     BOOLEAN
,UNIQUE(wl_movie_id)
);
`

var createTableMovieLearn = `
CREATE TABLE IF NOT EXISTS movie_learns (
 ml_id          INTEGER PRIMARY KEY AUTOINCREMENT
,ml_movie_id    INTEGER
,ml_created     INTEGER
,ml_updated     INTEGER
,ml_liked       BOOLEAN
,UNIQUE(ml_movie_id)
);
`

var createTableMovieParentsGuide = `
CREATE TABLE IF NOT EXISTS movie_parentsguides (
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

var createTableMovieSearchHistory = `
CREATE TABLE IF NOT EXISTS movie_searches (
 search_id       INTEGER PRIMARY KEY AUTOINCREMENT
,search_movie_id INTEGER
,search_created  INTEGER
,search_from     TEXT
);
`

var createTableMovieNote = `
CREATE TABLE IF NOT EXISTS movie_notes (
 note_id       INTEGER PRIMARY KEY AUTOINCREMENT
,note_movie_id INTEGER
,note_created  INTEGER
,note_updated  INTEGER
,note_from     INTEGER
,note_season   INTEGER
,note_episode  INTEGER
,note_hour     INTEGER
,note_minute   INTEGER
,note_second   INTEGER
,note_text     TEXT
);
`
