package request_validation

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	"log"
	"net/http"
	"strconv"
	"time"
)

var layout = "2006-01-02"

func ExpensesRequestValiation(writer http.ResponseWriter, request *http.Request) model.Expenses {
	request.ParseForm()

	var expID int
	expIdStr := request.FormValue("exp_id")
	if expIdStr != "" {
		expID, _ = strconv.Atoi(expIdStr)
		if expID == 0 {
			log.Fatal("Invalid expense id value")
		}
	}

	userIdStr := request.FormValue("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		log.Fatal("Invalid user id value")
	}

	title := request.Form["title"][0]
	if title == "" {
		log.Fatal("Invalid title value")
	}

	description := request.Form["desc"][0]

	cateStr := request.FormValue("category")
	category, err := strconv.Atoi(cateStr)
	if err != nil {
		log.Fatal("Invalid category value")
	}

	amoutStr := request.FormValue("amount")
	amount, err := strconv.Atoi(amoutStr)
	if err != nil {
		log.Fatal("Invalid amount value")
	}

	dateStr := request.FormValue("date")
	date, err := time.Parse(layout, dateStr)
	if err != nil {
		log.Fatal("Invalid date value")
	}

	validatedData := model.Expenses{
		ID:          expID,
		UserID:      userId,
		Title:       title,
		Description: description,
		CateID:      category,
		Amount:      amount,
		Date:        date,
	}
	return validatedData
}
