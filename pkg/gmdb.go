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

	"github.com/Dentrax/GMDB/learner"
	"github.com/Dentrax/GMDB/models"
	"github.com/Dentrax/GMDB/pkg/cache"
	"github.com/Dentrax/GMDB/pkg/config"
	"github.com/Dentrax/GMDB/services"
	"github.com/Dentrax/GMDB/services/torrent"
	"github.com/Dentrax/GMDB/store"
	"github.com/Dentrax/GMDB/store/database"
	"github.com/Dentrax/GMDB/store/movie"

	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
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
func AskYNQuestion(question string, tries int, defaultInput bool, red bool) bool {
	styleRed := chalk.Red.NewStyle().
		WithBackground(chalk.ResetColor).
		WithTextStyle(chalk.Bold)

	for ; tries > 0; tries-- {
		if defaultInput {
			if red {
				fmt.Printf("%s:: %s [Y/n]%s", styleRed, question, chalk.Reset)
			} else {
				fmt.Printf(":: %s [Y/n]", question)
			}
		} else {
			if red {
				fmt.Printf("%s:: %s [y/N]%s", styleRed, question, chalk.Reset)
			} else {
				fmt.Printf(":: %s [y/N]", question)
			}
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

func AskYNTorrentLegalUsage() bool {
	legal := AskYNQuestion("I have read and agree to the LEGAL_DISCLAIMER.md for torrent usage.", 3, true, true)
	return legal
}

func (a *App) HandleUpdateRequest(c *cli.Context) {
	request := new(models.UpdateRequest)
	request.NOP = true

	DefaultPrinter = services.NewUpdatePrinter(UseNoResultFilter(), *request)

	DefaultPrinter.PrintBanner()

	movies, err := a.Store.Movies.GetMovies(noContext)
	if err != nil {
		fmt.Println(err)
	}

	searchRequest := new(models.SearchRequest)
	searchRequest.ScanIMDB = true

	updateNow := AskYNQuestion("Do you wanna update movies now?", 3, true, false)
	fmt.Println()

	totalUpdated := 0

	if updateNow {
		for _, movie := range movies {
			StartSpinner(4, 100)

			title := strings.Replace(movie.Title, " ", "+", -1)
			searchRequest.Title = title

			imdb := *services.NewSearcher(*searchRequest)

			responses := imdb.GetSearchResponses()

			var res *models.Movie
			res = imdb.GetMovie("IMDB", responses[0].Searches[0], true)
			res.Info.Created = movie.Created

			//Search until correct movie found (only guess)
			//TODO: Find a better solution, this is so F-bad
			if len(res.Info.Year) == 0 || len(res.Info.Released) == 0 || len(res.Info.Rating) == 0 || len(res.Info.Duration) == 0 {
				for i, search := range responses[0].Searches {
					next := imdb.GetMovie("IMDB", search, true)
					next.Info.Created = movie.Created

					if i <= 5 {
						if len(next.Info.Year) != 0 && len(next.Info.Released) != 0 && len(next.Info.Rating) != 0 && len(next.Info.Duration) != 0 {
							res = next
							break
						}
					} else {
						break
					}
				}
			}

			if strings.Compare(res.Info.Title, movie.Title) == 0 {
				totalUpdated = totalUpdated + 1

				a.AddToSearchHistoryDB(res.Info, "Update")

				StopSpinner()
				DefaultPrinter.PrintUpdateMovieSuccess(res.Info.Title)
			} else {
				StopSpinner()
				DefaultPrinter.PrintUpdateMovieFailed(res.Info.Title, movie.Title)

			}
		}
	}

	StopSpinner()

	if totalUpdated > 0 {
		fmt.Println()
		fmt.Printf("Movies updated successfuly: [%d/%d]\n", totalUpdated, len(movies))
	}

	os.Exit(0)
}

func (a *App) HandleNoteRequest(c *cli.Context) {
	request := new(models.NoteRequest)
	request.NOP = true

	historyRequest := new(models.HistoryRequest)
	historyRequest.ScanSearches = true
	historyRequest.ScanWatches = true

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

	if len(responses) == 0 {
		fmt.Printf("No movie found! Please update your movie list using 'update' before take a note.")
		os.Exit(0)
	}

	DefaultPrinter.PrintHistoryResponses(responses)

	fmt.Println("(Enter 'q' to exit)")

	fmt.Printf("Please select your choice: ")
	choiceMovie := WaitInputIntFromCLI()

	if choiceMovie <= 0 || choiceMovie > len(responses) {
		fmt.Printf("You must choose between %d and %d!", 1, len(responses))
		os.Exit(2)
	}

	movie, err := a.Store.Movies.FindByTitle(noContext, responses[choiceMovie-1].MovieTitle)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("\nYou are selected: %s\n", movie.Title)
	fmt.Println()

	fmt.Println(" 1) Add new note")
	fmt.Println(" 2) Update a note")
	fmt.Println(" 3) Get notes")

	fmt.Printf("\nPlease select your operation: ")
	choiceOperation := WaitInputIntFromCLI()

	if choiceOperation <= 0 || choiceOperation > 3 {
		fmt.Printf("You must choose between %d and %d!", 1, 3)
		os.Exit(2)
	}

	fmt.Println()

	if choiceOperation == 1 {
		note := handleOperationMovieNoteInfo(*movie)
		err = a.Store.Movies.CreateNI(noContext, movie, &note)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		os.Exit(0)
	} else if choiceOperation == 2 {
		note := handleOperationMovieNoteInfo(*movie)
		err = a.Store.Movies.UpdateNI(noContext, movie, &note)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		os.Exit(0)
	} else if choiceOperation == 3 {
		DefaultPrinter = services.NewNotePrinter(*request)

		StartSpinner(4, 100)

		responses := []models.NoteResponse{}

		notes, err := a.Store.Movies.GetMovieNoteList(noContext)
		if err != nil {
			fmt.Println(err)
		}
		for i, data := range notes {
			movie, err := a.Store.Movies.FindByID(noContext, data.MovieID)
			if err != nil {
				fmt.Println(err)
			}
			response := models.NoteResponse{
				Note:       *notes[i],
				MovieTitle: movie.Title,
				MovieYear:  movie.Year,
			}
			responses = append(responses, response)
		}

		StopSpinner()

		DefaultPrinter.PrintNoteResponses(responses)

		os.Exit(2)
	}
}

func handleOperationMovieNoteInfo(movie models.MovieInfo) models.MovieNoteInfo {
	choiceSeason := 0
	choiceEpisode := 0

	if movie.IsTVSeries {
		fmt.Printf("Enter Season: ")
		choiceSeason = WaitInputIntFromCLI()

		fmt.Printf("Enter Episode: ")
		choiceEpisode = WaitInputIntFromCLI()
	}

	fmt.Printf("Enter the Hour: ")
	choiceHour := WaitInputIntFromCLI()

	fmt.Printf("Enter the Minute: ")
	choiceMinute := WaitInputIntFromCLI()

	fmt.Printf("Enter the Second: ")
	choiceSecond := WaitInputIntFromCLI()

	fmt.Printf("Enter the your Note: ")
	text := WaitInputStringFromCLI()

	note := models.MovieNoteInfo{
		Season:  uint8(choiceSeason),
		Episode: uint8(choiceEpisode),
		Hour:    uint8(choiceHour),
		Minute:  uint8(choiceMinute),
		Second:  uint8(choiceSecond),
		Text:    text,
	}

	return note
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
	request := new(models.LearnRequest)
	request.Filename = strings.Join(c.Args(), " ")

	DefaultPrinter = services.NewLearnPrinter(UseNoResultFilter(), *request)

	DefaultPrinter.PrintBanner()

	DefaultPrinter.PrintPhaseStart(1, "Checking file format...")

	provider, err := learner.CheckFileFormat(*request)

	if err != nil {
		DefaultPrinter.PrintPhaseFail(1, err.Error())
		os.Exit(1)
	} else {
		DefaultPrinter.PrintPhaseDone(1)
	}

	fmt.Printf("---> %s detected!\n", provider)

	DefaultPrinter.PrintPhaseStart(2, "Scanning movies...")

	learns, err := learner.ScanMovies(*request)

	if err != nil {
		DefaultPrinter.PrintPhaseFail(2, err.Error())
		os.Exit(1)
	} else {
		DefaultPrinter.PrintPhaseDone(2)
	}

	DefaultPrinter.PrintPhaseStart(3, "Learning movies...")

	DefaultPrinter.PrintPhaseDone(3)

	fmt.Println()

	//Map for Title (Liked or Not)
	likes := make(map[string]bool)

	for _, learn := range learns {
		question := fmt.Sprintf("Did you like the %s movie?", learn.Result.Title)
		likedMovie := AskYNQuestion(question, 3, true, false)

		likes[learn.Result.Title] = likedMovie
	}

	DefaultPrinter.PrintPhaseStart(3, "Implementing movies to database...")

	StartSpinner(4, 100)

	for key, value := range likes {
		movie := models.MovieInfo{
			Title: key,
		}
		a.AddToMovieLikeDB(movie, value)
		a.AddToWatchLaterDB(movie, true)
	}

	StopSpinner()

	DefaultPrinter.PrintPhaseDone(3)
}

func (a *App) HandleTorrentRequest(c *cli.Context, movie *models.MovieInfo) {
	request := new(models.SearchTorrentRequest)

	if c != nil && movie == nil {
		request.Title = strings.Join(c.Args(), "+")

		DefaultPrinter = services.NewTorrentPrinter(UseNoResultFilter(), *request)

		DefaultPrinter.PrintBanner()
	} else if movie != nil {
		request.Title = movie.Title + " " + movie.Year

		DefaultPrinter = services.NewTorrentPrinter(UseNoResultFilter(), *request)
	} else {
		fmt.Println("Wrong call")
		os.Exit(2)
	}

	if !AskYNTorrentLegalUsage() {
		os.Exit(0)
	}

	StartSpinner(4, 100)

	engine := torrent.New("1337x.to", *request)

	response := engine.SearchTorrent(request)

	responses := []models.SearchTorrentResponse{}
	responses = append(responses, *response)

	StopSpinner()

	DefaultPrinter.PrintTorrentResponses(0, 10, false, responses)

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
			DefaultPrinter.PrintTorrentResponses(10, 200, isShowMoreOptionSelected, responses)
			continue

		} else if (!isShowMoreOptionSelected && choice > 0 && choice <= maxChoice) ||
			(isShowMoreOptionSelected && choice > 0 && choice <= 200) {

			fmt.Println("\nGetting information...")
			fmt.Println()

			year := strconv.Itoa(responses[0].Searches[choice].Info.Year)

			season := responses[0].Searches[choice].Info.Season
			episode := responses[0].Searches[choice].Info.Episode

			isTVSeies := season > 0 || episode > 0

			movie := models.MovieInfo{
				Title:      responses[0].Searches[choice].Info.Title,
				Year:       year,
				IsTVSeries: isTVSeies,
			}

			err := a.AddToSearchHistoryDB(movie, "Torrent")

			magnet, err := engine.GetMagnet(&responses[0].Searches[choice])
			if err != nil {
				fmt.Println(err)
			}

			var watched = false
			if len(magnet) > 0 {
				watched = HandleWatchTorrentMagnet(magnet)
			}
			if watched {

			}

			//a.PostTrailerOperations(movie.Info, watched)

			break
		} else {
			os.Exit(0)
		}
	}

	os.Exit(0)
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
					movie = engine.GetMovie("RottenTomatoes", responses[1].Searches[choice-1-10], false)
				} else if (choice - 1) >= 0 {
					movie = engine.GetMovie("IMDB", responses[0].Searches[choice-1], false)
				}
			} else {
				if c.Bool("imdb") {
					movie = engine.GetMovie("IMDB", responses[0].Searches[choice-1], false)
				} else if c.Bool("rottentomatoes") {
					movie = engine.GetMovie("RottenTomatoes", responses[0].Searches[choice-1], false)
				} else {
					movie = engine.GetMovie("IMDB", responses[0].Searches[choice-1], false)
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

			a.AddToSearchHistoryDB(movie.Info, "Search")

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

	watch := AskYNQuestion("Do you want to watch Trailer?", 3, true, false)

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

func HandleWatchTorrentMagnet(magnet string) bool {
	watch := AskYNQuestion("Do you want to watch Movie using Peerflix?", 3, true, false)

	if watch {
		rgxMagnet, _ := regexp.Compile("magnet:\\?xt=urn:[a-z0-9]+:[a-zA-Z0-9]{32,40}")
		fmt.Println()
		fmt.Printf("Peerflix: (MPV) Player loading...: %s", rgxMagnet.FindString(magnet))

		cmd := exec.Command("/usr/bin/peerflix", "--mpv", magnet)

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
		likedTrailer = AskYNQuestion("Did you like the Trailer?", 3, true, false)
	}

	watchedBefore := AskYNQuestion("Did you watch this movie before?", 3, false, false)

	if watchedBefore {
		likedMovie = AskYNQuestion("Did you like this movie?", 3, true, false)

		err := a.AddToMovieLikeDB(info, likedMovie)
		if err != nil {
			fmt.Println(err)
		}
	}
	watchNow := AskYNQuestion("Do you want to watch this movie now?", 3, true, false)
	if watchNow {

		a.HandleTorrentRequest(nil, &info)

		request := new(models.SearchTorrentRequest)
		request.Title = info.Title + " " + info.Year

		DefaultPrinter = services.NewTorrentPrinter(UseNoResultFilter(), *request)

		if !AskYNTorrentLegalUsage() {
			os.Exit(0)
		}

		StartSpinner(4, 100)

		engine := torrent.New("1337x.to", *request)

		response := engine.SearchTorrent(request)

		responses := []models.SearchTorrentResponse{}
		responses = append(responses, *response)

		StopSpinner()

		DefaultPrinter.PrintTorrentResponses(0, 10, false, responses)

	} else {
		watchlater := AskYNQuestion("Do you want to add this movie to Watch Later list?", 3, true, false)

		if watchlater {
			err := a.AddToWatchLaterDB(info, false)
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

func (a *App) AddToWatchLaterDB(info models.MovieInfo, watched bool) error {
	exist, err := a.Store.Movies.FindByTitle(noContext, info.Title)
	if exist.ID > 0 {
		existWL, _ := a.Store.Movies.FindWL(noContext, exist.ID)
		if existWL.ID > 0 {
			err := a.Store.Movies.UpdateWL(noContext, &info, true)
			if err != nil {
				return err
			}
		} else if existWL.ID == 0 {
			err := a.Store.Movies.CreateWL(noContext, exist, watched)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (a *App) AddToSearchHistoryDB(info models.MovieInfo, from string) error {
	now := time.Now().Unix()
	exist, err := a.Store.Movies.FindByTitle(noContext, info.Title)
	if exist.ID == 0 {
		info.Created = now
		err := a.Store.Movies.Create(noContext, &info)
		if err != nil {
			return err
		}
		added, err := a.Store.Movies.FindByTitle(noContext, info.Title)
		if err != nil {
			return err
		}
		err = a.Store.Movies.CreateSearch(noContext, added, from)
		if err != nil {
			return err
		}
	} else if exist.ID > 0 {
		info.ID = exist.ID
		info.Updated = now
		err := a.Store.Movies.Update(noContext, &info)
		if err != nil {
			return err
		}
		err = a.Store.Movies.CreateSearch(noContext, exist, from)
		if err != nil {
			return err
		}
	}
	return err
}

func (a *App) AddToMovieLikeDB(info models.MovieInfo, liked bool) error {
	now := time.Now().Unix()
	exist, err := a.Store.Movies.FindByTitle(noContext, info.Title)
	if exist.ID == 0 {
		info.Created = now
		err := a.Store.Movies.Create(noContext, &info)
		if err != nil {
			return err
		}
		existNew, err := a.Store.Movies.FindByTitle(noContext, info.Title)
		if err != nil {
			return err
		}
		err = a.Store.Movies.CreateML(noContext, existNew, liked)
		if err != nil {
			return err
		}
	} else if exist.ID > 0 {
		info.ID = exist.ID
		err := a.Store.Movies.UpdateML(noContext, &info, liked)
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
