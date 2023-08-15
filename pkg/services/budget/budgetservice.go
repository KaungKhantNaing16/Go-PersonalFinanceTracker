package budgetservice

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	budgetRepo "Go-PersonalFinanceTracker/pkg/repository/budget"
	cateservice "Go-PersonalFinanceTracker/pkg/services/categories"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")
var categoriesService = cateservice.CategoriesService{}

type BudgetService struct {
	budgetRepo budgetRepo.BudgetRepository
}

func (i *BudgetService) GetBudgetsList() []models.Budget {
	return i.budgetRepo.GetBudgetsList()
}

func (i *BudgetService) CreateBudgetPlan(budgetPlan models.Budget) error {
	if err := categoriesService.CreateCategory(budgetPlan.Category); err != nil {
		return err
	}

	return i.budgetRepo.CreateBudgetPlan(budgetPlan)
}

func (i *BudgetService) DeleteBudgetPlan(id int) error {
	return i.budgetRepo.DeleteBudgetPlan(id)
}
