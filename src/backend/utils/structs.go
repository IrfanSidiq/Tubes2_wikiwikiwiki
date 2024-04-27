package scraper

import "sync"

/*
	STRUCT, GLOBAL-VARIABLE, DAN METHOD YANG DIGUNAKAN
*/

// Struct
type Tree struct {
	prev    []string
	judul   string
	link    string
	nextArr []*Tree
	depth   int
}

type stringBoolMap struct {
	Map map[string]bool
	sync.RWMutex
}

type IDSTree struct {
	Map map[string][]string
	sync.RWMutex
}

// Global Variables
var (
	linkAsal 		string
	linkTujuan 		string
	visitedAsync 	stringBoolMap
	
	cntAsync 		int
	cntMutex 		sync.Mutex

	foundAsync 		bool
	foundMutex 		sync.Mutex

	doneBFS 		bool
	doneMutex 		sync.Mutex
	
	panjangRute int
	pjMutex sync.Mutex
)

// Method
func setPJ(val int) {
	pjMutex.Lock()
	defer pjMutex.Unlock()
	if panjangRute == -99 {
		panjangRute = val
	}
}

func getPJ() int {
	pjMutex.Lock()
	defer pjMutex.Unlock()
	return panjangRute
}

func setDone(val bool) {
	doneMutex.Lock()
	defer doneMutex.Unlock()
	doneBFS = val
}

func getDone() bool {
	doneMutex.Lock()
	defer doneMutex.Unlock()
	return doneBFS
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

func (m *stringBoolMap) keyExists(key string) bool {
	m.RLock()
	defer m.RUnlock()

	_, ok := m.Map[key]
	return ok 
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