package exprepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
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
		err := rows.Scan(&exp.ID, &exp.UserID, &exp.CateID, &exp.Amount, &exp.Title, &exp.Description, &exp.Date, &exp.CreatedAt, &exp.UpdatedAt)
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

func (e *ExpensesRepository) GetExpensesById(expId int) (model.Expenses, error) {
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM expenses WHERE id=?", expId)

	expense := model.Expenses{}
	err := row.Scan(&expense.ID, &expense.UserID, &expense.CateID, &expense.Amount, &expense.Title, &expense.Description, &expense.Date, &expense.CreatedAt, &expense.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return expense, ErrExpensesNotFound
		}
	}
	fmt.Println(expense)
	fmt.Println("Passed - Get by ID")
	return expense, nil
}

func (e *ExpensesRepository) CreateExpenses(expenses model.Expenses) error {
	currentDateTime := time.Now()
	CreateAt := currentDateTime
	UpdatedAt := currentDateTime

	DB := config.NewDatabase()
	_, err := DB.Query("INSERT INTO expenses (uid, cate_id, amount, title, description, date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", expenses.UserID, expenses.CateID, expenses.Amount, expenses.Title, expenses.Description, expenses.Date, CreateAt, UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (e *ExpensesRepository) UpdateExpenses(Id int, expenses model.Expenses) error {
	fmt.Println(Id)
	oldExp, err := e.GetExpensesById(Id)
	if err != nil {
		return ErrExpensesNotFound
	}

	oldExp.Title = expenses.Title
	oldExp.Description = expenses.Description
	oldExp.CateID = expenses.CateID
	oldExp.Amount = expenses.Amount
	oldExp.Date = expenses.Date

	DB := config.NewDatabase()
	updateQuery := "UPDATE expenses SET  uid=?, cate_id=?, amount=?, title=?, description=?, date=? WHERE id=?"
	_, err = DB.Exec(updateQuery, oldExp.UserID, oldExp.CateID, oldExp.Amount, oldExp.Title, oldExp.Description, oldExp.Date, oldExp.ID)
	if err != nil {
		return err
	}
	fmt.Println("Passed Repository")
	return nil
}

func (e *ExpensesRepository) DeleteExpenses() {
	// Delete the expenses record from the database
}
