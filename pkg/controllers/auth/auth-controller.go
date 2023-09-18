package authcontroller

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	userservice "Go-PersonalFinanceTracker/pkg/services/users"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var tmpl *template.Template
var detailService = userservice.UserDetailService{}

type FlashMessage struct {
	SuccessMsg, ErrorMsg bool
}

func loadTemplates(fileName string) {
	templatesDir := "templates/auth/"
	templatePartialDir := "templates/partials/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
		"templates/error.html",
	))
}

func Login(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("login")
	flashMsg := FlashMessage{
		SuccessMsg: false,
		ErrorMsg:   false,
	}
	err := tmpl.ExecuteTemplate(writer, "login.html", flashMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CheckCredentials(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method", http.StatusBadRequest)
		return
	}

	var credsErr []string

	request.ParseForm()
	pwd := request.FormValue("password")
	email := request.FormValue("email")
	if email == "" {
		http.Error(writer, "Invalid Email Value", http.StatusInternalServerError)
		return
	}

	userData, err := detailService.GetUserDetailByEmail(email)
	if err != nil {
		credsErr = append(credsErr, "Invalid Email Address")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(pwd)); err != nil {
		credsErr = append(credsErr, "Invalid Password")
	}

	if len(credsErr) > 0 {
		loadTemplates("login")
		flashMsg := FlashMessage{
			SuccessMsg: false,
			ErrorMsg:   true,
		}
		tmpl.ExecuteTemplate(writer, "login.html", flashMsg)
	}

	http.SetCookie(writer, &http.Cookie{Name: "UserID", Value: strconv.Itoa(userData.ID)})
	http.Redirect(writer, request, "/dashboard/", http.StatusFound)
}

func Signup(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("signup")
	ttl := "Sign-up"
	err := tmpl.ExecuteTemplate(writer, "signup.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Registration(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		tmpl.ExecuteTemplate(writer, "error.html", "Invalid Request Method")
	}

	if err := request.ParseForm(); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	imgFileName, err := HandleUploadImg(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), 8)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	formData := model.UserDetail{
		Status:   1,
		Name:     request.FormValue("name"),
		Email:    request.FormValue("email"),
		Job:      request.FormValue("job"),
		Profile:  imgFileName,
		Password: hashedPwd,
	}

	if err := detailService.CreateUserDetails(formData); err != nil {
		tmpl.ExecuteTemplate(writer, "error.html", err.Error())
	}

	loadTemplates("login")
	flashMsg := FlashMessage{
		SuccessMsg: true,
		ErrorMsg:   false,
	}
	err = tmpl.ExecuteTemplate(writer, "login.html", flashMsg)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleUploadImg(request *http.Request) (string, error) {
	if err := request.ParseMultipartForm(32 << 20); err != nil {
		return "", err
	}

	file, header, err := request.FormFile("profileImg")
	if err != nil {
		return "", err
	}

	defer file.Close()

	err = os.MkdirAll("static/img/profile", os.ModePerm)
	if err != nil {
		return "", err
	}

	imgName := fmt.Sprintf("img-profile_%d%s", time.Now().Unix(), filepath.Ext(header.Filename))
	f, err := os.Create(fmt.Sprintf("static/img/profile/%s", imgName))
	if err != nil {
		return "", err
	}

	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", err
	}

	return imgName, nil
}

func ConfirmLogout(writer http.ResponseWriter, request *http.Request) {
	loadTemplates("logout-confirm")
	ttl := "logout-confirm"
	err := tmpl.ExecuteTemplate(writer, "logout-confirm.html", ttl)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Logout(writer http.ResponseWriter, request *http.Request) {
	expire := time.Now().Add(-7 * 24 * time.Hour)
	http.SetCookie(writer, &http.Cookie{Name: "UserID", Value: "", Expires: expire, HttpOnly: true})
	http.Redirect(writer, request, "/login", http.StatusFound)
}
