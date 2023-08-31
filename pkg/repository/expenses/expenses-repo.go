package exprepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"context"
	"database/sql"
	"errors"
	"log"
)

var ErrInvalidUserID = errors.New("FromRepository - Invalid ID Value")
var ErrExpensesNotFound = errors.New("FromRepository - expenses not found")

type ExpensesRepository struct{}

// Retrieve the list of Incomes record from the database
func (i *ExpensesRepository) GetExpenses(userId int) []model.Expenses {
	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT * FROM expenses WHERE uid = ?", userId)
	if err != nil {
		log.Println("Error querying the row", err)
	}
	defer rows.Close()
	var expenses []model.Expenses
	for rows.Next() {
		var exp model.Expenses
		err := rows.Scan(&exp.ID, &exp.UserID, &exp.CateID, &exp.Amount, &exp.Title, &exp.Description, &exp.Date, &exp.CreatedAt, &exp.UpdatedAt)
		if err != nil {
			log.Println("Error scaning the row", err)
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

	return expense, nil
}

func (e *ExpensesRepository) CreateExpenses(expenses model.Expenses) error {
	DB := config.NewDatabase()
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	insertQuery, err := transaction.PrepareContext(ctx, "INSERT INTO expenses (uid, cate_id, amount, title, description, date, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	_, err = insertQuery.Exec(expenses.UserID, expenses.CateID, expenses.Amount, expenses.Title, expenses.Description, expenses.Date, config.CreateAt, config.UpdatedAt)
	if err != nil {
		if err = transaction.Rollback(); err != nil {
			log.Println("Error rolling back transaction on the insertion", err)
		}
	}

	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}

func (e *ExpensesRepository) UpdateExpenses(expenses model.Expenses) error {
	oldExp, err := e.GetExpensesById(expenses.ID)
	if err != nil {
		return ErrExpensesNotFound
	}

	oldExp.Title = expenses.Title
	oldExp.Description = expenses.Description
	oldExp.CateID = expenses.CateID
	oldExp.Amount = expenses.Amount
	oldExp.Date = expenses.Date

	DB := config.NewDatabase()
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	updateQuery, err := transaction.PrepareContext(ctx, "UPDATE expenses SET  uid=?, cate_id=?, amount=?, title=?, description=?, date=? WHERE id=?")
	_, err = updateQuery.Exec(oldExp.UserID, oldExp.CateID, oldExp.Amount, oldExp.Title, oldExp.Description, oldExp.Date, oldExp.ID)
	if err != nil {
		if err = transaction.Rollback(); err != nil {
			log.Println("Error rolling back transaction on the insertion", err)
		}
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (e *ExpensesRepository) GetTotalAmountByCate(cateID int) (model.CateTotalAmount, error) {
	var cateAmount model.CateTotalAmount

	if cateID == 0 {
		log.Fatal("Invalid category Id value")
	}

	DB := config.NewDatabase()
	query := "SELECT budget.category AS category, SUM(expenses.amount) AS total_amount, budget.amount AS budget_amount FROM expenses INNER JOIN categories ON expenses.cate_id = categories.id INNER JOIN budget ON categories.title = budget.category WHERE expenses.cate_id = ?"
	row := DB.QueryRow(query, cateID)

	if err := row.Scan(&cateAmount.Category, &cateAmount.TotalAmount, &cateAmount.BudgetAmount); err != nil {
		if err == sql.ErrNoRows {
			return cateAmount, err
		}
	}

	return cateAmount, nil
}
