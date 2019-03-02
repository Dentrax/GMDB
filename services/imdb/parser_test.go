// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"os"
	"strings"
	"testing"

	"gmdb/models"
	"gmdb/services/common"

	. "github.com/smartystreets/goconvey/convey"
)

const PATH_TEST = "../../testdata/IMDB/"

//TODO: Move datas to testdata folder for better test struct
func TestIMDBParseMovieInfo(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.MovieInfo
	}{
		{"deadpool", "home.html",
			&models.MovieInfo{
				Title:      "Deadpool",
				Year:       "2016",
				Rating:     "8.0/10",
				Votes:      "779,471",
				Genres:     []string{"Action", "Adventure", "Comedy"},
				Duration:   "1h 48min",
				Released:   "12 February 2016 (Turkey)",
				IsTVSeries: false,
				Seasons:    "0",
				Episodes:   "0",
				Summary:    "A fast-talking mercenary with a morbid sense of humor is subjected to a rogue experiment that leaves him with accelerated healing powers and a quest for revenge.",
				Credit: models.CreditInfo{
					Directors: []string{"Tim Miller"},
					Writers:   []string{"Rhett Reese", "Paul Wernick"},
					Stars:     []string{"Ryan Reynolds", "Morena Baccarin", "T.J. Miller"},
				},
				Metascore:         "",
				ReviewCountUser:   "",
				ReviewCountCritic: "",
			},
		},
		{"ghostbusters", "home.html",
			&models.MovieInfo{
				Title:      "Ghost Busters",
				Year:       "1984",
				Rating:     "7.8/10",
				Votes:      "314,760",
				Genres:     []string{"Action", "Comedy", "Fantasy"},
				Duration:   "1h 45min",
				Released:   "8 June 1984 (USA)",
				IsTVSeries: false,
				Seasons:    "0",
				Episodes:   "0",
				Summary:    "Three former parapsychology professors set up shop as a unique ghost removal service.",
				Credit: models.CreditInfo{
					Directors: []string{"Ivan Reitman"},
					Writers:   []string{"Dan Aykroyd", "Harold Ramis"},
					Stars:     []string{"Bill Murray", "Dan Aykroyd", "Sigourney Weaver"},
				},
				Metascore:         "",
				ReviewCountUser:   "",
				ReviewCountCritic: "",
			},
		},
		{"replicas", "home.html",
			&models.MovieInfo{
				Title:      "Replicas",
				Year:       "2018",
				Rating:     "6.1/10",
				Votes:      "726",
				Genres:     []string{"Crime", "Mystery", "Sci-Fi"},
				Duration:   "1h 47min",
				Released:   "23 November 2018 (China)",
				IsTVSeries: false,
				Seasons:    "0",
				Episodes:   "0",
				Summary:    "A scientist becomes obsessed with bringing back his family members who died in a traffic accident.",
				Credit: models.CreditInfo{
					Directors: []string{"Jeffrey Nachmanoff"},
					Writers:   []string{"Chad St. John", "Stephen Hamel"},
					Stars:     []string{"Keanu Reeves", "Alice Eve", "Emily Alyn Lind"},
				},
				Metascore:         "",
				ReviewCountUser:   "",
				ReviewCountCritic: "",
			},
		},
		{"The100", "home.html",
			&models.MovieInfo{
				Title:      "The 100",
				Year:       "2014",
				Rating:     "7.8/10",
				Votes:      "169,134",
				Genres:     []string{"Drama", "Mystery", "Sci-Fi"},
				Duration:   "43min",
				Released:   "TV Series (2014– )",
				IsTVSeries: true,
				Seasons:    "6",
				Episodes:   "84",
				Summary:    "Set ninety-seven years after a nuclear war has destroyed civilization, when a spaceship housing humanity's lone survivors sends one hundred juvenile delinquents back to Earth, in hopes of possibly re-populating the planet.",
				Credit: models.CreditInfo{
					Directors: []string{"Jason Rothenberg"},
					Stars:     []string{"Eliza Taylor", "Bob Morley", "Marie Avgeropoulos"},
				},
				Metascore:         "",
				ReviewCountUser:   "",
				ReviewCountCritic: "",
			},
		},
	}

	Convey("[ParseMovieInfo:IMDB]: Parse movie info", t, func() {
		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestHome := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestHome)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			current := ParseMovieInfo(services.GetDocumentFromFile(currentTestHome))
			testName := strings.ToUpper(data.FolderName)

			So(current, ShouldNotBeNil)

			Convey("Testing movie: "+testName, func() {
				So(current, ShouldNotBeNil)

				So(current.Title, ShouldNotBeBlank)
				So(current.Title, ShouldEqual, data.Info.Title)

				So(current.Year, ShouldNotBeBlank)
				So(current.Year, ShouldEqual, data.Info.Year)

				So(current.Rating, ShouldNotBeBlank)
				So(current.Rating, ShouldEqual, data.Info.Rating)

				So(current.Votes, ShouldNotBeBlank)
				So(current.Votes, ShouldEqual, data.Info.Votes)

				So(current.Duration, ShouldNotBeBlank)
				So(current.Duration, ShouldEqual, data.Info.Duration)

				So(current.Released, ShouldNotBeBlank)
				So(current.Released, ShouldEqual, data.Info.Released)

				So(current.Genres, ShouldHaveLength, len(data.Info.Genres))

				So(current.IsTVSeries, ShouldEqual, data.Info.IsTVSeries)

				So(current.Seasons, ShouldEqual, data.Info.Seasons)

				So(current.Episodes, ShouldEqual, data.Info.Episodes)

				So(current.Summary, ShouldNotBeBlank)
				So(current.Summary, ShouldEqual, data.Info.Summary)

				So(current.Credit.Directors, ShouldHaveLength, len(data.Info.Credit.Directors))
				So(current.Credit.Directors, ShouldResemble, data.Info.Credit.Directors)

				So(current.Credit.Writers, ShouldHaveLength, len(data.Info.Credit.Writers))
				So(current.Credit.Writers, ShouldResemble, data.Info.Credit.Writers)

				So(current.Credit.Stars, ShouldHaveLength, len(data.Info.Credit.Stars))
				So(current.Credit.Stars, ShouldResemble, data.Info.Credit.Stars)
			})
		}
	})
}

