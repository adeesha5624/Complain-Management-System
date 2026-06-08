package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/crm_db?sslmode=disable"
	}

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// API Routing (implemented by other members)
	http.HandleFunc("/api/auth/login", handleLogin) // Member 2
	http.HandleFunc("/api/auth/user", handleGetUser) // Member 2
	http.HandleFunc("/api/complaints", handleComplaints) // Members 3, 4, 5

	fmt.Println("Backend server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// CORS Handler Wrapper
func corsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func getTokenFromRequest(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}
	// Simplified parsing
	return strings.TrimPrefix(authHeader, "Bearer ")
}
