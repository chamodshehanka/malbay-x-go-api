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
	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, _ := adminCollection.InsertOne(ctx, admin)

	RespondwithJSON(w, 200, result)
}

func AdminUpdate(w http.ResponseWriter, r *http.Request) {
	var admin models.Admin
	_ = json.NewDecoder(r.Body).Decode(&admin)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := adminCollection.UpdateOne(ctx, models.Admin{ID: admin.ID}, admin)
	if err != nil {
		RespondWithError(w, 404, err.Error())

		return
	}

	RespondwithJSON(w, 200, result)
}

func AdminDelete(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	result, err := adminCollection.DeleteOne(ctx, models.Admin{ID: id})
	if err != nil {
		RespondWithError(w, 404, err.Error())

		return
	}

	RespondwithJSON(w, 200, result)
}

func AdminList(w http.ResponseWriter, _ *http.Request) {
	var admins []models.Admin

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	cursor, err := productCollection.Find(ctx, bson.M{})
	if err != nil {
		RespondWithError(w, 404, err.Error())

		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var a models.Admin
		_ = cursor.Decode(&a)
		admins = append(admins, a)
	}

	if err := cursor.Err(); err != nil {
		RespondWithError(w, 404, err.Error())

		return
	}

	RespondwithJSON(w, 200, admins)
}

func AdminGetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(chi.URLParam(r, "id"))

	var a models.Admin

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err := adminCollection.FindOne(ctx, models.Admin{ID: id}).Decode(&a)
	if err != nil {
		RespondWithError(w, 404, err.Error())

		return
	}

	RespondwithJSON(w, 200, a)
}
