package handlers

import (
	"context"
	"encoding/json"
	"github.com/Shehanka/malbay-x-go-api/db"
	"github.com/Shehanka/malbay-x-go-api/models"
	"github.com/go-chi/chi"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"time"
)

var productCollection = db.GetProductCollection()

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := productCollection.InsertOne(ctx, product)

	ResponseWithJSON(w, http.StatusOK, result)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {

}

func ProductDelete(w http.ResponseWriter, r *http.Request) {

}

func ProductList(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var p models.Product
		_ = cursor.Decode(&p)
		products = append(products, p)
	}

	if err := cursor.Err(); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	ResponseWithJSON(w, http.StatusOK, products)
}

func ProductGetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	var p models.Product

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	f := bson.D{{"_id", id}}

	if err := productCollection.FindOne(ctx, f).Decode(&p); err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())

		return
	}

	ResponseWithJSON(w, http.StatusOK, p)
}
