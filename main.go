package main

import (
	"./api"
	"log"
	"net/http"
)

func main() {
	routes, err := api.LoadAPI()
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":4000", routes)
	if err != nil {
		log.Fatal(err)
	}
}
