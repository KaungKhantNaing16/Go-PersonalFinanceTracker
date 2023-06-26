package authcontroller

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
	templatesDir := "templates/auth/"
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
		t, err := template.ParseFiles(templatesDir + name + ".html")
		if err == nil {
			tmpl[name] = t
			fmt.Println("Load Template", index, name)
		} else {
			fmt.Printf("%s error occurred while parse files.", err)
			return
		}
	}
}

func Login(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["login"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["signup"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LogoutConfrim(writer http.ResponseWriter, request *http.Request) {
	loadTemplates()
	err := tmpl["logout-confirm"].Execute(writer, nil)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
