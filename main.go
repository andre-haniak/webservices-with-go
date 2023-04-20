package main

import (
	"log"
	"net/http"

	"github.com/andre-haniak/webservices-with-go/database"
	"github.com/andre-haniak/webservices-with-go/product"
	_ "github.com/go-sql-driver/mysql"
)

const apiBasePath = "/api"

func main() {
	database.SetupDB()
	product.SetupRoutes(apiBasePath)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
