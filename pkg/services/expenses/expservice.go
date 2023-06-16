package expservice

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	exprepo "Go-PersonalFinanceTracker/pkg/repository/expenses"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type IExpensesServices interface {
	GetExpenses() ([]models.Expenses, error)
	GetExpensesById(id int) (models.Expenses, error)
	CreateExpenses(expenses models.Expenses) error
	UpdateExpenses(id int, expenses models.Expenses) error
	DeleteExpenses(id int) error
}

type ExpensesService struct {
	expRepo exprepo.IExpensesRepository
}

func (e *ExpensesService) GetExpenses() {

}

func (e *ExpensesService) GetExpensesById() {

}

func (e *ExpensesService) CreateExpenses() {

}

func (e *ExpensesService) UpdateExpenses() {

}

func (e *ExpensesService) DeleteExpenses() {

}
