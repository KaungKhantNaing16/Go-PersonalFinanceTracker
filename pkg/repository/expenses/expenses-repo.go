package exprepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"errors"
	"log"
)

var ErrExpensesNotFound = errors.New("FromRepository - expenses not found")

type ExpensesRepository struct{}

// Retrieve the list of Incomes record from the database
func (i *ExpensesRepository) GetExpenses() []model.Expenses {
	DB := config.NewDatabase()

	rows, err := DB.Query("SELECT * FROM expenses")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var expenses []model.Expenses
	for rows.Next() {
		var exp model.Expenses
		err := rows.Scan(&exp.ID, &exp.UserID, &exp.CateID, &exp.Title, &exp.Amount, &exp.Description, &exp.Date, &exp.CreatedAt, &exp.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		expenses = append(expenses, exp)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return expenses
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
