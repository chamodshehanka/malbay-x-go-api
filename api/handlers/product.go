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

var collection = db.GetProductCollection()

func ProductCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	_ = json.NewDecoder(r.Body).Decode(&product)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := collection.InsertOne(ctx, product)

	_ = json.NewEncoder(w).Encode(result)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {

}

func ProductDelete(w http.ResponseWriter, r *http.Request) {

}

func ProductList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []models.Product

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var p models.Product
		_ = cursor.Decode(&p)
		products = append(products, p)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	_ = json.NewEncoder(w).Encode(products)

}

func ProductGetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	var p models.Product

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := collection.FindOne(ctx, models.Product{ID: id}).Decode(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}
	_ = json.NewEncoder(w).Encode(p)
}
