// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

type ParentsGuide struct {
	TotalVote   uint
	Nudity      SeverityRate
	Violence    SeverityRate
	Profanity   SeverityRate
	Alcohol     SeverityRate
	Frightening SeverityRate
}
