package database

import (
	"log"

	"github.com/ivanauliaa/shortener-go/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() *mongo.Collection {
	clientOptions := options.Client().ApplyURI(utils.DB_URI)

	client, err := mongo.Connect(utils.GLOBAL_CONTEXT, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(utils.GLOBAL_CONTEXT, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("DB connected")

	return client.Database(utils.DB_NAME).Collection(utils.DB_COLLECTION)
}
