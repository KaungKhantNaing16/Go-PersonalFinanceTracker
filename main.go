package main

import (
	"Go-PersonalFinanceTracker/pkg/routes"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func dashboardHandler(writer http.ResponseWriter, req *http.Request) {
	templatePartialDir := "templates/partials/"
	tmpl := template.Must(template.ParseFiles(templatePartialDir+"layout.html", templatePartialDir+"dataTable.html", "templates/index.html"))

	err := tmpl.ExecuteTemplate(writer, "layout.html", nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/dashboard", dashboardHandler)
	routes.RegisterAuthRoutes(r)
	routes.RegisterExpensesRoutes(r)
	routes.RegisterIncomeRoutes(r)
	routes.RegisterLoanRoutes(r)
	routes.RegisterBudgetRoutes(r)
	routes.RegisterMediaRoutes(r)
	err := http.ListenAndServe(":8080", r)
	fmt.Println(err)
}
