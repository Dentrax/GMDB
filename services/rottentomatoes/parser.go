package rottentomatoes

import (
	"fmt"
	"strings"

	"gmdb/models"

	"github.com/puerkitobio/goquery"
)

func ParseSearchMovies(doc *goquery.Document) *models.SearchResponse {
	//^http://www\.rottentomatoes\.com/tv/[^/]+/?$
	result := new(models.SearchResponse)

	finder0 := doc.Find("div.search-results-root")
	fmt.Println(finder0.Text())

	if len(finder0.Nodes) > 0 {
		doc.Find("ul.results_ul li.bottom_divider.clearfix").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
			node := s.Find("a")

			fmt.Println(node.Text())
		})
	}

	result.TotalFound = uint(len(finder0.Nodes))

	return result
}

func ParseMovieInfo(doc *goquery.Document) *models.MovieInfo {
	movieInfo := new(models.MovieInfo)

	rating := doc.Find(".meter-value.superPageFontColor span").Text()
	if len(rating) == 0 {
		movieInfo.Rating = "-1"
	} else if len(rating) > 4 {
		movieInfo.Rating = rating[:3]
	} else {
		movieInfo.Rating = rating[:2]
	}

	return movieInfo
}

func ParseMovieReviews(doc *goquery.Document) string {
	finder := doc.Find("#reviews .review_quote")
	if len(finder.Nodes) > 0 {
		fmt.Println("Reviews from RT!")
		doc.Find("#reviews .review_quote").Each(func(i int, s *goquery.Selection) {
			review := s.Find("p").Text()
			fmt.Println(strings.TrimSpace(review))
			fmt.Println("-------------------")
		})
	} else {
		fmt.Println("Looks like Rt also needs the year argument!")
	}

	return ""
}
