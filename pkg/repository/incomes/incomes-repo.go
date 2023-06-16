package inrepo

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	"errors"
)

var ErrExpensesNotFound = errors.New("FromRepository - expenses not found")

type IIncomeRepository interface {
	GetIncomes() ([]models.Income, error)
	GetIncomeById() (models.Income, error)
	CreateIncome() error
	UpdateIncome() error
	DeleteIncome() error
}

type IncomeRepository struct {
	Incomes []models.Income
}

func (i *IncomeRepository) GetIncomes() {
	// Retrieve the list of Incomes record from the database
}

func (i *IncomeRepository) GetIncomeById() {
	// Retrieve Income record from the database
}

func (i *IncomeRepository) CreateIncome() {
	// Insert the Income record into the database
}

func (i *IncomeRepository) UpdateIncome() {
	// Update the Income record from the database
}

func (i *IncomeRepository) DeleteIncome() {
	// Delete the Income record from the database
}
