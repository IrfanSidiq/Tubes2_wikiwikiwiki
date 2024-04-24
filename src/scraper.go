package main

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

/*
	Ambil semua link & judul dari bagian artikel halaman
*/
func linkScraping(link string, links *[]string, judul *string) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org", "id.wikipedia.org"),
	)

	c.OnHTML("div[id=mw-content-text] p a[href]", func(h *colly.HTMLElement) {
		var link string = h.Attr("href")
		// Ambil link wikipedia yang bukan file / template / wikipedia: 
		if (strings.HasPrefix(link, "/wiki/") && 
			!strings.HasPrefix(link, "/wiki/File:") && 
			!strings.HasPrefix(link, "/wiki/Template:") &&
			!strings.HasPrefix(link, "/wiki/Wikipedia:")) {
			link = "https://en.wikipedia.org" + link 
			*links = append(*links, link)
		}
	})

	c.OnHTML("h1[id=firstHeading]", func(h *colly.HTMLElement) {
		*judul = h.Text
	})

	c.Visit(link)
}

/*
	Encode judul untuk dipake di url
*/
func convertJudul(judul string) string {
	encoded := url.QueryEscape(judul)
	encoded = strings.ReplaceAll(encoded, "+", "_")
	return encoded
}

/*
	Ubah judul menjadi link en.wikipedia.org
*/
func judulToLink(judul string) string {
	encoded := convertJudul(judul)
	return "https://en.wikipedia.org/wiki/" + encoded
}

/*
	Ubah link menjadi judul artikel
*/
func linkTojudul(link string) string {
	encoded := strings.TrimPrefix(link, "https://en.wikipedia.org/wiki/")
	encoded = strings.ReplaceAll(encoded, "_", "+")
	decoded, err := url.QueryUnescape(encoded)
	if err == nil {
		return decoded
	} else {
		// harusnya ga pernah gagal
		return "GAGAL DECODE"
	}
}

func main() {
	/* Terima Input */

	// var jenisAlgoritma int
	var judulArtikelAwal string = "Adolf Hitler"
	var judulArtikelTujuan string = "Stardew Valley"

	/* Algoritma */
	startTime := time.Now()

	// a := Tree{
	// 	prev: nil,
	// 	link:	judulToLink(judulArtikelAwal),
	// 	nextArr: []*Tree{},
	// 	depth: 0,
	// }
	
	BFS_Async(judulArtikelAwal, judulArtikelTujuan)

	// IDS_async(judulArtikelAwal, judulArtikelTujuan)
	endTime := time.Now()
	duration := endTime.Sub(startTime)

	fmt.Print("Runtime: ")
	fmt.Println(duration)
}