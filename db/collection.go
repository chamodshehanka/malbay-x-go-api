package db

import (
	"github.com/Shehanka/malbay-x-go-api/config"
	"go.mongodb.org/mongo-driver/mongo"
)

var client = config.GetMongoDBConnection()

func getDB() *mongo.Database {
	return client.Database("MalbayX")
}

func GetProductCollection() *mongo.Collection {
	return getDB().Collection("products")
}

func GetAdminCollection() *mongo.Collection {
	return getDB().Collection("admins")
}

func GetUserCollection() *mongo.Collection {
	return getDB().Collection("user")
}
