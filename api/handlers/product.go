package handlers

import (
	"context"
	"encoding/json"
	"github.com/Shehanka/malbay-x-go-api/config"
	"github.com/Shehanka/malbay-x-go-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"time"
)

var products []models.Product

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {

}

func ProductDelete(w http.ResponseWriter, r *http.Request) {

}

func ProductList(w http.ResponseWriter, r *http.Request) {
	loadProducts()

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(products)
}

func ProductGetByID(w http.ResponseWriter, r *http.Request) {

}

// TODO: Move this to a separate layer
func loadProducts() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client := config.GetMongoDBConnection()

	testDB := client.Database("ShopDB")
	podcastsCollection := testDB.Collection("products")

	cursor, err := podcastsCollection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatalf("%s", err)
	}

	var productsList []models.Product

	if err = cursor.All(ctx, &products); err != nil {
		log.Fatal(err)
	}

	products = productsList
}
