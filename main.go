package main

import (
	"github.com/Shehanka/malbay-x-go-api/api"
	"github.com/Shehanka/malbay-x-go-api/config"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"

	cowsay "github.com/Code-Hex/Neo-cowsay"
)

func main() {

	routes, err := api.LoadAPI()

	if err != nil {

		log.Fatal(err)

	}

	port := config.GetEnv("server.port")

	say, _ := cowsay.Say(
		cowsay.Phrase("Malbay server listening on "+port),
		cowsay.Type("default"),
		cowsay.BallonWidth(40),
	)

	logrus.Print("\n" + say)

	log.Println(http.ListenAndServe(":"+port, routes))

}
