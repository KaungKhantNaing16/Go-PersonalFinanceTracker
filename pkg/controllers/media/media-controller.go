package mediacontroller

import (
	mediaservice "Go-PersonalFinanceTracker/pkg/services/media"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var tmpl *template.Template

const MAX_UPLOAD_SIZE = 1024 * 1024

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

var mediaService = mediaservice.MediaService{}

func GetMedia(writer http.ResponseWriter, request *http.Request) {
	userID, _ := request.Cookie("UserID")
	AuthorizeID, err := strconv.Atoi(userID.Value)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	mediaDataArr, err := mediaService.GetMedia(AuthorizeID)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	folderPath := "static/img/uploads"
	var imgArray = []string{}
	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Println(info.Name())
			imgArray = append(imgArray, info.Name())
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(mediaDataArr)

	loadTemplates()
	err = tmpl.ExecuteTemplate(writer, "index.html", mediaDataArr)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleUploadFile(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method", http.StatusBadRequest)
		return
	}

	// Limit the file size to 10MB;
	if err := request.ParseMultipartForm(10 << 20); err != nil {
		http.Error(writer, "The uploaded file is too big.", http.StatusInternalServerError)
		return
	}

	uploadFiles := request.MultipartForm.File["imgfiles"]
	expIdStr := request.FormValue("expID")
	expenseId, err := strconv.Atoi(expIdStr)
	if err != nil {
		http.Error(writer, "Invalid Expense ID value", http.StatusMethodNotAllowed)
		return
	}

	var fileNameArr = []string{}

	for index, fileHeader := range uploadFiles {

		if fileHeader.Size > MAX_UPLOAD_SIZE {
			http.Error(writer, fmt.Sprintf("The uploaded image is too big: %s. Please use an images less than 1MB in size", fileHeader.Filename), http.StatusBadRequest)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		defer file.Close()

		buff := make([]byte, 512)
		_, err = file.Read(buff)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		fileType := http.DetectContentType(buff)
		if fileType != "image/jpg" && fileType != "image/jpeg" && fileType != "image/png" {
			http.Error(writer, "The provided file format is not allowed. Please upload a JPEG or JPG or PNG image", http.StatusBadRequest)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		err = os.MkdirAll("static/img/uploads", os.ModePerm)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		tmpImgName := fmt.Sprintf("img-media0%d_%d%s", index, time.Now().Unix(), filepath.Ext(fileHeader.Filename))
		f, err := os.Create(fmt.Sprintf("static/img/uploads/%s", tmpImgName))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
		fileNameArr = append(fileNameArr, tmpImgName)
		defer f.Close()

		_, err = io.Copy(f, file)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if err = mediaService.CreateMedia(fileNameArr, expenseId); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(writer, request, "/media", http.StatusFound)
}
