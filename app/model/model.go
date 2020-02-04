package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"Users/pingjing/docker/goPractice/owning/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Shopping struct {
	Accounts []User `bson:"accounts" json:"accounts"`
}

type User struct {
	Name     string    `bson:"name" json:"name"`
	Password string    `bson:"password" json:"password"`
	Phone    string    `bson:"phone" json:"phone"`
	Products []Product `bson:"products" json:"products"`
}

// uppercase an exported to json
type Product struct {
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
func (p *Product) FindAll() error {
	var shoppingLists []*Product

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
		list := Product{}

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
func (p *Product) FindOne(query bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	err := c.FindOne(ctx, query).Decode(&p)

	return err
}

// Insert a new single product
func (p *Product) Insert(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	result, err := c.InsertOne(ctx, product)

	if err != nil {
		log.Fatal(err)
	}

	objectID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(objectID)

	return err
}

// Update a product
func (p *Product) Update(product Product) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c := database.Connect(collection)

	opts := options.Update().SetUpsert(true)

	// find the document for which the _id field matches id and set the email to "newemail@example.com"
	filter := bson.M{"productId": product.ProductId}
	update := bson.M{"$set": product}

	_, err := c.UpdateOne(ctx, filter, update, opts)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

// Delete a product
