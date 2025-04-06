package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://rashidassef71:Password@cluster0.g2ta8do.mongodb.net/"
const dbName = "todolst"
const colName = "todo"

var Collection *mongo.Collection

func Mongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOption := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		log.Fatal("MongoDB connection Failed:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping Failed:", err)
	}

	fmt.Println("MongoDB connection successfull")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection is ready")
}
