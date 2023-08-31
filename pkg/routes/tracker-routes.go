package routes

import (
	authcontroller "Go-PersonalFinanceTracker/pkg/controllers/auth"
	budgetcontroller "Go-PersonalFinanceTracker/pkg/controllers/budget"
	dashboardcontroller "Go-PersonalFinanceTracker/pkg/controllers/dashboard"
	expcontroller "Go-PersonalFinanceTracker/pkg/controllers/expenses"
	incontroller "Go-PersonalFinanceTracker/pkg/controllers/incomes"
	mediacontroller "Go-PersonalFinanceTracker/pkg/controllers/media"
	mw "Go-PersonalFinanceTracker/pkg/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRouters(router *mux.Router) {
	router.Handle("/", http.RedirectHandler("/dashboard/", http.StatusSeeOther))

	// RegisterAuthRoutes
	router.HandleFunc("/login", authcontroller.Login)
	router.HandleFunc("/check", authcontroller.CheckCredentials)
	router.HandleFunc("/signup", authcontroller.Signup)
	router.HandleFunc("/register", authcontroller.Registration)
	router.HandleFunc("/logout", authcontroller.Logout)

	subrouter := router.PathPrefix("/dashboard").Subrouter()
	subrouter.Use(mw.AuthMiddleware)

	// RegisterDashboardroutes
	subrouter.HandleFunc("/", dashboardcontroller.DashboardHandler)

	// RegisterIncomeroutes
	subrouter.HandleFunc("/incomes", incontroller.GetIncomes).Methods("GET")
	subrouter.HandleFunc("/incomes/upload", incontroller.HandleUploadFile).Methods("POST")
	subrouter.HandleFunc("/incomes/{id}", incontroller.GetIncomeDetail).Methods("GET")
	subrouter.HandleFunc("/incomes/edit/{id}", incontroller.EditIncome)
	subrouter.HandleFunc("/incomes/confirm", incontroller.ConfirmIncome).Methods("POST")
	subrouter.HandleFunc("/incomes/submit", incontroller.SubmitIncome).Methods("POST")

	// RegisterExpensesroutes
	subrouter.HandleFunc("/expenses", expcontroller.GetExpenses).Methods("GET")
	subrouter.HandleFunc("/expenses/create", expcontroller.CreateExpenses).Methods("GET")
	subrouter.HandleFunc("/expenses/{id}", expcontroller.GetExpenseDetail).Methods("GET")
	subrouter.HandleFunc("/expenses/edit/{id}", expcontroller.EditExpenses).Methods("GET")
	subrouter.HandleFunc("/expenses/confirm", expcontroller.ConfirmExpense).Methods("POST")
	subrouter.HandleFunc("/expenses/submit", expcontroller.SubmitExpenses).Methods("POST")

	// RegisterBudgetroutes
	subrouter.HandleFunc("/budget", budgetcontroller.GetBudgetsList).Methods("GET")
	subrouter.HandleFunc("/budget/create", budgetcontroller.CreateBudgetPlan).Methods("POST")
	subrouter.HandleFunc("/budget/delete/{id}", budgetcontroller.DeleteBudgetPlan).Methods("GET")

	// RegisterMediarroutes
	subrouter.HandleFunc("/media", mediacontroller.GetMedia).Methods("GET")
	subrouter.HandleFunc("/media/upload", mediacontroller.HandleUploadFile).Methods("POST")

	// RegisterLogoutRoutes
	subrouter.HandleFunc("/confirm", authcontroller.ConfirmLogout)
}
