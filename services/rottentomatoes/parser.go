package rottentomatoes

import (
	"encoding/json"
	"fmt"
	"log"
	"regexp"
	"strings"

	"gmdb/models"

	"github.com/puerkitobio/goquery"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FixSpace(input string) string {
	input = strings.Replace(input, "<br> \n", "", -1)
	input = strings.TrimSpace(input)
	input = strings.Replace(input, " ", " ", -1)
	return input
}

func ParseSearchMovies(doc *goquery.Document) *models.SearchResponse {
	//^http://www\.rottentomatoes\.com/tv/[^/]+/?$
	r := regexp.MustCompile(`(\{{*.*})`)

	firstScript := doc.Find("#main_container").First()
	matches := r.FindString(FixSpace(firstScript.Text()))

	data := &models.RTSearchResult{}
	err := json.Unmarshal([]byte(matches), data)

	if err != nil {
		log.Fatalln(err)
	}

	result := new(models.SearchResponse)

	for _, info := range data.MovieInfos {
		resp := models.SearchResult{}

		resp.URL = info.URL
		resp.Title = info.Name
		resp.Year = fmt.Sprintf("(%v)", info.Year)
		result.Searches = append(result.Searches, resp)
	}

	result.TotalFound = uint(len(data.MovieInfos))

	return result
}

func ParseMovieInfo(doc *goquery.Document) *models.MovieInfo {
	movieInfo := new(models.MovieInfo)

	title := doc.Find("#heroImageContainer > a > h1")
	year := doc.Find("#heroImageContainer > a > span")

	duration := doc.Find("#mainColumn > section.panel.panel-rt.panel-box.movie_info.media > div > div.panel-body.content_body > ul > li:nth-child(8) > div.meta-value > time")
	releaseDate := doc.Find("#mainColumn > section.panel.panel-rt.panel-box.movie_info.media > div > div.panel-body.content_body > ul > li.meta-row.clearfix.js-theater-release-dates > div.meta-value > time")

	rtMeter := doc.Find("div.tomato-left .meter-value.superPageFontColor span").Text()
	rateLeft := doc.Find(".tomato-left div.hidden-xs").First()
	scoreAudince := doc.Find("div.media-body div.meter-value span")
	votes := doc.Find("#scorePanel > div.col-sm-8.col-xs-12.audience-panel > div.audience-info.hidden-xs.superPageFontColor > div:nth-child(2)").Before("span")

	//scoreAvgRating := s.After("span").Text()

	c := 0
	rateLeft.Find("span.subtle.superPageFontColor").Next().Each(func(i int, s *goquery.Selection) {
		if c == 1 { //Reviews Counted
			movieInfo.Reviews = FixSpace(s.Text())
		} else if c == 2 { //Fresh

		} else if c == 3 { //Rotten

		}
	})

	if len(rtMeter) == 0 {
		movieInfo.RTMeter = "-1"
	} else if len(rtMeter) > 4 {
		movieInfo.RTMeter = rtMeter[:3]
	} else {
		movieInfo.RTMeter = rtMeter[:2]
	}

	movieInfo.Title = FixSpace(title.Text())
	movieInfo.Year = FixSpace(year.Text())

	movieInfo.Duration = FixSpace(duration.Text())
	movieInfo.Released = FixSpace(releaseDate.Text())

	movieInfo.Rating = FixSpace(scoreAudince.Text())
	movieInfo.Votes = FixSpace(votes.Text())

	creditInfo := new(models.CreditInfo)

	doc.Find("#mainColumn > section.panel.panel-rt.panel-box.movie_info.media > div > div.panel-body.content_body > ul > li:nth-child(2) > div.meta-value").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/browse/opening/?genres") {
				genre := FixSpace(l.Text())
				movieInfo.Genres = append(movieInfo.Genres, genre)
			}
		})
		c++
	})

	doc.Find("#mainColumn > section.panel.panel-rt.panel-box.movie_info.media > div > div.panel-body.content_body > ul > li:nth-child(4) > div.meta-value").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/celebrity") {
				name := FixSpace(l.Text())
				creditInfo.Writers = append(creditInfo.Writers, name)
			}
		})
		c++
	})

	doc.Find("#mainColumn > section.panel.panel-rt.panel-box.movie_info.media > div > div.panel-body.content_body > ul > li:nth-child(3) > div.meta-value").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/celebrity") {
				name := FixSpace(l.Text())
				creditInfo.Directors = append(creditInfo.Directors, name)
			}
		})
		c++
	})

	movieInfo.Credit = *creditInfo

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
