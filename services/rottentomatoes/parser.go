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

	finder0 := doc.Find("section > section > ul")
	finder1 := doc.Find("section.movieSection > ul")
	finder2 := doc.Find("section.movieSection > ul > li")
	finder3 := doc.Find("movieSection")
	finder4 := doc.Find("section.movieSection")
	finder5 := doc.Find("section.movieSection ul.results_ul")
	finder6 := doc.Find("div.Search.Results li.bottom_divider.clearfix")
	finder7 := doc.Find("div.search-results-root span.bold")
	finder8 := doc.Find("div.search-results-root #details")
	finder9 := doc.Find("div.search-results-root li.bottom_divider.clearfix")
	finder10 := doc.Find("#movieSection > ul > li:nth-child(1)")
	finder11 := doc.Find("#movieSection > ul")
	finder12 := doc.Find("#movieSection > ul .bottom_divider.clearfix")
	finder13 := doc.Find("#movieSection > ul > li")
	finder14 := doc.Find("section.Search.Results .movieSection")
	finder15 := doc.Find("#search-results-root section")
	finder16 := doc.Find("#search-results-root .section")
	finder17 := doc.Find("#search-results-root .section .section")
	finder18 := doc.Find("#search-results-root > section > section")
	finder19 := doc.Find("#search-results-root > section .section")
	finder20 := doc.Find("#movieSection > .results_ul > li > div.details")
	finder21 := doc.Find("#movieSection > .results_ul > li")
	finder22 := doc.Find("#movieSection > .results_ul")
	finder23 := doc.Find("#movieSection")

	fmt.Printf("%d", len(finder0.Nodes))
	fmt.Printf("%d", len(finder1.Nodes))
	fmt.Printf("%d", len(finder2.Nodes))
	fmt.Printf("%d", len(finder3.Nodes))
	fmt.Printf("%d", len(finder4.Nodes))
	fmt.Printf("%d", len(finder5.Nodes))
	fmt.Printf("%d", len(finder6.Nodes))
	fmt.Printf("%d", len(finder7.Nodes))
	fmt.Printf("%d", len(finder8.Nodes))
	fmt.Printf("%d", len(finder9.Nodes))
	fmt.Printf("%d", len(finder10.Nodes))
	fmt.Printf("%d", len(finder11.Nodes))
	fmt.Printf("%d", len(finder12.Nodes))
	fmt.Printf("%d", len(finder13.Nodes))
	fmt.Printf("%d", len(finder14.Nodes))
	fmt.Printf("%d", len(finder15.Nodes))
	fmt.Printf("%d", len(finder16.Nodes))
	fmt.Printf("%d", len(finder17.Nodes))
	fmt.Printf("%d", len(finder18.Nodes))
	fmt.Printf("%d", len(finder19.Nodes))
	fmt.Printf("%d", len(finder20.Nodes))
	fmt.Printf("%d", len(finder21.Nodes))
	fmt.Printf("%d", len(finder22.Nodes))
	fmt.Printf("%d", len(finder23.Nodes))

	if len(finder0.Nodes) > 0 {
		doc.Find("ul.results_ul li.bottom_divider.clearfix").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
			node := s.Find("a")

			fmt.Println(node.Text())
		})
	}

	result.RTCount = uint(len(finder0.Nodes))

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
