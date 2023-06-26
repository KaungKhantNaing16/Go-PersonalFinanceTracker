package mediacontroller

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

var tmpl = make(map[string]*template.Template)
var fileNames []string

func loadTemplates() {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/media/"
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
		t, err := template.ParseFiles(templatePartialDir+"layout.html", templatePartialDir+"dataTable.html", templatesDir+name+".html")
		if err == nil {
			tmpl[name] = t
			fmt.Println("Load Template", index, name)
		} else {
			fmt.Printf("%s error occurred while parse files.", err)
			return
		}
	}
}

func GetMedia(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["index"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
