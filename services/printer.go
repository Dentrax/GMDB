// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package services

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"gmdb/models"
	"gmdb/services/imdb"
	"gmdb/services/rottentomatoes"

	"github.com/ttacon/chalk"
)

type Printer struct {
	Filter  models.ResultFilter
	Request models.SearchRequest
}

func New(filter models.ResultFilter, request models.SearchRequest) *Printer {
	return &Printer{
		Filter:  filter,
		Request: request,
	}
}

func NewSearchPrinter(request models.SearchRequest) *Printer {
	return &Printer{
		Filter:  models.ResultFilter{},
		Request: request,
	}
}

func (p *Printer) PrintSearchResponse() {
	engineIMDB := imdb.New("IMDB", p.Request)
	engineRT := rottentomatoes.New("RT", p.Request)

	//Scan IMDB
	var okIMDB = false
	var resIMDB *models.SearchResponse
	if p.Request.ScanIMDB {
		resIMDB = engineIMDB.SearchMovie(&p.Request)
		okIMDB = resIMDB != nil
	}

	//Scan RT
	var okRT = false
	var resRT *models.SearchResponse
	if p.Request.ScanRT {
		resRT = engineRT.SearchMovie(&p.Request)
		okRT = resRT != nil
	}

	counter := 0

	if okIMDB {
		if len(resIMDB.Searches) > 0 {
			for i := range resIMDB.Searches[:10] {
				fmt.Printf("%2d) ", counter+1)
				p.printInfo(resIMDB.Searches[i].Title, resIMDB.Searches[i].Year)
				counter++
			}
			if len(resIMDB.Searches) > 10 {
				moreCount := len(resIMDB.Searches) - 10
				fmt.Printf("%2d) ", 0)
				p.printInfo(fmt.Sprintf("%v", moreCount), "more...")
			}
		}
	}

	if okRT {
		if len(resRT.Searches) > 0 {
			for i := range resRT.Searches[:10] {
				fmt.Printf("%2d) ", counter+1)
				p.printInfo(resRT.Searches[i].Title, resRT.Searches[i].Year)
				counter++
			}
			if len(resRT.Searches) > 10 {
				moreCount := len(resRT.Searches) - 10
				fmt.Printf("%2d) ", 0)
				p.printInfo(fmt.Sprintf("%v", moreCount), "more...")
			}
		}
	}
}

func (p *Printer) GetPrint() {

	//test := rottentomatoes.New("RottenTomatoes", "https://www.rottentomatoes.com")
	//test.GetMovie()

	r, _ := regexp.Compile("^(?i)(https?)://(www.imdb.com/title/)(tt(\\d)).*$")

	if r.MatchString(strings.TrimSpace(p.Request.URL)) && !strings.HasSuffix(p.Request.URL, "/") {
		cl := imdb.New("IMDB", p.Request)
		res, err := cl.GetMovie()
		if err != nil {
			log.Fatalln("nil")
		}
		p.printBanner()
		p.printFullMovie(*res)
	} else {
		log.Fatalln("URL format MUST be in https://www.imdb.com/title/ttID and musn't end with '/'")
	}
}

