package main

import (
	"Go-PersonalFinanceTracker/pkg/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterExpensesRoutes(r)
	routes.RegisterIncomeRoutes(r)

	err := http.ListenAndServe(":8080", r)
	fmt.Println(err)
}
