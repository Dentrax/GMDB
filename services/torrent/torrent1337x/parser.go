// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package torrent1337x

import (
	"errors"
	"regexp"
	"strings"

	"github.com/Dentrax/GMDB/models"

	"github.com/middelink/go-parse-torrent-name"
	"github.com/puerkitobio/goquery"
)

func FixSpace(input string) string {
	input = strings.Replace(input, "<br> \n", "", -1)
	input = strings.TrimSpace(input)
	input = strings.Replace(input, " ", " ", -1)
	input = strings.Replace(input, "»", "", -1)
	return input
}

func IsMagnetURI(input string) bool {
	rgxMagnet, _ := regexp.Compile("magnet:\\?xt=urn:[a-z0-9]+:[a-zA-Z0-9]{32,40}&dn=.+&tr=.+")
	return rgxMagnet.MatchString(input)
}

func ParseSearchTorrents(doc *goquery.Document) *models.SearchTorrentResponse {
	result := new(models.SearchTorrentResponse)
	finder := doc.Find("body > main > div > div > div > div.box-info-detail.inner-table > div.table-list-wrap > table > tbody")

	count := 0
	if len(finder.Nodes) > 0 {
		doc.Find("tr").Each(func(i int, s *goquery.Selection) {
			colName := s.Find("td.coll-1.name > a")

			name := colName.NextFilteredUntil("a", "span")
			title := FixSpace(name.Text())
			url, ok := name.Attr("href")

			colSeeds := s.Find("td.coll-2.seeds")
			colLeechers := s.Find("td.coll-3.leeches")
			colDate := s.Find("td.coll-date")
			colSize1 := s.Find("td.coll-4.size.mob-uploader").Children().Remove().End()
			colSize2 := s.Find("td.coll-4.size.mob-user").Children().Remove().End()
			colSize3 := s.Find("td.coll-4.size.mob-vip").Children().Remove().End()
			colUploader1 := s.Find("td.coll-5.uploader")
			colUploader2 := s.Find("td.coll-5.user")
			colUploader3 := s.Find("td.coll-5.vip")

			var colSize string
			var colUploader string

			if colSize1.Text() != "" {
				colSize = colSize1.Text()
			} else if colSize2.Text() != "" {
				colSize = colSize2.Text()
			} else if colSize3.Text() != "" {
				colSize = colSize3.Text()
			} else {
				colSize = ""
			}

			if colUploader1.Text() != "" {
				colUploader = colUploader1.Text()
			} else if colUploader2.Text() != "" {
				colUploader = colUploader2.Text()
			} else if colUploader3.Text() != "" {
				colUploader = colUploader3.Text()
			} else {
				colUploader = ""
			}

			if ok {
				info, _ := parsetorrentname.Parse(title)
				item := models.SearchTorrentResult{
					Info:     *info,
					Name:     title,
					URL:      url,
					Seeders:  FixSpace(colSeeds.Text()),
					Leechers: FixSpace(colLeechers.Text()),
					Date:     FixSpace(colDate.Text()),
					Size:     FixSpace(colSize),
					Uploader: FixSpace(colUploader),
				}
				result.Searches = append(result.Searches, item)
				count = count + 1
			}
		})
	}

	result.SearchEngine = "1337x.to"
	result.TotalFound = uint(count)

	return result
}

func ParseMagnetString(doc *goquery.Document) (string, error) {
	finder := doc.Find("body > main > div > div > div > div:nth-child(2) > div > ul > li > a")
	if len(finder.Nodes) > 0 {
		text, exist := finder.Attr("href")
		if exist {
			rgxMagnet, _ := regexp.Compile("magnet:\\?xt=urn:[a-z0-9]+:[a-zA-Z0-9]{32,40}&dn=.+&tr=.+")
			return rgxMagnet.FindString(text), nil
		}
	}
	return "", errors.New("Magnet not found")
}
