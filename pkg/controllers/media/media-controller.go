package mediacontroller

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func loadTemplates() {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/media/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"dataTable.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+"index.html",
	))
}

func GetMedia(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	ttl := "Media List"
	err := tmpl.ExecuteTemplate(writer, "index.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
