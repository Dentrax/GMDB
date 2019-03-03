// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

type LearnRequest struct {
	Filename string
}

type LearnResponse struct {
	Success bool
	Error   string
	Result  LearnResult
}

type LearnResult struct {
	Title      string
	IsTVSeries bool
	WatchDate  string
}
