// main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"panda-merch/internal/handlers"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
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

func initDB() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("\nloading .env file error:\n %w", err)
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("\nopening connection error:\n %w", err)
	}
	err = DB.Ping()
	if err != nil {
		DB.Close()
		return nil, fmt.Errorf("\ndatabase pinging error:\n %w", err)
	}

	fmt.Println("Connected to MySQL database: ", dbName)
	return DB, nil
}

func main() {
	log.Println("Starting server...")

	db, err := initDB()
	if err != nil {
		log.Fatalf("Database initialization error: %v", err)
	}
	defer db.Close()
	http.HandleFunc("/", rootHandler)      // Route for "/"
	http.HandleFunc("/users", dataHandler) // Route for "/"
	http.HandleFunc("/merch", handlers.MerchHandler)

	port := ":3000"
	log.Printf("Server is running on http://localhost%s\n", port)
	err = http.ListenAndServe(port, nil) // Start server
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
