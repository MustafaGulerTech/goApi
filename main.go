package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Person struct represents the structure of our user data
type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate string `json:"birth_date"`
}

var database []Person
var apiKey string

func init() {
	// Load the secret API key from an environment variable for better security.
	// For this example, we'll use a hardcoded fallback if the variable is not set.
	apiKey = os.Getenv("API_KEY")
	if apiKey == "" {
		apiKey = "your-secret-api-key" // Fallback for demonstration
		log.Println("Warning: API_KEY environment variable not set. Using default key.")
	}

	// Load the database from the JSON file
	jsonFile, err := os.Open("db.json")
	if err != nil {
		log.Fatalf("Failed to open db.json: %v", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &database)
}

// apiKeyMiddleware checks if the request has a valid API key.
func apiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.Header.Get("X-API-KEY")

		if key != apiKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// If the key is valid, call the next handler
		next.ServeHTTP(w, r)
	})
}

// getUsersHandler handles requests to the /users endpoint.
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(database)
}

func main() {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Wrap the handler with the API key middleware
	usersHandler := http.HandlerFunc(getUsersHandler)
	mux.Handle("/users", apiKeyMiddleware(usersHandler))

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
