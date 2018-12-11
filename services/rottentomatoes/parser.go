package rottentomatoes

import (
	"fmt"
	"github.com/puerkitobio/goquery"
	"strings"
)

func ParseMovieScore(doc *goquery.Document) string {
	rating := doc.Find(".meter-value.superPageFontColor span").Text()
	if len(rating) == 0 {
		return "-1"
	} else if len(rating) > 4 {
		return rating[:3]
	} else {
		return rating[:2]
	}
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
