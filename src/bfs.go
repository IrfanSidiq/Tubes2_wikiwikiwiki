package main

import "fmt"

type Tree struct {
	judul   string
	link    string
	nextArr []*Tree
}

var isFound bool

func BFS(depth int, currentTree Tree, judulArtikelTujuan string) {
	// Cari judul halaman
	getJudul(currentTree.link, func(judul string) {
		currentTree.judul = judul
	})

	// Base Case, berhenti kalo sudah ketemu
	isFound = (currentTree.judul == judulArtikelTujuan)
	if isFound {
		return
	}

	// Ambil semua link yang terdapat pada halaman
	arr := linkScraping(currentTree.link)
	for _, link := range arr {
		newTree := Tree{
			link:    link,
			nextArr: []*Tree{},
		}
		currentTree.nextArr = append(currentTree.nextArr, &newTree)
	}

	// fmt.Println(currentTree.judul, currentTree.link)

	// masih bukan bfs
	depth++
	for i := 0; i < len(currentTree.nextArr); i++ {
		fmt.Println(currentTree.nextArr[i].link)
		// BFS(depth, *currentTree.nextArr[i], judulArtikelTujuan)
	}

}