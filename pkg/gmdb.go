// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package gmdb

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gmdb/learner"
	"gmdb/models"
	"gmdb/pkg/cache"
	"gmdb/pkg/config"
	"gmdb/services"
	"gmdb/store"
	"gmdb/store/database"
	"gmdb/store/movie"

	"github.com/briandowns/spinner"
	"github.com/urfave/cli"
)

var DefaultSpinner *spinner.Spinner

var DefaultPrinter *services.Printer

var Config *config.Config

type App struct {
	Config *config.Config
	DB     *database.DB
	Store  store.Stores
	CTX    context.Context
}

func (a *App) Initialize(config *config.Config) {
	a.Config = config
	conn, err := database.Connect("sqlite3", "db")

	if err != nil {
		fmt.Println("Database connection failed!")
	}

	a.DB = conn
	a.Store.Movies = movie.New(conn)

	//dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 1234, "dentrax", "1", "gmdb")

	//a.DB = database.New("sqlite3", "db")
}

func StartSpinner(id int, speed int) {
	DefaultSpinner = spinner.New(spinner.CharSets[4], 100*time.Millisecond)
	DefaultSpinner.UpdateCharSet(spinner.CharSets[id])
	DefaultSpinner.UpdateSpeed(time.Duration(speed) * time.Millisecond)
	DefaultSpinner.Start()
}

func StopSpinner() {
	DefaultSpinner.Stop()
}

//Wait a string input from User
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

//Wait an int input from User
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

//Write an Ask(Y/N) question and wait an input
func AskYNQuestion(question string, tries int, defaultInput bool) bool {
	for ; tries > 0; tries-- {
		if defaultInput {
			fmt.Printf(":: %s [Y/n]", question)
		} else {
			fmt.Printf(":: %s [y/N]", question)
		}

		response := WaitInputStringFromCLI()
		response = strings.ToLower(response)

		if len(response) == 0 {
			return defaultInput
		}

		if response == "y" || response == "yes" || response == "1" {
			return true
		} else if response == "n" || response == "no" || response == "0" {
			return false
		} else {
			continue
		}
	}

	//Tries exceeded the limit and answer is weird?
	os.Exit(2)
	return false
}

func (a *App) HandleHistoryRequest(c *cli.Context) {
	historyRequest := new(models.HistoryRequest)
	//TODO: Add new parameters
	historyRequest.NOP = true

	if c.Bool("all") {
		historyRequest.ScanSearches = true
		historyRequest.ScanWatches = true
	} else {
		if c.Bool("search") {
			historyRequest.ScanSearches = true
		} else if c.Bool("watch") {
			historyRequest.ScanWatches = true
		} else {
			historyRequest.ScanSearches = true
			historyRequest.ScanWatches = true
		}
	}

	//Initialize Printer to print responses
	DefaultPrinter = services.NewHistoryPrinter(*historyRequest)

	DefaultPrinter.PrintBanner()

	StartSpinner(4, 100)

	searches, err := a.Store.Movies.GetSearches(noContext)

	if err != nil {
		fmt.Println(err)
	}

	responses := []models.HistoryResponse{}

	for i, data := range searches {
		movie, err := a.Store.Movies.FindByID(noContext, data.MovieID)
		if err != nil {
			fmt.Println(err)
		}

		response := models.HistoryResponse{
			Search:     *searches[i],
			MovieTitle: movie.Title,
			MovieYear:  movie.Year,
		}

		responses = append(responses, response)
	}

	StopSpinner()

	DefaultPrinter.PrintHistoryResponses(responses)

	os.Exit(0)
}

