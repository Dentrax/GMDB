// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package torrent1337x

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/Dentrax/GMDB/models"
	"github.com/Dentrax/GMDB/services/common"

	. "github.com/smartystreets/goconvey/convey"
)

const PATH_TEST = "../../../testdata/TORRENT/"

func TestTorrent1337XHome(t *testing.T) {

	var testDatas = []struct {
		FolderName string
		File       string
		Response   *models.SearchTorrentResponse
	}{
		{"1337X", "home.html",
			&models.SearchTorrentResponse{
				SearchEngine: "1337x.to",
				Error:        "",
				TotalFound:   20,
				Searches: []models.SearchTorrentResult{
					{
						Name:     "Deadpool 2 (2018) [BluRay] [1080p] [YTS] [YIFY]",
						URL:      "/torrent/3163584/Deadpool-2-2018-BluRay-1080p-YTS-YIFY/",
						Seeders:  "17387",
						Leechers: "5620",
						Date:     "Aug. 9th '18",
						Size:     "2.1 GB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Deadpool 2 (2018) [BluRay] [720p] [YTS] [YIFY]",
						URL:      "/torrent/3163455/Deadpool-2-2018-BluRay-720p-YTS-YIFY/",
						Seeders:  "10288",
						Leechers: "3848",
						Date:     "Aug. 8th '18",
						Size:     "1.1 GB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Deadpool 2 (2018) [WEBRip] [1080p] [YTS] [YIFY]",
						URL:      "/torrent/3162916/Deadpool-2-2018-WEBRip-1080p-YTS-YIFY/",
						Seeders:  "7378",
						Leechers: "3188",
						Date:     "Aug. 8th '18",
						Size:     "2.1 GB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Deadpool 2 (2018) [WEBRip] [720p] [YTS] [YIFY]",
						URL:      "/torrent/3162062/Deadpool-2-2018-WEBRip-720p-YTS-YIFY/",
						Seeders:  "6247",
						Leechers: "3102",
						Date:     "Aug. 8th '18",
						Size:     "1.1 GB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Deadpool 2 (2018) English HDTS- 720p - x264 - AC3 - 2.4GB-TamilRockers",
						URL:      "/torrent/2999484/Deadpool-2-2018-English-HDTS-720p-x264-AC3-2-4GB-TamilRockers/",
						Seeders:  "5780",
						Leechers: "1804",
						Date:     "May. 17th '18",
						Size:     "2.4 GB",
						Uploader: "mazemaze16",
					},
					{
						Name:     "Deadpool.2.2018.720p.KORSUB.HDRip.x264-NON",
						URL:      "/torrent/3136775/Deadpool-2-2018-720p-KORSUB-HDRip-x264-NON/",
						Seeders:  "4515",
						Leechers: "978",
						Date:     "Jul. 25th '18",
						Size:     "3.2 GB",
						Uploader: "MrStark",
					},
					{
						Name:     "Deadpool 2 2018 1080p KORSUB WEBRip-x264-AC3-Zi$t",
						URL:      "/torrent/3137322/Deadpool-2-2018-1080p-KORSUB-WEBRip-x264-AC3-Zi-t/",
						Seeders:  "2630",
						Leechers: "643",
						Date:     "Jul. 26th '18",
						Size:     "2.0 GB",
						Uploader: "Zist313",
					},
					{
						Name:     "Deadpool.2.2018.1080p.WEB-DL.DD5.1.H264-CMRG[EtHD]",
						URL:      "/torrent/3160028/Deadpool-2-2018-1080p-WEB-DL-DD5-1-H264-CMRG-EtHD/",
						Seeders:  "2055",
						Leechers: "494",
						Date:     "Aug. 7th '18",
						Size:     "4.1 GB",
						Uploader: "EtHD",
					},
					{
						Name:     "Deadpool 2 (2018) English HDCAM-Rip - 720p - x264 - MP3 - 800MB",
						URL:      "/torrent/2999183/Deadpool-2-2018-English-HDCAM-Rip-720p-x264-MP3-800MB/",
						Seeders:  "1884",
						Leechers: "329",
						Date:     "May. 17th '18",
						Size:     "806.6 MB",
						Uploader: "rock3",
					},
					{
						Name:     "Deadpool 2016 1080p BluRay x264 DTS-JYK",
						URL:      "/torrent/1572854/Deadpool-2016-1080p-BluRay-x264-DTS-JYK/",
						Seeders:  "1819",
						Leechers: "345",
						Date:     "May. 8th '16",
						Size:     "2.7 GB",
						Uploader: "jjblack2",
					},
					{
						Name:     "Deadpool 2 2018 NEW 720p HD-TC X264-CPG",
						URL:      "/torrent/3030369/Deadpool-2-2018-NEW-720p-HD-TC-X264-CPG/",
						Seeders:  "1593",
						Leechers: "621",
						Date:     "Jun. 1st '18",
						Size:     "2.5 GB",
						Uploader: "Silmarillion",
					},
					{
						Name:     "Deadpool (2016) [720p] [YTS.AG] - YIFY",
						URL:      "/torrent/2100438/Deadpool-2016-720p-YTS-AG-YIFY/",
						Seeders:  "1576",
						Leechers: "356",
						Date:     "Mar. 8th '17",
						Size:     "798.6 MB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Once Upon A Deadpool (2018) 720p WEB-DL 950MB - MkvCage",
						URL:      "/torrent/3522044/Once-Upon-A-Deadpool-2018-720p-WEB-DL-950MB-MkvCage/",
						Seeders:  "1405",
						Leechers: "376",
						Date:     "Jan. 15th '19",
						Size:     "959.2 MB",
						Uploader: "MkvCage",
					},
					{
						Name:     "Deadpool 2 (2018) English 720p HC HDRip x264 1GB - Team TR",
						URL:      "/torrent/3136841/Deadpool-2-2018-English-720p-HC-HDRip-x264-1GB-Team-TR/",
						Seeders:  "1206",
						Leechers: "378",
						Date:     "Jul. 25th '18",
						Size:     "1,005.4 MB",
						Uploader: "RockerS",
					},
					{
						Name:     "Deadpool (2016) [1080p] [YTS.AG] - YIFY",
						URL:      "/torrent/2100442/Deadpool-2016-1080p-YTS-AG-YIFY/",
						Seeders:  "1171",
						Leechers: "143",
						Date:     "Mar. 8th '17",
						Size:     "1.6 GB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Deadpool.2.2018.720p.KORSUB.HDRip.x264-NON[EtHD]",
						URL:      "/torrent/3136852/Deadpool-2-2018-720p-KORSUB-HDRip-x264-NON-EtHD/",
						Seeders:  "1164",
						Leechers: "225",
						Date:     "Jul. 25th '18",
						Size:     "3.2 GB",
						Uploader: "EtHD",
					},
					{
						Name:     "Deadpool.2016.720p.HC.HDRip.x264.AAC-ETRG",
						URL:      "/torrent/1511041/Deadpool-2016-720p-HC-HDRip-x264-AAC-ETRG/",
						Seeders:  "995",
						Leechers: "40",
						Date:     "Mar. 24th '16",
						Size:     "924.6 MB",
						Uploader: "SaM",
					},
					{
						Name:     "Deadpool 2 (2018) [WEBRip] [720p] [YTS] [YIFY]",
						URL:      "/torrent/3161487/Deadpool-2-2018-WEBRip-720p-YTS-YIFY/",
						Seeders:  "828",
						Leechers: "66",
						Date:     "Aug. 7th '18",
						Size:     "2.1 GB",
						Uploader: "YTSAGx",
					},
					{
						Name:     "Deadpool 2 2018 720p HDTS x264 Dual Audio [Hindi - English] [MW]",
						URL:      "/torrent/3000277/Deadpool-2-2018-720p-HDTS-x264-Dual-Audio-Hindi-English-MW/",
						Seeders:  "804",
						Leechers: "437",
						Date:     "May. 17th '18",
						Size:     "876.9 MB",
						Uploader: "moviezworldz",
					},
					{
						Name:     "Deadpool.2.2018.Super.Duper.Cut.UNRATED.1080p.AMZN.WEBRip.DDP5.1.x264-ION10",
						URL:      "/torrent/3160311/Deadpool-2-2018-Super-Duper-Cut-UNRATED-1080p-AMZN-WEBRip-DDP5-1-x264-ION10/",
						Seeders:  "789",
						Leechers: "195",
						Date:     "Aug. 7th '18",
						Size:     "9.5 GB",
						Uploader: "SeekNDstroy",
					},
				},
			},
		},
	}

	Convey("Torrent:1337X Parse Home test", t, func() {

		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestHome := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestHome)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			current := ParseSearchTorrents(services.GetDocumentFromFile(currentTestHome))
			testName := strings.ToUpper(data.FolderName)

			So(current, ShouldNotBeNil)

			Convey("Testing model: "+testName, func() {

				So(current, ShouldNotBeNil)

				Convey("Testing search engine type", func() {
					So(current.SearchEngine, ShouldEqual, data.Response.SearchEngine)
				})

				Convey("Testing errors", func() {
					So(current.Error, ShouldEqual, "")
				})

				Convey("Testing total founds", func() {
					So(current.TotalFound, ShouldEqual, data.Response.TotalFound)
				})

				Convey("Testing searches", func() {

					So(current.Searches, ShouldNotBeEmpty)
					So(len(current.Searches), ShouldEqual, len(data.Response.Searches))

					for i, search := range current.Searches {

						So(search.Name, ShouldNotBeBlank)
						So(search.Name, ShouldEqual, data.Response.Searches[i].Name)

						t := fmt.Sprintf("%02s", strconv.Itoa(i))
						Convey("Testing torrent: ["+t+"]: "+search.Name, func() {
							So(search.URL, ShouldNotBeBlank)
							So(search.URL, ShouldStartWith, "/torrent/")
							So(search.URL, ShouldEndWith, "/")
							So(search.URL, ShouldEqual, data.Response.Searches[i].URL)

							seeders, _ := strconv.Atoi(search.Seeders)
							So(search.Seeders, ShouldNotBeBlank)
							So(seeders, ShouldBeBetween, 0, 1000000)
							So(search.Seeders, ShouldEqual, data.Response.Searches[i].Seeders)

							leechers, _ := strconv.Atoi(search.Leechers)
							So(search.Leechers, ShouldNotBeBlank)
							So(leechers, ShouldBeBetween, 0, 1000000)
							So(search.Leechers, ShouldEqual, data.Response.Searches[i].Leechers)

							So(search.Date, ShouldNotBeBlank)
							So(search.Date, ShouldContainSubstring, ".")
							So(search.Date, ShouldContainSubstring, "'")
							So(search.Date, ShouldEqual, data.Response.Searches[i].Date)

							So(search.Size, ShouldNotBeBlank)
							So(search.Size, ShouldEqual, data.Response.Searches[i].Size)

							So(search.Uploader, ShouldNotBeBlank)
							So(search.Uploader, ShouldEqual, data.Response.Searches[i].Uploader)
						})
					}
				})
			})
		}
	})
}

