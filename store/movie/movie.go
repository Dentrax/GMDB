// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package movie

import (
	"context"
	"time"

	"gmdb/models"
	"gmdb/store/database"
)

// New returns a new UserStore.
func New(db *database.DB) models.MovieStore {
	return &movieStore{db}
}

type movieStore struct {
	db *database.DB
}

func (s *movieStore) Create(ctx context.Context, movie *models.MovieInfo) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		params := toParams(movie)
		stmt, args, err := binder.BindNamed(queryInsert, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(stmt, args...)
		if err != nil {
			return err
		}
		movie.ID, err = res.LastInsertId()
		return err
	})
}

func (s *movieStore) CreateSearch(ctx context.Context, movie *models.MovieInfo, from string) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		params := map[string]interface{}{
			"search_movie_id": movie.ID,
			"search_created":  time.Now().Unix(),
			"search_from":     from,
		}
		query, args, err := binder.BindNamed(queryInsertSearch, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(query, args...)
		if err != nil {
			return err
		}
		movie.ID, err = res.LastInsertId()
		return err
	})
}

func (s *movieStore) CreateNI(ctx context.Context, movie *models.MovieInfo, note *models.MovieNoteInfo) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		now := time.Now().Unix()
		params := map[string]interface{}{
			"note_movie_id": movie.ID,
			"note_created":  now,
			"note_updated":  now,
			"note_from":     note.From,
			"note_season":   note.Season,
			"note_episode":  note.Episode,
			"note_hour":     note.Hour,
			"note_minute":   note.Minute,
			"note_second":   note.Second,
			"note_text":     note.Text,
		}
		query, args, err := binder.BindNamed(queryInsertNI, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(query, args...)
		if err != nil {
			return err
		}
		movie.ID, err = res.LastInsertId()
		return err
	})
}

func (s *movieStore) CreateML(ctx context.Context, movie *models.MovieInfo, liked bool) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		now := time.Now().Unix()
		params := map[string]interface{}{
			"ml_movie_id": movie.ID,
			"ml_created":  now,
			"ml_updated":  now,
			"ml_liked":    liked,
		}
		query, args, err := binder.BindNamed(queryInsertML, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(query, args...)
		if err != nil {
			return err
		}
		movie.ID, err = res.LastInsertId()
		return err
	})
}

func (s *movieStore) CreateWL(ctx context.Context, movie *models.MovieInfo, watched bool) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		now := time.Now().Unix()
		params := map[string]interface{}{
			"wl_movie_id": movie.ID,
			"wl_created":  now,
			"wl_updated":  now,
			"wl_watched":  watched,
		}
		query, args, err := binder.BindNamed(queryInsertWL, params)
		if err != nil {
			return err
		}
		res, err := execer.Exec(query, args...)
		if err != nil {
			return err
		}
		movie.ID, err = res.LastInsertId()
		return err
	})
}

func (s *movieStore) GetMovies(ctx context.Context) ([]*models.MovieInfo, error) {
	out := []*models.MovieInfo{}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		rows, err := queryer.Query(queryMoviesAll)
		if err != nil {
			return err
		}
		out, err = scanRows(rows)
		return err
	})
	return out, err
}

func (s *movieStore) GetSearches(ctx context.Context) ([]*models.MovieSearchInfo, error) {
	out := []*models.MovieSearchInfo{}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		rows, err := queryer.Query(queryMovieSearchesAll)
		if err != nil {
			return err
		}
		out, err = scanRowsSI(rows)
		return err
	})
	return out, err
}

func (s *movieStore) GetMovieNoteList(ctx context.Context) ([]*models.MovieNoteInfo, error) {
	out := []*models.MovieNoteInfo{}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		rows, err := queryer.Query(queryMovieNoteAll)
		if err != nil {
			return err
		}
		out, err = scanRowsNI(rows)
		return err
	})
	return out, err
}

