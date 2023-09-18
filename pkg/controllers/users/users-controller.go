package userscontroller

import (
	userservice "Go-PersonalFinanceTracker/pkg/services/users"
	"html/template"
)

var detailService = userservice.UserDetailService{}

var tmpl *template.Template

func loadTemplates() {
	templatePartialDir := "templates/partials/"
	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
	))
}

// func GetUserDetail(writer http.ResponseWriter, authId int) error {
// 	userDetail, err := detailService.GetUserDetailByID(authId)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	return nil
// }
