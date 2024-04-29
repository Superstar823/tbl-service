package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/tbl-server/internal/database"
	"example.com/tbl-server/internal/handlers"
)

func main() {
	client, err := database.ConnectToDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	defer client.Disconnect(nil)

	// Set up handlers with database client
	http.HandleFunc("/api/transactions", handlers.TransactionHandler(client))
	http.HandleFunc("/api/transactions/{id}", handlers.TransactionHandler(client))
	http.HandleFunc("/api/customers", handlers.CustomerHandler(client))
	http.HandleFunc("/api/items", handlers.ItemHandler(client))

	// Add CORS middleware
	corsHandler := addCorsHeaders(http.DefaultServeMux)

	// Start the server on port 8080 with CORS middleware
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

// addCorsHeaders adds CORS headers to the response
func addCorsHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from any origin during development
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow the OPTIONS method (preflight requests)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			return
		}

		// Call the original handler
		handler.ServeHTTP(w, r)
	})
}
