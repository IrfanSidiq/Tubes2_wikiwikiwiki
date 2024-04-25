package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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

	// Access form data
	startPage := requestData["start-page"]
	endPage := requestData["end-page"]
	algorithm := requestData["algorithm"]

	// Process the data...
	fmt.Printf("Received data: From=%s, To=%s, Algorithm=%s\n", startPage, endPage, algorithm)

	// Send response
	response := map[string]string{"message": "Data received successfully"}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/data", processHandler)
	http.ListenAndServe(":8080", nil)
}
