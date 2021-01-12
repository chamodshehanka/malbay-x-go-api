package routes

import (
	"github.com/Shehanka/malbay-x-go-api/api/handlers"
	"github.com/go-chi/chi"
)

func ProductRoutes(r chi.Router) {
	r.Post("/add", handlers.ProductCreate)
	r.Put("/update", handlers.ProductUpdate)
	r.Delete("/delete/{id}", handlers.ProductDelete)
	r.Get("/get/{id}", handlers.ProductGetByID)
	r.Get("/list", handlers.ProductList)
}
