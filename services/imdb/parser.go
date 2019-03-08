// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Dentrax/GMDB/models"

	"github.com/puerkitobio/goquery"
)

func ParseSearchMovies(doc *goquery.Document) *models.SearchResponse {
	result := new(models.SearchResponse)
	finder := doc.Find("div.findSection td.result_text")

	rgxURL, _ := regexp.Compile("(tt.*?)\\w+")
	rgxTITLE, _ := regexp.Compile("\\(([^()]*)\\)")

	if len(finder.Nodes) > 0 {
		doc.Find("div.findSection td.result_text").Each(func(i int, s *goquery.Selection) {
			node := s.Find("a")
			title := FixSpace(node.Text())
			url, ok := node.Attr("href")
			if ok {
				item := models.SearchResult{
					ID:     rgxURL.FindString(FixSpace(url)),
					Title:  title,
					Year:   rgxTITLE.FindString(FixSpace(strings.Replace(s.Text(), title, "", -1))),
					TVShow: false,
				}
				result.Searches = append(result.Searches, item)
			}
		})
	}

	result.TotalFound = uint(len(finder.Nodes))

	return result
}

func ParseMovieInfo(doc *goquery.Document) *models.MovieInfo {
	movieInfo := new(models.MovieInfo)

	doc.Find("div > div > div > div .title_block .title_bar_wrapper").Each(func(i int, s *goquery.Selection) {
		stars := s.Find(".ratings_wrapper .imdbRating .ratingValue")
		rated := s.Find(".ratings_wrapper .imdbRating .small")

		title := s.Find(".titleBar .title_wrapper > h1")
		year := s.Find(".titleBar .title_wrapper > h1").Children()

		duration := s.Find(".subtext > time")
		releaseDate := s.Find(".subtext").Children().Last()

		s.Find("div > a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/search") {
				movieInfo.Genres = append(movieInfo.Genres, strings.TrimSpace(l.Text()))
			}
		})

		//TODO: add Metascore, ReviewCountUser, ReviewCountCritic
		yearStr := strings.Replace(year.Text(), "(", "", -1)
		yearStr = strings.Replace(yearStr, ")", "", -1)

		titleStr := strings.Replace(title.Text(), year.Text(), "", -1)

		movieInfo.Title = FixSpace(titleStr)
		movieInfo.Year = FixSpace(yearStr)

		movieInfo.Rating = FixSpace(stars.Text())
		movieInfo.Votes = FixSpace(rated.Text())

		movieInfo.Duration = FixSpace(duration.Text())
		movieInfo.Released = FixSpace(releaseDate.Text())

		movieInfo.IsTVSeries = false
		movieInfo.Seasons = "0"
		movieInfo.Episodes = "0"
	})

	doc.Find("#title-episode-widget > div").Each(func(i int, s *goquery.Selection) {
		seasonFind := s.Find("div:nth-child(4)")
		seasons := seasonFind.Children().Length()

		yearFind := s.Find("div:nth-child(5)")
		year := yearFind.Children().Last()

		if seasons > 0 {
			movieInfo.Year = FixSpace(year.Text())

			movieInfo.IsTVSeries = true
			movieInfo.Seasons = strconv.Itoa(seasons)
			movieInfo.Episodes = "0"
		}
	})

	if movieInfo.IsTVSeries {
		episodesFind := doc.Find("#title-overview-widget > div.vital > div.button_panel.navigation_panel > a > div > div > span")
		episodesText := FixSpace(strings.Replace(episodesFind.Text(), "episodes", "", -1))
		episodes, _ := strconv.Atoi(episodesText)

		if episodes > 0 {
			movieInfo.Episodes = episodesText
		}
	}

	doc.Find("div > div > div > div .vital .slate_wrapper").Each(func(i int, s *goquery.Selection) {
		slate := s.Find("div.slate")
		urlTrailer, ok := slate.Find("a").Attr("data-video")

		rgxURL, _ := regexp.Compile("(vi.*?)\\w+")

		if ok {
			id := rgxURL.FindString(FixSpace(urlTrailer))
			movieInfo.URLTrailerIMDB = "https://www.imdb.com/videoplayer/" + id
		}
	})

	doc.Find("div > div > div > div .plot_summary_wrapper .summary_text").Each(func(i int, s *goquery.Selection) {

		s.Contents().Each(func(i int, s *goquery.Selection) {
			if goquery.NodeName(s) == "#text" {
				movieInfo.Summary += FixSpace(s.Text())
			}
		})

	})

	creditInfo := new(models.CreditInfo)

	doc.Find("div > div > div > div .plot_summary_wrapper .credit_summary_item").Each(func(i int, s *goquery.Selection) {
		typo := s.Find("h4")

		s.Find("a").Each(func(j int, l *goquery.Selection) {
			links, _ := l.Attr("href")
			if strings.HasPrefix(links, "/name") {
				name := FixSpace(l.Text())

				//FIXME: Creator and Director? Seperate them.
				if strings.Contains(FixSpace(typo.Text()), "Creator") {
					creditInfo.Directors = append(creditInfo.Directors, name)
				} else if strings.Contains(FixSpace(typo.Text()), "Director") {
					creditInfo.Directors = append(creditInfo.Directors, name)
				} else if strings.Contains(FixSpace(typo.Text()), "Star") {
					creditInfo.Stars = append(creditInfo.Stars, name)
				} else if strings.Contains(FixSpace(typo.Text()), "Writer") {
					creditInfo.Writers = append(creditInfo.Writers, name)
				}
			}
		})
	})

	movieInfo.Credit = *creditInfo

	return movieInfo
}

