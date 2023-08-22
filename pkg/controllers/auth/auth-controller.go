package authcontroller

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func loadTemplates(fileName string) {
	templatesDir := "templates/auth/"
	templatePartialDir := "templates/partials/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
	))
}

func Login(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("login")
	ttl := "login"
	err := tmpl.ExecuteTemplate(writer, "login.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("signup")
	ttl := "signup"
	err := tmpl.ExecuteTemplate(writer, "signup.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func LogoutConfrim(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("logout-confirm")
	ttl := "logout-confirm"
	err := tmpl.ExecuteTemplate(writer, "logout-confirm.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}
