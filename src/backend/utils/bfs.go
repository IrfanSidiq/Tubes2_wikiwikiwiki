package scraper

import (
	"fmt"
	"sync"
)

func initTree(currentTree *Tree, c chan Tree, singleSolution bool) bool {
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
			prev: append(currentTree.prev, currentTree.link),
			link:    link,
			nextArr: []*Tree{},
			depth:   currentTree.depth + 1,
		}
		
		// Jika link sama dengan link yang di cari, return
		if link == linkTujuan {
			newTree.judul = LinkTojudul(linkTujuan)
			if singleSolution {
				if !getFound() {
					setFound(true)
		
					select {
					case c <- newTree:
					default:
					}
				}
			} else {
				if len(c) < 1 {
					if getPJ() == -99  {
						setPJ(newTree.depth)
					}
					select {
					case c <- newTree:
					default:
					}
				} else {
					if getPJ() == newTree.depth {
						select {
						case c <- newTree:
						default:
						}
					}
				}
			}

			return false
		}

		currentTree.nextArr = append(currentTree.nextArr, &newTree)
	}
	return true
}

func BFS(from string, to string, singleSolution bool) (int, int, [][]string) {
	// Return jika judul asal dan tujuan sama
	if (from == to) {
		fmt.Println("Judul awal dan judul tujuan harus berbeda!")
		return 0, 0, [][]string{}
	}

	// Inisialisasi
	var wg sync.WaitGroup
	
	setFound(false)
	setDone(false)
	panjangRute = -99
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
	resultTree := make(chan Tree, 10000)
	
	initTree(&a, resultTree, singleSolution)
	for _, val := range a.nextArr {
		queueA <- val
	}

	visitedAsync.set(linkAsal)
	
	// Start
	setCnt(0)
	for {
		wg.Add(100)
		for i := 0; i < 100; i++ {
			go func() {
				defer wg.Done()
				
				currentTree, isOpen := <- queueA
				if (!isOpen) {
					return
				}
				
				// Return jika sudah di visit / sudah ketemu
				isNotVisited := initTree(currentTree, resultTree, singleSolution)
				if !isNotVisited {
					return
				} else {
					visitedAsync.set(currentTree.link)
					incCnt()

					if singleSolution {
						if getFound() {
							return
						}
					} else {
						if getDone() {
							return
						} 

						if getPJ() != -99 {
							if currentTree.depth >= getPJ() {
								setDone(true)
								return
							}
						}
					}
				}
				
				// Print
				fmt.Print("Searching: ")
				for _, val := range currentTree.prev {
					fmt.Print(val)
					fmt.Print(" -> ")
				}
				fmt.Println(currentTree.link, len(queueA), len(visitedAsync.Map)) 
				
				// Set found ke true kalo judul sama dg to
				if (currentTree.judul == to) {
					if singleSolution {
						if !getFound() {
							setFound(true)
							select {
							case resultTree <- *currentTree:
							default:
							}
						}
					} else {
						if len(resultTree) == 0 {
							if getPJ() == -99 {
								fmt.Println("setPJ", currentTree.link, currentTree.depth)
								setPJ(currentTree.depth)
							}

							select {
							case resultTree <- *currentTree: 
							default:
							}
						} else {
							if currentTree.depth == getPJ() {
								select {
								case resultTree <- *currentTree:
								default:
								}									
							} else {
								setDone(true)
							}
						}
					}
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

		if singleSolution {
			if getFound() {
				break
			}
		} else {
			if getDone() {
				break
			}
		}
	}
	
	// Output
	if singleSolution {
		resT := <- resultTree
		path := append(resT.prev, resT.link)
		return cntAsync, resT.depth, [][]string{path}
	} else {
		allPath := [][]string{}
		lenC := len(resultTree)

		for i := 0; i < lenC; i++ {
			val := <-resultTree
			path := append(val.prev, val.link)
			if !isInSlice(allPath, path) {
				allPath = append(allPath, path)
			}
		}

		return cntAsync, getPJ(), allPath
	}
}

func isEqualSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i, val := range s1 {
		if val != s2[i] {
			return false
		}
	}
	return true
}

func isInSlice(s1 [][]string, s2 []string) bool {
	if len(s1) == 0 {
		return false
	}

	for _, val := range s1 {
		if isEqualSlice(val, s2) {
			return true
		}
	}
	return false
}