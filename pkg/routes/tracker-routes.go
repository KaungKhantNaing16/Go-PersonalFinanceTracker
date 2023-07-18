package routes

import (
	authcontroller "Go-PersonalFinanceTracker/pkg/controllers/auth"
	budgetcontroller "Go-PersonalFinanceTracker/pkg/controllers/budget"
	expcontroller "Go-PersonalFinanceTracker/pkg/controllers/expenses"
	incontroller "Go-PersonalFinanceTracker/pkg/controllers/incomes"
	loancontroller "Go-PersonalFinanceTracker/pkg/controllers/loan"
	mediacontroller "Go-PersonalFinanceTracker/pkg/controllers/media"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	// RegisterIncomeRoutes
	router.HandleFunc("/incomes", incontroller.GetIncomes).Methods("GET")
	// router.HandleFunc("/income/create", incontroller.CreateIncome).Methods("POST")
	// router.HandleFunc("/income/:id", incontroller.GetIncomeById).Methods("GET")
	// router.HandleFunc("/income/:id", incontroller.UpdateIncome).Methods("PUT")
	// router.HandleFunc("/income/:id", incontroller.DeleteIncome).Methods("DELETE")

	// RegisterExpensesRoutes
	router.HandleFunc("/expenses", expcontroller.GetExpenses).Methods("GET")
	// 	router.HandleFunc("/expenses/create", expcontroller.CreateExpenses).Methods("POST")
	// 	router.HandleFunc("/expenses/:id", expcontroller.GetExpensesById).Methods("GET")
	// 	router.HandleFunc("/expenses/:id", expcontroller.UpdateExpenses).Methods("PUT")
	// 	router.HandleFunc("/expenses/:id", expcontroller.DeleteExpenses).Methods("DELETE")

	// RegisterBudgetRoutes
	router.HandleFunc("/budget", budgetcontroller.GetBudgetsList)

	// var RegisterLoanRoutes
	router.HandleFunc("/loan/give", loancontroller.GetGiveLoan)
	router.HandleFunc("/loan/recieve", loancontroller.GetReceiveLoan)

	// var RegisterMediaRoutes
	router.HandleFunc("/media", mediacontroller.GetMedia)

	// var RegisterAuthRoutes
	router.HandleFunc("/", authcontroller.Login)
	router.HandleFunc("/signup", authcontroller.Signup)
	router.HandleFunc("/logout", authcontroller.LogoutConfrim)
}
