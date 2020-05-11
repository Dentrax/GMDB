// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"log"

	"github.com/Dentrax/GMDB/models"
	"github.com/Dentrax/GMDB/pkg/cache"
	"github.com/Dentrax/GMDB/services/common"

	"github.com/PuerkitoBio/goquery"
)

type IMDB struct {
	Name    string
	Request models.SearchRequest
}

func New(name string, request models.SearchRequest) *IMDB {
	return &IMDB{
		Name:    name,
		Request: request,
	}
}

func (s *IMDB) SearchMovie(request *models.SearchRequest) *models.SearchResponse {
	url := "https://www.imdb.com/find?q=" + request.Title + "&s=tt"

	if cache.IsFileExist(s.Name, "searches", request.Title) {
		result, err := cache.GetSearchResponse(s.Name, "searches", request.Title)
		if err != nil {
			log.Fatalln(err)
		}
		return result
	}

	doc := services.GetDocumentFromURL(url)
	result, err := GetSearchMovies(doc)
	if err != nil {
		log.Fatalln("nil")
	}
	return result
}

func (s *IMDB) GetMovie(onlyHome bool) (*models.Movie, error) {
	if cache.IsFileExist(s.Name, "movies", s.Request.ID) {
		result, err := cache.GetMovie(s.Name, "movies", s.Request.ID)
		if err != nil {
			log.Fatalln(err)
		}
		return result, nil
	}

	movie := new(models.Movie)

	urlTL := s.Request.URL + "/taglines"
	urlPS := s.Request.URL + "/plotsummary"
	urlPK := s.Request.URL + "/keywords"
	urlPG := s.Request.URL + "/parentalguide"

	//TODO: Optimization and err handling required immediately !!!

	mi, err := GetMovieInfo(services.GetDocumentFromURL(s.Request.URL))

	if err != nil {
		log.Fatalln("nil")
	}

	movie.Info = *mi

	if !onlyHome {
		tl, err := GetTagline(services.GetDocumentFromURL(urlTL))
		ps, err := GetPlotSummary(services.GetDocumentFromURL(urlPS))
		pk, err := GetPlotKeywords(services.GetDocumentFromURL(urlPK))
		pg, err := GetParentsGuide(services.GetDocumentFromURL(urlPG))

		if err != nil {
			log.Fatalln("nil")
		}

		movie.TL = *tl
		movie.PS = *ps
		movie.PK = *pk
		movie.PG = *pg
	}

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

func GetTagline(doc *goquery.Document) (*models.Tagline, error) {
	tags := ParseTagline(doc)
	return tags, nil
}

func GetPlotKeywords(doc *goquery.Document) (*models.PlotKeywords, error) {
	keywords := ParsePlotKeywords(doc)
	return keywords, nil
}

func GetPlotSummary(doc *goquery.Document) (*models.PlotSummary, error) {
	summary := ParsePlotSummary(doc)
	return summary, nil
}

func GetParentsGuide(doc *goquery.Document) (*models.ParentsGuide, error) {
	rates := ParseParentsGuide(doc)
	return rates, nil
}
