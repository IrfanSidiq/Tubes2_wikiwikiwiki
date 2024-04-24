package main

import (
	"fmt"
)

var queue []*Tree

var visited map[string]bool

func initTree(currentTree *Tree) bool {
	var links []string
	
	linkScraping(currentTree.link, &links, &currentTree.judul)
	
	// Cek apakah halaman sudah pernah di-visit
	_, ok := visited[currentTree.judul]
	if (ok)  {
		return false
	} 

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

func BFS(root Tree, judulArtikelTujuan string) {
	queue = append(queue, &root)
	visited = make(map[string]bool)
	
	var isFound bool
	var currentTree *Tree
	var cnt int = 0 // Jumlah artikel yang diperiksa
	for {
		currentTree = queue[0]
		
		// Ambil semua link yang terdapat pada halaman
		isNotVisited := initTree(currentTree)
		
		if isNotVisited {
			cnt++
			// Berhenti ketika sudah ketemu / link habis
			isFound = (currentTree.judul == judulArtikelTujuan)
			if isFound || len(queue) == 0 {
				break
			}
	
			// Add semua link kecuali yang sudah ada dalam visited ke queue
			for _, val := range currentTree.nextArr {
				_, ok := visited[val.judul]
				if (!ok)  {
					queue = append(queue, val)
				} 
			}
	
			// Tambahkan ke visited
			visited[currentTree.judul] = true
		} 

		// Hapus elemen pertama queue
		queue = queue[1:]
	}
	
	// Output
	fmt.Print("\nLen Queue: ")
	fmt.Println(len(queue))
	output(currentTree, cnt)
}

/*
	Print OUTPUT
*/
func output(currentTree *Tree, cnt int) {
	cpyTree := currentTree

	fmt.Println("\nRute:")
	var route []string
	for {
		if cpyTree == nil {
			break
		}

		route = append(route, cpyTree.link)
		cpyTree = cpyTree.prev
	}

	for i := len(route) - 1; i >= 0; i-- {
		fmt.Println(route[i])
	}

	fmt.Print("\nJumlah artikel yang dilalui: ")
	fmt.Println(currentTree.depth)

	fmt.Print("Jumlah artikel yang diperiksa: ")
	fmt.Println(cnt)
}