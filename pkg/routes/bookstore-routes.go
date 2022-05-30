package routes

import (
	"github.com/gorilla/mux"
	"github.com/raducl/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.ListBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