func (a *App) HandleMyListRequest(c *cli.Context) {
	request := new(models.MyListRequest)

	if c.Bool("all") {
		request.ScanWLs = true
		request.ScanMLs = true
	} else {
		if c.Bool("watch") {
			request.ScanWLs = true
		} else if c.Bool("like") {
			request.ScanMLs = true
		} else {
			request.ScanWLs = true
			request.ScanMLs = true
		}
	}

	//Initialize Printer to print responses
	DefaultPrinter = services.NewMyListPrinter(UseNoResultFilter(), *request)

	DefaultPrinter.PrintBanner()

	StartSpinner(4, 100)

	responses := []models.MyListResponse{}

	if request.ScanWLs {
		wls, err := a.Store.Movies.GetWatchLaterList(noContext)
		if err != nil {
			fmt.Println(err)
		}
		for i, data := range wls {
			movie, err := a.Store.Movies.FindByID(noContext, data.MovieID)
			if err != nil {
				fmt.Println(err)
			}
			response := models.MyListResponse{
				WL:         *wls[i],
				MovieTitle: movie.Title,
				MovieYear:  movie.Year,
			}
			responses = append(responses, response)
		}
	}

	if request.ScanMLs {
		mls, err := a.Store.Movies.GetMovieLearnList(noContext)
		if err != nil {
			fmt.Println(err)
		}
		for i, data := range mls {
			movie, err := a.Store.Movies.FindByID(noContext, data.MovieID)
			if err != nil {
				fmt.Println(err)
			}
			response := models.MyListResponse{
				ML:         *mls[i],
				MovieTitle: movie.Title,
				MovieYear:  movie.Year,
			}
			responses = append(responses, response)
		}
	}

	StopSpinner()

	DefaultPrinter.PrintMyListResponses(responses)

	os.Exit(0)
}

func (a *App) HandleLearnRequest(c *cli.Context) {
	learnRequest := new(models.LearnRequest)
	learnRequest.Filename = strings.Join(c.Args(), " ")

	learn, err := learner.Learn(learnRequest)

	if err != nil {
		os.Exit(1)
	}

	fmt.Println(learn.Success)
}

//FIXME: Too long and hardcoded function, make it better
func (a *App) HandleSearchTitleRequest(c *cli.Context) {
	searchRequest := new(models.SearchRequest)
	searchRequest.Title = strings.Join(c.Args(), "+")

	if c.Bool("all") {
		searchRequest.ScanRT = true
		searchRequest.ScanIMDB = true
	} else {
		if c.Bool("imdb") {
			searchRequest.ScanIMDB = true
		} else if c.Bool("rottentomatoes") {
			searchRequest.ScanRT = true
		} else {
			//Default engine
			searchRequest.ScanIMDB = true

		}
	}

	//Initialize Printer to print responses
	DefaultPrinter = services.NewPrinter(UseNoResultFilter(), *searchRequest)

	DefaultPrinter.PrintBanner()

	StartSpinner(4, 100)

	//Initialize SearchEngine for Query
	engine := *services.NewSearcher(*searchRequest)

	//Get Responses from URLs
	responses := engine.GetSearchResponses()

	//Add to cache if available
	if config.Cache.UseCache {
		if config.Cache.UseSearchCache {
			searchJSON, err := json.Marshal(responses[0])
			if err != nil {
				fmt.Println(err)
				return
			}
			cache.WriteFile("IMDB", "searches", searchRequest.Title, string(searchJSON))
			//TODO: else timeout 1 day write
		}
	}

	StopSpinner()

	//Print the responses
	//Start Index: 0
	//End Index: 10
	//ShowMore Selected: false
	DefaultPrinter.PrintSearchResponses(0, 10, false, responses)

	var maxChoice = 10 * len(responses)

	isShowMoreOptionSelected := false

	fmt.Println()
	fmt.Println("(Enter 'q' to exit)")

	for {
		fmt.Printf("Please select your choice: ")

		choice := WaitInputIntFromCLI()

		if !isShowMoreOptionSelected && choice == 0 && len(responses) == 1 {
			isShowMoreOptionSelected = true

			//Print the other results that didn't show first
			//Start Index: 10
			//End Index: 250
			//ShowMore Selected: true
			DefaultPrinter.PrintSearchResponses(10, 200, isShowMoreOptionSelected, responses)
			continue

		} else if (!isShowMoreOptionSelected && choice > 0 && choice <= maxChoice) ||
			(isShowMoreOptionSelected && choice > 0 && choice <= 200) {

			fmt.Println("\nGetting information...")
			fmt.Println()

			StartSpinner(37, 150)

			movie := new(models.Movie)

			if c.Bool("all") {
				//FIXME: Default engine IMDB, make array? for all results?
				if (choice - 1) >= 10 {
					movie = engine.GetMovie("RottenTomatoes", responses[1].Searches[choice-1-10])
				} else if (choice - 1) >= 0 {
					movie = engine.GetMovie("IMDB", responses[0].Searches[choice-1])
				}
			} else {
				if c.Bool("imdb") {
					movie = engine.GetMovie("IMDB", responses[0].Searches[choice-1])
				} else if c.Bool("rottentomatoes") {
					movie = engine.GetMovie("RottenTomatoes", responses[0].Searches[choice-1])
				} else {
					movie = engine.GetMovie("IMDB", responses[0].Searches[choice-1])
				}
			}

			//Add to cache if available
			if config.Cache.UseCache {
				if config.Cache.UseMovieCache {
					movieJSON, err := json.Marshal(movie)
					if err != nil {
						fmt.Println(err)
						return
					}
					cache.WriteFile("IMDB", "movies", responses[0].Searches[choice-1].ID, string(movieJSON))
					//TODO: else timeout 1 day write
				}
			}

			StopSpinner()

			//Print the movie info that we had get above
			DefaultPrinter.PrintMovie(*movie)

			a.AddToSearchHistoryDB(movie.Info)

			var watched = false
			if len(movie.Info.URLTrailerIMDB) > 0 {
				watched = HandleWatchMovie(movie.Info.URLTrailerIMDB)
			}

			a.PostTrailerOperations(movie.Info, watched)

			break
		} else {
			os.Exit(0)
		}
	}

}

