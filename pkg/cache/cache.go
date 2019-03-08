// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package cache

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gmdb/models"
)

const pathRoot string = "./cache/gmdb"
const fileExt string = ".cache"

func IsFileExist(service string, rootName string, title string) bool {
	path := pathRoot + "/" + service + "/" + rootName + "/" + title + fileExt
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func GetSearchResponse(service string, rootName string, title string) (*models.SearchResponse, error) {
	path := pathRoot + "/" + service + "/" + rootName + "/" + title + fileExt
	result := new(models.SearchResponse)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	json.Unmarshal(file, &result)

	return result, nil
}

func GetMovie(service string, rootName string, title string) (*models.Movie, error) {
	path := pathRoot + "/" + service + "/" + rootName + "/" + title + fileExt
	result := new(models.Movie)

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	json.Unmarshal(file, &result)

	return result, nil
}

func GetFile(service string, rootName string, title string) (*models.CacheData, error) {
	path := pathRoot + "/" + service + "/" + rootName + "/" + title + fileExt
	result := new(models.CacheData)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	result.Service = service
	result.Title = title
	result.File = file

	return result, nil
}

func WriteFile(service string, rootName string, title string, data string) bool {
	path := pathRoot + "/" + service + "/" + rootName + "/"
	filepath := path + title + fileExt

	if err := os.MkdirAll(path, 0755); err != nil {
		log.Fatalln(err)
		return false
	}

	file, err := os.Create(filepath)

	if err != nil {
		return false
	}

	_, err = io.Copy(file, strings.NewReader(data))
	if err != nil {
		return false
	}

	defer file.Close()

	return true
}
