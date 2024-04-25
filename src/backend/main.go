package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"time"
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

	// Find path
	var jumlahArtikelDiperiksa int = 100 // testing with placeholder value
	var jumlahArtikelDilalui int = 100   // testing with placeholder value
	var routes [][]string = [][]string{  // testing with placeholder value
		{"apple.com", "banana.com", "egg.com"},
		{"apple.com", "banana.com", "cherry.com", "egg.com"},
		{"apple.com", "egg.com"},
		{"apple.com", "durian.com", "egg.com"},
	}

	startTime := time.Now()

	if algorithm == "bfs" {
		fmt.Println(startPage, endPage, "BFS") // testing
		// BFS_Async(startPage, endPage)
	} else {
		fmt.Println(startPage, endPage, "IDS") // testing
		// IDS_Async(startPage, endPage)
	}

	endTime := time.Now()
	searchDuration := endTime.Sub(startTime).Milliseconds()

	fmt.Println("Jumlah artikel diperiksa:", jumlahArtikelDiperiksa) // testing
	fmt.Println("Jumlah artikel dilalui:", jumlahArtikelDilalui)     // testing
	fmt.Println("Routes:")                                           // testing
	fmt.Println(routes)                                              // testing
	fmt.Println("Search duration:", searchDuration, "ms")            // testing

	// Send response
	response := map[string]any{
		"jumlahArtikelDiperiksa": jumlahArtikelDiperiksa,
		"jumlahArtikelDilalui":   jumlahArtikelDilalui,
		"routes":                 routes,
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
