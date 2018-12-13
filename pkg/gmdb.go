package gmdb

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gmdb/models"
	"gmdb/services"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli"
)

var DefaultSpinner *spinner.Spinner

var DefaultPrinter *services.Printer

func StartSpinner(id int, speed int) {
	DefaultSpinner = spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	DefaultSpinner.UpdateCharSet(spinner.CharSets[id])
	DefaultSpinner.UpdateSpeed(time.Duration(speed) * time.Millisecond)
	DefaultSpinner.Start()
}

func StopSpinner() {
	DefaultSpinner.Stop()
}

func WaitInputStringFromCLI() string {
	//Create a Scanner to get CLI input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	response := scanner.Text()
	return strings.TrimSpace(response)
}

func WaitInputIntFromCLI() int {
	strInput := WaitInputStringFromCLI()

	if strings.EqualFold(strInput, "q") {
		os.Exit(0)
	}

	isIntInput, _ := regexp.MatchString("(^[0-9]*$)", strInput)
	if !isIntInput || len(strInput) == 0 {
		fmt.Println("Invalid input")
		os.Exit(1)
	}

	intInput, err := strconv.Atoi(strInput)
	if err != nil {
		log.Fatal(err)
	}

	return intInput
}

func AskYNQuestion(question string, tries int, defaultInput bool) bool {
	for ; tries > 0; tries-- {
		if defaultInput {
			fmt.Printf(":: %s [Y/n]", question)
		} else {
			fmt.Printf(":: %s [y/N]", question)
		}

		response := WaitInputStringFromCLI()
		response = strings.ToLower(response)

		if len(response) == 0 && defaultInput {
			return true
		}

		return (response == "y" || response == "yes" || response == "1")
	}
	return false
}

func HandleSearchTitleRequest(c *cli.Context) {
	searchRequest := new(models.SearchRequest)
	searchRequest.Title = strings.Join(c.Args(), "+")

	if c.Bool("imdb") {
		searchRequest.ScanIMDB = true
	}

	if c.Bool("rottentomatoes") {
		searchRequest.ScanRT = true
	}

	StartSpinner(4, 100)

	//Initialize SearchEngine for Query
	engine := *services.NewSearcher(*searchRequest)

	//Get Responses from URLs
	responses := engine.GetSearchResponses()

	StopSpinner()

	//Initialize Printer to print responses
	DefaultPrinter = services.NewPrinter(UseNoResultFilter(), *searchRequest)

	//Print the responses
	//Start Index: 0
	//End Index: 10
	//ShowMore Selected: false
	DefaultPrinter.PrintSearchResponses(0, 10, false, responses)

	isShowMoreOptionSelected := false

	fmt.Println("(Enter 'q' to exit.)")

	for {
		fmt.Printf("Please select your choice: ")

		choice := WaitInputIntFromCLI()

		if !isShowMoreOptionSelected && choice == 0 {
			isShowMoreOptionSelected = true

			//Print the other results that didn't show first
			//Start Index: 10
			//End Index: 250
			//ShowMore Selected: true
			DefaultPrinter.PrintSearchResponses(10, 200, isShowMoreOptionSelected, responses)
			continue

		} else if (!isShowMoreOptionSelected && choice > 0 && choice <= 10) ||
			(isShowMoreOptionSelected && choice > 0 && choice <= 200) {

			fmt.Println("\nGetting information...\n")

			StartSpinner(37, 150)

			//Get the movie info from IMDB
			//TODO: Add multiple search engine instead of only IMDB
			movie := engine.GetMovie("IMDB", responses[0].Searches[choice-1])

			StopSpinner()

			//Print the movie info that we had get above
			DefaultPrinter.PrintMovie(*movie)

			if len(movie.Info.URLTrailerIMDB) > 0 {
				HandleWatchMovie(movie.Info.URLTrailerIMDB)
			}

			break
		} else {
			os.Exit(0)
		}
	}

}

func HandleWatchMovie(url string) {
	fmt.Println()

	watch := AskYNQuestion("Do you want to watch Trailer?", 3, true)

	if watch {
		fmt.Printf("MPV Player loading...: %s", url)

		cmd := exec.Command("/usr/bin/mpv", url)

		if err := cmd.Start(); err != nil {
			fmt.Printf("Failed to start cmd: %v", err)
			os.Exit(2)
		}

		if err := cmd.Wait(); err != nil {
			fmt.Printf("Cmd returned error: %v", err)
			os.Exit(2)
		}
	}
	os.Exit(0)
}

func UseNoResultFilter() models.ResultFilter {
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
