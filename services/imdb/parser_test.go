// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package imdb

import (
	"os"
	"testing"

	"gmdb/models"
)

type GTestHomeInfo struct {
	File string
	Info *models.MovieInfo
}

type GTestTaglineInfo struct {
	File string
	Info *models.Tagline
}

type GTestPlotSummaryInfo struct {
	File string
	Info *models.PlotSummary
}

type GTestPlotKeywordsInfo struct {
	File string
	Info *models.PlotKeywords
}

type GTestParentsGuideInfo struct {
	File string
	Info *models.ParentsGuide
}

func TestParseMovieInfo(t *testing.T) {
	var testPath = "../../testdata/"

	var testDatas = []struct {
		FolderName   string
		Home         *GTestHomeInfo
		Tagline      *GTestTaglineInfo
		PlotSummary  *GTestPlotSummaryInfo
		PlotKeywords *GTestPlotKeywordsInfo
		ParentsGuide *GTestParentsGuideInfo
	}{
		{"deadpool",
			&GTestHomeInfo{
				File: "home.html",
				Info: &models.MovieInfo{
					Title:    "Test",
					Year:     "Test",
					Rating:   "Test",
					Votes:    "Test",
					Genres:   nil,
					Duration: "Test",
					Released: "Test",
					Summary:  "Test",
				},
			},
			&GTestTaglineInfo{
				File: "tagline.html",
				Info: &models.Tagline{
					Tags: nil,
				},
			},
			&GTestPlotSummaryInfo{
				File: "summary.html",
				Info: &models.PlotSummary{
					Summaries: nil,
					Total:     uint(0),
				},
			},
			&GTestPlotKeywordsInfo{
				File: "keywords.html",
				Info: &models.PlotKeywords{
					Keywords: nil,
					Total:    0,
				},
			},
			&GTestParentsGuideInfo{
				File: "parental.html",
				Info: &models.ParentsGuide{
					Nudity: models.SeverityRate{
						NONE:      0,
						MILD:      0,
						MODERATE:  0,
						SEVERE:    0,
						TotalRate: 0,
						FinalRate: "Test",
					},
					Violence: models.SeverityRate{
						NONE:      0,
						MILD:      0,
						MODERATE:  0,
						SEVERE:    0,
						TotalRate: 0,
						FinalRate: "Test",
					},
					Profanity: models.SeverityRate{
						NONE:      0,
						MILD:      0,
						MODERATE:  0,
						SEVERE:    0,
						TotalRate: 0,
						FinalRate: "Test",
					},
					Alcohol: models.SeverityRate{
						NONE:      0,
						MILD:      0,
						MODERATE:  0,
						SEVERE:    0,
						TotalRate: 0,
						FinalRate: "Test",
					},
					Frightening: models.SeverityRate{
						NONE:      0,
						MILD:      0,
						MODERATE:  0,
						SEVERE:    0,
						TotalRate: 0,
						FinalRate: "Test",
					},
				},
			},
		},
	}

	for _, data := range testDatas {
		currentTestPath := testPath + data.FolderName
		currentTestHome := currentTestPath + "/" + data.Home.File

		fileHome, err1 := os.Open(currentTestHome)
		if err1 != nil {
			t.Errorf("[Test::ParseMovieInfo]: File %s open error: %s", currentTestHome, err1)
		}
		defer fileHome.Close()

		info := ParseMovieInfo(GetDocumentFromFile(currentTestHome))

		//TODO: Test yaz
		t.Errorf("ok %s", info.Title)

	}
}

func TestParseTagline(t *testing.T) {

}

func TestParsePlotSummary(t *testing.T) {

}

func TestParsePlotKeywords(t *testing.T) {

}

func TestParseParentsGuide(t *testing.T) {

}
