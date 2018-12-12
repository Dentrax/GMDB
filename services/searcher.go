// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package services

import (
	"log"

	"gmdb/models"
	"gmdb/services/imdb"
	"gmdb/services/rottentomatoes"
)

type Searcher struct {
	MinSelectable uint8
	MaxSelectable uint8
	Request       models.SearchRequest
}

func NewSearcher(request models.SearchRequest) *Searcher {
	return &Searcher{
		Request: request,
	}
}

func (s *Searcher) GetSearchResponses() []models.SearchResponse {
	responses := []models.SearchResponse{}

	engineIMDB := imdb.New("IMDB", s.Request)
	engineRT := rottentomatoes.New("RT", s.Request)

	//Scan IMDB
	var okIMDB = false
	var resIMDB *models.SearchResponse
	if s.Request.ScanIMDB {
		resIMDB = engineIMDB.SearchMovie(&s.Request)
		resIMDB.SearchEngine = "IMDB"
		okIMDB = resIMDB != nil

		if okIMDB {
			responses = append(responses, *resIMDB)
		}
	}

	//Scan RT
	var okRT = false
	var resRT *models.SearchResponse
	if s.Request.ScanRT {
		resRT = engineRT.SearchMovie(&s.Request)
		resRT.SearchEngine = "RottenTomatoes"

		okRT = resRT != nil
		if okRT {
			responses = append(responses, *resRT)
		}
	}

	return responses
}

func (p *Searcher) GetMovie(engine string, response models.SearchResult) *models.Movie {
	//	r, _ := regexp.Compile("^(?i)(https?)://(www.imdb.com/title/)(tt(\\d)).*$")
	//	if r.MatchString(strings.TrimSpace(p.Request.URL)) && !strings.HasSuffix(p.Request.URL, "/") {
	if engine == "IMDB" {
		url := "https://www.imdb.com/title/" + response.ID

		request := models.SearchRequest{
			Title: response.Title,
			Year:  response.Year,
			ID:    response.ID,
			URL:   url,
		}
		client := imdb.New("IMDB", request)
		movieInfo, err := client.GetMovie()
		if err != nil {
			log.Fatalln("nil")
		}
		return movieInfo
	}

	return nil
}
