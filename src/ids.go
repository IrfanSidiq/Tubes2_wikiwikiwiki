package main

import (
	"fmt"
)

func IDS(judulArtikelAwal string, judulArtikelTujuan string) {
	if (judulArtikelAwal == judulArtikelTujuan) {
		return
	}
	
	/* Deklarasi variabel */
	linkAsal := judulToLink(judulArtikelAwal)
	nextLinks := []string{}
	path := []string{}
	var currentJudul string
	maxDepth := 0
	found := false

	linkScraping(linkAsal, &nextLinks, &currentJudul)

	for !found {
		maxDepth++
		depth := 1
		for !found && depth <= maxDepth {
			fmt.Println("\nDepth:", depth)
			found = DFS(linkAsal, judulArtikelTujuan, depth, &path)
			fmt.Println()
			depth++
		}
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


/* ----- Helper type untuk map visited ---- */

// type stringBoolMap struct {
// 	Map map[string]bool
// 	sync.RWMutex
// }

// func newStringBoolMap() stringBoolMap {
// 	return stringBoolMap{map[string]bool{}, sync.RWMutex{}}
// }

// func (m *stringBoolMap) get(key string) bool {
// 	m.RLock()
// 	defer m.RUnlock()
// 	return m.Map[key]
// }

// func (m *stringBoolMap) set(key string) {
// 	m.Lock()
// 	defer m.Unlock()
// 	m.Map[key] = true
// }



	// for depth > 0 {
	// 	currentLink := stack[len(stack)-1]
	// 	stack = stack[:len(stack)-1]

	// 	var currentJudul string
	// 	nextLinks := []string{}
	// 	linkScraping(currentLink, &nextLinks, &currentJudul)

	// 	if (currentJudul == judulArtikelTujuan) {
	// 		resultChannel <- true
	// 		return
	// 	}

	// 	if (!visited.get(currentJudul)) {
	// 		visited.set(currentJudul)
	// 		stack = append(stack, nextLinks...)
	// 	}
		
	// }