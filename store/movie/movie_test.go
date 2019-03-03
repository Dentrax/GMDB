package movie

import (
	"context"
	"testing"

	"gmdb/models"
	"gmdb/store/database/dbtest"

	. "github.com/smartystreets/goconvey/convey"
)

var SampleMovie = &models.MovieInfo{
	Title:             "Deadpool",
	Year:              "2016",
	Rating:            "8.0/10",
	Votes:             "707,707",
	Reviews:           "Reviews",
	Duration:          "1h 48min",
	Seasons:           "0",
	Episodes:          "0",
	Summary:           "Summary",
	Metascore:         "65",
	ReviewCountUser:   "7",
	ReviewCountCritic: "7",
	RTMeter:           "7",
	URLTrailerIMDB:    "URLTrailerIMDB",
	URLPosterIMDB:     "URLPosterIMDB",
}

var SampleMovieNote = &models.MovieNoteInfo{
	From:    1,
	Season:  3,
	Episode: 5,
	Hour:    2,
	Minute:  34,
	Second:  44,
	Text:    "My Awesome Note",
}

var noContext = context.TODO()

//So(Version(), ShouldEqual, version)
//So(inSlice("a", ss), ShouldBeTrue)
//So(inSlice("d", ss), ShouldBeFalse)
//So(err, ShouldBeNil)
//So(f, ShouldNotBeNil)

