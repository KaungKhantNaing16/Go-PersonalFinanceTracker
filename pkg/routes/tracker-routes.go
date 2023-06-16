package routes

import (
	expcontroller "Go-PersonalFinanceTracker/pkg/controllers/expenses"
	incontroller "Go-PersonalFinanceTracker/pkg/controllers/incomes"

	"github.com/gorilla/mux"
)

var RegisterIncomeRoutes = func(router *mux.Router) {
	income := router.PathPrefix("/tracker").Subrouter()
	income.HandleFunc("/incomes/", incontroller.GetIncomes).Methods("GET")
	income.HandleFunc("/income/", incontroller.CreateIncome).Methods("POST")
	income.HandleFunc("/income/:id", incontroller.GetIncomeById).Methods("GET")
	income.HandleFunc("/income/:id", incontroller.UpdateIncome).Methods("PUT")
	income.HandleFunc("/income/:id", incontroller.DeleteIncome).Methods("DELETE")
}

var RegisterExpensesRoutes = func(router *mux.Router) {
	expenses := router.PathPrefix("/tracker").Subrouter()
	expenses.HandleFunc("/expenses/", expcontroller.GetExpenses).Methods("GET")
	expenses.HandleFunc("/expenses/", expcontroller.CreateExpenses).Methods("POST")
	expenses.HandleFunc("/expenses/:id", expcontroller.GetExpensesById).Methods("GET")
	expenses.HandleFunc("/expenses/:id", expcontroller.UpdateExpenses).Methods("PUT")
	expenses.HandleFunc("/expenses/:id", expcontroller.DeleteExpenses).Methods("DELETE")
}
