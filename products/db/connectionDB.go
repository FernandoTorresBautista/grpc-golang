package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN connection object to db
var MongoCN = ConectarDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://goproject:j7f8aG7Hx0amdIIl@cluster0.rdvbe.mongodb.net/goproducs")

// init: init function
func init() {
}

// ConectarDB: return the connection to db
func ConectarDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return client
	}
	log.Println("Successful connection with mongodb")
	return client
}

// CheckConecction: check if the connection it's ok
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	return 1
}
