package main

import (
	"fmt"
	"sync"
)

func initTreeAsync(currentTree *Tree) bool {
	var links []string
	
	// Scraping link & judul
	linkScraping(currentTree.link, &links, &currentTree.judul)
	
	// Cek apakah halaman sudah pernah di-visit
	if visitedAsync.keyExists(currentTree.judul) {
		return false
	}

	// Masukkan link ke dalam nextArr
	for _, link := range links {
		newTree := Tree{
			prev: currentTree,
			link:    link,
			nextArr: []*Tree{},
			depth:   currentTree.depth + 1,
		}
		currentTree.nextArr = append(currentTree.nextArr, &newTree)
	}
	return true
}

func BFS_Async(from string, to string) {
	// Return jika judul asal dan tujuan sama
	if (from == to) {
		return
	}

	// Init
	linkAsal = judulToLink(from)
	setFound(false)
	queueAsync = newTreeArr()
	visitedAsync = newStringBoolMap()

	var wg sync.WaitGroup

	a := Tree{
		prev: nil,
		link:	linkAsal,
		nextArr: []*Tree{},
		depth: 0,
	}

	queueAsync.apd(&a)

	// Start
	go func() {
		var cnt int = 0
		for !getFound() {
			cnt++
			wg.Add(1)
			
			currentTree := queueAsync.getFirstElement()
			queueAsync.removeFirstElement()
			
			fmt.Println(currentTree.link)
			
			// Return jika sudah di visit / sudah ketemu
			isNotVisited := initTreeAsync(currentTree)
			if (!isNotVisited || getFound()) {
				return
			} else {
				visitedAsync.set(currentTree.link)
			}
				
			// Set found ke true kalo judul sama
			if currentTree.judul == to {
				setFound(true)
				output(currentTree, cnt)
			}
			
			for _, val := range currentTree.nextArr {
				if !visitedAsync.keyExists(val.judul) {
					queueAsync.apd(val)
				}
			}

			defer wg.Done()
		}
	} ()
	
	wg.Wait()
}