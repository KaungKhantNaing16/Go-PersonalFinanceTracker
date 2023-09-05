package dashboardcontroller

import (
	mail "Go-PersonalFinanceTracker/pkg/mails"
	dashboardservice "Go-PersonalFinanceTracker/pkg/services/dashboard"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl *template.Template

func loadTemplates() {
	templatePartialDir := "templates/partials/"
	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		"templates/index.html",
	))
}

func DashboardHandler(writer http.ResponseWriter, req *http.Request) {
	userID, _ := req.Cookie("UserID")
	AuthorizeID, err := strconv.Atoi(userID.Value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	loadTemplates()
	oaDataArr, err := dashboardservice.DashboardService(AuthorizeID)
	if err != nil {
		tmpl.ExecuteTemplate(writer, "error.html", err.Error())
	}

	err = tmpl.ExecuteTemplate(writer, "index.html", oaDataArr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SendAlert(expAmt int, inAmt int) {
	floorAmt := expAmt - inAmt

	if floorAmt < 10000 {
		mail.SendAlertMail()
	}
}
