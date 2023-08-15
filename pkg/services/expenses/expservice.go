package expservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	exprepository "Go-PersonalFinanceTracker/pkg/repository/expenses"
	cateservice "Go-PersonalFinanceTracker/pkg/services/categories"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ErrExpensesNotFound = errors.New("Expense not found")
var ErrIDIsNotValid = errors.New("Id is not valid")

var categoriesService = cateservice.CategoriesService{}

type ExpensesService struct {
	expRepo exprepository.ExpensesRepository
}

func (e *ExpensesService) GetExpenses() ([]model.Expenses, error) {
	var Expenses []model.Expenses
	expenses := e.expRepo.GetExpenses()
	categories, err := categoriesService.GetCategories()
	if err != nil {
		return nil, err
	}

	for _, expense := range expenses {
		for _, category := range categories {
			if expense.CateID == category.ID {
				expense.CateName = category.Title
			}
		}
		Expenses = append(Expenses, expense)
	}
	return Expenses, nil
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

	categories, err := categoriesService.GetCategories()
	if err != nil {
		return expense, err
	}

	for _, category := range categories {
		if expense.CateID == category.ID {
			expense.CateName = category.Title
		}
	}

	return expense, nil
}

func (e *ExpensesService) CreateExpenses(expenses model.Expenses) error {
	return e.expRepo.CreateExpenses(expenses)
}

func (e *ExpensesService) UpdateExpenses(expense model.Expenses) error {
	if expense.ID == 0 {
		return ErrIDIsNotValid
	}
	fmt.Println("Passed Services")
	return e.expRepo.UpdateExpenses(expense)
}
