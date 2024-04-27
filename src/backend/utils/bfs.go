package scraper

import (
	"fmt"
	"sync"
)

func initTree(currentTree *Tree, c chan Tree) bool {
	if getFound() {
		return false
	} 

	var links []string
	
	// Scraping link & judul
	linkScraping(currentTree.link, &links, &currentTree.judul)
	// Cek apakah halaman sudah pernah di-visit
	if visitedAsync.keyExists(currentTree.link) {
		return false
	}
	
	// Masukkan semua link yang di-scrape ke dalam nextArr
	for _, link := range links {
		newTree := Tree{
			prev: append(currentTree.prev, currentTree.judul),
			link:    link,
			nextArr: []*Tree{},
			depth:   currentTree.depth + 1,
		}
		
		// Jika link sama dengan link yang di cari, return
		if link == linkTujuan && !getFound() {
			setFound(true)

			newTree.judul = linkTojudul(linkTujuan)
			select {
			case c <- newTree:
			default:
			}
			// if (len(c) < 1) {
			// 	c <- newTree
			// }

			return false
		}

		currentTree.nextArr = append(currentTree.nextArr, &newTree)
	}
	return true
}

func BFS(from string, to string) (int, int, [][]string) {
	// Return jika judul asal dan tujuan sama
	if (from == to) {
		fmt.Println("Judul awal dan judul tujuan harus berbeda!")
		return 0, 0, [][]string{}
	}

	// Inisialisasi
	var wg sync.WaitGroup
	
	setFound(false)
	linkAsal = judulToLink(from)
	linkTujuan = judulToLink(to)
	
	visitedAsync = newStringBoolMap()
	
	a := Tree{
		prev: []string{},
		link:	linkAsal,
		nextArr: []*Tree{},
		depth: 0,
	}
	
	queueA := make(chan *Tree, 10000000)
	resultTree := make(chan Tree, 1)
	
	initTree(&a, resultTree)
	for _, val := range a.nextArr {
		queueA <- val
	}

	visitedAsync.set(linkAsal)
	
	// Start
	setCnt(0)
	go func () {
		for !getFound() {
			wg.Add(100)
			for i := 0; i < 100; i++ {
				go func() {
					defer wg.Done()
					
					currentTree, isOpen := <- queueA
					if (!isOpen) {
						return
					}
	
					// Return jika sudah di visit / sudah ketemu
					isNotVisited := initTree(currentTree, resultTree)
					if !isNotVisited {
						return
					} else if getFound() {
						return
					} else {
						visitedAsync.set(currentTree.link)
						incCnt()
					}
					
					// Print
					fmt.Print("Searching: ")
					for _, val := range currentTree.prev {
						fmt.Print(val)
						fmt.Print(" -> ")
					}
					fmt.Println(currentTree.judul, len(queueA), len(visitedAsync.Map)) 
					
					// Set found ke true kalo judul sama dg to
					if (currentTree.judul == to) && !getFound() {
						setFound(true)
						select {
						case resultTree <- *currentTree:
						default:
						}
						// if (len(resultTree) < 1) {
						// 	resultTree <- *currentTree
						// }
						return
					}
					
					// Masukkan link dalam halaman kedalam queue, asumsi queue tidak akan penuh
					for _, val := range currentTree.nextArr {
						if !visitedAsync.keyExists(val.link) {
							select {
							case queueA <- val:
							default: 
								fmt.Println("Channel penuh")
							}
						}
					}
					
				} ()
			}
	
			wg.Wait()
		}
	} ()
	
	// Output
	resT := <- resultTree
	path := append(resT.prev, resT.judul)

	return cntAsync, resT.depth, [][]string{path}
}