// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

import (
	"context"
)

// MovieInfo stores all of the informations about the movie
type MovieInfo struct {
	ID                int64
	Title             string
	Year              string
	Rating            string
	Votes             string
	Reviews           string
	Genres            []string
	Duration          string
	Released          string
	IsTVSeries        bool
	Seasons           string
	Episodes          string
	Summary           string
	Credit            CreditInfo
	Metascore         string
	ReviewCountUser   string
	ReviewCountCritic string
	RTMeter           string
	URLTrailerIMDB    string
	URLPosterIMDB     string
	Created           int64
	Updated           int64
}

// Movie stores the master informations using MovieInfo, TagLine, PlotSummary, PlotKeywords, ParentsGuide
type Movie struct {
	Info MovieInfo
	TL   Tagline
	PS   PlotSummary
	PK   PlotKeywords
	PG   ParentsGuide
}

// CreditInfo stores credit informations like (Directors, Writers, Stars)
type CreditInfo struct {
	Directors []string
	Writers   []string
	Stars     []string
}

// MovieLearnInfo represents a database row from the table (movie_learns)
type MovieLearnInfo struct {
	ID      int64
	MovieID int64
	Created int64
	Updated int64
	Liked   bool
}

// MovieWatchLaterInfo represents a database row from the table (movie_watchlist)
type MovieWatchLaterInfo struct {
	ID      int64
	MovieID int64
	Created int64
	Updated int64
	Watched bool
}

// MovieSearchInfo represents a database row from the table (movie_searches)
type MovieSearchInfo struct {
	ID      int64
	MovieID int64
	Created int64
	From    string
}

// MovieNoteInfo represents a database row from the table (movie_notes)
type MovieNoteInfo struct {
	ID      int64
	MovieID int64
	Created int64
	Updated int64
	From    int64
	Season  uint8
	Episode uint8
	Hour    uint8
	Minute  uint8
	Second  uint8
	Text    string
}

// MovieStore defines operations for working with movies.
type MovieStore interface {
	// Create persists a new movie info to the datastore.
	Create(context.Context, *MovieInfo) error

	// CreateSearch persists a new movie search info to the datastore.
	CreateSearch(context.Context, *MovieInfo, string) error

	// CreateNI persists a new note info to the datastore.
	CreateNI(context.Context, *MovieInfo, *MovieNoteInfo) error

	// CreateML persists a new movie like to the datastore.
	CreateML(context.Context, *MovieInfo, bool) error

	// CreateWL persists a new watch later to the datastore.
	CreateWL(context.Context, *MovieInfo, bool) error

	// Update persists an updated movie to the datastore.
	Update(context.Context, *MovieInfo) error

	// UpdateNI persists an updated note info to the datastore.
	UpdateNI(context.Context, *MovieInfo, *MovieNoteInfo) error

	// UpdateWL persists an updated watch later to the datastore by watched
	UpdateWL(context.Context, *MovieInfo, bool) error

	// UpdateML persists an updated movie info to the datastore by like.
	UpdateML(context.Context, *MovieInfo, bool) error

	// Delete deletes a movie from the datastore.
	Delete(context.Context, *MovieInfo) error

	// GetMovies resutns a list of MovieInfo from the datastore.
	GetMovies(context.Context) ([]*MovieInfo, error)

	// GetSearches resutns a list of MovieSearchInfo from the datastore.
	GetSearches(context.Context) ([]*MovieSearchInfo, error)

	// GetMovieNoteList resutns a list of MovieNoteInfo from the datastore.
	GetMovieNoteList(context.Context) ([]*MovieNoteInfo, error)

	// GetWatchLaterList resutns a list of MovieWatchLaterInfo from the datastore.
	GetWatchLaterList(context.Context) ([]*MovieWatchLaterInfo, error)

	// GetMovieLearnList resutns a list of MovieLearnInfo from the datastore.
	GetMovieLearnList(context.Context) ([]*MovieLearnInfo, error)

	// FindByTitle returns a MovieInfo from the datastore by movie Title.
	FindByTitle(context.Context, string) (*MovieInfo, error)

	// FindByID returns a MovieInfo from the datastore by movie ID.
	FindByID(context.Context, int64) (*MovieInfo, error)

	// FindNI returns a MovieNoteInfo from the datastore by movie ID.
	FindNI(context.Context, int64) (*MovieNoteInfo, error)

	// FindWL returns a MovieWatchLaterInfo from the datastore by movie ID.
	FindWL(context.Context, int64) (*MovieWatchLaterInfo, error)

	// FindML returns a MovieLearnInfo from the datastore by movie ID.
	FindML(context.Context, int64) (*MovieLearnInfo, error)

	// Count returns a a count of movie infos.
	Count(context.Context) (int64, error)

	// CountNI returns a a count of note infos.
	CountNI(context.Context) (int64, error)

	// CountWL returns a a count of watch later.
	CountWL(context.Context) (int64, error)

	// CountML returns a a count of movies like.
	CountML(context.Context) (int64, error)
}