func (s *movieStore) GetWatchLaterList(ctx context.Context) ([]*models.MovieWatchLaterInfo, error) {
	out := []*models.MovieWatchLaterInfo{}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		rows, err := queryer.Query(queryMovieWatchLaterAll)
		if err != nil {
			return err
		}
		out, err = scanRowsWL(rows)
		return err
	})
	return out, err
}

func (s *movieStore) GetMovieLearnList(ctx context.Context) ([]*models.MovieLearnInfo, error) {
	out := []*models.MovieLearnInfo{}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		rows, err := queryer.Query(queryMovieLearnAll)
		if err != nil {
			return err
		}
		out, err = scanRowsML(rows)
		return err
	})
	return out, err
}

func (s *movieStore) FindByTitle(ctx context.Context, title string) (*models.MovieInfo, error) {
	out := &models.MovieInfo{Title: title}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryKey, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err
}

func (s *movieStore) FindByID(ctx context.Context, id int64) (*models.MovieInfo, error) {
	out := &models.MovieInfo{ID: id}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		params := toParams(out)
		query, args, err := binder.BindNamed(queryMovieFindByIDKey, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRow(row, out)
	})
	return out, err
}

func (s *movieStore) FindNI(ctx context.Context, id int64) (*models.MovieNoteInfo, error) {
	out := &models.MovieNoteInfo{MovieID: id}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		params := toParamsNI(out)
		query, args, err := binder.BindNamed(queryMovieNIKey, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRowNI(row, out)
	})
	return out, err
}

func (s *movieStore) FindWL(ctx context.Context, id int64) (*models.MovieWatchLaterInfo, error) {
	out := &models.MovieWatchLaterInfo{MovieID: id}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		params := toParamsWL(out)
		query, args, err := binder.BindNamed(queryMovieWLKey, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRowWL(row, out)
	})
	return out, err
}

func (s *movieStore) FindML(ctx context.Context, id int64) (*models.MovieLearnInfo, error) {
	out := &models.MovieLearnInfo{MovieID: id}
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		params := toParamsML(out)
		query, args, err := binder.BindNamed(queryMovieMLKey, params)
		if err != nil {
			return err
		}
		row := queryer.QueryRow(query, args...)
		return scanRowML(row, out)
	})
	return out, err
}

// Update persists an updated movie to the datastore.
func (s *movieStore) Update(ctx context.Context, movie *models.MovieInfo) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		params := toParams(movie)
		stmt, args, err := binder.BindNamed(queryUpdate, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Update persists an updated Note Info to the datastore.
func (s *movieStore) UpdateNI(ctx context.Context, movie *models.MovieInfo, note *models.MovieNoteInfo) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		now := time.Now().Unix()
		params := map[string]interface{}{
			"note_movie_id": movie.ID,
			"note_updated":  now,
			"note_season":   note.Season,
			"note_episode":  note.Episode,
			"note_hour":     note.Hour,
			"note_minute":   note.Minute,
			"note_second":   note.Second,
			"note_text":     note.Text,
		}
		stmt, args, err := binder.BindNamed(queryUpdateNI, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Update persists an updated Watch Later to the datastore.
func (s *movieStore) UpdateWL(ctx context.Context, movie *models.MovieInfo, watched bool) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		now := time.Now().Unix()
		params := map[string]interface{}{
			"wl_movie_id": movie.ID,
			"wl_updated":  now,
			"wl_watched":  watched,
		}
		stmt, args, err := binder.BindNamed(queryUpdateWL, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Update persists an updated Watch Later to the datastore.
func (s *movieStore) UpdateML(ctx context.Context, movie *models.MovieInfo, liked bool) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		now := time.Now().Unix()
		params := map[string]interface{}{
			"ml_movie_id": movie.ID,
			"ml_updated":  now,
			"ml_liked":    liked,
		}
		stmt, args, err := binder.BindNamed(queryUpdateML, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Delete deletes a user from the datastore.
func (s *movieStore) Delete(ctx context.Context, movie *models.MovieInfo) error {
	return s.db.Lock(func(execer database.Execer, binder database.Binder) error {
		params := toParams(movie)
		stmt, args, err := binder.BindNamed(queryDelete, params)
		if err != nil {
			return err
		}
		_, err = execer.Exec(stmt, args...)
		return err
	})
}

// Count returns a count of movies.
func (s *movieStore) Count(ctx context.Context) (int64, error) {
	var out int64
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		return queryer.QueryRow(queryCount).Scan(&out)
	})
	return out, err
}

// Count returns a count of Watch Laters.
func (s *movieStore) CountNI(ctx context.Context) (int64, error) {
	var out int64
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		return queryer.QueryRow(queryCountNI).Scan(&out)
	})
	return out, err
}

