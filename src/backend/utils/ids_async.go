package scraper

import (
	"fmt"
	"sync"
	"golang.org/x/exp/slices"
)


var (
	fromLink          string
	toLink            string
	toTitle			  string
	resultPaths       chan []string
	done			  chan bool
	found             bool
	visitedLinks      stringBoolMap
	cache             IDSTree
)


func IDS(fromTitle string, judulArtikelTujuan string, singleSolution bool) (int, int, [][]string) {
	fromLink 	= judulToLink(fromTitle)
	toLink		= judulToLink(judulArtikelTujuan)
	toTitle     = judulArtikelTujuan
	resultPaths = make(chan []string, 1000)
	done		= make(chan bool)
	cache		= newIDSTree()
	found 		= false

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	go func() {
		maxDepth := 0
		for !found {
			fmt.Println("\nDepth:", maxDepth)
			fmt.Println()

			visitedLinks = newStringBoolMap()
			var wg sync.WaitGroup
			wg.Add(1)
			
			DLS(fromLink, maxDepth, []string{}, &wg)
			
			if (!found) {
				fmt.Println("\nSearch with maxDepth", maxDepth, "didn't find a result. Trying again with maxDepth", maxDepth+1, "...")
				maxDepth++
			} else if (!singleSolution) {
				done <- true
			}
		}
	} ()

	// If searching for multiple paths, wait until all goroutines finishes
	if (!singleSolution) {
		<- done
	}

	
	var result [][]string
	result = append(result, <- resultPaths)
	close(resultPaths)
	close(done)

	for i := 0; i < len(resultPaths); i++ {
		var path = <- resultPaths
		result = append(result, path)
	}

	return len(visitedLinks.Map), len(result[0]) - 1, result
}


func DLS(currentLink string, depth int, path []string, wg *sync.WaitGroup) {
	defer func() {
		if err := recover(); err != nil {
			// do nothing
		}
	}()
	defer wg.Done()

	currentJudul := linkTojudul(currentLink)
	
	if (visitedLinks.get(currentJudul)) {
		return
	}
	visitedLinks.set(currentJudul)

	var nextLinks []string = cache.get(currentJudul)
	if (len(nextLinks) == 0) {
		linkScraping(currentLink, &nextLinks, &currentJudul)
		cache.set(currentJudul, nextLinks)
	}
	
	path = append(path, currentJudul)

	// fmt.Print("SEARCHING: \"" + path[0] + "\"")
	// for i := 1; i < len(path); i++ {
	// 	fmt.Print(" -> \"" + path[i] + "\"")
	// }
	// fmt.Println()

	if (slices.Contains(nextLinks, toLink)) {
		path = append(path, toTitle)
		found = true
		resultPaths <- path

		return
	}
	
	if (currentJudul == toTitle) {
		found = true
		resultPaths <- path

		return
	}

	if (depth == 0) {
		return
	}

	xthreads := 4
	var linkChannel = make(chan string, len(nextLinks))
	var currentWg sync.WaitGroup
	currentWg.Add(len(nextLinks) + xthreads)
	

	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				nextLink, isOpen := <- linkChannel
				if (!isOpen) {
					currentWg.Done()
					return
				}
				DLS(nextLink, depth - 1, path, &currentWg)
			}
		} ()
	}
	
	for _, nextLink := range nextLinks {
		linkChannel <- nextLink
	}

	close(linkChannel)
	currentWg.Wait()
}