func TestTorrent1337XMagnetLink(t *testing.T) {
	var testDatas = []struct {
		FolderName string
		File       string
		Magnet     string
	}{
		{"1337X", "page.html", "magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=Deadpool+2+%282018%29+%5BBluRay%5D+%5B1080p%5D+%5BYTS%5D+%5BYIFY%5D&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2F9.rarbg.com%3A2710%2Fannounce&tr=udp%3A%2F%2Fp4p.arenabg.com%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce"},
	}

	Convey("Torrent:1337X Fetch Magnet link test", t, func() {
		for _, data := range testDatas {
			currentTestPath := PATH_TEST + data.FolderName
			currentTestHome := currentTestPath + "/" + data.File

			file, err := os.Open(currentTestHome)

			So(file, ShouldNotBeNil)
			So(err, ShouldBeNil)

			defer file.Close()

			dataValid := IsMagnetURI(data.Magnet)

			So(dataValid, ShouldBeTrue)

			Convey("Fetching Magnet link from torrent page: "+data.File, func() {
				current, err := ParseMagnetString(services.GetDocumentFromFile(currentTestHome))

				So(err, ShouldBeNil)

				So(current, ShouldNotBeBlank)
				So(current, ShouldEqual, data.Magnet)
			})
		}
	})
}

