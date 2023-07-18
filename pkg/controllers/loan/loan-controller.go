package loancontroller

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func loadTemplates(fileName string) {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/loan/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"dataTable.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
	))
}

func GetGiveLoan(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("giveList")
	ttl := "Giving List"
	err := tmpl.ExecuteTemplate(writer, "giveList.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetReceiveLoan(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("receiveList")
	ttl := "Receive List"
	err := tmpl.ExecuteTemplate(writer, "receiveList.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
