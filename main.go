// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package main

import (
	"fmt"
	"os"
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
			Usage: "Usage a - ignore others",
		},
		cli.BoolFlag{
			Name:  "title, t",
			Usage: "Usage t",
		},
		cli.BoolFlag{
			Name:  "year, y",
			Usage: "Usage y",
		},
		cli.BoolFlag{
			Name:  "released, r",
			Usage: "Usage r",
		},
		cli.BoolFlag{
			Name:  "duration, z",
			Usage: "Usage z",
		},
		cli.BoolFlag{
			Name:  "summary, k",
			Usage: "Usage k",
		},
		cli.BoolFlag{
			Name:  "directors, d",
			Usage: "Usage d",
		},
		cli.BoolFlag{
			Name:  "writers, w",
			Usage: "Usage w",
		},
		cli.BoolFlag{
			Name:  "stars, p",
			Usage: "Usage p",
		},
		cli.BoolFlag{
			Name:  "genres, g",
			Usage: "Usage s",
		},
		cli.BoolFlag{
			Name:  "tagline, T",
			Usage: "Usage T",
		},
		cli.BoolFlag{
			Name:  "summaries, S",
			Usage: "Usage S",
		},
		cli.BoolFlag{
			Name:  "parental, P",
			Usage: "Usage P",
		},
	}
	app.Action = func(c *cli.Context) error {
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
				filter.Summaries = true
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

		DefaultPrinter = services.New(filter, c.String("url"), "test")
		DefaultPrinter.GetPrint()

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
