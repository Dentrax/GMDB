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
	Summary           string
	Credit            CreditInfo
	Metascore         string
	ReviewCountUser   string
	ReviewCountCritic string
	RTMeter           string
	URLTrailerIMDB    string
	URLPosterIMDB     string
}

type Movie struct {
	Info MovieInfo
	TL   Tagline
	PS   PlotSummary
	PK   PlotKeywords
	PG   ParentsGuide
}

//Load the struct from database table (movie_learns)
type MovieLearnInfo struct {
	ID      int64
	MovieID int64
	Created int64
	Updated int64
	Liked   bool
}

//Load the struct from database table (movie_watchlaters)
type MovieWatchLaterInfo struct {
	ID      int64
	MovieID int64
	Created int64
	Updated int64
	Watched bool
}

//Load the struct from database table (movie_searches)
type MovieSearchInfo struct {
	ID      int64
	MovieID int64
	Created int64
	From    int64
}

//Load the struct from database table (movie_notes)
type MovieNoteInfo struct {
	ID      int64
	MovieID int64
	Created int64
	Updated int64
	From    int64
	Hour    uint8
	Minute  uint8
	Second  uint8
	Text    string
}

//Operations interface for database functions
type MovieStore interface {
	//Create movie info
	Create(context.Context, *MovieInfo) error
	//Create search history
	CreateSearch(context.Context, *MovieInfo) error
	//Create Movie Note Info
	CreateNI(context.Context, *MovieInfo, *MovieNoteInfo) error
	//Create likes and dislikes for the movie
	CreateML(context.Context, *MovieInfo, bool) error
	//Create watch later
	CreateWL(context.Context, *MovieInfo) error
	// Update persists an updated movie to the datastore.
	Update(context.Context, *MovieInfo) error
	// Update persists an updated Movie Note Info to the datastore.
	UpdateNI(context.Context, *MovieInfo, *MovieNoteInfo) error
	// Update persists an updated Movie Learn to the datastore.
	UpdateWL(context.Context, *MovieInfo, bool) error
	// Update persists an updated Movie Learn to the datastore.
	UpdateML(context.Context, *MovieInfo, bool) error
	// Delete deletes a movie from the datastore.
	Delete(context.Context, *MovieInfo) error
	GetSearches(context.Context) ([]*MovieSearchInfo, error)
	GetMovieNoteList(context.Context) ([]*MovieNoteInfo, error)
	GetWatchLaterList(context.Context) ([]*MovieWatchLaterInfo, error)
	GetMovieLearnList(context.Context) ([]*MovieLearnInfo, error)
	FindByTitle(context.Context, string) (*MovieInfo, error)
	FindByID(context.Context, int64) (*MovieInfo, error)
	FindNI(context.Context, int64) (*MovieNoteInfo, error)
	FindWL(context.Context, int64) (*MovieWatchLaterInfo, error)
	FindML(context.Context, int64) (*MovieLearnInfo, error)
	Count(context.Context) (int64, error)
	CountNI(context.Context) (int64, error)
	CountWL(context.Context) (int64, error)
	CountML(context.Context) (int64, error)
}
