// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package torrent

import (
	"testing"

	"github.com/Dentrax/GMDB/models"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTorrent1337XLive(t *testing.T) {
	request := new(models.SearchTorrentRequest)
	request.Title = "Deadpool"

	engine := New("1337x.to", *request)

	response := engine.SearchTorrent(request)

	responses := []models.SearchTorrentResponse{}
	responses = append(responses, *response)

	Convey("Testing Live", t, func() {

		So(len(responses), ShouldBeGreaterThan, 0)

		Convey("Testing search response count", func() {
			So(responses[0].TotalFound, ShouldBeGreaterThan, 0)
		})

		Convey("Testing first response", func() {
			first := responses[0].Searches[0]

			Convey("Check response count", func() {
				So(len(responses), ShouldBeGreaterThan, 0)
			})

			Convey("Check Name", func() {
				So(first.Name, ShouldNotBeBlank)
			})

			Convey("Check URL", func() {
				So(first.URL, ShouldNotBeBlank)
			})

			Convey("Check Seeders", func() {
				So(first.Seeders, ShouldNotBeBlank)
			})

			Convey("Check Leechers", func() {
				So(first.Leechers, ShouldNotBeBlank)
			})

			Convey("Check Date", func() {
				So(first.Date, ShouldNotBeBlank)
			})

			Convey("Check Size", func() {
				So(first.Size, ShouldNotBeBlank)
			})

			Convey("Check Uploader", func() {
				So(first.Uploader, ShouldNotBeBlank)
			})

			Convey("Check Magnet URI", func() {
				magnet, err := engine.GetMagnet(&first)

				So(err, ShouldBeNil)
				So(magnet, ShouldNotBeBlank)
				So(magnet, ShouldStartWith, "magnet:?")
			})
		})

	})

}
