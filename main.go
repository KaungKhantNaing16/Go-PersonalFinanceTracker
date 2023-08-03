package main

import (
	"Go-PersonalFinanceTracker/pkg/routes"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func dashboardHandler(writer http.ResponseWriter, req *http.Request) {
	templatePartialDir := "templates/partials/"
	tmpl, err := template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		"templates/index.html",
	)

	if err != nil {
		log.Fatal(err)
	}

	ttl := "Dashboard"
	err = tmpl.ExecuteTemplate(writer, "index.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/dashboard", dashboardHandler)
	routes.RegisterRoutes(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