// Count returns a count of Watch Laters.
func (s *movieStore) CountWL(ctx context.Context) (int64, error) {
	var out int64
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		return queryer.QueryRow(queryCountWL).Scan(&out)
	})
	return out, err
}

// Count returns a count of Movie Learns
func (s *movieStore) CountML(ctx context.Context) (int64, error) {
	var out int64
	err := s.db.View(func(queryer database.Queryer, binder database.Binder) error {
		return queryer.QueryRow(queryCountML).Scan(&out)
	})
	return out, err
}

const queryCount = `
SELECT COUNT(*)
FROM movies
`

const queryCountNI = `
SELECT COUNT(*)
FROM movie_notes
`

const queryCountWL = `
SELECT COUNT(*)
FROM movie_watchlist
`

const queryCountML = `
SELECT COUNT(*)
FROM movie_learns
`

const queryInsert = `
INSERT INTO movies (
 movie_title
,movie_year
,movie_rating
,movie_votes
,movie_reviews
,movie_duration
,movie_released
,movie_istvseries
,movie_seasons
,movie_episodes
,movie_summary
,movie_metascore
,movie_review_count_user
,movie_review_count_critic
,movie_rtmeter
,movie_url_trailer_imdb
,movie_url_poster_imdb
,movie_created
,movie_updated
) VALUES (
 :movie_title
,:movie_year
,:movie_rating
,:movie_votes
,:movie_reviews
,:movie_duration
,:movie_released
,:movie_istvseries
,:movie_seasons
,:movie_episodes
,:movie_summary
,:movie_metascore
,:movie_review_count_user
,:movie_review_count_critic
,:movie_rtmeter
,:movie_url_trailer_imdb
,:movie_url_poster_imdb
,:movie_created
,:movie_updated
)
`

const queryMovieSearchesBase = `
SELECT
 search_id
,search_movie_id
,search_created
,search_from
`

const queryMovieBase = `
SELECT
 movie_id
,movie_title
,movie_year
,movie_rating
,movie_votes
,movie_reviews
,movie_duration
,movie_released
,movie_istvseries
,movie_seasons
,movie_episodes
,movie_summary
,movie_metascore
,movie_review_count_user
,movie_review_count_critic
,movie_rtmeter
,movie_url_trailer_imdb
,movie_url_poster_imdb
,movie_created
,movie_updated
`

const queryMovieNoteBase = `
SELECT
 note_id
,note_movie_id
,note_created
,note_updated
,note_from
,note_season
,note_episode
,note_hour
,note_minute
,note_second
,note_text
`

const queryMovieWLBase = `
SELECT
 wl_id
,wl_movie_id
,wl_created
,wl_updated
,wl_watched
`

const queryMovieMLBase = `
SELECT
 ml_id
,ml_movie_id
,ml_created
,ml_updated
,ml_liked
`

const queryKey = queryMovieBase + `
FROM movies
WHERE movie_title = :movie_title
`

const queryMovieFindByIDKey = queryMovieBase + `
FROM movies
WHERE movie_id = :movie_id
`

const queryMovieNIKey = queryMovieNoteBase + `
FROM movie_notes
WHERE note_movie_id = :note_movie_id
`

