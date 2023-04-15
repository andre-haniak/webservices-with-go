package main

import "net/http"

type fooHandler struct {
	Message string
}

func (h *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(h.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "Hello World!"})

	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
