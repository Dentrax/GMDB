// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

type MovieInfo struct {
	Title             string
	Year              string
	Rating            string
	Votes             string
	Reviews           string
	Genres            []string
	Duration          string
	Released          string
	Summary           string
	Credit            CreditInfo
	Metascore         string
	ReviewCountUser   string
	ReviewCountCritic string
	RTMeter           string
	URLTrailerIMDB    string
	URLPosterIMDB     string
}

type Movie struct {
	Info MovieInfo
	TL   Tagline
	PS   PlotSummary
	PK   PlotKeywords
	PG   ParentsGuide
}
