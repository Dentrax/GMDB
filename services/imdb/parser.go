// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"fmt"
	"strconv"
	"strings"

	"gmdb/models"

	"github.com/puerkitobio/goquery"
)

func ParseMovieInfo(doc *goquery.Document) *models.MovieInfo {
	movieInfo := new(models.MovieInfo)

	doc.Find("div > div > div > div .title_block .title_bar_wrapper").Each(func(i int, s *goquery.Selection) {
		stars := s.Find(".ratings_wrapper .imdbRating .ratingValue")
		rated := s.Find(".ratings_wrapper .imdbRating .small")

		title := s.Find(".titleBar .title_wrapper > h1").Not("span")
		year := s.Find(".titleBar .title_wrapper > h1").Children()

		duration := s.Find(".subtext").Children().First().Next()
		releaseDate := s.Find(".subtext").Children().Last()

		s.Find("div > a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/search") {
				movieInfo.Genres = append(movieInfo.Genres, strings.TrimSpace(l.Text()))
			}
		})

		movieInfo.Title = strings.TrimSpace(title.Text())
		movieInfo.Year = strings.TrimSpace(year.Text())

		movieInfo.Rating = strings.TrimSpace(stars.Text())
		movieInfo.Votes = strings.TrimSpace(rated.Text())

		movieInfo.Duration = strings.TrimSpace(duration.Text())
		movieInfo.Released = strings.TrimSpace(releaseDate.Text())

	})

	doc.Find("div > div > div > div .plot_summary_wrapper .summary_text").Each(func(i int, s *goquery.Selection) {
		summary := s.First()
		movieInfo.Summary = strings.TrimSpace(summary.Text())
	})

	creditInfo := new(models.CreditInfo)

	c := 0 //0: Directors, 1: Writers, 2: Stars
	doc.Find("div > div > div > div .plot_summary_wrapper .credit_summary_item").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/name") {
				name := strings.TrimSpace(l.Text())
				if c == 0 {
					creditInfo.Directors = append(creditInfo.Directors, name)
				}
				if c == 1 {
					creditInfo.Writers = append(creditInfo.Writers, name)
				}
				if c == 2 {
					creditInfo.Stars = append(creditInfo.Stars, name)
				}
			}
		})
		c++
	})

	movieInfo.Credit = *creditInfo

	return movieInfo
}

func ParseTagline(doc *goquery.Document) *models.Tagline {
	tagLine := new(models.Tagline)

	doc.Find("div > div > div").ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
		if s.HasClass("soda odd") || s.HasClass("soda even") {
			curr := strings.TrimSpace(s.Text())
			tagLine.Tags = append(tagLine.Tags, curr)
		}
	})

	return tagLine
}

func ParsePlotSummary(doc *goquery.Document) *models.PlotSummary {
	plotSummary := new(models.PlotSummary)

	c := uint(0)
	doc.Find("div > div > section > ul .ipl-zebra-list__item ").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Find("p").Text())
		author := strings.TrimSpace(s.Find("em").Text())

		summary := models.Summary{
			Author: author,
			Text:   text,
		}

		plotSummary.Summaries = append(plotSummary.Summaries, summary)
		c++
	})

	plotSummary.Total = c
	return plotSummary
}

func ParsePlotKeywords(doc *goquery.Document) *models.PlotKeywords {
	plotKeywords := new(models.PlotKeywords)

	c := uint(0)
	doc.Find("div > div > div > table > tbody .sodatext").Each(func(i int, s *goquery.Selection) {
		curr := strings.TrimSpace(s.Text())
		plotKeywords.Keywords = append(plotKeywords.Keywords, curr)
		c++
	})

	plotKeywords.Total = c
	return plotKeywords
}

func ParseParentsGuide(doc *goquery.Document) *models.ParentsGuide {
	rates := new(models.ParentsGuide)

	c := 0

	//for sub rows
	c2 := 0

	//for main rows
	c3 := 0
	cTotal := uint(0)
	//counts := SeverityRate{}

	var rts [5][4]uint

	doc.Find("div > div > section").ChildrenFiltered("section").Find("span").Each(func(i int, s *goquery.Selection) {
		if s.HasClass("ipl-status-pill") {
			curr := strings.TrimSpace(s.Text())
			//fmt.Printf("C: %d, Data: %s \n", c, curr)
			if c == 0 {
				rates.Nudity.FinalRate = curr
			}
			if c == 2 {
				rates.Violence.FinalRate = curr
			}
			if c == 4 {
				rates.Profanity.FinalRate = curr
			}
			if c == 6 {
				rates.Alcohol.FinalRate = curr
			}
			if c == 8 {
				rates.Frightening.FinalRate = curr
			}
			c++
		}
		if s.HasClass("ipl-vote-button__details") {
			curr := string(strings.TrimSpace(s.Text()))

			data, err := strconv.ParseUint(curr, 10, 32)

			if err != nil {
				fmt.Println(err)
			}

			//fmt.Printf("C2: %d, C3: %d, Data: %d\n", c2, c3, data)

			cTotal += uint(data)

			rts[c3][c2] = uint(data)

			if c2 != 0 && c2%3 == 0 {
				c3++
				c2 = 0
			} else {
				c2++
			}
		}
	})

	//fmt.Println(rts[3][1])

	//fmt.Printf("N: %d, V: %d, P: %d, A: %d, F: %d", rates.Nudity, rates.Violence, rates.Profanity, rates.Alcohol, rates.Frightening)

	return rates
}
