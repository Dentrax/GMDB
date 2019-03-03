// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

// ResultFilter represents a CLI output filter
type ResultFilter struct {
	NoBanner     bool
	NoColor      bool
	Title        bool
	Year         bool
	Released     bool
	Rating       bool
	Duration     bool
	Summary      bool
	Directors    bool
	Writers      bool
	Stars        bool
	Genres       bool
	Tagline      bool
	Summaries    bool
	Keywords     bool
	ParentsGuide bool
	ShowWLs      bool
	ShowMLs      bool
}
