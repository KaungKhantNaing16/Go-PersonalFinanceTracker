package incontroller

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	inservice "Go-PersonalFinanceTracker/pkg/services/incomes"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"github.com/xuri/excelize/v2"
)

var tmpl *template.Template

func loadTemplates(fileName string) {
	templatePartialDir := "templates/partials/"
	templatesDir := "templates/incomes/"

	tmpl = template.Must(template.ParseFiles(
		templatePartialDir+"sideBar.html",
		templatePartialDir+"js.html",
		templatePartialDir+"css.html",
		templatesDir+fileName+".html",
	))
}

var incomeService = inservice.IncomeService{}

func GetIncomes(writer http.ResponseWriter, request *http.Request) {
	incomes := incomeService.GetIncomes()
	loadTemplates("list")
	err := tmpl.ExecuteTemplate(writer, "list.html", incomes)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleUploadFile(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalid Request Method, method should be POST.", http.StatusMethodNotAllowed)
		return
	}
	uploadFile, header, err := request.FormFile("file")
	if err != nil {
		http.Error(writer, "Error reading file", http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}

	defer uploadFile.Close()
	var incomes []model.Income

	switch filepath.Ext(header.Filename) {
	case ".csv":
		incomes, err = parseCSVFile(uploadFile)
	case ".xlsx":
		incomes, err = parseXLSXFile(uploadFile)
	default:
		http.Error(writer, "Unsupported file formt", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(writer, "Error parsing file data", http.StatusInternalServerError)
		return
	}
	fmt.Print("Handle Upload File: ")
	fmt.Println(incomes)
	if err = incomeService.CreateIncomes(incomes); err != nil {
		http.Error(writer, "Error inserting incomes to Database", http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/incomes", http.StatusFound)
}

func parseCSVFile(uploadFile io.Reader) ([]model.Income, error) {
	csvReader := csv.NewReader(uploadFile)
	csvReader.Comma = ','

	if _, err := csvReader.Read(); err != nil && err != io.EOF {
		return nil, err
	}

	var incomes []model.Income
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		userID, _ := strconv.Atoi(record[0])
		amount, _ := strconv.Atoi(record[1])
		currentDate := time.Now()
		income := model.Income{
			UserID:      userID,
			Amount:      amount,
			Title:       record[2],
			Description: record[3],
			FileURL:     record[4],
			CreatedAt:   currentDate,
			UpdatedAt:   currentDate,
		}

		incomes = append(incomes, income)
		fmt.Print("Parse CSV File:")
		fmt.Println(incomes)
	}
	return incomes, nil
}

func parseXLSXFile(uploadFile io.Reader) ([]model.Income, error) {
	file, err := excelize.OpenReader(uploadFile)
	if err != nil {
		return nil, err
	}

	var incomes []model.Income
	sheetName := file.GetSheetName(0)
	rows, err := file.GetRows(sheetName)
	if err != nil {
		return nil, err
	}
	currentDate := time.Now()
	rows = rows[1:]
	fmt.Print("Parse Excel File: ")
	for _, row := range rows {
		userID, _ := strconv.Atoi(row[0])
		amount, _ := strconv.Atoi(row[1])

		income := model.Income{
			UserID:      userID,
			Amount:      amount,
			Title:       row[2],
			Description: row[3],
			FileURL:     row[4],
			CreatedAt:   currentDate,
			UpdatedAt:   currentDate,
		}

		incomes = append(incomes, income)
		fmt.Println(incomes)
	}
	return incomes, nil
}

func GetIncomeDetail(writer http.ResponseWriter, request *http.Request) {
	var income model.Income
	income, err := incomeService.GetIncomeById(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	loadTemplates("detail")
	err = tmpl.ExecuteTemplate(writer, "detail.html", income)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func EditIncome(writer http.ResponseWriter, request *http.Request) {
	var income model.Income
	income, err := incomeService.GetIncomeById(request)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	loadTemplates("edit")
	err = tmpl.ExecuteTemplate(writer, "edit.html", income)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ConfirmIncome(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalide Request Method, method should be POST.", http.StatusBadRequest)
		return
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	incomeIdStr := request.FormValue("incomeId")
	incomeId, _ := strconv.Atoi(incomeIdStr)

	userIdStr := request.FormValue("user_id")
	userId, _ := strconv.Atoi(userIdStr)

	amountStr := request.FormValue("amount")
	amount, _ := strconv.Atoi(amountStr)

	income := model.Income{
		ID:          incomeId,
		UserID:      userId,
		Amount:      amount,
		Title:       request.FormValue("title"),
		Description: request.FormValue("desc"),
		FileURL:     request.FormValue("file_url"),
	}

	loadTemplates("confirm")
	err = tmpl.ExecuteTemplate(writer, "confirm.html", income)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SubmitIncome(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(writer, "Invalide Request Method, method should be POST.", http.StatusBadRequest)
		return
	}

	err := request.ParseForm()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	incomeIdStr := request.FormValue("incomeId")
	incomeId, _ := strconv.Atoi(incomeIdStr)

	userIdStr := request.FormValue("user_id")
	userId, _ := strconv.Atoi(userIdStr)

	amountStr := request.FormValue("amount")
	amount, _ := strconv.Atoi(amountStr)

	currentDate := time.Now()

	income := model.Income{
		ID:          incomeId,
		UserID:      userId,
		Amount:      amount,
		Title:       request.FormValue("title"),
		Description: request.FormValue("desc"),
		FileURL:     request.FormValue("file_url"),
		CreatedAt:   currentDate,
		UpdatedAt:   currentDate,
	}

	if err := incomeService.UpdateIncome(income); err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(writer, request, "/incomes", http.StatusFound)
}
