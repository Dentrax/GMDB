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

	"github.com/Dentrax/GMDB/models"
	"github.com/Dentrax/GMDB/pkg"
	"github.com/Dentrax/GMDB/pkg/config"
	"github.com/Dentrax/GMDB/services"

	"github.com/urfave/cli"
)

var DefaultPrinter *services.Printer

func main() {
	app := cli.NewApp()
	app.Name = "GMDB"
	app.Version = "Alpha v0.0.0"
	app.Compiled = time.Now()
	app.Author = "Furkan Türkal"
	app.Copyright = "(c) 2019 - Dentrax"
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
			},
			Action: func(c *cli.Context) error {
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
				if len(c.Args()) > 0 {
					config, err := config.LoadConfig()
					if err != nil {
						return cli.NewExitError("Failed to load config", 1)
					}
					gmdb := &gmdb.App{}
					gmdb.Initialize(config)
					gmdb.HandleSearchTitleRequest(c)
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
		cli.Command{
			Name:        "learn",
			Usage:       "usg",
			UsageText:   "usg text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Action: func(c *cli.Context) error {
				if len(c.Args()) > 0 {
					config, err := config.LoadConfig()
					if err != nil {
						return cli.NewExitError("Failed to load config", 1)
					}
					gmdb := &gmdb.App{}
					gmdb.Initialize(config)
					gmdb.HandleLearnRequest(c)
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
		cli.Command{
			Name:        "history",
			Usage:       "hst",
			UsageText:   "hst text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Get all histories both of search and watch",
				},
				cli.BoolFlag{
					Name:  "search, s",
					Usage: "Get search histories",
				},
				cli.BoolFlag{
					Name:  "watch, w",
					Usage: "Get watch histories",
				},
			},
			Action: func(c *cli.Context) error {
				config, err := config.LoadConfig()
				if err != nil {
					return cli.NewExitError("Failed to load config", 1)
				}
				gmdb := &gmdb.App{}
				gmdb.Initialize(config)
				gmdb.HandleHistoryRequest(c)
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Wrong usage: %q \n", err)
				return err
			},
		},
		cli.Command{
			Name:        "list",
			Usage:       "lst",
			UsageText:   "lst text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Get all lists both of Watch Later(s) and Like(s)",
				},
				cli.BoolFlag{
					Name:  "watch, w",
					Usage: "Get Watch Later list",
				},
				cli.BoolFlag{
					Name:  "like, l",
					Usage: "Get Movie Likes list",
				},
			},
			Action: func(c *cli.Context) error {
				config, err := config.LoadConfig()
				if err != nil {
					return cli.NewExitError("Failed to load config", 1)
				}
				gmdb := &gmdb.App{}
				gmdb.Initialize(config)
				gmdb.HandleMyListRequest(c)
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Wrong usage: %q \n", err)
				return err
			},
		},
		cli.Command{
			Name:        "note",
			Usage:       "nt",
			UsageText:   "nt text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Action: func(c *cli.Context) error {
				config, err := config.LoadConfig()
				if err != nil {
					return cli.NewExitError("Failed to load config", 1)
				}
				gmdb := &gmdb.App{}
				gmdb.Initialize(config)
				gmdb.HandleNoteRequest(c)
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Wrong usage: %q \n", err)
				return err
			},
		},
		cli.Command{
			Name:        "torrent",
			Usage:       "trrnt",
			UsageText:   "trrnt text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Action: func(c *cli.Context) error {
				config, err := config.LoadConfig()
				if err != nil {
					return cli.NewExitError("Failed to load config", 1)
				}
				gmdb := &gmdb.App{}
				gmdb.Initialize(config)
				gmdb.HandleTorrentRequest(c, nil)
				return nil
			},
			OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
				fmt.Fprintf(c.App.Writer, "Wrong usage: %q \n", err)
				return err
			},
		},
		cli.Command{
			Name:        "update",
			Usage:       "updt",
			UsageText:   "updt text",
			Description: "desc",
			ArgsUsage:   "[arg]",
			Action: func(c *cli.Context) error {
				config, err := config.LoadConfig()
				if err != nil {
					return cli.NewExitError("Failed to load config", 1)
				}
				gmdb := &gmdb.App{}
				gmdb.Initialize(config)
				gmdb.HandleUpdateRequest(c)
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
			Name:  "nobanner, q",
			Usage: "Usage q",
		},
		cli.StringFlag{
			Name:  "nocolor, x",
			Usage: "Usage x",
		},
	}
	app.Action = func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			cli.ShowAppHelpAndExit(c, 1)
		}

		filter := models.ResultFilter{}

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
