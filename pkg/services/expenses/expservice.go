package expservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	exprepository "Go-PersonalFinanceTracker/pkg/repository/expenses"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type ExpensesService struct {
	expRepo exprepository.ExpensesRepository
}

func (e *ExpensesService) GetExpenses() []model.Expenses {
	return e.expRepo.GetExpenses()
}

func (e *ExpensesService) GetExpensesById() {

}

func (e *ExpensesService) CreateExpenses() {

}

func (e *ExpensesService) UpdateExpenses() {

}

func (e *ExpensesService) DeleteExpenses() {

}
