// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

type AppConfig struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string

	TimeFormat string
}

var App = &AppConfig{}

type CacheConfig struct {
	UseCache        bool
	UseSearchCache  bool
	UseMovieCache   bool
	UseTrailerCache bool
}

var Cache = &CacheConfig{}

type Config struct {
	App   *AppConfig
	Cache *CacheConfig
}

func LoadConfig() (*Config, error) {
	path := "config/gmdb.conf"
	cfg, err := ini.Load(path)

	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return nil, err
	}

	configApp := &AppConfig{
		LogSavePath: cfg.Section("log").Key("path").String(),
		LogSaveName: cfg.Section("log").Key("name").String(),
		LogFileExt:  cfg.Section("log").Key("ext").String(),
		TimeFormat:  cfg.Section("log").Key("format").String(),
	}

	useCache, err1 := cfg.Section("cache").Key("use-cache").Bool()
	useSearchCache, err2 := cfg.Section("cache").Key("search-cache").Bool()
	useMovieCache, err3 := cfg.Section("cache").Key("movie-cache").Bool()
	useTrailerCache, err4 := cfg.Section("cache").Key("trailer-cache").Bool()

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		fmt.Printf("[Config] cache parse error: %v", err)
		return nil, err
	}

	configCache := &CacheConfig{
		UseCache:        useCache,
		UseSearchCache:  useSearchCache,
		UseMovieCache:   useMovieCache,
		UseTrailerCache: useTrailerCache,
	}

	App = configApp
	Cache = configCache

	return &Config{
		App:   configApp,
		Cache: configCache,
	}, nil
}
