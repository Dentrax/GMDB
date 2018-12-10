// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"log"
	"net/http"
	"os"
	"unicode/utf8"

	"gmdb/models"

	"github.com/puerkitobio/goquery"
)

type Service struct {
	Name string
	URL  string
}

func New(name string, url string) *Service {
	return &Service{
		Name: name,
		URL:  url,
	}
}

func (s *Service) GetMovie() (*models.Movie, error) {

	urlTL := s.URL + "/taglines"
	urlPS := s.URL + "/plotsummary"
	urlPK := s.URL + "/keywords"
	urlPG := s.URL + "/parentalguide"

	movie := new(models.Movie)

	//TODO: Optimization for spesific arguments

	mi, err1 := GetMovieInfo(GetDocumentFromURL(s.URL))

	tl, err2 := GetTagline(GetDocumentFromURL(urlTL))
	ps, err3 := GetPlotSummary(GetDocumentFromURL(urlPS))
	pk, err4 := GetPlotKeywords(GetDocumentFromURL(urlPK))
	pg, err5 := GetParentsGuide(GetDocumentFromURL(urlPG))

	if err1 != nil {
		log.Fatalln("nil")
	}
	if err2 != nil {
		log.Fatalln("nil")
	}
	if err3 != nil {
		log.Fatalln("nil")
	}
	if err4 != nil {
		log.Fatalln("nil")
	}
	if err5 != nil {
		log.Fatalln("nil")
	}

	movie.Info = *mi

	movie.TL = *tl
	movie.PS = *ps
	movie.PK = *pk
	movie.PG = *pg

	return movie, nil
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

func GetDocumentFromURL(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return doc
}

func GetDocumentFromFile(filename string) *goquery.Document {
	file, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
		return nil
	}

	defer file.Close()
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(e)
		return nil
	}
	if !utf8.ValidString(doc.Text()) {
		log.Fatalf("DOC: %s", "NOT UTF-8 FORMAT")
		return nil
	}
	return doc
}
