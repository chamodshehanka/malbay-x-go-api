package main

import (
	"context"
	"fmt"
	"github.com/Shehanka/malbay-x-go-api/config"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"

	"github.com/Shehanka/malbay-x-go-api/api"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := config.GetMongoDBConnection()

	testDB := client.Database("ShopDB")
	podcastsCollection := testDB.Collection("admins")

	cursor,err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Println(cursor)

	routes, err := api.LoadAPI()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(http.ListenAndServe(":4000", routes))
}
