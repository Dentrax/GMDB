// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package rottentomatoes

import (
	"fmt"

	"gmdb/models"
	"gmdb/services/common"

	"github.com/puerkitobio/goquery"
)

type RottenTomatoes struct {
	Name string
	URL  string
}

func New(name string, url string) *RottenTomatoes {
	return &RottenTomatoes{
		Name: name,
		URL:  url,
	}
}

func (s *RottenTomatoes) GetMovie() (*models.Movie, error) {

	url := s.URL + "/m/" + "ghostbusters_2016"

	fmt.Println(url)
	//movie := new(models.Movie)

	//TODO: Optimization for spesific arguments

	mi, err1 := GetMovieInfo(services.GetDocumentFromURL(url))
	mo, err1 := GetMovieReviews(services.GetDocumentFromURL(url))

	if err1 != nil {
		fmt.Println("nil")
		return nil, nil
	}
	//movieName := strings.Replace(mname, " ", "_", -1)

	//if len(year) == 4 {
	//	urlis = urlis + "_" + year
	//}
	fmt.Println(mi)
	fmt.Println(mo)

	return nil, nil
}

func GetMovieInfo(doc *goquery.Document) (string, error) {
	info := ParseMovieScore(doc)
	return info, nil
}

func GetMovieReviews(doc *goquery.Document) (string, error) {
	info := ParseMovieReviews(doc)
	return info, nil
}
