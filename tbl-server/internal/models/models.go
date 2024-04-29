// internal/models/models.go

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TblCustomer struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CustomerName string             `json:"customer_name" bson:"customer_name"`
	Balance      uint               `json:"balance" bson:"balance"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}

type TblItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	ItemName  string             `json:"item_name" bson:"item_name"`
	Cost      uint               `json:"cost" bson:"cost"`
	Price     uint               `json:"price" bson:"price"`
	Sort      int                `json:"sort" bson:"sort"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}

type TblTransaction struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	ItemId     primitive.ObjectID `json:"item_id" bson:"item_id"`
	Qty        int                `json:"qty" bson:"qty"`
	Price      uint               `json:"price" bson:"price"`
	Amount     uint               `json:"amount" bson:"amount"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at"`
}

type TblTransactionView struct {
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	CustomerID   primitive.ObjectID `json:"customer_id" bson:"customer_id"`
	ItemId       primitive.ObjectID `json:"item_id" bson:"item_id"`
	CustomerName string             `json:"customer_name" bson:"customer_name"`
	ItemName     string             `json:"item_name" bson:"item_name"`
	Qty          int                `json:"qty" bson:"qty"`
	Price        uint               `json:"price" bson:"price"`
	Amount       uint               `json:"amount" bson:"amount"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
