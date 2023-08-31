package budgetRepo

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"context"
	"errors"
	"fmt"
	"log"
)

var ErrInvalidUserID = errors.New("FromRepository - Invalid ID Value")
var ErrBudgetNotFound = errors.New("FromRepository - budget not found")

type BudgetRepository struct{}

func (b *BudgetRepository) GetBudgetsList(id int) ([]model.Budget, error) {
	if id == 0 {
		return []model.Budget{}, ErrInvalidUserID
	}

	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT * FROM budget WHERE uid = ?", id)
	if err != nil {
		log.Println("Error querying the row", err)
	}

	defer rows.Close()
	var budgetPlans []model.Budget
	for rows.Next() {
		var budget model.Budget
		err = rows.Scan(&budget.ID, &budget.UserID, &budget.Title, &budget.Category, &budget.Amount, &budget.CreatedAt, &budget.UpdatedAt, &budget.DeletedAt)
		if err != nil {
			log.Println("Error scaning the row", err)
		}
		budgetPlans = append(budgetPlans, budget)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return budgetPlans, nil
}

func (b *BudgetRepository) CreateBudgetPlan(budgetPlan model.Budget) (model.Budget, error) {
	DB := config.NewDatabase()
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	insertQuery, err := transaction.PrepareContext(ctx, "INSERT INTO budget (uid, title, category, amount, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)")
	_, err = insertQuery.Exec(budgetPlan.UserID, budgetPlan.Title, budgetPlan.Category, budgetPlan.Amount, config.CreateAt, config.UpdatedAt)
	if err != nil {
		if err = transaction.Rollback(); err != nil {
			log.Println("Error rolling back transaction on the insertion", err)
		}
	}

	if err = transaction.Commit(); err != nil {
		return model.Budget{}, err
	}

	row := DB.QueryRow("SELECT budget.* FROM budget ORDER BY budget.id DESC LIMIT 1;")
	var budget model.Budget
	err = row.Scan(&budget.ID, &budget.UserID, &budget.Title, &budget.Category, &budget.Amount, &budget.CreatedAt, &budget.UpdatedAt, &budget.DeletedAt)
	if err != nil {
		log.Println("Error when scaning budget data", err)
	}

	return budget, nil
}

func (b *BudgetRepository) GetBudgetPlanById(id int) (model.Budget, error) {
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM budget WHERE id = ?", id)
	var budget model.Budget
	err := row.Scan(&budget.ID, &budget.UserID, &budget.Title, &budget.Category, &budget.Amount, &budget.CreatedAt, &budget.UpdatedAt, &budget.DeletedAt)
	if err != nil {
		return model.Budget{}, err
	}

	return budget, nil
}

func (b *BudgetRepository) DeleteBudgetPlan(id int) error {
	DB := config.NewDatabase()
	query := "DELETE FROM budget WHERE id = ?"
	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	fmt.Println("Delete budget plan successfully")
	return nil
}