func (p *Printer) printFullMovie(movie models.Movie) {
	if p.Filter.Title {
		p.printInfo("Movie Name: ", movie.Info.Title)
	}
	if p.Filter.Year {
		p.printInfo("Year: ", movie.Info.Year)
	}
	if p.Filter.Released {
		p.printInfo("Released: ", movie.Info.Released)
	}
	if p.Filter.Rating {
		p.printInfo("Rating: ", movie.Info.Rating+" ("+movie.Info.Votes+")")
	}
	if p.Filter.Genres {
		p.printInfo("Genres: ", strings.Join(movie.Info.Genres, ", "))
	}
	if p.Filter.Duration {
		p.printInfo("Duration: ", movie.Info.Duration)
	}
	if p.Filter.Summary {
		p.printInfo("Summary: ", movie.Info.Summary)
	}
	if p.Filter.Directors {
		p.printInfo("Directors: ", strings.Join(movie.Info.Credit.Directors, ", "))
	}
	if p.Filter.Writers {
		p.printInfo("Writers: ", strings.Join(movie.Info.Credit.Writers, ", "))
	}
	if p.Filter.Stars {
		p.printInfo("Stars: ", strings.Join(movie.Info.Credit.Stars, ", "))
	}
	if p.Filter.Tagline {
		p.printInfo("Tagline: ", strings.Join(movie.TL.Tags, ", "))
	}
	if p.Filter.Summaries {
		p.printInfo("Summaries: ", string(movie.PS.Total))
		for i := range movie.PS.Summaries {
			fmt.Println()
			p.printInfo(movie.PS.Summaries[i].Author, movie.PS.Summaries[i].Text)
		}
	}
	if p.Filter.Keywords {
		sums := movie.PK.Keywords[0:10]
		p.printInfo("Keywords: ", strings.Join(sums, ", "))
	}
	if p.Filter.ParentsGuide {
		p.printInfo("ParentsGuide: ", "")

		fmt.Printf(chalk.Bold.TextStyle("- Sex & Nudity: "))
		p.printForRate(movie.PG.Nudity.FinalRate)

		fmt.Printf(chalk.Bold.TextStyle("- Violence & Gore: "))
		p.printForRate(movie.PG.Violence.FinalRate)

		fmt.Printf(chalk.Bold.TextStyle("- Profanity: "))
		p.printForRate(movie.PG.Profanity.FinalRate)

		fmt.Printf(chalk.Bold.TextStyle("- Alcohol & Drugs & Smoking: "))
		p.printForRate(movie.PG.Alcohol.FinalRate)

		fmt.Printf(chalk.Bold.TextStyle("- Frightening & Intense: "))
		p.printForRate(movie.PG.Frightening.FinalRate)
	}
}

func (p *Printer) printInfo(s1 string, s2 string) {
	if p.Filter.NoColor {
		fmt.Println(s1, s2)
	} else {
		fmt.Println(chalk.Magenta.Color(s1), chalk.Bold.TextStyle(s2))
	}
}

func (p *Printer) printForRate(rate string) {
	blackOnRed := chalk.Black.NewStyle().
		WithBackground(chalk.Red).
		WithTextStyle(chalk.Bold)

	blackOnYellow := chalk.Black.NewStyle().
		WithBackground(chalk.Yellow).
		WithTextStyle(chalk.Bold)

	blackOnGreen := chalk.Black.NewStyle().
		WithBackground(chalk.Green).
		WithTextStyle(chalk.Bold)

	blackOnWhite := chalk.Black.NewStyle().
		WithBackground(chalk.White).
		WithTextStyle(chalk.Bold)

	switch rate {
	case "EMPTY":
		if p.Filter.NoColor {
			fmt.Printf("EMPTY\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnWhite, "EMPTY", chalk.Reset)
		}
	case "NONE":
		if p.Filter.NoColor {
			fmt.Printf("NONE\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnWhite, "NONE", chalk.Reset)
		}
	case "MILD":
		if p.Filter.NoColor {
			fmt.Printf("MILD\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnGreen, "MILD", chalk.Reset)
		}
	case "MODERATE":
		if p.Filter.NoColor {
			fmt.Printf("MODERATE\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnYellow, "MODERATE", chalk.Reset)
		}
	case "SEVERE":
		if p.Filter.NoColor {
			fmt.Printf("SEVERE\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnRed, "SEVERE", chalk.Reset)
		}
	}
}

func (p *Printer) printBanner() {
	banner := `
 ██████╗ ███╗   ███╗██████╗ ██████╗
██╔════╝ ████╗ ████║██╔══██╗██╔══██╗
██║  ███╗██╔████╔██║██║  ██║██████╔╝
██║   ██║██║╚██╔╝██║██║  ██║██╔══██╗
╚██████╔╝██║ ╚═╝ ██║██████╔╝██████╔╝
 ╚═════╝ ╚═╝     ╚═╝╚═════╝ ╚═════╝
`
	if !p.Filter.NoBanner {
		if p.Filter.NoColor {
			fmt.Println(banner)
		} else {
			fmt.Println(chalk.Cyan, banner)
		}
	}
}