func HandleWatchMovie(url string) bool {
	fmt.Println()

	watch := AskYNQuestion("Do you want to watch Trailer?", 3, true)

	if watch {
		fmt.Printf("MPV Player loading...: %s", url)

		//TODO: Windows?
		cmd := exec.Command("/usr/bin/mpv", url)

		if err := cmd.Start(); err != nil {
			fmt.Printf("Failed to start cmd: %v", err)
			return false
		}

		if err := cmd.Wait(); err != nil {
			fmt.Printf("Cmd returned error: %v", err)
			return false
		}
	}

	return watch
}

func (a *App) PostTrailerOperations(info models.MovieInfo, watched bool) {
	var likedTrailer bool
	var likedMovie bool

	if watched {
		likedTrailer = AskYNQuestion("Did you like the Trailer?", 3, true)
	}

	watchedBefore := AskYNQuestion("Did you watch this movie before?", 3, false)

	if watchedBefore {
		likedMovie = AskYNQuestion("Did you like this movie?", 3, true)

		err := a.AddToMovieLikeDB(info, likedMovie)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		watchlater := AskYNQuestion("Do you want to add this movie to Watch Later list?", 3, true)

		if watchlater {
			err := a.AddToWatchLaterDB(info)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	if likedMovie || likedTrailer {
		os.Exit(2)
	}
}

var noContext = context.TODO()

func (a *App) AddToWatchLaterDB(info models.MovieInfo) error {
	exist, err := a.Store.Movies.FindByTitle(noContext, info.Title)
	if exist.ID != 0 {
		err := a.Store.Movies.CreateWL(noContext, exist)
		if err != nil {
			return err
		}
	}
	return err
}

func (a *App) AddToSearchHistoryDB(info models.MovieInfo) error {
	exist, err := a.Store.Movies.FindByTitle(noContext, info.Title)
	if exist.ID == 0 {
		err := a.Store.Movies.Create(noContext, &info)
		if err != nil {
			return err
		}
		added, err := a.Store.Movies.FindByTitle(noContext, info.Title)
		if err != nil {
			return err
		}
		err = a.Store.Movies.CreateSearch(noContext, added)
		if err != nil {
			return err
		}
	} else {
		err = a.Store.Movies.CreateSearch(noContext, exist)
		if err != nil {
			return err
		}
	}
	return err
}

func (a *App) AddToMovieLikeDB(info models.MovieInfo, liked bool) error {
	exist, err := a.Store.Movies.FindByTitle(noContext, info.Title)
	if exist.ID != 0 {
		err := a.Store.Movies.CreateML(noContext, exist, liked)
		if err != nil {
			return err
		}
	}
	return err
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
		ShowWLs:      true,
		ShowMLs:      true,
	}
}
