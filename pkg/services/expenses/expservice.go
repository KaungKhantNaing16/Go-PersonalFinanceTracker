package expservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	exprepository "Go-PersonalFinanceTracker/pkg/repository/expenses"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ErrExpensesNotFound = errors.New("Expense not found")
var ErrIDIsNotValid = errors.New("Id is not valid")

type ExpensesService struct {
	expRepo exprepository.ExpensesRepository
}

func (e *ExpensesService) GetExpenses() []model.Expenses {
	return e.expRepo.GetExpenses()
}

func (e *ExpensesService) GetExpensesById(request *http.Request) (model.Expenses, error) {
	params := mux.Vars(request)
	expId, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Print(err)
	}

	expense, err := e.expRepo.GetExpensesById(expId)
	if err != nil {
		return expense, ErrExpensesNotFound
	}

	return expense, nil
}

func (e *ExpensesService) CreateExpenses(expenses model.Expenses) error {
	return e.expRepo.CreateExpenses(expenses)
}

func (e *ExpensesService) UpdateExpenses(Id int, expense model.Expenses) error {
	if Id == 0 {
		return ErrIDIsNotValid
	}
	fmt.Println("Passed Services")
	return e.expRepo.UpdateExpenses(Id, expense)
}

func (e *ExpensesService) DeleteExpenses() {

}
