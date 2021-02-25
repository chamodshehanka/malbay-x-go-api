package routes

import (
	"github.com/Shehanka/malbay-x-go-api/api/handlers"
	"github.com/go-chi/chi"
)

func UserRouter(r chi.Router) {
	r.Post("/signin", handlers.Signin)
	r.Post("/signup", handlers.Signup)
	r.Get("/welcome", handlers.Welcome)
	r.Post("/refresh", handlers.Refresh)
	r.Post("/forgotpassword", handlers.ForgotPassword)
	r.Get("/resetpassword/{id}", handlers.ResetPassword)
}
