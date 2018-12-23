// ====================================================
// GMDB Copyright(C) 2019 Furkan Türkal
// This program comes with ABSOLUTELY NO WARRANTY; This is free software,
// and you are welcome to redistribute it under certain conditions; See
// file LICENSE, which is part of this source code package, for details.
// ====================================================

package services

import (
	"log"
	"net/http"
	"os"
	"unicode/utf8"

	"gmdb/pkg/cache"
	"gmdb/pkg/config"

	"github.com/puerkitobio/goquery"
)

func GetDocumentFile(service string, rootName string, title string, url string) *goquery.Document {
	var doc *goquery.Document

	if cache.IsFileExist(service, rootName, title) {
		result, err := cache.GetFile(service, rootName, title)
		if err != nil {
			log.Fatal(err)
		}
		doc, err = goquery.NewDocumentFromReader(result.File)
		if err != nil {
			log.Fatal(err)
		}
		return doc
	}

	doc = GetDocumentFromURL(url)

	if config.Cache.UseCache {
		if config.Cache.UseMovieCache {
			str, err := doc.Html()
			if err != nil {
				log.Fatal(err)
			}
			cache.WriteFile(service, rootName, title, str)
			//TODO: else timeout 1 day write
		}
	}
	return doc
}

func GetDocumentFromURL(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return doc
}

func GetDocumentFromFile(filename string) *goquery.Document {
	file, e := os.Open(filename)
	if e != nil {
		log.Fatal(e)
		return nil
	}

	defer file.Close()
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatal(e)
		return nil
	}
	if !utf8.ValidString(doc.Text()) {
		log.Fatalf("DOC: %s", "NOT UTF-8 FORMAT")
		return nil
	}
	return doc
}
