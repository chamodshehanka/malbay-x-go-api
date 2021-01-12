package api

import (
	"./routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

var r *chi.Mux

func LoadAPI()(*chi.Mux, error)  {
	r = chi.NewRouter()

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(cors.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/admin", routes.AdminRoutes)
	})

	return r, nil
}