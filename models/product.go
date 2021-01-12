package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Name         string             `bson:"name,omitempty"`
	Description  string             `bson:"description,omitempty"`
	Stocks       int                `bson:"stocks"`
	Price        int                `bson:"price"`
	ImageGallery string             `bson:"imageGallery"`
}
