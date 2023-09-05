package dashboardservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	budgetservice "Go-PersonalFinanceTracker/pkg/services/budget"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	inservice "Go-PersonalFinanceTracker/pkg/services/incomes"
	userservice "Go-PersonalFinanceTracker/pkg/services/users"
)

var detailService = userservice.UserDetailService{}
var incomeService = inservice.IncomeService{}
var expensesService = expservice.ExpensesService{}
var budgetPlanService = budgetservice.BudgetService{}

func DashboardService(AuthorizeID int) (model.OAData, error) {
	userProfileData := model.OAData{}
	totalExp, err := detailService.GetExpAmtByUserId(AuthorizeID)
	if err != nil {
		return userProfileData, err
	}

	totalIncomes, err := detailService.GetIncomesAmtByUserId(AuthorizeID)
	if err != nil {
		return userProfileData, err
	}

	totalAmount := model.TotalAmountData{
		Expenses: totalExp,
		Incomes:  totalIncomes,
	}

	userDetail, err := detailService.GetUserDetailByID(AuthorizeID)
	if err != nil {
		return userProfileData, err
	}

	ie := ChartService(AuthorizeID)

	userProfileData.TotalAmount = totalAmount
	userProfileData.UserDetail = userDetail
	userProfileData.IEB = ie

	return userProfileData, nil
}

func ChartService(AuthorizeID int) model.IEBData {
	incomeData := incomeService.GetIncomes(AuthorizeID)
	expData, _ := expensesService.GetExpenses(AuthorizeID)
	planData, _ := budgetPlanService.GetBudgetsList(AuthorizeID)

	ie := model.IEBData{
		Incomes:  &incomeData,
		Expenses: &expData,
		Budget:   &planData,
	}

	return ie
}
