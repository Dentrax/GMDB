// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gmdb/models"
	"gmdb/services"

	"github.com/urfave/cli"
)

var DefaultPrinter *services.Printer

func main() {
	app := cli.NewApp()
	app.Name = "GMDB"
	app.Version = "0.0.0"
	app.Compiled = time.Now()
	app.Author = "Furkan Türkal"
	app.Copyright = "(c) 2018 - Dentrax"
	app.Usage = "gmdb"
	app.ArgsUsage = "[args and such]"
	app.HideHelp = false
	app.HideVersion = false
	app.Commands = []cli.Command{
		cli.Command{
			Name:        "search",
			Usage:       "usg",
			UsageText:   "usg text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Search all supported sites (ignore others)",
				},
				cli.BoolFlag{
					Name:  "imdb, i",
					Usage: "Search in IMDB",
				},
				cli.BoolFlag{
					Name:  "rottentomatoes, r",
					Usage: "Search in RottenTomatoes",
				},
			},
			Action: func(c *cli.Context) error {

				searchRequest := new(models.SearchRequest)

				if c.Bool("imdb") {
					searchRequest.ScanIMDB = true
				}

				if c.Bool("rottentomatoes") {
					searchRequest.ScanRT = true
				}

				if len(c.Args()) > 0 {
					searchRequest.Title = strings.Join(c.Args(), "+")

					//Initialize SearchEngine for Query
					engine := *services.NewSearcher(*searchRequest)

					//Get Responses from URLs
					responses := engine.GetSearchResponses()

					//Initialize Printer to print responses
					DefaultPrinter = services.NewPrinter(NoFilter(), *searchRequest)

					//Print the responses
					//Start Index: 0
					//End Index: 10
					//ShowMore Selected: false
					DefaultPrinter.PrintSearchResponses(0, 10, false, responses)

					isShowMoreOptionSelected := false

					fmt.Println("(Enter 'q' to exit.)")

					for {
						//Create a Scanner to get CLI input
						scanner := bufio.NewScanner(os.Stdin)
						fmt.Printf("Please select your choice: ")
						scanner.Scan()
						text := scanner.Text()

						//Exit if given key is 'q'
						if text == "q" {
							os.Exit(0)
						} else {
							//Otherwise, check if it was in integer type
							ok, _ := regexp.MatchString("(^[0-9]*$)", text)
							if !ok {
								fmt.Println("Invalid selection")
								continue
							} else {
								if len(text) == 0 {
									fmt.Println("Invalid selection")
									continue
								}
								selected, _ := strconv.Atoi(text)
								if selected == 0 {
									isShowMoreOptionSelected = true
									//Print the other results that didn't show first
									//Start Index: 10
									//End Index: 250
									//ShowMore Selected: true
									DefaultPrinter.PrintSearchResponses(10, 200, true, responses)
									continue
								} else if (!isShowMoreOptionSelected && selected > 0 && selected <= 10) || (isShowMoreOptionSelected && selected > 0 && selected <= 200) {
									fmt.Println("\nGetting information...\n")

									//Get the movie info from IMDB
									//TODO: Add multiple search engine instead of only IMDB
									movie := engine.GetMovie("IMDB", responses[0].Searches[selected-1])

									//Print the movie info that we had get above
									DefaultPrinter.PrintMovie(*movie)

									if len(movie.Info.URLTrailerIMDB) > 0 {
										fmt.Println()
										fmt.Printf("%s [Y/n]: ", "Do you want to watch Trailer now?")

										scanner.Scan()
										text := scanner.Text()

										if len(text) == 0 || text == "y" || text == "Y" {
											//watch
											fmt.Printf("MPV Player loading...: %s", movie.Info.URLTrailerIMDB)
											out, err := exec.Command("/usr/bin/mpv", movie.Info.URLTrailerIMDB).Output()
											if err != nil {
												os.Exit(2)
											}
											fmt.Printf("%s\n\n", out)
										} else {
											os.Exit(0)
										}
									}

									//Success
									os.Exit(0)
								} else {
									fmt.Println("Invalid selection")
									continue
								}
							}
						}
					}
				} else {
					return cli.NewExitError("No keywords provided", 1)
				}

				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Wrong usage: %q \n", err)
				return err
			},
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url, u",
			Usage: "Usage u",
		},
		cli.StringFlag{
			Name:  "search, s",
			Usage: "Usage s",
		},
		cli.StringFlag{
			Name:  "filename, f",
			Usage: "Usage f",
		},
		cli.StringFlag{
			Name:  "nobanner, q",
			Usage: "Usage q",
		},
		cli.StringFlag{
			Name:  "nocolor, x",
			Usage: "Usage x",
		},

		//For ResultFilter
		cli.BoolFlag{
			Name:  "all, a",
			Usage: "Show all (ignores all other filters)",
		},
		cli.BoolFlag{
			Name:  "title, ft",
			Usage: "Filter the output by Title",
		},
		cli.BoolFlag{
			Name:  "year, fy",
			Usage: "Filter the output by Year",
		},
		cli.BoolFlag{
			Name:  "released, fr",
			Usage: "Filter the output by Released",
		},
		cli.BoolFlag{
			Name:  "duration, fz",
			Usage: "Filter the output by Duration",
		},
		cli.BoolFlag{
			Name:  "summary, fk",
			Usage: "Filter the output by Summary",
		},
		cli.BoolFlag{
			Name:  "directors, fd",
			Usage: "Filter the output by Directors",
		},
		cli.BoolFlag{
			Name:  "writers, fw",
			Usage: "Filter the output by Writers",
		},
		cli.BoolFlag{
			Name:  "stars, fp",
			Usage: "Filter the output by Stars",
		},
		cli.BoolFlag{
			Name:  "genres, fg",
			Usage: "Filter the output by Genres",
		},
		cli.BoolFlag{
			Name:  "tagline, fT",
			Usage: "Filter the output by Tagline",
		},
		cli.BoolFlag{
			Name:  "summaries, fS",
			Usage: "Filter the output by Summaries",
		},
		cli.BoolFlag{
			Name:  "keywords, fK",
			Usage: "Filter the output by Keywords",
		},
		cli.BoolFlag{
			Name:  "parental, fP",
			Usage: "Filter the output by Parental",
		},
	}
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		filter := models.ResultFilter{}
		leastOne := false
		if c.Bool("title") {
			filter.Title = true
			leastOne = true
		}
		if c.Bool("year") {
			filter.Year = true
			leastOne = true
		}
		if c.Bool("released") {
			filter.Released = true
			leastOne = true
		}
		if c.Bool("rating") {
			filter.Rating = true
			leastOne = true
		}
		if c.Bool("duratin") {
			filter.Duration = true
			leastOne = true
		}
		if c.Bool("summary") {
			filter.Summary = true
			leastOne = true
		}
		if c.Bool("directors") {
			filter.Directors = true
			leastOne = true
		}
		if c.Bool("writers") {
			filter.Writers = true
			leastOne = true
		}
		if c.Bool("stars") {
			filter.Stars = true
			leastOne = true
		}
		if c.Bool("genres") {
			filter.Genres = true
			leastOne = true
		}
		if c.Bool("tagline") {
			filter.Tagline = true
			leastOne = true
		}
		if c.Bool("summaries") {
			filter.Summaries = true
			leastOne = true
		}
		if c.Bool("parental") {
			filter.ParentsGuide = true
			leastOne = true
		}
		if c.Bool("keywords") {
			filter.Keywords = true
			leastOne = true
		}
		if c.Bool("all") {
			filter.Title = true
			filter.Year = true
			filter.Released = true
			filter.Rating = true
			filter.Duration = true
			filter.Summary = true
			filter.Directors = true
			filter.Writers = true
			filter.Stars = true
			filter.Genres = true
			filter.Tagline = true
			filter.Summaries = true
			filter.Keywords = true
			filter.ParentsGuide = true
		} else {
			if !leastOne {
				filter.Title = true
				filter.Year = true
				filter.Released = true
				filter.Rating = true
				filter.Duration = true
				filter.Summary = true
				filter.Directors = true
				filter.Writers = true
				filter.Stars = true
				filter.Genres = true
				filter.Tagline = true
				filter.Summaries = false
				filter.Keywords = true
				filter.ParentsGuide = true
			}
		}

		//TODO: Usage explain

		filter.NoBanner = false
		filter.NoColor = false

		if c.Bool("nobanner") {
			filter.NoBanner = true
		}

		if c.Bool("nocolor") {
			filter.NoColor = true
		}

		return nil
	}
	app.CommandNotFound = func(c *cli.Context, command string) {
		fmt.Fprintf(c.App.Writer, "Wrong command: %q \n", command)
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func NoFilter() models.ResultFilter {
	return models.ResultFilter{
		Title:        true,
		Year:         true,
		Released:     true,
		Rating:       true,
		Duration:     true,
		Summary:      true,
		Directors:    true,
		Writers:      true,
		Stars:        true,
		Genres:       true,
		Tagline:      true,
		Summaries:    false,
		Keywords:     true,
		ParentsGuide: true,
	}
}
