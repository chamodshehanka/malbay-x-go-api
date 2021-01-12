package main

import (
	"log"
	"net/http"

	"github.com/Shehanka/malbay-x-go-api/api"
)

func main() {
	routes, err := api.LoadAPI()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(http.ListenAndServe(":4000", routes))
}
