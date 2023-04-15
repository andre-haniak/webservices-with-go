package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/andre-haniak/webservices-with-go/models"
)

var productList []models.Product

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)
	case http.MethodPost:
		var newProduct models.Product
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(body, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

}

func getNextID() int {
	highestID := -1
	for _, p := range productList {
		if p.ProductID > highestID {
			highestID = p.ProductID
		}
	}
	return highestID + 1
}