func TestIMDBParseTagline(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.Tagline
	}{
		{"deadpool", "tagline.html",
			&models.Tagline{
				Tags: []string{
					"With great power comes great irresponsibility.",
					"Witness the beginning of a happy ending",
					"Bad ass. Smart ass. Great ass.",
					"A new class of superhero.",
					"Sit on this",
					"Wait 'til you get a load of me",
					"Size matters",
					"Feel the love Valentine's Day",
					"Justice has a new face"},
			},
		},
		{"ghostbusters", "tagline.html",
			&models.Tagline{
				Tags: []string{
					"Here to save the world again. [re-release]",
					"Let's get slimed one more time. [re-release]",
					"They're here to save the world.",
					"Coming to save the world this summer.",
					"We're ready to believe you.",
					"Who ya gonna call? Ghostbusters!",
					"The supernatural spectacular.",
					"They ain't afraid of no ghost.",
					"The world's most successful comedy."},
			},
		},
		{"replicas", "tagline.html",
			&models.Tagline{
				Tags: nil,
			},
		},
	}

	Convey("[ParseTagLine:IMDB]: Parse lag line", t, func() {
		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestFile := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestFile)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			current := ParseTagline(services.GetDocumentFromFile(currentTestFile))
			testName := strings.ToUpper(data.FolderName)

			So(current, ShouldNotBeNil)

			Convey("Testing movie: "+testName, func() {

				So(current.Tags, ShouldHaveLength, len(data.Info.Tags))

				for i := 0; i < len(current.Tags); i++ {
					So(current.Tags[i], ShouldResemble, data.Info.Tags[i])
				}
			})
		}
	})
}

