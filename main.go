package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/andre-haniak/webservices-with-go/products"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("before handler; middleware start")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished; %s", time.Since(start))
	})
}

const apiBasePath = "/api"

func main() {
	products.SetupRoutes(apiBasePath)

	http.ListenAndServe(":5000", nil)
}
