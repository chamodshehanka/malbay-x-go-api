package main

import (
	"github.com/Shehanka/malbay-x-go-api/api"
	"log"
	"net/http"
)

func main() {

	routes, err := api.LoadAPI()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(http.ListenAndServe(":4000", routes))
}
