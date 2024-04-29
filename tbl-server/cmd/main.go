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

	//Set up handlers with database client
	http.HandleFunc("/api/transactions", handlers.TransactionHandler(client))
	http.HandleFunc("/api/customers", handlers.CustomerHandler(client))
	http.HandleFunc("/api/items", handlers.ItemHandler(client))

	//Start the server on port 8080
	fmt.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
