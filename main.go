package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andre-haniak/webservices-with-go/controllers"
	"github.com/andre-haniak/webservices-with-go/models"
)

var productList []models.Product

func init() {
	productsJSON := `[
		{
			"productId": 1,
			"manufacturer": "Johns-Jenkins",
			"sku": "p5w705",
			"upc": "939851000000",
			"pricePerUnit": "81.29",
			"quantityOnHand": 981,
			"productName": "sticky note"
		},
		{
			"productId": 2,
			"manufacturer": "Kshlerin, Reinger and Beier",
			"sku": "q7t864",
			"upc": "939851000000",
			"pricePerUnit": "85.60",
			"quantityOnHand": 998,
			"productName": "leg warmers"
		}
	]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/products", controllers.ProductsHandler)
	http.HandleFunc("/products/", controllers.ProductHandler)
	http.ListenAndServe(":5000", nil)
}
