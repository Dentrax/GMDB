// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package torrent

import (
	"log"

	"github.com/Dentrax/GMDB/models"
	"github.com/Dentrax/GMDB/services/common"
	"github.com/Dentrax/GMDB/services/torrent/torrent1337x"

	"github.com/puerkitobio/goquery"
)

type Torrent struct {
	Name    string
	Request models.SearchTorrentRequest
}

func New(name string, request models.SearchTorrentRequest) *Torrent {
	return &Torrent{
		Name:    name,
		Request: request,
	}
}

func (s *Torrent) SearchTorrent(request *models.SearchTorrentRequest) *models.SearchTorrentResponse {
	url := "https://1337x.to/search/" + request.Title + "/1/"

	doc := services.GetDocumentFromURL(url)
	result, err := GetSearchTorrentsFrom1337x(doc)

	if err != nil {
		log.Fatalln("nil")
	}

	return result
}

func (s *Torrent) GetMagnet(result *models.SearchTorrentResult) (string, error) {
	url := "https://1337x.to" + result.URL
	doc := services.GetDocumentFromURL(url)
	magnet, err := GetMagnetStringFrom1337x(doc)
	return magnet, err
}

func GetSearchTorrentsFrom1337x(doc *goquery.Document) (*models.SearchTorrentResponse, error) {
	searches := torrent1337x.ParseSearchTorrents(doc)
	return searches, nil
}

func GetMagnetStringFrom1337x(doc *goquery.Document) (string, error) {
	result, err := torrent1337x.ParseMagnetString(doc)
	return result, err
}
