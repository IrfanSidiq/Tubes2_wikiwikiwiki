package main

import "sync"

/*
	STRUCT, GLOBAL-VARIABLE, DAN METHOD
	YANG DIGUNAKAN UNTUK BFS_ASYNC.GO
*/

// Struct
type Tree struct {
	prev    []string
	judul   string
	link    string
	nextArr []*Tree
	depth   int
}

// Global Variables
var (
	linkAsal string
	linkTujuan string
	visitedAsync stringBoolMap
	
	cntAsync int
	cntMutex sync.Mutex

	foundAsync bool
	foundMutex sync.Mutex
)

// Method
func setFound(val bool) {
	foundMutex.Lock()
	defer foundMutex.Unlock()
	foundAsync = val
}

func getFound() bool {
	foundMutex.Lock()
	defer foundMutex.Unlock()
	return foundAsync
}

func setCnt(val int) {
	cntMutex.Lock()
	defer cntMutex.Unlock()
	cntAsync = val
}

func incCnt() {
	cntMutex.Lock()
	defer cntMutex.Unlock()
	cntAsync++
}

func (m *stringBoolMap) keyExists(key string) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.Map[key]
	return ok 
}