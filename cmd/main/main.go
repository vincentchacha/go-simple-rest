package main

import (
	"log"
	"net/http"
	"github.com/example/go-simple-rest/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
	r := mux.NewRouter()
	routes.RegisterUserRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}