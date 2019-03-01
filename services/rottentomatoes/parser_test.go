// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package rottentomatoes

import (
	"os"
	"reflect"
	"strings"
	"testing"

	"gmdb/models"
	"gmdb/services/common"
)

const PATH_TEST = "../../testdata/RT/"

func TestRTParseMovieInfo(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Info       *models.MovieInfo
	}{
		{"deadpool", "home.html",
			&models.MovieInfo{
				Title:    "Deadpool",
				Year:     "2016",
				Rating:   "90%",
				Votes:    "186,305",
				Genres:   []string{"Action & Adventure", "Comedy", "Science Fiction & Fantasy"},
				Duration: "103 minutes",
				Released: "Feb 12, 2016",
				Summary:  "Based upon Marvel Comics' most unconventional anti-hero, DEADPOOL tells the origin story of former Special Forces operative turned mercenary Wade Wilson, who after being subjected to a rogue experiment that leaves him with accelerated healing powers, adopts the alter ego Deadpool. Armed with his new abilities and a dark, twisted sense of humor, Deadpool hunts down the man who nearly destroyed his life. (C) Fox",
				Credit: models.CreditInfo{
					Directors: []string{"Tim Miller"},
					Writers:   []string{"Rhett Reese", "Paul Wernick"},
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
