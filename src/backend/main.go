package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"
	scraper "wikipediaScraper/backend/utils"
)

func processHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Parse JSON data from request body
	var requestData map[string]string
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Get all components of the received data
	fmt.Println(requestData) // testing
	startPage := requestData["start-page"]
	endPage := requestData["end-page"]
	algorithm := requestData["algorithm"]
	numberOfPath := requestData["number-of-path"]

	// Find path

	var (
		jumlahArtikelDiperiksa int
		jumlahArtikelDilalui   int
		routes                 [][]string
	)

	startTime := time.Now()

	if algorithm == "bfs" {
		if numberOfPath == "single" {
			fmt.Println(startPage, endPage, "BFS", "Single") // testing
			// single bfs
			jumlahArtikelDiperiksa, jumlahArtikelDilalui, routes = scraper.BFS(startPage, endPage, true)
		} else {
			fmt.Println(startPage, endPage, "BFS", "Multiple") // testing
			// multiple bfs
			jumlahArtikelDiperiksa, jumlahArtikelDilalui, routes = scraper.BFS(startPage, endPage, false)
		}
	} else {
		if numberOfPath == "single" {
			fmt.Println(startPage, endPage, "IDS", "Single") // testing
			// single ids
			jumlahArtikelDiperiksa, jumlahArtikelDilalui, routes = scraper.IDS(startPage, endPage, true)
		} else {
			fmt.Println(startPage, endPage, "IDS", "Multiple") // testing
			jumlahArtikelDiperiksa, jumlahArtikelDilalui, routes = scraper.IDS(startPage, endPage, false)
		}
	}

	endTime := time.Now()
	searchDuration := endTime.Sub(startTime).Milliseconds()

	fmt.Println("Jumlah artikel diperiksa:", jumlahArtikelDiperiksa) // testing
	fmt.Println("Jumlah artikel dilalui:", jumlahArtikelDilalui)     // testing
	fmt.Println("Routes:")                                           // testing
	fmt.Println(routes)                                              // testing
	fmt.Println("Search duration:", searchDuration, "ms")            // testing
	
	// Convert routes to include titles
	routesWithTitle := make([]map[string]string, len(routes))
	for i, route := range routes {
		routesWithTitle[i] = make(map[string]string, len(route))
		for _, link := range route {
			routesWithTitle[i][link] = scraper.LinkTojudul(link)
		}
	}
	
	fmt.Println("routesWithTitle:")
	for _, route := range routesWithTitle {
		fmt.Println(route)
	}

	// Send response
	response := map[string]any{
		"jumlahArtikelDiperiksa": jumlahArtikelDiperiksa,
		"jumlahArtikelDilalui":   jumlahArtikelDilalui,
		"routes":                 routesWithTitle,
		"searchDuration":         searchDuration,
	}

	// Encode response into JSON and send it
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Get the path of the frontend directory from the backend directory
	frontendDir := filepath.Join("..", "frontend")

	// Serve static files from the frontend directory
	http.Handle("/", http.FileServer(http.Dir(frontendDir)))

	// Handle data endpoint
	http.HandleFunc("/data", processHandler)

	// Start HTTP server
	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// How to run:
// cd src/backend
// go run main.go
// Open browser and go to http://localhost:8080
