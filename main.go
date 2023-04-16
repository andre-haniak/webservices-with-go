package main

import (
	"log"
	"net/http"

	"github.com/andre-haniak/webservices-with-go/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
