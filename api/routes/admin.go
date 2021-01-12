package routes

import (
	"../handlers"
	"github.com/go-chi/chi"
)

func AdminRoutes(r chi.Router) {
	r.Post("/add", handlers.AppCreate)
}
