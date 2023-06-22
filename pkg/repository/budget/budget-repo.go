package budgetRepo

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	"errors"
)

var ErrBudgetNotFound = errors.New("FromRepository - budget not found")

type IBudgetRepository interface {
	GetBudgetsList() ([]models.Budget, error)
}

type BudgetRepository struct {
	Budget []models.Budget
}

func (b *BudgetRepository) GetBudgetsList() {
	// Retrieve the list of Budget record from the database
}
