// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package learner

import (
	"bufio"
	"errors"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/Dentrax/GMDB/models"
	"github.com/Dentrax/GMDB/pkg/utils"
)

var providers = []struct {
	Provider       string
	FirstLine      string
	LastLine       string
	RegexGroup     string
	RegexMovieName string
	RegexMovieDate string
}{
	{"Netflix", "Title,Date", "", "(\".*?\",\"\\d+\\/\\d+\\/\\d+\")", "([^\"][a-zA-Z0-9_ ]+[^\\:\"])", "(((0)?[0-9])|((1)[0-2]))(\\/)([0-2]?[0-9]|(3)[0-1])(\\/)\\d{2}"},
}

var netflixRegex = regexp.MustCompile("(\".*?\",\"\\d+\\/\\d+\\/\\d+\")")

// CheckFileFormat returns a string and checks the given file format.
func CheckFileFormat(request models.LearnRequest) (string, error) {
	file, err := os.OpenFile(request.Filename, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return "", err
	}

	defer file.Close()

	rd := bufio.NewReader(file)

	filePrefix, _, err := rd.ReadLine()

	for _, provider := range providers {
		if strings.Compare(string(filePrefix), provider.FirstLine) == 0 {
			return provider.Provider, nil
		}
	}

	return "", nil
}

// ScanMovies scans given LearnRequest and returns LearnResponse.
// Scans the providers array item by item and find the correct regex pattern,
// then use it to parse the movie informations.
func ScanMovies(request models.LearnRequest, unique bool) ([]models.LearnResponse, error) {
	responses := new([]models.LearnResponse)

	file, err := os.OpenFile(request.Filename, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	rd := bufio.NewReader(file)

	filePrefix, _, err := rd.ReadLine()

	for _, provider := range providers {
		if strings.Compare(string(filePrefix), provider.FirstLine) == 0 {
			//filePrefix, _, err := rd.ReadLine()

			var rgxGroup = regexp.MustCompile(provider.RegexGroup)
			var rgxMovieName = regexp.MustCompile(provider.RegexMovieName)
			var rgxMovieDate = regexp.MustCompile(provider.RegexMovieDate)
			//TODO: Implement this
			//var rgxMovieDate = regexp.MustCompile(provider.RegexMovieDate)

			//For remove duplicate Titles and TV Series
			askedTitles := []string{}

			for {
				line, err := rd.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					return nil, err
				}
				match := rgxGroup.MatchString(line)
				var result models.LearnResult
				if match {
					name := rgxMovieName.FindString(line)
					date := rgxMovieDate.FindString(line)
					result = models.LearnResult{
						Title:      name,
						IsTVSeries: false,
						WatchDate:  date,
					}
				}
				response := models.LearnResponse{
					Success: match,
					Error:   "",
					Result:  result,
				}

				//TODO: title de : olabilie journer 2: jungle gibi

				if unique {
					if !utils.IsContains(askedTitles, result.Title) {
						askedTitles = append(askedTitles, result.Title)
						*responses = append(*responses, response)
					}
				} else {
					*responses = append(*responses, response)
				}
			}

			return *responses, nil
		}
	}
	return nil, errors.New("Null result")
}
