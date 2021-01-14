package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

func GetMongoDBConnection() *mongo.Client {

	user := GetEnv("database.user")
	password := GetEnv("database.password")
	database := GetEnv("database.name")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://" + user + ":" + password + "@afcluster-mig7i.gcp.mongodb.net/" + database + "?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("%s", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("%s", err)
	}
	//defer client.Disconnect(ctx)

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("%s", err)
	}

	return client
}