func TestIMDBParsePlotSummary(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.PlotSummary
	}{
		{"deadpool", "summary.html",
			&models.PlotSummary{
				Summaries: []models.Summary{
					{
						Author: "Bill Smith",
						Text:   "A fast-talking mercenary with a morbid sense of humor is subjected to a rogue experiment that leaves him with accelerated healing powers and a quest for revenge."},
					{
						Author: "20th Century Fox",
						Text:   "This is the origin story of former Special Forces operative turned mercenary Wade Wilson, who after being subjected to a rogue experiment that leaves him with accelerated healing powers, adopts the alter ego Deadpool. Armed with his new abilities and a dark, twisted sense of humor, Deadpool hunts down the man who nearly destroyed his life."},
					{
						Author: "grantss",
						Text:   "Wade Wilson is a small-time mercenary. He meets Vanessa and falls in love. Life is idyllic until one day he is diagnosed with terminal cancer. Things look bleak but a man appears who says he can be cured, through a treatment that gives him superhero powers. After initially turning him down, Wilson agrees, and meets the man behind the project, Ajax. While undergoing the treatment he discovers that it will involve him becoming a mutant, and he will need to undergo several painful tests to discover what his mutant abilities are. Plus, Ajax is a sadist. The treatment results in Wilson getting powers of accelerated healing but also leaves him terribly disfigured. Ajax tries to keep him a prisoner but he escapes. He now has two aims: find Vanessa and make Ajax pay for what he did to him. Killing Ajax may not be as easy it seems, as Ajax is also a mutant and the meddling superhero Colossus and his sidekick Negasonic Teenage Warhead keep getting in the way. Oh, and Wade Wilson is now operating under an alias: Deadpool."},
					{
						Author: "rcs0411@yahoo.com",
						Text:   "Wade Wilson, a soldier turned mercenary who's devoid of emotions meets a woman named Vanessa and he decides to settle down. But when he learns he has cancer, he's worried about dying on her. But a man approaches him and says he can give him a cure and also powers and abilities. he agrees and the man in charge of him is a guy named Francis. Wade has the procedure and he is endowed with healing abilities but is also disfigured. Francis says he can fix his disfigurement but doesn't. Wade gets away from him and tries to go back to Vanessa but doesn't because of how he looks. So he sets out to find Francis by going after everyone who knows him. He makes a suit to hide his disfigurement and adopts the name Deadpool."},
					{
						Author: "MadMovieManiac",
						Text:   "After a fast-talking mercenary is diagnosed with terminal cancer, he finds the possibility of healing in a scientific experience of a covert organization. Recovered, with accelerated healing factor and an unusual sense of humor, he adopts the alter-ego Deadpool to seek revenge against the man who destroyed his life (and his face)."},
					{
						Author: "ahmetkozan",
						Text:   "After being diagnosed with terminal cancer on the same day his girlfriend (Monica Baccarin) accepts his marriage proposal, former Special Ops agent Wade Wilson (Ryan Reynolds) is approached by a shady organization offering a cure to his cancer. Wade accepts and meets a psychopath, Ajax, actually Francis Freeman, (Ed Skrein) who injects Wade with a strange serum, supposedly the cure to his cancer. Francis, along with his assistant Angel Dust (Gina Cariano) proceed to put Wade through unbearable torture, leaving him with horrible burn-like scars in his entire body and an accelerated healing factor. But Wade doesn't forget, as he's pissed and wants his face back to normal. Armed with two katanas, pistols, and a red suit and mask, Wade becomes Deadpool, and makes it his mission to hunt Francis down and get revenge."},
				},
				Total: uint(6),
			},
		},
	}

	Convey("[ParsePlotSummary:IMDB]: Parse plot summary", t, func() {
		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestFile := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestFile)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			current := ParsePlotSummary(services.GetDocumentFromFile(currentTestFile))
			testName := strings.ToUpper(data.FolderName)

			So(current, ShouldNotBeNil)

			Convey("Testing movie: "+testName, func() {
				So(current.Total-1, ShouldEqual, data.Info.Total)
				So(data.Info.Total, ShouldEqual, len(data.Info.Summaries))

				for i := uint(0); i < data.Info.Total; i++ {
					So(current.Summaries[i].Author, ShouldNotBeBlank)
					So(current.Summaries[i].Author, ShouldEqual, data.Info.Summaries[i].Author)

					So(current.Summaries[i].Text, ShouldNotBeBlank)
					So(current.Summaries[i].Text, ShouldEqual, data.Info.Summaries[i].Text)
				}
			})
		}
	})
}

