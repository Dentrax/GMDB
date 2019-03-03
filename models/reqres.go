// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

// UpdateRequest represents a connection between CLI and Printer
type UpdateRequest struct {
	NOP bool
}

// UpdateResponse stores an information about list of MovieInfo
type UpdateResponse struct {
	UpdateList []MovieInfo
}

// MyListRequest represents a connection between CLI and Printer
type MyListRequest struct {
	NOP     bool
	ScanWLs bool
	ScanMLs bool
}

// MyListResponse stores an information about both of MovieWatchLaterInfo and MovieLearnInfo responses
type MyListResponse struct {
	WL         MovieWatchLaterInfo
	ML         MovieLearnInfo
	MovieTitle string
	MovieYear  string
}

// NoteRequest represents a connection between CLI and Printer
type NoteRequest struct {
	NOP bool
}

// NoteResponse stores an information about MovieNoteInfo response
type NoteResponse struct {
	Note       MovieNoteInfo
	MovieTitle string
	MovieYear  string
}

// HistoryRequest represents a connection between CLI and Printer
type HistoryRequest struct {
	NOP          bool
	ScanSearches bool
	ScanWatches  bool
}

// HistoryResponse stores an information about MovieSearchInfo response
type HistoryResponse struct {
	Search     MovieSearchInfo
	MovieTitle string
	MovieYear  string
}

// LearnRequest represents a connection between CLI and Printer
type LearnRequest struct {
	Filename string
}

// LearnResponse stores an information about LearnResult
type LearnResponse struct {
	Success bool
	Error   string
	Result  LearnResult
}

// LearnResult stores an information about movie
type LearnResult struct {
	Title      string
	IsTVSeries bool
	WatchDate  string
}
