package expcontroller

import (
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	"html/template"
	"net/http"
)

var tmpl *template.Template

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

func GetExpenses(writer http.ResponseWriter, request *http.Request) {
	// expenses := expensesService.GetExpenses()
	ttl := "Expenses"
	loadTemplates("list")
	err := tmpl.ExecuteTemplate(writer, "list.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func GetExpensesById(writer http.ResponseWriter, request *http.Request) {
// loadTemplates("list")
// err := tmpl.ExecuteTemplate(writer, "list.html", expenses)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func CreateExpenses(writer http.ResponseWriter, request *http.Request) {
// loadTemplates("list")
// err := tmpl.ExecuteTemplate(writer, "list.html", expenses)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func UpdateExpenses(writer http.ResponseWriter, request *http.Request) {
// loadTemplates("list")
// err := tmpl.ExecuteTemplate(writer, "list.html", expenses)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func DeleteExpenses(writer http.ResponseWriter, request *http.Request) {
// loadTemplates("list")
// err := tmpl.ExecuteTemplate(writer, "list.html", expenses)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
