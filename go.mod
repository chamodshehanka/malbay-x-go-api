module github.com/Shehanka/malbay-x-go-api

go 1.15

replace github.com/Shehanka/malbay-x-go-api/api => ./api

replace github.com/Shehanka/malbay-x-go-api/api/routes => ./api/routes

replace github.com/Shehanka/malbay-x-go-api/api/handlers => ./api/handlers

require (
	github.com/cosmtrek/air v1.21.2 // indirect
	github.com/creack/pty v1.1.11 // indirect
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-chi/chi v1.5.1
	github.com/go-chi/cors v1.1.1
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/pelletier/go-toml v1.8.1 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.7.1
	go.mongodb.org/mongo-driver v1.4.4
	golang.org/x/sys v0.0.0-20210113181707-4bcb84eeeb78 // indirect
)
