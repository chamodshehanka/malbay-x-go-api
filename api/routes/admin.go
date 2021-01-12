package routes

import (
	"github.com/Shehanka/malbay-x-go-api/api/handlers"
	"github.com/go-chi/chi"
)

func AdminRoutes(r chi.Router) {
	r.Post("/add", handlers.AppCreate)
}
