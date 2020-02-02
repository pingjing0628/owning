package database

import (
	"context"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload" // _ means load this file's func init()

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	url      = os.Getenv("DATABASE_URL")
	database = os.Getenv("DB_DATABASE")
	err      error
	client   *mongo.Client
)

// Remember to export DATABASE_URL first
// export DATABASE_URL=mongodb://localhost:27017
func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Set client options
	clientOpts := options.Client().ApplyURI(url)

	// Connect to MongoDB
	if client, err = mongo.Connect(ctx, clientOpts); err != nil {
		log.Fatalln(err.Error())
	}

	// fmt.Println("Connected to MongoDB!")
}

// Create connect to database
func Connect(collection string) *mongo.Collection {
	return client.Database(database).Collection(collection)
}
