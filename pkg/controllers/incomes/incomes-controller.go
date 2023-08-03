package incontroller

import (
	inservice "Go-PersonalFinanceTracker/pkg/services/incomes"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func loadTemplates(fileName string) {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/incomes/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
	))
}

var incomeService = inservice.IncomeService{}

func GetIncomes(writer http.ResponseWriter, request *http.Request) {
	incomes := incomeService.GetIncomes()
	loadTemplates("list")
	err := tmpl.ExecuteTemplate(writer, "list.html", incomes)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func GetIncomeById(writer http.ResponseWriter, request *http.Request) {
// 	loadTemplates()
// 	err := tmpl["detail"].Execute(writer, nil)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func CreateIncome(writer http.ResponseWriter, request *http.Request) {
// 	loadTemplates()
// 	err := tmpl["create"].Execute(writer, nil)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func UpdateIncome(writer http.ResponseWriter, request *http.Request) {
// 	loadTemplates()
// 	err := tmpl["edit"].Execute(writer, nil)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }

// func DeleteIncome(writer http.ResponseWriter, request *http.Request) {
// 	loadTemplates()
// 	err := tmpl["index"].Execute(writer, nil)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }
