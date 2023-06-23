package routes

import (
	expcontroller "Go-PersonalFinanceTracker/pkg/controllers/expenses"
	incontroller "Go-PersonalFinanceTracker/pkg/controllers/incomes"

	"github.com/gorilla/mux"
)

var RegisterIncomeRoutes = func(router *mux.Router) {
	router.HandleFunc("/incomes/", incontroller.GetIncomes).Methods("GET")
	router.HandleFunc("/income/create", incontroller.CreateIncome).Methods("POST")
	router.HandleFunc("/income/detail", incontroller.GetIncomeById).Methods("GET")
	router.HandleFunc("/income/update", incontroller.UpdateIncome).Methods("PUT")
	router.HandleFunc("/income/delete", incontroller.DeleteIncome).Methods("DELETE")
}

var RegisterExpensesRoutes = func(router *mux.Router) {
	router.HandleFunc("/expenses/", expcontroller.GetExpenses).Methods("GET")
	router.HandleFunc("/expenses/create", expcontroller.CreateExpenses).Methods("POST")
	router.HandleFunc("/expenses/detail", expcontroller.GetExpensesById).Methods("GET")
	router.HandleFunc("/expenses/update", expcontroller.UpdateExpenses).Methods("PUT")
	router.HandleFunc("/expenses/delete", expcontroller.DeleteExpenses).Methods("DELETE")
}
