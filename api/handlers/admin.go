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

var adminCollection = db.GetAdminCollection()

func AdminCreate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := adminCollection.InsertOne(ctx, admin)

	_ = json.NewEncoder(w).Encode(result)
}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := adminCollection.UpdateOne(ctx, models.Admin{ID: admin.ID}, admin)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func AdminDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := adminCollection.DeleteOne(ctx, models.Admin{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	_ = json.NewEncoder(w).Encode(result)
}

func AdminList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var admins []models.Admin

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var a models.Admin
		_ = cursor.Decode(&a)
		admins = append(admins, a)
	}

	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}

	_ = json.NewEncoder(w).Encode(admins)
}

func AdminGetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	var a models.Admin

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := adminCollection.FindOne(ctx, models.Admin{ID: id}).Decode(&a)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"message": "` + err.Error() + `"}`))

		return
	}
	_ = json.NewEncoder(w).Encode(a)
}
