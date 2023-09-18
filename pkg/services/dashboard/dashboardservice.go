package dashboardservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	budgetservice "Go-PersonalFinanceTracker/pkg/services/budget"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	inservice "Go-PersonalFinanceTracker/pkg/services/incomes"
	mediaservice "Go-PersonalFinanceTracker/pkg/services/media"
	userservice "Go-PersonalFinanceTracker/pkg/services/users"
)

var detailService = userservice.UserDetailService{}
var incomeService = inservice.IncomeService{}
var expensesService = expservice.ExpensesService{}
var budgetPlanService = budgetservice.BudgetService{}
var mediaService = mediaservice.MediaService{}

// var LineChartData = map[string]model.TotalAmountData{}
type OAData struct {
	TotalAmount model.TotalAmountData
	UserDetail  model.UserDetail
	BI          model.BIData
}

func DashboardService(AuthorizeID int) (OAData, error) {
	userProfileData := OAData{}
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

	bi := BIDataService(AuthorizeID)

	userProfileData.TotalAmount = totalAmount
	userProfileData.UserDetail = userDetail
	userProfileData.BI = bi

	return userProfileData, nil
}

func BIDataService(AuthorizeID int) model.BIData {
	planData, _ := budgetPlanService.GetBudgetsList(AuthorizeID)
	mediaData, _ := mediaService.GetMedia(AuthorizeID)

	var imgURLs []string
	for _, media := range mediaData.Media {
		imgURLs = append(imgURLs, media.ImgURL)
	}

	bi := model.BIData{
		Budget:   &planData,
		ImageSrc: imgURLs[0],
	}
	return bi
}

func ChartDataService(AuthorizeID int) map[string]model.TotalAmountData {
	incomes := incomeService.GetAmountByDay(AuthorizeID)
	expenses := expensesService.GetAmountByDay(AuthorizeID)

	var LineChartData = map[string]model.TotalAmountData{}
	LineChartData["Monday"] = model.TotalAmountData{}
	LineChartData["Tuesday"] = model.TotalAmountData{}
	LineChartData["Wednesday"] = model.TotalAmountData{}
	LineChartData["Thursday"] = model.TotalAmountData{}
	LineChartData["Friday"] = model.TotalAmountData{}
	LineChartData["Saturday"] = model.TotalAmountData{}
	LineChartData["Sunday"] = model.TotalAmountData{}

	for _, income := range incomes {
		switch income.Day {
		case "Monday":
			LineChartData["Monday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		case "Tuesday":
			LineChartData["Tuesday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		case "Wednesday":
			LineChartData["Wednesday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		case "Thursday":
			LineChartData["Thursday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		case "Friday":
			LineChartData["Friday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		case "Saturday":
			LineChartData["Saturday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		case "Sunday":
			LineChartData["Sunday"] = model.TotalAmountData{
				Incomes: income.Amount,
			}
		}
	}

	for _, expense := range expenses {
		switch expense.Day {
		case "Monday":
			LineChartData["Monday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		case "Tuesday":
			LineChartData["Tuesday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		case "Wednesday":
			LineChartData["Wednesday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		case "Thursday":
			LineChartData["Thursday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		case "Friday":
			LineChartData["Friday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		case "Saturday":
			LineChartData["Saturday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		case "Sunday":
			LineChartData["Sunday"] = model.TotalAmountData{
				Expenses: expense.Amount,
			}
		}
	}
	return LineChartData
}
