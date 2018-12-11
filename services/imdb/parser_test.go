// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"gmdb/models"
	"gmdb/services/common"
)

const PATH_TEST = "../../testdata/"

func TestParseMovieInfo(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.MovieInfo
	}{
		{"deadpool", "home.html",
			&models.MovieInfo{
				Title:    "Deadpool (2016)",
				Year:     "(2016)",
				Rating:   "8.0/10",
				Votes:    "779,471",
				Genres:   []string{"Action", "Adventure", "Comedy"},
				Duration: "1h 48min",
				Released: "12 February 2016 (Turkey)",
				Summary:  "A fast-talking mercenary with a morbid sense of humor is subjected to a rogue experiment that leaves him with accelerated healing powers and a quest for revenge.",
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
				Title:    "Ghost Busters (1984)",
				Year:     "(1984)",
				Rating:   "7.8/10",
				Votes:    "314,760",
				Genres:   []string{"Action", "Comedy", "Fantasy"},
				Duration: "1h 45min",
				Released: "8 June 1984 (USA)",
				Summary:  "Three former parapsychology professors set up shop as a unique ghost removal service.",
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
				Title:    "Replicas (2018)",
				Year:     "(2018)",
				Rating:   "6.1/10",
				Votes:    "726",
				Genres:   []string{"Crime", "Mystery", "Sci-Fi"},
				Duration: "1h 47min",
				Released: "23 November 2018 (China)",
				Summary:  "A scientist becomes obsessed with bringing back his family members who died in a traffic accident.",
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
	}

	for _, data := range testDatas {
		currentTestPath := PATH_TEST + data.FolderName
		currentTestHome := currentTestPath + "/" + data.File

		fileHome, err1 := os.Open(currentTestHome)
		if err1 != nil {
			t.Errorf("[Test::ParseMovieInfo]: File %s open error: %s", currentTestHome, err1)
		}
		defer fileHome.Close()

		info := ParseMovieInfo(services.GetDocumentFromFile(currentTestHome))

		testName := strings.ToUpper(data.FolderName)

		expectedTitle := data.Info.Title
		actualTitle := info.Title
		if !reflect.DeepEqual(expectedTitle, actualTitle) {
			t.Errorf("[Test::%s::MI::Title]: Expected: '%s', Actual: '%s'", testName, expectedTitle, actualTitle)
		}

		expectedYear := data.Info.Year
		actualYear := info.Year
		if !reflect.DeepEqual(expectedYear, actualYear) {
			t.Errorf("[Test::%s::MI::Year]: Expected: '%s', Actual: '%s'", testName, expectedYear, actualYear)
		}

		expectedRating := data.Info.Rating
		actualRating := info.Rating
		if !reflect.DeepEqual(expectedRating, actualRating) {
			t.Errorf("[Test::%s::MI::Rating]: Expected: '%s', Actual: '%s'", testName, expectedRating, actualRating)
		}

		expectedVotes := data.Info.Votes
		actualVotes := info.Votes
		if !reflect.DeepEqual(expectedVotes, actualVotes) {
			t.Errorf("[Test::%s::MI::Votes]: Expected: '%s', Actual: '%s'", testName, expectedVotes, actualVotes)
		}

		expectedDuration := data.Info.Duration
		actualDuration := info.Duration
		if !reflect.DeepEqual(expectedDuration, actualDuration) {
			t.Errorf("[Test::%s::MI::Duration]: Expected: '%s', Actual: '%s'", testName, expectedDuration, actualDuration)
		}

		expectedReleased := data.Info.Released
		actualReleased := info.Released
		if !reflect.DeepEqual(expectedReleased, actualReleased) {
			t.Errorf("[Test::%s::MI::Released]: Expected: '%s', Actual: '%s'", testName, expectedReleased, actualReleased)
		}

		expectedGenres := data.Info.Genres
		actualGenres := info.Genres
		if !reflect.DeepEqual(expectedGenres, actualGenres) {
			t.Errorf("[Test::%s::MI::Genres]: Expected: '%s', Actual: '%s'", testName, expectedGenres, actualGenres)
		}

		expectedSummary := data.Info.Summary
		actualSummary := info.Summary
		if !reflect.DeepEqual(expectedSummary, actualSummary) {
			t.Errorf("[Test::%s::MI::Summary]: Expected: '%s', Actual: '%s'", testName, expectedSummary, actualSummary)
		}

		expectedCreditDirectors := data.Info.Credit.Directors
		actualCreditDirectors := info.Credit.Directors
		if !reflect.DeepEqual(expectedCreditDirectors, actualCreditDirectors) {
			t.Errorf("[Test::%s::MI::Directors]: Expected: '%s', Actual: '%s'", testName, expectedCreditDirectors, actualCreditDirectors)
		}
		expectedCreditWriters := data.Info.Credit.Writers
		actualCreditWriters := info.Credit.Writers
		if !reflect.DeepEqual(expectedCreditWriters, actualCreditWriters) {
			t.Errorf("[Test::%s::MI::Writers]: Expected: '%s', Actual: '%s'", testName, expectedCreditWriters, actualCreditWriters)
		}
		expectedCreditStars := data.Info.Credit.Stars
		actualCreditStars := info.Credit.Stars
		if !reflect.DeepEqual(expectedCreditStars, actualCreditStars) {
			t.Errorf("[Test::%s::MI::Stars]: Expected: '%s', Actual: '%s'", testName, expectedCreditStars, actualCreditStars)
		}
	}
}

func TestParseTagline(t *testing.T) {
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
	for _, data := range testDatas {
		currentTestPath := PATH_TEST + data.FolderName
		currentTestFile := currentTestPath + "/" + data.File

		file, err := os.Open(currentTestFile)
		if err != nil {
			t.Errorf("[Test::Tagline]: File %s open error: %s", currentTestFile, err)
		}
		defer file.Close()

		info := ParseTagline(services.GetDocumentFromFile(currentTestFile))

		testName := strings.ToUpper(data.FolderName)

		expectedTagsCount := len(data.Info.Tags)
		actualTagsCount := len(info.Tags)
		if !reflect.DeepEqual(expectedTagsCount, actualTagsCount) {
			t.Errorf("[Test::%s::TL::Count]: Expected: '%d', Actual: '%d'", testName, expectedTagsCount, actualTagsCount)
		}

		for i := 0; i < expectedTagsCount; i++ {
			expectedTag := data.Info.Tags[i]
			actualTag := info.Tags[i]

			if !reflect.DeepEqual(expectedTag, actualTag) {
				t.Errorf("[Test::%s::TL::Tag]: Expected: '%s', Actual: '%s'", testName, expectedTag, actualTag)
			}
		}
	}
}

func TestParsePlotSummary(t *testing.T) {
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
	for _, data := range testDatas {
		currentTestPath := PATH_TEST + data.FolderName
		currentTestFile := currentTestPath + "/" + data.File

		file, err := os.Open(currentTestFile)
		if err != nil {
			t.Errorf("[Test::PlotSummary]: File %s open error: %s", currentTestFile, err)
		}
		defer file.Close()

		info := ParsePlotSummary(services.GetDocumentFromFile(currentTestFile))

		testName := strings.ToUpper(data.FolderName)

		expectedTotal := data.Info.Total
		actualTotal := info.Total - 1
		if !reflect.DeepEqual(expectedTotal, actualTotal) {
			t.Errorf("[Test::%s::PS::Total]: Expected: '%d', Actual: '%d'", testName, expectedTotal, actualTotal)
		}

		if data.Info.Total != uint(len(data.Info.Summaries)) {
			t.Errorf("[Test::%s::PS::MatchArrayCount]: Expected: '%d', Actual: '%d'", testName, actualTotal, len(data.Info.Summaries))
		}

		for i := uint(0); i < expectedTotal; i++ {
			expectedAuthor := data.Info.Summaries[i].Author
			actualAuthor := info.Summaries[i].Author
			if !reflect.DeepEqual(expectedAuthor, actualAuthor) {
				t.Errorf("[Test::%s::PS::Author]: Expected: '%s', Actual: '%s'", testName, expectedAuthor, actualAuthor)
			}

			expectedText := data.Info.Summaries[i].Text
			actualText := info.Summaries[i].Text
			if !reflect.DeepEqual(expectedText, actualText) {
				t.Errorf("[Test::%s::PS::Text]: Expected: '%s', Actual: '%s'", testName, expectedText, actualText)
			}

		}
	}
}

func TestParsePlotKeywords(t *testing.T) {
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

	for _, data := range testDatas {
		currentTestPath := PATH_TEST + data.FolderName
		currentTestFile := currentTestPath + "/" + data.File

		file, err := os.Open(currentTestFile)
		if err != nil {
			t.Errorf("[Test::PlotKeywords]: File %s open error: %s", currentTestFile, err)
		}
		defer file.Close()

		info := ParsePlotKeywords(services.GetDocumentFromFile(currentTestFile))

		testName := strings.ToUpper(data.FolderName)

		expectedTotal := data.Info.Total
		actualTotal := info.Total
		if !reflect.DeepEqual(expectedTotal, actualTotal) {
			t.Errorf("[Test::%s::PK::Total]: Expected: '%d', Actual: '%d'", testName, expectedTotal, actualTotal)
		}

		var Top10 = 10

		if info.Total != 0 {
			for i := 0; i < Top10; i++ {
				expectedKeyword := data.Info.Keywords[i]
				actualKeyword := info.Keywords[i]
				if !reflect.DeepEqual(expectedKeyword, actualKeyword) {
					t.Errorf("[Test::%s::PK::Keyword]: Expected: '%s', Actual: '%s'", testName, expectedKeyword, actualKeyword)
				}
			}
		}
	}
}

func TestParseParentsGuide(t *testing.T) {
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
	for _, data := range testDatas {
		currentTestPath := PATH_TEST + data.FolderName
		currentTestFile := currentTestPath + "/" + data.File

		file, err := os.Open(currentTestFile)
		if err != nil {
			t.Errorf("[Test::ParentsGuide]: File %s open error: %s", currentTestFile, err)
		}
		defer file.Close()

		info := ParseParentsGuide(services.GetDocumentFromFile(currentTestFile))

		testName := strings.ToUpper(data.FolderName)

		expectedNudityFinal := data.Info.Nudity.FinalRate
		actualNudityFinal := info.Nudity.FinalRate
		if !reflect.DeepEqual(expectedNudityFinal, actualNudityFinal) {
			t.Errorf("[Test::%s::PG::NudityFinalRate]: Expected: '%s', Actual: '%s'", testName, expectedNudityFinal, actualNudityFinal)
		}

		expectedViolenceFinal := data.Info.Violence.FinalRate
		actualViolenceFinal := info.Violence.FinalRate
		if !reflect.DeepEqual(expectedViolenceFinal, actualViolenceFinal) {
			t.Errorf("[Test::%s::PG::ViolenceFinalRate]: Expected: '%s', Actual: '%s'", testName, expectedViolenceFinal, actualViolenceFinal)
		}

		expectedProfanityFinal := data.Info.Profanity.FinalRate
		actualProfanityFinal := info.Profanity.FinalRate
		if !reflect.DeepEqual(expectedProfanityFinal, actualProfanityFinal) {
			t.Errorf("[Test::%s::PG::ProfanityFinalRate]: Expected: '%s', Actual: '%s'", testName, expectedProfanityFinal, actualProfanityFinal)
		}

		expectedAlcoholFinal := data.Info.Alcohol.FinalRate
		actualAlcoholFinal := info.Alcohol.FinalRate
		if !reflect.DeepEqual(expectedAlcoholFinal, actualAlcoholFinal) {
			t.Errorf("[Test::%s::PG::AlcoholFinalRate]: Expected: '%s', Actual: '%s'", testName, expectedAlcoholFinal, actualAlcoholFinal)
		}

		expectedFrighteningFinal := data.Info.Frightening.FinalRate
		actualFrighteningFinal := info.Frightening.FinalRate
		if !reflect.DeepEqual(expectedFrighteningFinal, actualFrighteningFinal) {
			t.Errorf("[Test::%s::PG::FrighteningFinalRate]: Expected: '%s', Actual: '%s'", testName, expectedFrighteningFinal, actualFrighteningFinal)
		}
	}
}
