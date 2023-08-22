package dashboardcontroller

import (
	mail "Go-PersonalFinanceTracker/pkg/mails"
	model "Go-PersonalFinanceTracker/pkg/models"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	inservice "Go-PersonalFinanceTracker/pkg/services/incomes"
	userservice "Go-PersonalFinanceTracker/pkg/services/users"
	"fmt"
	"net/http"
	"text/template"
)

var tmpl *template.Template
var expensesService = expservice.ExpensesService{}
var incomeService = inservice.IncomeService{}
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
	loadTemplates()
	expTotal, err := expensesService.GetTotalAmount()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	inTotal, err := incomeService.GetTotalAmount()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	userDetail, err := detailService.GetUserDetailById(2)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	totalAmount := TotalAmountData{
		Expenses: expTotal,
		Incomes:  inTotal,
	}

	userProfileData := ProfileData{
		TotalAmount: totalAmount,
		UserDetail:  userDetail,
	}
	fmt.Println(userProfileData)
	err = tmpl.ExecuteTemplate(writer, "index.html", userProfileData)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	// SendAlert(expTotal, inTotal)
}

func SendAlert(expAmt int, inAmt int) {
	floorAmt := expAmt - inAmt

	if floorAmt < 10000 {
		mail.SendAlertMail()
	}
}
