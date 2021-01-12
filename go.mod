module github.com/Shehanka/malbay-x-go-api

go 1.15

replace github.com/Shehanka/malbay-x-go-api/api => ./api

replace github.com/Shehanka/malbay-x-go-api/api/routes => ./api/routes

replace github.com/Shehanka/malbay-x-go-api/api/handlers => ./api/handlers

require (
	github.com/go-chi/chi v1.5.1
	github.com/go-chi/cors v1.1.1
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.4
)