func TestMovie(t *testing.T) {
	Convey("Testing Movie database operations", t, func() {
		Convey("Connect to SQLite database", func() {
			conn, err := dbtest.Connect()

			So(err, ShouldBeNil)
			So(conn, ShouldNotBeNil)

			defer func() {
				dbtest.Reset(conn)
				dbtest.Disconnect(conn)
			}()

			store := New(conn).(*movieStore)

			So(store, ShouldNotBeNil)

			Convey("Create a movie model in the database", func() {
				err := store.Create(noContext, SampleMovie)

				So(err, ShouldBeNil)
				So(SampleMovie.ID, ShouldEqual, 1)

				Convey("Create a movie model that already exist", func() {
					err := store.Create(noContext, SampleMovie)

					So(err, ShouldNotBeNil)
				})

				Convey("Get all movies in the database", func() {
					movies, err := store.GetMovies(noContext)

					So(err, ShouldBeNil)
					So(movies, ShouldNotBeNil)
					So(len(movies), ShouldEqual, int64(1))
				})

				Convey("Count movie model in the database", func() {
					count, err := store.Count(noContext)

					So(err, ShouldBeNil)
					So(count, ShouldNotBeNil)
					So(count, ShouldEqual, int64(1))
				})

				Convey("Find the movie model by Title in the database", func() {
					movie, err := store.FindByTitle(noContext, SampleMovie.Title)

					So(err, ShouldBeNil)
					So(movie, ShouldNotBeNil)

					Convey("Test the found movie model", func() {
						So(movie.Title, ShouldEqual, "Deadpool")
						So(movie.Year, ShouldEqual, "2016")
						So(movie.Rating, ShouldEqual, "8.0/10")
					})
				})

				Convey("Find the movie model by ID in the database", func() {
					movie, err := store.FindByID(noContext, SampleMovie.ID)

					So(err, ShouldBeNil)
					So(movie, ShouldNotBeNil)

					Convey("Test the found movie model", func() {
						So(movie.Title, ShouldEqual, "Deadpool")
						So(movie.Year, ShouldEqual, "2016")
						So(movie.Rating, ShouldEqual, "8.0/10")
					})
				})

				Convey("Update the movie model in the database", func() {
					sampleUpdate := &models.MovieInfo{
						ID:     SampleMovie.ID,
						Title:  "Deadpool 2",
						Year:   "2018",
						Rating: "7.8/10",
					}
					err := store.Update(noContext, sampleUpdate)

					So(err, ShouldBeNil)

					updated, err := store.FindByTitle(noContext, sampleUpdate.Title)

					So(err, ShouldBeNil)
					So(updated, ShouldNotBeNil)

					Convey("Test the updated movie model", func() {
						So(updated.Title, ShouldEqual, sampleUpdate.Title)
						So(updated.Year, ShouldEqual, sampleUpdate.Year)
						So(updated.Rating, ShouldEqual, sampleUpdate.Rating)
					})
				})

				Convey("Find and Update the movie model in the database", func() {
					exist, err := store.FindByTitle(noContext, SampleMovie.Title)
					So(err, ShouldBeNil)

					sampleUpdate := &models.MovieInfo{
						ID:     exist.ID,
						Title:  "Deadpool 3",
						Year:   "202X",
						Rating: "8.8/10",
					}

					err = store.Update(noContext, sampleUpdate)

					So(err, ShouldBeNil)

					updated, err := store.FindByTitle(noContext, sampleUpdate.Title)

					So(err, ShouldBeNil)
					So(updated, ShouldNotBeNil)

					Convey("Test the updated movie model", func() {
						So(updated.Title, ShouldEqual, sampleUpdate.Title)
						So(updated.Year, ShouldEqual, sampleUpdate.Year)
						So(updated.Rating, ShouldEqual, sampleUpdate.Rating)
					})
				})

				Convey("Delete the movie model in the database", func() {
					count, err := store.Count(noContext)

					So(err, ShouldBeNil)
					So(count, ShouldNotBeNil)
					So(count == 1, ShouldBeTrue)

					err = store.Delete(noContext, &models.MovieInfo{ID: SampleMovie.ID})
					So(err, ShouldBeNil)

					count, err = store.Count(noContext)

					So(err, ShouldBeNil)
					So(count, ShouldNotBeNil)
					So(count == 0, ShouldBeTrue)
				})
			})

			Convey("Create a Search model in the database", func() {
				err := store.CreateSearch(noContext, SampleMovie, "Test")

				So(err, ShouldBeNil)

				Convey("Get Search List model in the database", func() {
					searches, err := store.GetSearches(noContext)

					So(err, ShouldBeNil)
					So(searches, ShouldNotBeNil)
					So(len(searches), ShouldEqual, int64(1))
				})
			})

			Convey("Create a Watch Later model in the database", func() {
				err := store.CreateWL(noContext, SampleMovie, false)

				So(err, ShouldBeNil)

				Convey("Count Watch Later model in the database", func() {
					count, err := store.CountWL(noContext)

					So(err, ShouldBeNil)
					So(count, ShouldNotBeNil)
					So(count, ShouldEqual, int64(1))
				})

				Convey("Find Watch Later model in the database", func() {
					movieWL, err := store.FindWL(noContext, SampleMovie.ID)

					So(err, ShouldBeNil)
					So(movieWL, ShouldNotBeNil)

					Convey("Test the found WL model", func() {
						So(movieWL.Watched, ShouldBeFalse)
					})
				})

				Convey("Get Watch Later List model in the database", func() {
					movieWLs, err := store.GetWatchLaterList(noContext)

					So(err, ShouldBeNil)
					So(movieWLs, ShouldNotBeNil)
					So(len(movieWLs), ShouldEqual, int64(1))
				})

				Convey("Update Watch Later model in the database", func() {
					err := store.UpdateWL(noContext, SampleMovie, true)

					So(err, ShouldBeNil)

					updated, err := store.FindWL(noContext, SampleMovie.ID)

					So(err, ShouldBeNil)
					So(updated, ShouldNotBeNil)

					Convey("Test the updated Watch Later model", func() {
						So(updated.Watched, ShouldEqual, true)
					})
				})
			})

			Convey("Create a Movie Learn model in the database", func() {
				err := store.CreateML(noContext, SampleMovie, true)

				So(err, ShouldBeNil)

				Convey("Count Movie Learn model in the database", func() {
					count, err := store.CountML(noContext)

					So(err, ShouldBeNil)
					So(count, ShouldNotBeNil)
					So(count, ShouldEqual, int64(1))
				})

				Convey("Find Movie Learn model in the database", func() {
					movieML, err := store.FindML(noContext, SampleMovie.ID)

					So(err, ShouldBeNil)
					So(movieML, ShouldNotBeNil)

					Convey("Test the found ML model", func() {
						So(movieML.Liked, ShouldBeTrue)
					})
				})

				Convey("Get Movie Learn List model in the database", func() {
					movieMLs, err := store.GetMovieLearnList(noContext)

					So(err, ShouldBeNil)
					So(movieMLs, ShouldNotBeNil)
					So(len(movieMLs), ShouldEqual, int64(1))
				})

				Convey("Update Movie Learn model in the database", func() {
					err := store.UpdateML(noContext, SampleMovie, true)

					So(err, ShouldBeNil)

					err = store.UpdateML(noContext, SampleMovie, false)

					So(err, ShouldBeNil)
				})
			})

			Convey("Create a Movie Note model in the database", func() {
				err := store.CreateNI(noContext, SampleMovie, SampleMovieNote)

				So(err, ShouldBeNil)

				Convey("Count Movie Note model in the database", func() {
					count, err := store.CountNI(noContext)

					So(err, ShouldBeNil)
					So(count, ShouldNotBeNil)
					So(count, ShouldEqual, int64(1))
				})

				Convey("Find Movie Note model in the database", func() {
					movieNI, err := store.FindNI(noContext, SampleMovie.ID)

					So(err, ShouldBeNil)
					So(movieNI, ShouldNotBeNil)

					Convey("Test the found NI model", func() {
						So(movieNI.Season, ShouldEqual, SampleMovieNote.Season)
						So(movieNI.Episode, ShouldEqual, SampleMovieNote.Episode)
						So(movieNI.Hour, ShouldEqual, SampleMovieNote.Hour)
						So(movieNI.Minute, ShouldEqual, SampleMovieNote.Minute)
						So(movieNI.Second, ShouldEqual, SampleMovieNote.Second)
						So(movieNI.Text, ShouldEqual, SampleMovieNote.Text)
					})
				})

				Convey("Get Movie Note List model in the database", func() {
					movieNIs, err := store.GetMovieNoteList(noContext)

					So(err, ShouldBeNil)
					So(movieNIs, ShouldNotBeNil)
					So(len(movieNIs), ShouldEqual, int64(1))
				})

				Convey("Update Movie Note model in the database", func() {
					sampleUpdate := &models.MovieNoteInfo{
						Season:  5,
						Episode: 4,
						Hour:    2,
						Minute:  43,
						Second:  55,
						Text:    "My Awesome Note Updated",
					}

					err := store.UpdateNI(noContext, SampleMovie, sampleUpdate)

					So(err, ShouldBeNil)

					movieNI, err := store.FindNI(noContext, SampleMovie.ID)

					So(err, ShouldBeNil)
					So(movieNI, ShouldNotBeNil)

					Convey("Test the updated NI model", func() {
						So(movieNI.Season, ShouldEqual, sampleUpdate.Season)
						So(movieNI.Episode, ShouldEqual, sampleUpdate.Episode)
						So(movieNI.Hour, ShouldEqual, sampleUpdate.Hour)
						So(movieNI.Minute, ShouldEqual, sampleUpdate.Minute)
						So(movieNI.Second, ShouldEqual, sampleUpdate.Second)
						So(movieNI.Text, ShouldEqual, sampleUpdate.Text)
					})
				})
			})
		})

		Convey("Reset the SQLite database", func() {
			conn, err := dbtest.Connect()

			So(err, ShouldBeNil)
			So(conn, ShouldNotBeNil)

			err = dbtest.Reset(conn)

			So(err, ShouldBeNil)
		})

		Convey("Disconnect from SQLite database", func() {
			conn, err := dbtest.Connect()

			So(err, ShouldBeNil)
			So(conn, ShouldNotBeNil)

			err = dbtest.Disconnect(conn)

			So(err, ShouldBeNil)
		})
	})

}
