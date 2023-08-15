package budgetRepo

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"
)

var ErrBudgetNotFound = errors.New("FromRepository - budget not found")

type BudgetRepository struct{}

func (b *BudgetRepository) GetBudgetsList() []model.Budget {
	DB := config.NewDatabase()

	rows, err := DB.Query("SELECT * FROM budget")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var budgetPlans []model.Budget
	for rows.Next() {
		var budget model.Budget
		err = rows.Scan(&budget.ID, &budget.Title, &budget.Category, &budget.Amount, &budget.CreatedAt, &budget.UpdatedAt, &budget.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}

		budgetPlans = append(budgetPlans, budget)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return budgetPlans
}

func (b *BudgetRepository) CreateBudgetPlan(budgetPlan model.Budget) error {
	currentDateTime := time.Now()
	CreateAt := currentDateTime
	UpdatedAt := currentDateTime

	DB := config.NewDatabase()
	_, err := DB.Query("INSERT INTO budget (title, category, amount, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", budgetPlan.Title, budgetPlan.Category, budgetPlan.Amount, CreateAt, UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (b *BudgetRepository) DeleteBudgetPlan(id int) error {
	DB := config.NewDatabase()
	query := "DELETE FROM budget WHERE id = ?"
	_, err := DB.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Delete budget plan successfully")
	return nil
}
