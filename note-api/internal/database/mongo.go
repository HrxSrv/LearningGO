package database

import(
	"log"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LoadMongo(uri string) *mongo.Client{

	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx,options.Client().ApplyURI(uri))
    if err != nil {
		log.Fatal("Connecting to mongodb failed");
	}

	err = client.Ping(ctx,nil)

	if err != nil {
		log.Fatal("Pinging to mongodb failed")
	}

	log.Println("Voila Connected to mongodb");

	return client
}