func TestTorrent1337XMagnetValid(t *testing.T) {
	var testDatas = []struct {
		Magnet string
		Valid  bool
	}{
		{"magnet:?zt=urn:btih:E774B886539A3F7EBF1", false},
		{"magnet:?xt=urn:test:E774B886539A3F7EBF1", false},
		{"magnet:?xt=huh:btih:E774B886539A3F7EBF1", false},
		{"magnet:?xthuh:btih:E774B886539A3F7EBF1", false},
		{"magnet:?xthuh:btihE774B886539A3F7EBF1", false},
		{"magnet:?xthuhbtihE774B886539A3F7EBF1", false},
		{"magnet:xt=urn:btih:E774B886539A3F7EBF1", false},
		{"magnet:xturn:btih:E774B886539A3F7EBF1", false},
		{"magnet:xturnbtih:E774B886539A3F7EBF1", false},
		{"magnet:xturnbtihE774B886539A3F7EBF1", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=Test", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=&tr=Test", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=Test&tr=", false},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=Test&tr=Test", true},
		{"magnet:?xt=urn:btih:E774B886539A3F7EBF1FFE7CD01A107F73298248&dn=Deadpool+2+%282018%29+%5BBluRay%5D+%5B1080p%5D+%5BYTS%5D+%5BYIFY%5D&tr=udp%3A%2F%2Ftracker.coppersurfer.tk%3A6969%2Fannounce&tr=udp%3A%2F%2F9.rarbg.com%3A2710%2Fannounce&tr=udp%3A%2F%2Fp4p.arenabg.com%3A1337&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969&tr=udp%3A%2F%2Ftracker.internetwarriors.net%3A1337&tr=udp%3A%2F%2Ftracker.opentrackr.org%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.zer0day.to%3A1337%2Fannounce&tr=udp%3A%2F%2Ftracker.leechers-paradise.org%3A6969%2Fannounce&tr=udp%3A%2F%2Fcoppersurfer.tk%3A6969%2Fannounce",
			true},
	}

	Convey("Testing Magnet links", t, func() {
		for _, data := range testDatas {
			current := IsMagnetURI(data.Magnet)

			if data.Valid {
				So(current, ShouldBeTrue)
			} else {
				So(current, ShouldBeFalse)
			}
		}
	})
}
