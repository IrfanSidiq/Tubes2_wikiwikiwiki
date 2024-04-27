package scraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

/*
	Ambil semua link & judul dari bagian artikel halaman
*/
func linkScraping(link string, links *[]string, judul *string) {
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org", "id.wikipedia.org"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnHTML("div[id=mw-content-text] p a[href]", func(h *colly.HTMLElement) {
		var link string = h.Attr("href")
		// Ambil link wikipedia kecuali yang bukan artikel 
		if (strings.HasPrefix(link, "/wiki/") && 
			!strings.HasPrefix(link, "/wiki/File:") && 
			!strings.HasPrefix(link, "/wiki/Template:") &&
			!strings.HasPrefix(link, "/wiki/Wikipedia:")) &&
			!strings.HasPrefix(link, "/wiki/Help:") &&
			!strings.HasPrefix(link, "/wiki/Portal:") {
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
func LinkTojudul(link string) string {
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

// func main() {
// 	/* Terima Input */

// 	// var jenisAlgoritma int
// 	var judulArtikelAwal string = "Joko Widodo"
// 	var judulArtikelTujuan string = "Atheism"

// 	/* Algoritma */
// 	startTime := time.Now()

// 	BFS_Async(judulArtikelAwal, judulArtikelTujuan)
// 	// IDS_async(judulArtikelAwal, judulArtikelTujuan)

// 	duration := time.Since(startTime)

// 	fmt.Print("Runtime: ")
// 	fmt.Println(duration)
// }