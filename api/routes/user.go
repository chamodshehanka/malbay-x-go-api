package routes

import (
	"github.com/Shehanka/malbay-x-go-api/api/handlers"
	"github.com/go-chi/chi"
)

func UserRouter(r chi.Router) {
	r.Post("/refresh", handlers.Refresh)
	r.Post("/singin", handlers.Signin)
	r.Post("/singup", handlers.Singup)
	r.Get("/welcome", handlers.Welcome)
}
