package main

import (
	"log"
	"net/http"

	"github.com/goobric/goBookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// log  net/http  github.com/gorilla/mux  _ "github.com/jinzhu/gorm/dialects/mysql" "github.com/goobric/goBookstore/pkg/routes"

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
