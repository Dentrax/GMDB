// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package rottentomatoes

import (
	"log"

	"gmdb/models"
	"gmdb/services/common"

	"github.com/puerkitobio/goquery"
)

type RottenTomatoes struct {
	Name    string
	Request models.SearchRequest
}

func New(name string, request models.SearchRequest) *RottenTomatoes {
	return &RottenTomatoes{
		Name:    name,
		Request: request,
	}
}

func (s *RottenTomatoes) SearchMovie(request *models.SearchRequest) *models.SearchResponse {

	url := "https://www.rottentomatoes.com" + "/search/?search=" + request.Title

	rq, err := GetSearchMovies(services.GetDocumentFromURL(url))
	//year, id\ exactsearch

	if err != nil {
		log.Fatalln("nil")
	}

	return rq
}

func (s *RottenTomatoes) GetMovie() (*models.Movie, error) {

	url := s.Request.URL

	movie := new(models.Movie)

	//TODO: Optimization for spesific arguments

	mi, err1 := GetMovieInfo(services.GetDocumentFromURL(url))
	//mo, err1 := GetMovieReviews(services.GetDocumentFromURL(url))

	if err1 != nil {
		log.Fatalln("nil")
		return nil, nil
	}

	movie.Info = *mi

	return movie, nil
}

func GetSearchMovies(doc *goquery.Document) (*models.SearchResponse, error) {
	searches := ParseSearchMovies(doc)
	return searches, nil
}

func GetMovieInfo(doc *goquery.Document) (*models.MovieInfo, error) {
	info := ParseMovieInfo(doc)
	return info, nil
}

func GetMovieReviews(doc *goquery.Document) (string, error) {
	info := ParseMovieReviews(doc)
	return info, nil
}
