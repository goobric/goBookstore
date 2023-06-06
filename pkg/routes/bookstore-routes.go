package routes

import (
	"github.com/goobric/goBookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

//Golang has Absolute Paths
// "github.com/gorilla/mux" "github.com/goobric/goBookstore/pkg/controllers"

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
