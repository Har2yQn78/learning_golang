package main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/Har2yQn78/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegiesterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
