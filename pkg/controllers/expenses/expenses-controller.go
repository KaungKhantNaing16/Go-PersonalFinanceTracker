package expcontroller

import (
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"text/template"
)

var tmpl = make(map[string]*template.Template)
var fileNames []string

type expensesHandler struct {
	service expservice.IExpensesServices
}

func loadTemplates() {
	templatesDir := "templates/expenses/"
	pattern := templatesDir + "*.html"
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Printf("%s error occurred while get template name.", err)
		return
	}

	for _, match := range matches {
		fileName := filepath.Base(strings.TrimSuffix(match, ".html"))
		fileNames = append(fileNames, fileName)
	}

	for index, name := range fileNames {
		t, err := template.ParseFiles("templates/partials/layout.html", templatesDir+name+".html")
		if err == nil {
			tmpl[name] = t
			fmt.Println("Load Template", index, name)
		} else {
			fmt.Printf("%s error occurred while parse files.", err)
			return
		}
	}
}

func GetExpenses(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["list"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetExpensesById(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["detail"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreateExpenses(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["create"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateExpenses(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["update"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteExpenses(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["index"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
