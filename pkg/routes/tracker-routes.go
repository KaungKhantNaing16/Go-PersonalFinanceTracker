package routes

import (
	authcontroller "Go-PersonalFinanceTracker/pkg/controllers/auth"
	budgetcontroller "Go-PersonalFinanceTracker/pkg/controllers/budget"
	expcontroller "Go-PersonalFinanceTracker/pkg/controllers/expenses"
	incontroller "Go-PersonalFinanceTracker/pkg/controllers/incomes"
	mediacontroller "Go-PersonalFinanceTracker/pkg/controllers/media"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	// RegisterIncomeRoutes
	router.HandleFunc("/incomes", incontroller.GetIncomes).Methods("GET")
	router.HandleFunc("/incomes/upload", incontroller.HandleUploadFile).Methods("POST")
	router.HandleFunc("/incomes/{id}", incontroller.GetIncomeDetail).Methods("GET")
	router.HandleFunc("/incomes/edit/{id}", incontroller.EditIncome)
	router.HandleFunc("/incomes/confirm", incontroller.ConfirmIncome).Methods("POST")
	router.HandleFunc("/incomes/submit", incontroller.SubmitIncome).Methods("POST")

	// RegisterExpensesRoutes
	router.HandleFunc("/expenses", expcontroller.GetExpenses).Methods("GET")
	router.HandleFunc("/expenses/create", expcontroller.CreateExpenses).Methods("GET")
	router.HandleFunc("/expenses/{id}", expcontroller.GetExpenseDetail).Methods("GET")
	router.HandleFunc("/expenses/edit/{id}", expcontroller.EditExpenses).Methods("GET")
	router.HandleFunc("/expenses/confirm", expcontroller.ConfirmExpense).Methods("POST")
	router.HandleFunc("/expenses/submit", expcontroller.SubmitExpenses).Methods("POST")

	// RegisterBudgetRoutes
	router.HandleFunc("/budget", budgetcontroller.GetBudgetsList).Methods("GET")
	router.HandleFunc("/budget/create", budgetcontroller.CreateBudgetPlan).Methods("POST")
	router.HandleFunc("/budget/delete/{id}", budgetcontroller.DeleteBudgetPlan).Methods("GET")

	// RegisterMediaRoutes
	router.HandleFunc("/media", mediacontroller.GetMedia).Methods("GET")
	router.HandleFunc("/media/upload", mediacontroller.HandleUploadFile).Methods("POST")

	// RegisterAuthRoutes
	router.HandleFunc("/", authcontroller.Login)
	router.HandleFunc("/signup", authcontroller.Signup)
	router.HandleFunc("/logout", authcontroller.LogoutConfrim)
}
