// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

import "github.com/middelink/go-parse-torrent-name"

// SearchTorrentRequest represents a torrent search request
type SearchTorrentRequest struct {
	Title           string
	Year            string
	ResultLimit     int
	ScanTorrentsCSV bool
	Scan1337x       bool
}

// SearchTorrentResult represents a torrent search result
type SearchTorrentResult struct {
	Name     string
	URL      string
	Seeders  string
	Leechers string
	Date     string
	Size     string
	Uploader string
	Info     parsetorrentname.TorrentInfo
}

// SearchTorrentResponse represents a list of SearchTorrentResult
// that fetched from a torrent service
type SearchTorrentResponse struct {
	SearchEngine string
	Searches     []SearchTorrentResult
	Error        string
	TotalFound   uint
}
