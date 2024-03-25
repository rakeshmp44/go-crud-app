package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rakeshmp44/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
