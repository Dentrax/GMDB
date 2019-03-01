// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

//A MyListRequest taken from CLI
type MyListRequest struct {
	NOP     bool
	ScanWLs bool
	ScanMLs bool
}

//MyListResponse for printing
type MyListResponse struct {
	WL         MovieWatchLaterInfo
	ML         MovieLearnInfo
	MovieTitle string
	MovieYear  string
}
