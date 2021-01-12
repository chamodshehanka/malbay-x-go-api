package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Description  string             `bson:"description,omitempty"`
	stocks       int                `bson:"stocks"`
	price        int                `bson:"price"`
	imageGallery string             `bson:"imageGallery"`
}
