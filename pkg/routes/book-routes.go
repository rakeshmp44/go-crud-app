package routes

import (
	"github.com/gorilla/mux"
	"github.com/rakeshmp44/go-bookstore/pkg/controllers"
)

var BookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")

}