func TestIMDBParsePlotKeywords(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.PlotKeywords
	}{
		{"deadpool", "keywords.html",
			&models.PlotKeywords{
				Keywords: []string{
					"breaking the fourth wall",
					"anti hero",
					"mercenary",
					"revenge",
					"mutant",
					"sex scene",
					"scene after end credits",
					"cancer",
					"character name in title",
					"self healing"},
				Total: 524,
			},
		},
		{"ghostbusters", "keywords.html",
			&models.PlotKeywords{
				Keywords: []string{
					"ghost",
					"ghostbuster",
					"university",
					"paranormal",
					"scientist",
					"paranormal phenomenon",
					"paranormal investigation team",
					"ghostbusting",
					"supernatural being",
					"haunting"},
				Total: 163,
			},
		},
		{"replicas", "keywords.html",
			&models.PlotKeywords{
				Keywords: nil,
				Total:    0,
			},
		},
	}

	Convey("[ParsePlotKeywords:IMDB]: Parse plot keywords", t, func() {
		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestFile := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestFile)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			current := ParsePlotKeywords(services.GetDocumentFromFile(currentTestFile))
			testName := strings.ToUpper(data.FolderName)

			So(current, ShouldNotBeNil)

			Convey("Testing movie: "+testName, func() {
				So(current.Total, ShouldEqual, data.Info.Total)

				if current.Total > 0 {
					for i := 0; i < 10; i++ {
						So(current.Keywords[i], ShouldResemble, data.Info.Keywords[i])
					}
				}
			})

		}
	})
}

func TestIMDBParseParentsGuide(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.ParentsGuide
	}{
		{"deadpool", "parental.html",
			&models.ParentsGuide{
				Nudity:      models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "SEVERE"},
				Violence:    models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "SEVERE"},
				Profanity:   models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "SEVERE"},
				Alcohol:     models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MILD"},
				Frightening: models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MODERATE"},
			},
		},
		{"ghostbusters", "parental.html",
			&models.ParentsGuide{
				Nudity:      models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MILD"},
				Violence:    models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MILD"},
				Profanity:   models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MODERATE"},
				Alcohol:     models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MILD"},
				Frightening: models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "MILD"},
			},
		},
		{"replicas", "parental.html",
			&models.ParentsGuide{
				Nudity:      models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "EMPTY"},
				Violence:    models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "EMPTY"},
				Profanity:   models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "EMPTY"},
				Alcohol:     models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "EMPTY"},
				Frightening: models.SeverityRate{NONE: 0, MILD: 0, MODERATE: 0, SEVERE: 0, TotalRate: 0, FinalRate: "EMPTY"},
			},
		},
	}

	Convey("[ParseParentsGuide:IMDB]: Parse parents guide", t, func() {
		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestFile := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestFile)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			current := ParseParentsGuide(services.GetDocumentFromFile(currentTestFile))
			testName := strings.ToUpper(data.FolderName)

			So(current, ShouldNotBeNil)

			Convey("Testing movie: "+testName, func() {
				So(current.Nudity.FinalRate, ShouldEqual, data.Info.Nudity.FinalRate)
				So(current.Violence.FinalRate, ShouldEqual, data.Info.Violence.FinalRate)
				So(current.Profanity.FinalRate, ShouldEqual, data.Info.Profanity.FinalRate)
				So(current.Alcohol.FinalRate, ShouldEqual, data.Info.Alcohol.FinalRate)
				So(current.Frightening.FinalRate, ShouldEqual, data.Info.Frightening.FinalRate)
			})
		}
	})
}