func FixSpace(input string) string {
	input = strings.Replace(input, "<br> \n", "", -1)
	input = strings.TrimSpace(input)
	input = strings.Replace(input, " ", " ", -1)
	input = strings.Replace(input, "»", "", -1)
	return input
}

func ParseTagline(doc *goquery.Document) *models.Tagline {
	tagLine := new(models.Tagline)

	doc.Find("div > div > div").ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
		if s.HasClass("soda odd") || s.HasClass("soda even") {
			curr := FixSpace(s.Text())
			if !strings.Contains(curr, "we don't have any") {
				tagLine.Tags = append(tagLine.Tags, curr)
			}
		}
	})

	return tagLine
}

func ParsePlotSummary(doc *goquery.Document) *models.PlotSummary {
	plotSummary := new(models.PlotSummary)

	c := uint(0)
	doc.Find("div > div > section > ul .ipl-zebra-list__item ").Each(func(i int, s *goquery.Selection) {
		c++
		text := FixSpace(s.Find("p").Text())
		author := FixSpace(s.Find("em").Text())
		author = strings.Replace(author, "—", "", -1)

		summary := models.Summary{
			Author: author,
			Text:   text,
		}

		plotSummary.Summaries = append(plotSummary.Summaries, summary)
	})

	plotSummary.Total = c
	return plotSummary
}

func ParsePlotKeywords(doc *goquery.Document) *models.PlotKeywords {
	plotKeywords := new(models.PlotKeywords)

	c := uint(0)
	doc.Find("div > div > div > table > tbody .sodatext").Each(func(i int, s *goquery.Selection) {
		curr := FixSpace(s.Text())
		plotKeywords.Keywords = append(plotKeywords.Keywords, curr)
		c++
	})

	if c > 10 {
		plotKeywords.Keywords = plotKeywords.Keywords[:10]
	}
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
			curr := FixSpace(s.Text())
			curr = strings.ToUpper(curr)
			//fmt.Printf("C: %d, Data: %s \n", c, curr)
			if c == 0 {
				rates.Nudity.FinalRate = GetRateOrEmpty(curr)
			}
			if c == 2 {
				rates.Violence.FinalRate = GetRateOrEmpty(curr)
			}
			if c == 4 {
				rates.Profanity.FinalRate = GetRateOrEmpty(curr)
			}
			if c == 6 {
				rates.Alcohol.FinalRate = GetRateOrEmpty(curr)
			}
			if c == 8 {
				rates.Frightening.FinalRate = GetRateOrEmpty(curr)
			}
			c++
		}
		if s.HasClass("ipl-vote-button__details") {
			curr := string(FixSpace(s.Text()))

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

	rates.TotalVote = cTotal

	//TODO: Find better solution for empty PG areas
	if c <= 5 {
		rates.Alcohol.FinalRate = "EMPTY"
		rates.Frightening.FinalRate = "EMPTY"
	}

	//fmt.Println(rts[3][1])

	//fmt.Printf("N: %d, V: %d, P: %d, A: %d, F: %d", rates.Nudity, rates.Violence, rates.Profanity, rates.Alcohol, rates.Frightening)

	return rates
}

func GetRateOrEmpty(rate string) string {
	if len(FixSpace(rate)) < 4 || rate == "" {
		return "EMPTY"
	}
	return rate
}
