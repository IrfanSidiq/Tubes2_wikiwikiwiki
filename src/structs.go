package main

import "sync"

// Structs

type Tree struct {
	prev    *Tree
	judul   string
	link    string
	nextArr []*Tree
	depth   int
}

type TreeArr struct {
	arr []*Tree
	sync.RWMutex
}

// Global Variables

var (
	linkAsal string
	queueAsync TreeArr
	visitedAsync stringBoolMap

	foundAsync bool
	foundMutex sync.Mutex
)

// Method

func newTreeArr() TreeArr {
	return TreeArr{
		arr: []*Tree{},
	}
}

func (a *TreeArr) apd(val *Tree) {
	a.RLock()
	defer a.RUnlock()
	a.arr = append(a.arr, val)
}

func (a *TreeArr) len() int {
	a.RLock()
	defer a.RUnlock()
	return len(a.arr)
}

func (a *TreeArr) removeFirstElement() {
	a.RLock()
	defer a.RUnlock()
	a.arr = a.arr[1:]
}

func (a *TreeArr) getFirstElement() *Tree {
	a.RLock()
	defer a.RUnlock()
	return a.arr[0]
}

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

func (m *stringBoolMap) keyExists(key string) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.Map[key]
	return ok 
}