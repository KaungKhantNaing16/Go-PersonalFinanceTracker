package exprepo

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	"errors"
)

var ErrExpensesNotFound = errors.New("FromRepository - expenses not found")

type IExpensesRepository interface {
	GetExpenses() ([]models.Expenses, error)
	GetExpensesById(id int) (models.Expenses, error)
	CreateExpenses(expenses models.Expenses) error
	UpdateExpenses(id int, expenses models.Expenses) error
	DeleteExpenses(id int) error
}

type ExpensesRepository struct {
	Expenses []models.Expenses
}

func (e *ExpensesRepository) GetExpenses() {
	// Retrieve the list of expenses record from the database
}

func (e *ExpensesRepository) GetExpensesById() {
	// Retrieve expenses record from the database
}

func (e *ExpensesRepository) CreateExpenses() {
	// Insert the expenses record into the database
}

func (e *ExpensesRepository) UpdateExpenses() {
	// Update the expenses record from the database
}

func (e *ExpensesRepository) DeleteExpenses() {
	// Delete the expenses record from the database
}
