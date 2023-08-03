package expcontroller

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	cateservice "Go-PersonalFinanceTracker/pkg/services/categories"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

var tmpl *template.Template
var errors = []string{}
var layout = "2006-01-02"

type FormData struct {
	Expense    *model.Expenses
	Categories *[]model.Category
}

func loadTemplates(fileName string) {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/expenses/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
	))
}

var expensesService = expservice.ExpensesService{}
var categoriesService = cateservice.CategoriesService{}

func GetExpenses(writer http.ResponseWriter, request *http.Request) {
	expenses := expensesService.GetExpenses()

	loadTemplates("list")
	err := tmpl.ExecuteTemplate(writer, "list.html", expenses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateExpenses(writer http.ResponseWriter, request *http.Request) {
	categories, err := categoriesService.GetCategories()
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
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	loadTemplates("detail")
	err = tmpl.ExecuteTemplate(writer, "detail.html", expense)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EditExpenses(writer http.ResponseWriter, request *http.Request) {
	expense, err := expensesService.GetExpensesById(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	categories, err := categoriesService.GetCategories()
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
		http.Error(writer, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	expIdStr := request.FormValue("exp_id")
	var expId int
	if expIdStr != "" {
		if expId, err = strconv.Atoi(expIdStr); err != nil {
			http.Error(writer, "Invalid Expense Id value", http.StatusMethodNotAllowed)
			return
		}
	}

	userIdStr := request.FormValue("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(writer, "Invalid User Id value", http.StatusMethodNotAllowed)
		return
	}

	title := request.Form["title"][0]
	if title == "" {
		http.Error(writer, "Invalid Title value", http.StatusMethodNotAllowed)
		return
	}

	description := request.Form["desc"][0]
	if description == "" {
		http.Error(writer, "Invalid Description value", http.StatusMethodNotAllowed)
		return
	}

	cateStr := request.FormValue("category")
	category, err := strconv.Atoi(cateStr)
	if err != nil {
		http.Error(writer, "Invalid Category value", http.StatusMethodNotAllowed)
		return
	}

	amoutStr := request.FormValue("amount")
	amount, err := strconv.Atoi(amoutStr)
	if err != nil {
		http.Error(writer, "Invalid Amount value", http.StatusMethodNotAllowed)
		return
	}

	dateStr := request.FormValue("date")
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		// errors = append(errors, "Invalid date value")
		http.Error(writer, "Invalid Date value", http.StatusMethodNotAllowed)
		return
	}

	Expenses := model.Expenses{
		UserID:      userId,
		Title:       title,
		Description: description,
		CateID:      category,
		Amount:      amount,
		Date:        date,
	}

	if expId != 0 {
		Expenses.ID = expId
	}
	fmt.Print("Confirm page:")
	fmt.Println(Expenses)
	loadTemplates("confirm")
	err = tmpl.ExecuteTemplate(writer, "confirm.html", Expenses)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

}

func SubmitExpenses(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method, method should be POST.", http.StatusMethodNotAllowed)
		return
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userIdStr := request.FormValue("user_id")
	userId, _ := strconv.Atoi(userIdStr)

	title := request.FormValue("title")

	description := request.FormValue("desc")

	cateStr := request.FormValue("category")
	category, _ := strconv.Atoi(cateStr)

	amoutStr := request.FormValue("amount")
	amount, _ := strconv.Atoi(amoutStr)

	dateStr := request.FormValue("date")
	date, _ := time.Parse(layout, dateStr)

	responseData := model.Expenses{
		UserID:      userId,
		Title:       title,
		Description: description,
		CateID:      category,
		Amount:      amount,
		Date:        date,
	}

	expIdStr := request.FormValue("exp_id")
	fmt.Println(expIdStr)
	fmt.Println(reflect.TypeOf(expIdStr))
	if expIdStr != "0" {
		expId, err := strconv.Atoi(expIdStr)
		if err != nil {
			// http.Error(writer, "Invalid Expense Id value", http.StatusMethodNotAllowed)
			// return
			log.Fatal(err)
		}
		responseData.ID = expId
		fmt.Print("To edit:")
		fmt.Println(responseData)
		if err = expensesService.UpdateExpenses(expId, responseData); err != nil {
			// http.Error(writer, err.Error(), http.StatusInternalServerError)
			// return
			log.Fatal(err)
		}
		http.Redirect(writer, request, "/expenses", http.StatusFound)
	} else {
		fmt.Print("To create:")
		fmt.Println(responseData)
		expensesService.CreateExpenses(responseData)
		http.Redirect(writer, request, "/expenses", http.StatusFound)
	}

}

// func DeleteExpenses(writer http.ResponseWriter, request *http.Request) {
// 	loadTemplates("list")
// 	err := tmpl.ExecuteTemplate(writer, "list.html", expenses)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
