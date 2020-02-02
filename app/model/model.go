package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"Users/pingjing/docker/goPractice/owning/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Users struct
// type Users struct {
// 	name string `bson:"name" json:"name"`
// }

// ShoppingLists struct
// type Shopping_lists struct {
// 	// name string `bson:"name" json:"name"`
// 	lists []Products `bson:"shoppingLists" json:"shoppingLists"`
// }

// product struct
// uppercase an exported to json
type Products struct {
	ProductId    string `bson:"productId" json:"productId"`
	ProductName  string `bson:"productName" json:"productName"`
	Price        string `bson:"price" json:"price"`
	Category     string `bson:"category" json:"category"`
	PurchaseDate string `bson:"purchaseDate" json:"purchaseDate"`
}

const (
	collection = "shopping_lists"
)

// Get all product list
// 放置型別為 collection 的指標
func (p *Products) FindAll() error {
	var shoppingLists []*Products

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	// opts := options.Find().SetLimit(2)

	// Passing bson.M{} as the filter matches all documents in the collection
	cur, err := c.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Error on Finding all the documents ", err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		list := Products{}

		// create a value into which the single document can be decoded
		if err := cur.Decode(&list); err != nil {
			log.Fatal("Error on Decoding the document ", err)
		}

		shoppingLists = append(shoppingLists, &list)

	}

	if err := cur.Err(); err != nil {
		return err
	}
	return cur.Close(ctx)
}

// Get single product list
func (p *Products) FindOne(query bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	err := c.FindOne(ctx, query).Decode(&p)

	return err
}

// Insert a new single product
func (p *Products) Insert(query bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	list := Products{"2", "123", "333", "rrr", "2020"}

	result, err := c.InsertOne(ctx, list)

	if err != nil {
		log.Fatal(err)
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	return err
}

// Update a product

// Delete a product
