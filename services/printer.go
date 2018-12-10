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

	"github.com/ttacon/chalk"
	"gmdb/services/imdb"
)

type Printer struct {
	Filter models.ResultFilter
	URL    string
	ID     string
}

func New(filter models.ResultFilter, url string, id string) *Printer {
	return &Printer{
		Filter: filter,
		URL:    url,
		ID:     id,
	}
}

func (p *Printer) GetPrint() {

	r, _ := regexp.Compile("^(?i)(https?)://(www.imdb.com/title/)(tt(\\d)).*$")

	if r.MatchString(strings.TrimSpace(p.URL)) && !strings.HasSuffix(p.URL, "/") {
		cl := imdb.New("IMDB", p.URL)
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
	case "NONE":
		if p.Filter.NoColor {
			fmt.Printf("NONE\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnWhite, "NONE", chalk.Reset)
		}
	case "Mild":
		if p.Filter.NoColor {
			fmt.Printf("MILD\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnGreen, "MILD", chalk.Reset)
		}
	case "Moderate":
		if p.Filter.NoColor {
			fmt.Printf("MODERATE\n")
		} else {
			fmt.Printf("%s%s%s\n", blackOnYellow, "MODERATE", chalk.Reset)
		}
	case "Severe":
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
