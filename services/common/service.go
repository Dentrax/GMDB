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

	"github.com/puerkitobio/goquery"
)

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
