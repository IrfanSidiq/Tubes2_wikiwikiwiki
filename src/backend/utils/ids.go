package scraper

import (
	"fmt"
)

func IDS(judulArtikelAwal string, judulArtikelTujuan string) {
	if (judulArtikelAwal == judulArtikelTujuan) {
		return
	}
	
	/* Deklarasi variabel */
	linkAsal := judulToLink(judulArtikelAwal)
	path := []string{}
	maxDepth := 0
	found := false

	for !found {
		maxDepth++
		fmt.Println("\nDepth:", maxDepth)
		fmt.Println()
		found = DFS(linkAsal, judulArtikelTujuan, maxDepth, &path)
		if (!found) {
			fmt.Println("\nSearch with maxDepth", maxDepth, "didn't find a result. Trying again with maxDepth", maxDepth+1, "...")
		} 
	}
	
	fmt.Print("\nFOUND PATH: \"" + path[0] + "\"")
	for i := 1; i < len(path); i++ {
		fmt.Print(" -> \"" + path[i] + "\"")
	}
	fmt.Println()
}

func DFS(currentLink string, judulArtikelTujuan string, depth int, path *[]string) bool {
	var nextLinks []string
	var currentJudul string
	linkScraping(currentLink, &nextLinks, &currentJudul)
	*path = append(*path, currentJudul)
	
	fmt.Print("SEARCHING: \"" + (*path)[0] + "\"")
	for i := 1; i < len(*path); i++ {
		fmt.Print(" -> \"" + (*path)[i] + "\"")
	}
	fmt.Println()
	
	if (currentJudul == judulArtikelTujuan) {
		return true
	}

	found := false
	idx := 0

	if (depth > 1) {
		for !found && idx < len(nextLinks) {
			found = DFS(nextLinks[idx], judulArtikelTujuan, depth - 1, path)
			idx++
		}
	}

	if (found) {
		return true
	}

	*path = (*path)[:len(*path)-1]
	return false
}