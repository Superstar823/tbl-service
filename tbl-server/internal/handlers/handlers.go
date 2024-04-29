package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"example.com/tbl-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func ItemHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Handle GET request (list items)
			cursor, err := client.Database("mydb").Collection("tbl_item").Find(context.Background(), bson.D{})
			if err != nil {
				http.Error(w, "Failed to query the items: "+err.Error(), http.StatusInternalServerError)
				return
			}
			var items []models.TblItem
			if err = cursor.All(context.Background(), &items); err != nil {
				http.Error(w, "Failed to decode the items: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(items)

		case http.MethodPost:
			var item models.TblItem
			if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			item.ID = primitive.NewObjectID()
			item.CreatedAt = time.Now()
			item.UpdatedAt = time.Now()

			_, err := client.Database("mydb").Collection("tbl_item").InsertOne(context.Background(), item)
			if err != nil {
				http.Error(w, "Failed to insert the item: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)

		case http.MethodDelete:
			var item models.TblItem
			// Decode the request body into the item variable
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields() // Avoid decoding unknown fields
			if err := decoder.Decode(&item); err != nil {
				// Handle decoding error and respond with a Bad Request status
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Check if the provided ID is in the correct format
			objectID, err := primitive.ObjectIDFromHex(item.ID.Hex())
			if err != nil {
				http.Error(w, "Invalid item ID format: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Create the filter for the MongoDB query
			filter := bson.M{"_id": objectID}

			result, err := client.Database("mydb").Collection("tbl_item").DeleteOne(context.Background(), filter)
			if err != nil || result.DeletedCount == 0 {
				http.Error(w, "Failed to delete the item: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		case http.MethodPut:
			var item models.TblItem

			// Decode the request body into the item variable
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields() // Avoid decoding unknown fields
			if err := decoder.Decode(&item); err != nil {
				// Handle decoding error and respond with a Bad Request status
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Check if the provided ID is in the correct format
			objectID, err := primitive.ObjectIDFromHex(item.ID.Hex())
			if err != nil {
				http.Error(w, "Invalid item ID format: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Create the filter for the MongoDB query
			filter := bson.M{"_id": objectID}

			// Create the update map for the MongoDB query
			update := bson.M{"$set": bson.M{
				"item_name":  item.ItemName,
				"price":      item.Price,
				"cost":       item.Cost,
				"updated_at": time.Now(),
			}}

			// Perform the update operation on the MongoDB collection
			result, err := client.Database("mydb").Collection("tbl_item").UpdateOne(context.Background(), filter, update)

			// Check for errors in the update operation
			if err != nil {
				http.Error(w, "Failed to update the item: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Check the ModifiedCount to ensure the update was successful
			if result.ModifiedCount == 0 {
				http.Error(w, "No item was updated; the item ID might not exist", http.StatusNotFound)
				return
			}

			// Respond with a No Content status to indicate success
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func CustomerHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			//Handle GET request (list users)
			cursor, err := client.Database("mydb").Collection("tbl_customer").Find(context.Background(), bson.D{})

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			var customers []models.TblCustomer

			if err = cursor.All(context.Background(), &customers); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(customers)

		case http.MethodPost:
			var customer models.TblCustomer
			if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			customer.ID = primitive.NewObjectID()
			customer.CreatedAt = time.Now()
			customer.UpdatedAt = time.Now()

			_, err := client.Database("mydb").Collection("tbl_customer").InsertOne(context.Background(), customer)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusCreated)

		case http.MethodDelete:
			var customer models.TblCustomer
			// Decode the request body into the customer variable
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields() // Avoid decoding unknown fields
			if err := decoder.Decode(&customer); err != nil {
				// Handle decoding error and respond with a Bad Request status
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Check if the provided ID is in the correct format
			objectID, err := primitive.ObjectIDFromHex(customer.ID.Hex())
			if err != nil {
				http.Error(w, "Invalid customer ID format: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Create the filter for the MongoDB query
			filter := bson.M{"_id": objectID}

			result, err := client.Database("mydb").Collection("tbl_customer").DeleteOne(context.Background(), filter)
			if err != nil || result.DeletedCount == 0 {
				http.Error(w, "Failed to delete the customer: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		case http.MethodPut:
			var customer models.TblCustomer

			// Decode the request body into the customer variable
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields() // Avoid decoding unknown fields
			if err := decoder.Decode(&customer); err != nil {
				// Handle decoding error and respond with a Bad Request status
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Check if the provided ID is in the correct format
			objectID, err := primitive.ObjectIDFromHex(customer.ID.Hex())
			if err != nil {
				http.Error(w, "Invalid customer ID format: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Create the filter for the MongoDB query
			filter := bson.M{"_id": objectID}

			// Create the update map for the MongoDB query
			update := bson.M{"$set": bson.M{
				"customer_name": customer.CustomerName,
				"balance":       customer.Balance,
				"updated_at":    time.Now(),
			}}

			// Perform the update operation on the MongoDB collection
			result, err := client.Database("mydb").Collection("tbl_customer").UpdateOne(context.Background(), filter, update)

			// Check for errors in the update operation
			if err != nil {
				http.Error(w, "Failed to update the customer: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Check the ModifiedCount to ensure the update was successful
			if result.ModifiedCount == 0 {
				http.Error(w, "No customer was updated; the customer ID might not exist", http.StatusNotFound)
				return
			}

			// Respond with a No Content status to indicate success
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func TransactionHandler(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Get a handle to the tbl_transaction collection
			transCollection := client.Database("mydb").Collection("tbl_transaction")

			// Define the aggregation pipeline with the $lookup stage
			pipeline := mongo.Pipeline{
				bson.D{
					// $lookup stage for tbl_customer
					bson.E{
						Key: "$lookup",
						Value: bson.D{
							{Key: "from", Value: "tbl_customer"},      // Collection to join with
							{Key: "localField", Value: "customer_id"}, // Field in tbl_transaction
							{Key: "foreignField", Value: "_id"},       // Field in tbl_customer
							{Key: "as", Value: "customer_details"},    // Name of the new field in tbl_transaction
						},
					},
				},
				bson.D{
					// $lookup stage for tbl_item
					bson.E{
						Key: "$lookup",
						Value: bson.D{
							{Key: "from", Value: "tbl_item"},      // Collection to join with
							{Key: "localField", Value: "item_id"}, // Field in tbl_transaction
							{Key: "foreignField", Value: "_id"},   // Field in tbl_item
							{Key: "as", Value: "item_details"},    // Name of the new field in tbl_transaction
						},
					},
				},
				bson.D{
					// $addFields stage to add customer_name and item_name
					bson.E{
						Key: "$addFields",
						Value: bson.D{
							// Extract first element (index 0) of customer_name array, or provide a default value
							{Key: "customer_name", Value: bson.D{
								{Key: "$ifNull", Value: bson.A{
									bson.D{
										{Key: "$arrayElemAt", Value: bson.A{"$customer_details.customer_name", 0}},
									},
									"N/A", // Default value if customer_name array is empty
								}},
							}},
							// Extract first element (index 0) of item_name array, or provide a default value
							{Key: "item_name", Value: bson.D{
								{Key: "$ifNull", Value: bson.A{
									bson.D{
										{Key: "$arrayElemAt", Value: bson.A{"$item_details.item_name", 0}},
									},
									"N/A", // Default value if item_name array is empty
								}},
							}},
						},
					},
				},
			}

			// Handle GET request (list transactions)
			cursor, err := transCollection.Aggregate(context.Background(), pipeline)

			if err != nil {
				http.Error(w, "Failed to execute aggregation: "+err.Error(), http.StatusInternalServerError)
				return
			}
			defer cursor.Close(context.Background())

			var transactions []models.TblTransactionView
			if err = cursor.All(context.Background(), &transactions); err != nil {
				http.Error(w, "Failed to decode the transactions: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(transactions)

		case http.MethodPost:
			var transaction models.TblTransaction
			if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			log.Println("transaction: ", transaction)

			transaction.ID = primitive.NewObjectID()
			transaction.CreatedAt = time.Now()
			transaction.UpdatedAt = time.Now()

			_, err := client.Database("mydb").Collection("tbl_transaction").InsertOne(context.Background(), transaction)
			if err != nil {
				http.Error(w, "Failed to insert the transaction: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)

		case http.MethodDelete:
			var transaction models.TblTransaction
			// Decode the request body into the transaction variable
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields() // Avoid decoding unknown fields
			if err := decoder.Decode(&transaction); err != nil {
				// Handle decoding error and respond with a Bad Request status
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Check if the provided ID is in the correct format
			objectID, err := primitive.ObjectIDFromHex(transaction.ID.Hex())
			if err != nil {
				http.Error(w, "Invalid item ID format: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Create the filter for the MongoDB query
			filter := bson.M{"_id": objectID}

			result, err := client.Database("mydb").Collection("tbl_transaction").DeleteOne(context.Background(), filter)
			if err != nil || result.DeletedCount == 0 {
				http.Error(w, "Failed to delete the transaction: "+err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusNoContent)

		case http.MethodPut:
			var transaction models.TblTransaction

			// Decode the request body into the transaction variable
			decoder := json.NewDecoder(r.Body)
			decoder.DisallowUnknownFields() // Avoid decoding unknown fields
			if err := decoder.Decode(&transaction); err != nil {
				// Handle decoding error and respond with a Bad Request status
				http.Error(w, "Failed to decode the request body: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Check if the provided ID is in the correct format
			objectID, err := primitive.ObjectIDFromHex(transaction.ID.Hex())
			if err != nil {
				http.Error(w, "Invalid transaction ID format: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Create the filter for the MongoDB query
			filter := bson.M{"_id": objectID}

			// Create the update map for the MongoDB query
			update := bson.M{"$set": bson.M{
				"customer_id": transaction.CustomerID,
				"item_id":     transaction.ItemId,
				"qty":         transaction.Qty,
				"price":       transaction.Price,
				"amount":      transaction.Amount,
				"updated_at":  time.Now(),
			}}

			// Perform the update operation on the MongoDB collection
			result, err := client.Database("mydb").Collection("tbl_transaction").UpdateOne(context.Background(), filter, update)

			// Check for errors in the update operation
			if err != nil {
				http.Error(w, "Failed to update the transaction: "+err.Error(), http.StatusInternalServerError)
				return
			}

			// Check the ModifiedCount to ensure the update was successful
			if result.ModifiedCount == 0 {
				http.Error(w, "No transaction was updated; the transaction ID might not exist", http.StatusNotFound)
				return
			}

			// Respond with a No Content status to indicate success
			w.WriteHeader(http.StatusNoContent)

		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}
