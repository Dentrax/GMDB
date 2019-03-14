// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package learner

import (
	"os"
	"strings"
	"testing"

	"github.com/Dentrax/GMDB/models"

	. "github.com/smartystreets/goconvey/convey"
)

const PATH_TEST = "../testdata/NETFLIX/"

func TestNetflixParseFile(t *testing.T) {
	currentTestFile := PATH_TEST + "NetflixViewingHistory.csv"

	file, err := os.Open(currentTestFile)

	Convey("[Learner::NETFLIX]: Testing file parser", t, func() {
		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		defer file.Close()

		request := new(models.LearnRequest)
		request.Filename = currentTestFile

		check, err := CheckFileFormat(*request)

		So(check, ShouldNotBeNil)
		So(err, ShouldBeNil)

		So(check, ShouldNotBeNil)

		So(check, ShouldEqual, "Netflix")
	})
}

func TestNetflixParseMovies(t *testing.T) {
	var testDatas = []struct {
		Title     string
		WatchDate string
	}{
		{"The 100", "12/10/18"},
		{"Black Mirror", "12/7/18"},
		{"The Rain", "12/7/18"},
		{"ARQ", "11/28/18"},
		{"Journey 2", "11/28/18"},
		{"Stranger Things", "11/26/18"},
		{"Stranger Things", "11/25/18"},
		{"Gravity", "11/25/18"},
	}
	//TODO: Add support for Season, Episode and Episode Name like;
	//{"The 100", "Season 1", "Pilot", "12/10/18"}

	Convey("[Learner:NETFLIX]: Parse movie infos", t, func() {
		currentTestFile := PATH_TEST + "NetflixViewingHistory.csv"

		file, err := os.Open(currentTestFile)

		So(file, ShouldNotBeNil)
		So(err, ShouldBeNil)

		defer file.Close()

		request := new(models.LearnRequest)
		request.Filename = currentTestFile

		movies, err := ScanMovies(*request, false)

		So(movies, ShouldNotBeNil)
		So(movies, ShouldHaveLength, len(testDatas))

		for i, data := range testDatas {
			testName := strings.ToUpper(data.Title)

			Convey("Testing Netflix movie: "+testName+" ("+data.WatchDate+")", func() {
				So(movies[i].Result.Title, ShouldNotBeBlank)
				So(data.Title, ShouldEqual, movies[i].Result.Title)

				So(movies[i].Result.WatchDate, ShouldNotBeBlank)
				So(data.WatchDate, ShouldEqual, movies[i].Result.WatchDate)
			})
		}

	})

}
