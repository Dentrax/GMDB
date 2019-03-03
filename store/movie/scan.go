package movie

import (
	"database/sql"

	"gmdb/models"
	"gmdb/store/database"
)

//Ref: https://github.com/drone/drone/blob/master/store/user/scan.go

// helper function converts the User structure to a set
// of named query parameters.
func toParams(m *models.MovieInfo) map[string]interface{} {
	return map[string]interface{}{
		"movie_id":                  m.ID,
		"movie_title":               m.Title,
		"movie_year":                m.Year,
		"movie_rating":              m.Rating,
		"movie_votes":               m.Votes,
		"movie_reviews":             m.Reviews,
		"movie_duration":            m.Duration,
		"movie_released":            m.Released,
		"movie_istvseries":          m.IsTVSeries,
		"movie_seasons":             m.Seasons,
		"movie_episodes":            m.Episodes,
		"movie_summary":             m.Summary,
		"movie_metascore":           m.Metascore,
		"movie_review_count_user":   m.ReviewCountUser,
		"movie_review_count_critic": m.ReviewCountCritic,
		"movie_rtmeter":             m.RTMeter,
		"movie_url_trailer_imdb":    m.URLTrailerIMDB,
		"movie_url_poster_imdb":     m.URLPosterIMDB,
		"movie_created":             m.Created,
		"movie_updated":             m.Updated,
	}
}

func toParamsSI(m *models.MovieSearchInfo) map[string]interface{} {
	return map[string]interface{}{
		"search_id":       m.ID,
		"search_movie_id": m.MovieID,
		"search_created":  m.Created,
		"search_from":     m.From,
	}
}

func toParamsNI(m *models.MovieNoteInfo) map[string]interface{} {
	return map[string]interface{}{
		"note_id":       m.ID,
		"note_movie_id": m.MovieID,
		"note_created":  m.Created,
		"note_updated":  m.Updated,
		"note_from":     m.From,
		"note_season":   m.Season,
		"note_episode":  m.Episode,
		"note_hour":     m.Hour,
		"note_minute":   m.Minute,
		"note_second":   m.Second,
		"note_text":     m.Text,
	}
}

func toParamsWL(m *models.MovieWatchLaterInfo) map[string]interface{} {
	return map[string]interface{}{
		"wl_id":       m.ID,
		"wl_movie_id": m.MovieID,
		"wl_created":  m.Created,
		"wl_updated":  m.Updated,
		"wl_watched":  m.Watched,
	}
}

func toParamsML(m *models.MovieLearnInfo) map[string]interface{} {
	return map[string]interface{}{
		"ml_id":       m.ID,
		"ml_movie_id": m.MovieID,
		"ml_created":  m.Created,
		"ml_updated":  m.Updated,
		"ml_liked":    m.Liked,
	}
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRow(scanner database.Scanner, dest *models.MovieInfo) error {
	return scanner.Scan(
		&dest.ID,
		&dest.Title,
		&dest.Year,
		&dest.Rating,
		&dest.Votes,
		&dest.Reviews,
		&dest.Duration,
		&dest.Released,
		&dest.IsTVSeries,
		&dest.Seasons,
		&dest.Episodes,
		&dest.Summary,
		&dest.Metascore,
		&dest.ReviewCountUser,
		&dest.ReviewCountCritic,
		&dest.RTMeter,
		&dest.URLTrailerIMDB,
		&dest.URLPosterIMDB,
		&dest.Created,
		&dest.Updated,
	)
}

//Scanner for SearchInfo
func scanRowSI(scanner database.Scanner, dest *models.MovieSearchInfo) error {
	return scanner.Scan(
		&dest.ID,
		&dest.MovieID,
		&dest.Created,
		&dest.From,
	)
}

//Scanner for WatchMovieLearnInfo
func scanRowNI(scanner database.Scanner, dest *models.MovieNoteInfo) error {
	return scanner.Scan(
		&dest.ID,
		&dest.MovieID,
		&dest.Created,
		&dest.Updated,
		&dest.From,
		&dest.Season,
		&dest.Episode,
		&dest.Hour,
		&dest.Minute,
		&dest.Second,
		&dest.Text,
	)
}

//Scanner for WatchLaterInfo
func scanRowWL(scanner database.Scanner, dest *models.MovieWatchLaterInfo) error {
	return scanner.Scan(
		&dest.ID,
		&dest.MovieID,
		&dest.Created,
		&dest.Updated,
		&dest.Watched,
	)
}

//Scanner for WatchMovieLearnInfo
func scanRowML(scanner database.Scanner, dest *models.MovieLearnInfo) error {
	return scanner.Scan(
		&dest.ID,
		&dest.MovieID,
		&dest.Created,
		&dest.Updated,
		&dest.Liked,
	)
}

// helper function scans the sql.Row and copies the column
// values to the destination object.
func scanRows(rows *sql.Rows) ([]*models.MovieInfo, error) {
	defer rows.Close()

	movies := []*models.MovieInfo{}
	for rows.Next() {
		movie := new(models.MovieInfo)
		err := scanRow(rows, movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func scanRowsSI(rows *sql.Rows) ([]*models.MovieSearchInfo, error) {
	defer rows.Close()

	movies := []*models.MovieSearchInfo{}
	for rows.Next() {
		movie := new(models.MovieSearchInfo)
		err := scanRowSI(rows, movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func scanRowsWL(rows *sql.Rows) ([]*models.MovieWatchLaterInfo, error) {
	defer rows.Close()

	movies := []*models.MovieWatchLaterInfo{}
	for rows.Next() {
		movie := new(models.MovieWatchLaterInfo)
		err := scanRowWL(rows, movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func scanRowsML(rows *sql.Rows) ([]*models.MovieLearnInfo, error) {
	defer rows.Close()

	movies := []*models.MovieLearnInfo{}
	for rows.Next() {
		movie := new(models.MovieLearnInfo)
		err := scanRowML(rows, movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}

func scanRowsNI(rows *sql.Rows) ([]*models.MovieNoteInfo, error) {
	defer rows.Close()

	movies := []*models.MovieNoteInfo{}
	for rows.Next() {
		movie := new(models.MovieNoteInfo)
		err := scanRowNI(rows, movie)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
