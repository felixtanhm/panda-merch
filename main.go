// main.go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

func initDB() error {
	godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	fmt.Println("initDB")
	DB, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("Error Open")
		return err
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error Ping")
		return err
	}

	fmt.Println("Connected to MySQL database: ", dbName)
	return nil
}

func main() {
	fmt.Println("running")
	errInit := initDB()
	if errInit != nil {
		log.Fatalf("Error initializing database %v", errInit)
	}
	http.HandleFunc("/", rootHandler)      // Route for "/"
	http.HandleFunc("/users", dataHandler) // Route for "/"

	port := ":3000"
	fmt.Printf("Server is running on http://localhost%s\n", port)
	err := http.ListenAndServe(port, nil) // Start the server
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
