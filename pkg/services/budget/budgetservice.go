package budgetservice

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	budgetRepo "Go-PersonalFinanceTracker/pkg/repository/budget"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type IBudgetServices interface {
	GetBudgetsList() ([]models.Budget, error)
}

type BudgetService struct {
	budgetRepo budgetRepo.IBudgetRepository
}

func (i *BudgetService) GetBudgetsList() {

}
