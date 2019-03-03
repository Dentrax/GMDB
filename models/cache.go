// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package models

import (
	"os"
	"time"
)

// CacheData stores a cache info for a service
type CacheData struct {
	Service string
	Title   string
	File    *os.File
	Time    *time.Duration
}
