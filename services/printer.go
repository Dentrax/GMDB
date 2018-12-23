// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package services

import (
	"fmt"
	"strings"

	"gmdb/models"

	"github.com/ttacon/chalk"
)

type Printer struct {
	Filter  models.ResultFilter
	Request models.SearchRequest
}

func NewPrinter(filter models.ResultFilter, request models.SearchRequest) *Printer {
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

func (p *Printer) PrintSearchResponses(min uint8, max uint8, isMore bool, responses []models.SearchResponse) {
	totalResponse := len(responses)

	if totalResponse > 0 {
		for i := 0; i < totalResponse; i++ {
			currEngine := responses[i].SearchEngine
			totalResult := uint8(len(responses[i].Searches))

			if totalResult > 0 {
				if max > totalResult {
					max = totalResult
				}
				if !isMore {
					lime := chalk.Yellow.NewStyle().
						WithBackground(chalk.Black).
						WithTextStyle(chalk.Bold).
						Style
					fmt.Println()
					fmt.Printf("From%v", " ")
					fmt.Println(lime(currEngine))
				}
				filterCount := responses[i].Searches[min:max]
				indexCounter := min
				for j := range filterCount {
					if isMore {
						fmt.Printf("%2d) ", indexCounter+1)
						p.printInfo(responses[i].Searches[indexCounter].Title, responses[i].Searches[indexCounter].Year)
						indexCounter++
					} else {
						fmt.Printf("%2d) ", j+1)
						p.printInfo(responses[i].Searches[j].Title, responses[i].Searches[j].Year)
					}
				}
				if totalResult > max {
					if totalResult > 10 {
						moreCount := len(responses[i].Searches) - 10
						fmt.Printf("%2d) ", 0)
						p.printInfo(fmt.Sprintf("%v", moreCount), "more...")
					}
				}
			}
		}
	}
}

func (p *Printer) PrintMovie(movie models.Movie) {
	if p.Filter.Title && len(movie.Info.Title) > 0 {
		p.printInfo("Movie Name: ", movie.Info.Title)
	}
	if p.Filter.Year && len(movie.Info.Year) > 0 {
		p.printInfo("Year: ", movie.Info.Year)
	}
	if p.Filter.Released && len(movie.Info.Released) > 0 {
		p.printInfo("Released: ", movie.Info.Released)
	}
	if p.Filter.Rating && len(movie.Info.Rating) > 0 {
		p.printInfo("Rating: ", movie.Info.Rating+" ("+movie.Info.Votes+")")
	}
	if p.Filter.Genres && len(movie.Info.Genres) > 0 {
		p.printInfo("Genres: ", strings.Join(movie.Info.Genres, ", "))
	}
	if p.Filter.Duration && len(movie.Info.Duration) > 0 {
		p.printInfo("Duration: ", movie.Info.Duration)
	}
	if p.Filter.Summary && len(movie.Info.Summary) > 0 {
		p.printInfo("Summary: ", movie.Info.Summary)
	}
	if p.Filter.Directors && len(movie.Info.Credit.Directors) > 0 {
		p.printInfo("Directors: ", strings.Join(movie.Info.Credit.Directors, ", "))
	}
	if p.Filter.Writers && len(movie.Info.Credit.Writers) > 0 {
		p.printInfo("Writers: ", strings.Join(movie.Info.Credit.Writers, ", "))
	}
	if p.Filter.Stars && len(movie.Info.Credit.Stars) > 0 {
		p.printInfo("Stars: ", strings.Join(movie.Info.Credit.Stars, ", "))
	}
	if p.Filter.Tagline && len(movie.TL.Tags) > 0 {
		p.printInfo("Tagline: ", strings.Join(movie.TL.Tags, ", "))
	}
	if p.Filter.Summaries && movie.PS.Total > 0 {
		p.printInfo("Summaries: ", string(movie.PS.Total))
		for i := range movie.PS.Summaries {
			fmt.Println()
			p.printInfo(movie.PS.Summaries[i].Author, movie.PS.Summaries[i].Text)
		}
	}
	if p.Filter.Keywords && len(movie.PK.Keywords) > 0 {
		max := 10
		count := len(movie.PK.Keywords)
		if count > 10 {
			max = 10
		} else {
			max = count
		}
		sums := movie.PK.Keywords[0:max]
		p.printInfo("Keywords: ", strings.Join(sums, ", "))
	}
	if p.Filter.ParentsGuide && movie.PG.TotalVote > 0 {
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

func (p *Printer) PrintBanner() {
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
