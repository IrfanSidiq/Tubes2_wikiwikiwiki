package main

import (
	"fmt"
	"sync"
	"golang.org/x/exp/slices"
)

func IDS_async(fromTitle string, toTitle string) {
	if (fromTitle == toTitle) {
		fmt.Println("Judul artikel asal dan tujuan harus beda!")
		return
	}

	fromLink := judulToLink(fromTitle)
	resultPath := make(chan []string)
	found := false

	var wg sync.WaitGroup
	var visitedLinks stringBoolMap
	var cache IDSTree = newIDSTree()

	go func() {
		maxDepth := 0
		for !found {
			maxDepth++
			wg.Add(1)
			visitedLinks = newStringBoolMap()
			fmt.Println("\nDepth:", maxDepth - 1)
			fmt.Println()
			DLS(fromLink, toTitle, 0, maxDepth - 1, []string{}, resultPath, &found, &wg, &visitedLinks, &cache)
			if (!found) {
				fmt.Println("\nSearch with maxDepth", maxDepth, "didn't find a result. Trying again with maxDepth", maxDepth+1, "...")
			}
		}
	} ()

	result := <- resultPath
	fmt.Print("\nFound path              : \"" + result[0] + "\"")
	for i := 1; i < len(result); i++ {
		fmt.Print(" -> \"" + result[i] + "\"")
	}
	fmt.Println()
	fmt.Println("Number of links visited :", len(visitedLinks.Map))
	fmt.Println("Route length            :", len(result))
}

func DLS(fromLink string, toTitle string, depth int, maxDepth int, path []string, resultChan chan []string, found *bool, wg *sync.WaitGroup, visitedLinks *stringBoolMap, cache *IDSTree) {
	defer wg.Done()

	currentJudul := linkTojudul(fromLink)
	if (visitedLinks.get(currentJudul)) {
		return
	}
	visitedLinks.set(currentJudul)

	var nextLinks []string = cache.get(currentJudul)
	if (len(nextLinks) == 0) {
		linkScraping(fromLink, &nextLinks, &currentJudul)
		cache.set(currentJudul, nextLinks)
	}
	
	path = append(path, currentJudul)

	fmt.Print("SEARCHING: \"" + path[0] + "\"")
	for i := 1; i < len(path); i++ {
		fmt.Print(" -> \"" + path[i] + "\"")
	}
	fmt.Println()

	if (slices.Contains(nextLinks, judulToLink(toTitle))) {
		path = append(path, toTitle)
		*found = true
		resultChan <- path
		return
	}
	
	if (currentJudul == toTitle) {
		*found = true
		resultChan <- path
		return
	}

	if (depth + 1 > maxDepth) {
		return
	}

	var linkChannel = make(chan string, len(nextLinks))
	var currentWg sync.WaitGroup
	xthreads := 4

	currentWg.Add(len(nextLinks) + xthreads)
	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				nextLink, isOpen := <- linkChannel
				if (!isOpen) {
					currentWg.Done()
					return
				}
				DLS(nextLink, toTitle, depth + 1, maxDepth, path, resultChan, found, &currentWg, visitedLinks, cache)
			}
		} ()
	}
	
	for _, nextLink := range nextLinks {
		linkChannel <- nextLink
	}

	close(linkChannel)
	currentWg.Wait()
}


/* ----- Helper Data Types ----- */

type stringBoolMap struct {
	Map map[string]bool
	sync.RWMutex
}

func newStringBoolMap() stringBoolMap {
	return stringBoolMap{map[string]bool{}, sync.RWMutex{}}
}

func (m *stringBoolMap) get(key string) bool {
	m.RLock()
	defer m.RUnlock()
	return m.Map[key]
}

func (m *stringBoolMap) set(key string) {
	m.Lock()
	defer m.Unlock()
	m.Map[key] = true
}

type IDSTree struct {
	Map map[string][]string
	sync.RWMutex
}

func newIDSTree() IDSTree {
	return IDSTree{map[string][]string{}, sync.RWMutex{}}
}

func (m *IDSTree) get(key string) []string {
	m.RLock()
	defer m.RUnlock()
	return m.Map[key]
}

func (m *IDSTree) set(key string, value []string) {
	m.Lock()
	defer m.Unlock()
	m.Map[key] = value
}