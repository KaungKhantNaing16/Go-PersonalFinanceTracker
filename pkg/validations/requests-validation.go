package request_validation

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	"fmt"
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
			fmt.Println("Invalid expense id value")
		}
	}

	userIdStr := request.FormValue("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("Invalid user id value")
	}

	title := request.Form["title"][0]
	if title == "" {
		fmt.Println("Invalid title value")
	}

	description := request.Form["desc"][0]

	cateStr := request.FormValue("category")
	category, err := strconv.Atoi(cateStr)
	if err != nil {
		fmt.Println("Invalid category value")
	}

	amoutStr := request.FormValue("amount")
	amount, err := strconv.Atoi(amoutStr)
	if err != nil {
		fmt.Println("Invalid amount value")
	}

	dateStr := request.FormValue("date")
	date, err := time.Parse(layout, dateStr)
	if date.IsZero() {
		fmt.Println("Invalid date format")
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