const queryMovieMLKey = queryMovieMLBase + `
FROM movie_learns
WHERE ml_movie_id = :ml_movie_id
`

const queryMovieWLKey = queryMovieWLBase + `
FROM movie_watchlist
WHERE wl_movie_id = :wl_movie_id
`

const queryMoviesAll = queryMovieBase + `
FROM movies
`

const queryMovieSearchesAll = queryMovieSearchesBase + `
FROM movie_searches
`

const queryMovieNoteAll = queryMovieNoteBase + `
FROM movie_notes
`

const queryMovieWatchLaterAll = queryMovieWLBase + `
FROM movie_watchlist
`

const queryMovieLearnAll = queryMovieMLBase + `
FROM movie_learns
`

const queryMovieSearches = queryMovieBase + `
FROM movies LEFT OUTER JOIN movie_searches ON movie_id = (
 SELECT search_id FROM movie_searches
 WHERE movie_searches.search_movie_id = movie_id
)
`

const queryUpdate = `
UPDATE movies
SET
 movie_title               = :movie_title
,movie_year                = :movie_year
,movie_rating              = :movie_rating
,movie_votes               = :movie_votes
,movie_reviews             = :movie_reviews
,movie_duration            = :movie_duration
,movie_released            = :movie_released
,movie_istvseries          = :movie_istvseries
,movie_seasons             = :movie_seasons
,movie_episodes            = :movie_episodes
,movie_summary             = :movie_summary
,movie_metascore           = :movie_metascore
,movie_review_count_user   = :movie_review_count_user
,movie_review_count_critic = :movie_review_count_critic
,movie_rtmeter             = :movie_rtmeter
,movie_url_trailer_imdb    = :movie_url_trailer_imdb
,movie_url_poster_imdb     = :movie_url_poster_imdb
,movie_created             = :movie_created
,movie_updated             = :movie_updated
WHERE movie_id = :movie_id
`

const queryUpdateNI = `
UPDATE movie_notes
SET
 note_updated = :note_updated
,note_season  = :note_season
,note_episode = :note_episode
,note_hour    = :note_hour
,note_minute  = :note_minute
,note_second  = :note_second
,note_text    = :note_text
WHERE note_movie_id = :note_movie_id
`

const queryUpdateWL = `
UPDATE movie_watchlist
SET
 wl_updated = :wl_updated
,wl_watched = :wl_watched
WHERE wl_movie_id = :wl_movie_id
`

const queryUpdateML = `
UPDATE movie_learns
SET
 ml_updated  = :ml_updated
,ml_liked    = :ml_liked
WHERE ml_movie_id = :ml_movie_id
`

const queryDelete = `
DELETE FROM movies WHERE movie_id = :movie_id
`

//Search History
const queryInsertSearch = `
INSERT INTO movie_searches (
 search_movie_id
,search_created
,search_from
) VALUES (
 :search_movie_id
,:search_created
,:search_from
)
`

//Movie Note
const queryInsertNI = `
INSERT INTO movie_notes (
 note_movie_id
,note_created
,note_updated
,note_from
,note_season
,note_episode
,note_hour
,note_minute
,note_second
,note_text
) VALUES (
 :note_movie_id
,:note_created
,:note_updated
,:note_from
,:note_season
,:note_episode
,:note_hour
,:note_minute
,:note_second
,:note_text
)
`

//Watch Later
const queryInsertWL = `
INSERT INTO movie_watchlist (
 wl_movie_id
,wl_created
,wl_updated
,wl_watched
) VALUES (
 :wl_movie_id
,:wl_created
,:wl_updated
,:wl_watched
)
`

//Movie Learn [Liked or not]
const queryInsertML = `
INSERT INTO movie_learns (
 ml_movie_id
,ml_created
,ml_updated
,ml_liked
) VALUES (
 :ml_movie_id
,:ml_created
,:ml_updated
,:ml_liked
)
`
