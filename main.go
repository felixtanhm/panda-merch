// main.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler for the root route
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go!")
}

// Handler for a JSON response
func dataHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		response := map[string]string{"message": "This is a GET request!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	case http.MethodPost:
		response := map[string]string{"message": "This is a POST request!"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}

}

func main() {
	http.HandleFunc("/", rootHandler)         // Route for "/"
	http.HandleFunc("/api/data", dataHandler) // Route for "/api/data"

	port := ":3000"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) // Start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
