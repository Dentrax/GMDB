// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package learner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"gmdb/models"
)

var providers = []struct {
	Provider   string
	FirstLine  string
	LastLine   string
	RegexGroup string
	RegexMovie string
}{
	{"Netflix", "Title,Date", "", "(\".*?\",\"\\d+\\/\\d+\\/\\d+\")", "s"},
}

//"The 100: Season 1: Pilot","12/10/18"
var netflixPrefix = "Title,Date"
var netflixRegex = regexp.MustCompile("(\".*?\",\"\\d+\\/\\d+\\/\\d+\")")
var netflixNameRegex = regexp.MustCompile("(\".*?\",\"\\d+\\/\\d+\\/\\d+\")")

func Learn(request *models.LearnRequest) (*models.LearnResponse, error) {
	result := new(models.LearnResponse)

	startPhase(1, "Checking file format...")

	f, err := os.OpenFile(request.Filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		result.Success = false
		return result, err
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	filePrefix, _, err := rd.ReadLine()
	filePrefix2, _, err := rd.ReadLine()

	fmt.Println(string(filePrefix))
	fmt.Println(string(filePrefix2))

	if strings.Compare(string(filePrefix), netflixPrefix) == 0 {
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, err
			}

			match := netflixRegex.MatchString(line)

			if match {
				name := netflixNameRegex.FindString(line)
				fmt.Println(name)
			}

			fmt.Println(match)
		}

		donePhase(1)
		fmt.Println()
		fmt.Println("--> Netflix detected!")
	}

	startPhase(2, "Dedecting dictionary type...")

	donePhase(2)

	startPhase(3, "Learning dictionary...")

	failPhase(3, "Learn faild")

	result.Success = true

	return result, nil
}

func startPhase(phase int, description string) {
	fmt.Printf("\nPHASE[%d/3]: %s", phase, description)
}

func donePhase(phase int) {
	fmt.Printf("[DONE]")
}

func failPhase(phase int, description string) {
	fmt.Printf("[FAIL]")
	fmt.Println(description)
}

func matchNetflix(line string) bool {
	return netflixRegex.MatchString(line)
}
