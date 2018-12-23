// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

type RTSearchResult struct {
	ActorCount     uint                `json:"actorCount"`
	Actors         []RTActorInfo       `json:"actors"`
	CriticCount    uint                `json:"criticCount"`
	FranchiseCount uint                `json:"franchiseCount"`
	MovieCount     uint                `json:"movieCount"`
	MovieInfos     []RTSearchMovieInfo `json:"movies"`
	TVSeries       []RTTVSeriesInfo    `json:"tvSeries"`
	TVCount        uint                `json:"tvCount"`
}

type RTActorInfo struct {
	Name  string `json:"name"`
	URL   string `json:"url"`
	Image string `json:"image"`
}

type RTSearchMovieInfo struct {
	Name       string                  `json:"name"`
	Year       uint                    `json:"year"`
	URL        string                  `json:"url"`
	Image      string                  `json:"image"`
	MeterClass string                  `json:"meterClass"`
	MeterScore uint                    `json:"meterScore"`
	Casts      []RTSearchMovieCastItem `json:"castItems"`
	Subline    string                  `json:"subline"`
}

type RTSearchMovieCastItem struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type RTTVSeriesInfo struct {
	Title      string `json:"title"`
	StartYear  uint   `json:"startYear"`
	EndYear    uint   `json:"endYear"`
	URL        string `json:"url"`
	MeterClass string `json:"meterClass"`
	Image      string `json:"image"`
}
