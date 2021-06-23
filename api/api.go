package api

import (
	"github.com/Shehanka/malbay-x-go-api/api/routes"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

var r *chi.Mux

func LoadAPI() (*chi.Mux, error) {
	r = chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	})

	r.Use(c.Handler)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)

	r.Route("/api/v1", func(r chi.Router) {
		r.Route("/admin", routes.AdminRoutes)
		r.Route("/product", routes.ProductRoutes)
		r.Route("/user", routes.UserRouter)
	})

	return r, nil
}
