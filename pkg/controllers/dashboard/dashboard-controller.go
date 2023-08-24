package dashboardcontroller

import (
	mail "Go-PersonalFinanceTracker/pkg/mails"
	model "Go-PersonalFinanceTracker/pkg/models"
	userservice "Go-PersonalFinanceTracker/pkg/services/users"
	"net/http"
	"strconv"
	"text/template"
)

var tmpl *template.Template
var detailService = userservice.UserDetailService{}

type TotalAmountData struct {
	Expenses, Incomes int
}

type ProfileData struct {
	TotalAmount TotalAmountData
	UserDetail  model.UserDetail
}

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
	totalExp, err := detailService.GetExpAmtByUserId(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	totalIncomes, err := detailService.GetIncomesAmtByUserId(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userDetail, err := detailService.GetUserDetailByID(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	totalAmount := TotalAmountData{
		Expenses: totalExp,
		Incomes:  totalIncomes,
	}

	userProfileData := ProfileData{
		TotalAmount: totalAmount,
		UserDetail:  userDetail,
	}

	err = tmpl.ExecuteTemplate(writer, "index.html", userProfileData)
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
