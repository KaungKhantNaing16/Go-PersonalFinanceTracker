package budgetcontroller

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

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

func GetBudgetsList(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	ttl := "Budget List"
	err := tmpl.ExecuteTemplate(writer, "budget.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
