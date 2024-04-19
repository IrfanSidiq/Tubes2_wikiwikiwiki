package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

/*
	Ambil judul halaman
*/
func getJudul(link string, callback func(string)) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org", "id.wikipedia.org"),
	)
	
	c.OnHTML("h1[id=firstHeading]", func(h *colly.HTMLElement) {
		callback(h.Text)
		return
	})

	c.Visit(link)
}

/*
	Ambil semua link dari bagian artikel halaman
*/
func linkScraping(link string) []string {
	var links []string

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org", "id.wikipedia.org"),
	)

	c.OnHTML("div[id=mw-content-text] a[href]", func(h *colly.HTMLElement) {
		var link string = h.Attr("href")
		// Ambil link wikipedia yang bukan file
		if (strings.HasPrefix(link, "/wiki/") && !strings.HasPrefix(link, "/wiki/File:")) {
			fmt.Println(link)
			link = "https://en.wikipedia.org" + link 
			links = append(links, link)
		}
	})

	c.Visit(link)

	return links
}

func main() {
	/* Terima Input */

	// var jenisAlgoritma int
	// var judulArtikelAwal string = "aaa"
	var judulArtikelTujuan string = "World War I"

	/* Algoritma */
	a := Tree{
		link:	"https://en.wikipedia.org/wiki/World_War_II",
		nextArr: []*Tree{},
	}
	
	BFS(0, a, judulArtikelTujuan)
}