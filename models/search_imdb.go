// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

// SearchRequest represents a movie search request
// that to be sent to a service
type SearchRequest struct {
	Title       string
	Year        string
	ID          string
	URL         string
	ExactSearch bool
	ResultLimit int
	ScanIMDB    bool
	ScanRT      bool
}

// SearchResult represents a movie in the search list
type SearchResult struct {
	URL    string
	ID     string
	Title  string
	Year   string
	TVShow bool
}

// SearchResponse represents a list of SearchResult
// that fetched from a service
type SearchResponse struct {
	SearchEngine string
	Searches     []SearchResult
	Error        string
	TotalFound   uint
}
