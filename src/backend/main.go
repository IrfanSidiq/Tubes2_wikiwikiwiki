package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		// Sample slice of strings
		data := []string{"apple", "banana", "orange", ""}

		// Convert slice to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Set response content type
		w.Header().Set("Content-Type", "application/json")

		// Write JSON response
		w.Write(jsonData)
	})

	http.ListenAndServe(":8080", nil)
}
