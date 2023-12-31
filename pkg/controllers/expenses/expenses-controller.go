package expcontroller

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	cateservice "Go-PersonalFinanceTracker/pkg/services/categories"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	mediaservice "Go-PersonalFinanceTracker/pkg/services/media"
	request_validation "Go-PersonalFinanceTracker/pkg/validations"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl *template.Template

var expensesService = expservice.ExpensesService{}
var mediaService = mediaservice.MediaService{}
var categoriesService = cateservice.CategoriesService{}

type FormData struct {
	Expense    *model.Expenses
	Categories *[]model.Category
}

type DetailData struct {
	Expense    model.Expenses
	Media      []string
	IsHasMedia bool
}

func loadTemplates(fileName string) {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/expenses/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
		"templates/error.html",
	))
}

func GetExpenses(writer http.ResponseWriter, request *http.Request) {
	userID, _ := request.Cookie("UserID")
	AuthorizeID, err := strconv.Atoi(userID.Value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	expenses, err := expensesService.GetExpenses(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	loadTemplates("list")
	err = tmpl.ExecuteTemplate(writer, "list.html", expenses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateExpenses(writer http.ResponseWriter, request *http.Request) {
	userID, _ := request.Cookie("UserID")
	AuthorizeID, err := strconv.Atoi(userID.Value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := categoriesService.GetCategories(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	loadTemplates("create")
	err = tmpl.ExecuteTemplate(writer, "create.html", categories)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetExpenseDetail(writer http.ResponseWriter, request *http.Request) {
	var expense model.Expenses
	expense, err := expensesService.GetExpensesById(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	mediaArr, err := mediaService.GetMediaByExpId(expense.ID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	hasMedia := true

	if mediaArr == nil {
		hasMedia = false
	}

	detailData := DetailData{
		Expense:    expense,
		Media:      mediaArr,
		IsHasMedia: hasMedia,
	}

	loadTemplates("detail")
	err = tmpl.ExecuteTemplate(writer, "detail.html", detailData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EditExpenses(writer http.ResponseWriter, request *http.Request) {
	expense, err := expensesService.GetExpensesById(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userID, _ := request.Cookie("UserID")
	AuthorizeID, err := strconv.Atoi(userID.Value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	categories, err := categoriesService.GetCategories(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData := FormData{
		Expense:    &expense,
		Categories: &categories,
	}

	loadTemplates("edit")
	err = tmpl.ExecuteTemplate(writer, "edit.html", responseData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConfirmExpense(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method", http.StatusInternalServerError)
		return
	}

	validatedData := request_validation.ExpensesRequestValiation(writer, request)

	loadTemplates("confirm")
	err := tmpl.ExecuteTemplate(writer, "confirm.html", validatedData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SubmitExpenses(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method", http.StatusInternalServerError)
		return
	}

	validatedData := request_validation.ExpensesRequestValiation(writer, request)

	if validatedData.ID != 0 {
		fmt.Println("To update expense")
		if err := expensesService.UpdateExpenses(validatedData); err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(writer, request, "/dashboard/expenses", http.StatusFound)
	} else {
		fmt.Println("To create expense")
		expensesService.CreateExpenses(validatedData)
		http.Redirect(writer, request, "/dashboard/expenses", http.StatusFound)
	}
}
