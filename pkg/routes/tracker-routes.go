package routes

import (
	authcontroller "Go-PersonalFinanceTracker/pkg/controllers/auth"
	budgetcontroller "Go-PersonalFinanceTracker/pkg/controllers/budget"
	expcontroller "Go-PersonalFinanceTracker/pkg/controllers/expenses"
	incontroller "Go-PersonalFinanceTracker/pkg/controllers/incomes"
	loancontroller "Go-PersonalFinanceTracker/pkg/controllers/loan"

	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/", authcontroller.Login)
	router.HandleFunc("/signup", authcontroller.Signup)
	router.HandleFunc("/logout", authcontroller.LogoutConfrim)
}

var RegisterIncomeRoutes = func(router *mux.Router) {
	router.HandleFunc("/incomes/", incontroller.GetIncomes)
	router.HandleFunc("/income/create", incontroller.CreateIncome)
	router.HandleFunc("/income/detail", incontroller.GetIncomeById)
	router.HandleFunc("/income/update", incontroller.UpdateIncome)
	router.HandleFunc("/income/delete", incontroller.DeleteIncome)
}

var RegisterExpensesRoutes = func(router *mux.Router) {
	router.HandleFunc("/expenses/", expcontroller.GetExpenses)
	router.HandleFunc("/expenses/create", expcontroller.CreateExpenses)
	router.HandleFunc("/expenses/detail", expcontroller.GetExpensesById)
	router.HandleFunc("/expenses/update", expcontroller.UpdateExpenses)
	router.HandleFunc("/expenses/delete", expcontroller.DeleteExpenses)
}

var RegisterLoanRoutes = func(router *mux.Router) {
	router.HandleFunc("/loan/give", loancontroller.GetGiveLoan)
	router.HandleFunc("/loan/recieve", loancontroller.GetReceiveLoan)
}

var RegisterBudgetRoutes = func(router *mux.Router) {
	router.HandleFunc("/budget", budgetcontroller.GetBudgetsList)
}
