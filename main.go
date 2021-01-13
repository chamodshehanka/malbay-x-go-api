package main

import (
	"github.com/Shehanka/malbay-x-go-api/api"
	"github.com/Shehanka/malbay-x-go-api/config"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {

	routes, err := api.LoadAPI()

	if err != nil {

		log.Fatal(err)

	}

	port := config.GetEnv("server.port")

	logrus.Info("Malbay started listening on ", port)

	log.Println(http.ListenAndServe(":"+port, routes))

}
