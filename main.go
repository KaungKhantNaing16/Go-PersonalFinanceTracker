package main

import (
	"Go-PersonalFinanceTracker/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	routes.RegisterRoutes(r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
