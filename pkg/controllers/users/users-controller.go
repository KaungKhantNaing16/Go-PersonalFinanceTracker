package userscontroller

// import (
// 	model "Go-PersonalFinanceTracker/pkg/models"
// 	userservice "Go-PersonalFinanceTracker/pkg/services/users"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"time"

// 	"golang.org/x/crypto/bcrypt"
// )

// var detailService = userservice.UserDetailService{}

// func Registration(writer http.ResponseWriter, request *http.Request) {
// 	if request.Method != http.MethodPost {
// 		http.Error(writer, "Invalid Request Method", http.StatusInternalServerError)
// 		return
// 	}

// 	if err := request.ParseForm(); err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	imgFileName, err := HandleUploadImg(request)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(request.FormValue("password")), 8)
// 	if err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	formData := model.UserDetail{
// 		Status:   1,
// 		Name:     request.FormValue("name"),
// 		Email:    request.FormValue("email"),
// 		Profile:  imgFileName,
// 		Password: hashedPwd,
// 	}

// 	fmt.Println(formData)
// 	if err := detailService.CreateUserDetails(formData); err != nil {
// 		http.Error(writer, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(writer, request, "/signup", http.StatusFound)
// }

// func HandleUploadImg(request *http.Request) (string, error) {
// 	if err := request.ParseMultipartForm(32 << 20); err != nil {
// 		return "", err
// 	}

// 	file, header, err := request.FormFile("profileImg")
// 	if err != nil {
// 		return "", err
// 	}

// 	defer file.Close()

// 	err = os.MkdirAll("static/img/profile", os.ModePerm)
// 	if err != nil {
// 		return "", err
// 	}

// 	imgName := fmt.Sprintf("img-profile_%d%s", time.Now().Unix(), filepath.Ext(header.Filename))
// 	f, err := os.Create(fmt.Sprintf("static/img/profile/%s", imgName))
// 	if err != nil {
// 		return "", err
// 	}

// 	defer f.Close()

// 	_, err = io.Copy(f, file)
// 	if err != nil {
// 		return "", err
// 	}

// 	return imgName, nil
// }
