package routes

import (
	"github.com/Shehanka/malbay-x-go-api/api/handlers"
	"github.com/go-chi/chi"
)

func AdminRoutes(r chi.Router) {
	r.Post("/add", handlers.AdminCreate)
	r.Put("/update", handlers.AdminUpdate)
	r.Delete("/delete/{id}", handlers.AdminDelete)
	r.Get("/list", handlers.AdminList)
	r.Get("/get/{id}", handlers.AdminGetByID)
}
