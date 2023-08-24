package budgetcontroller

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	budgetservice "Go-PersonalFinanceTracker/pkg/services/budget"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tmpl *template.Template
var ErrIDIsNotValid = errors.New("Id is not valid")

func loadTemplates() {
	templatePartialDir := "templates/partials/"
	templateDir := "templates/budget/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"dataTable.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templateDir+"budget.html",
	))
}

var budgetPlanService = budgetservice.BudgetService{}

func GetBudgetsList(writer http.ResponseWriter, request *http.Request) {
	userID, _ := request.Cookie("UserID")
	AuthorizeID, err := strconv.Atoi(userID.Value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	budgetPlan, err := budgetPlanService.GetBudgetsList(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	loadTemplates()
	err = tmpl.ExecuteTemplate(writer, "budget.html", budgetPlan)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateBudgetPlan(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method", http.StatusBadRequest)
		return
	}

	if err := request.ParseForm(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userId, err := strconv.Atoi(request.FormValue("user_id"))
	if err != nil {
		http.Error(writer, "Invalid User Id value", http.StatusMethodNotAllowed)
		return
	}

	amount, err := strconv.Atoi(request.FormValue("amount"))
	if err != nil {
		http.Error(writer, "Invalid Amount value", http.StatusMethodNotAllowed)
		return
	}

	budgetPlan := model.Budget{
		UserID:   userId,
		Title:    request.FormValue("title"),
		Category: request.FormValue("category"),
		Amount:   amount,
	}

	if err = budgetPlanService.CreateBudgetPlan(budgetPlan); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(writer, request, "/budget", http.StatusFound)
}

func DeleteBudgetPlan(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(ErrIDIsNotValid)
	}

	if err = budgetPlanService.DeleteBudgetPlan(id); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(writer, request, "/budget", http.StatusFound)
}
