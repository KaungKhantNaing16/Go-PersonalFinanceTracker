package mediaservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	mediarepository "Go-PersonalFinanceTracker/pkg/repository/media"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
)

var expensesService = expservice.ExpensesService{}

type MediaService struct {
	mediaRepo mediarepository.MediaRepository
}

type MediaData struct {
	Expenses []model.Expenses
	Media    []model.ExpenseMediaData
}

func (m *MediaService) GetMedia(userID int) (MediaData, error) {
	var mediaDataArr MediaData
	expenses, err := expensesService.GetExpenses(userID)
	if err != nil {
		return mediaDataArr, err
	}

	mData, err := m.mediaRepo.GetMedia(userID)
	if err != nil {
		return mediaDataArr, err
	}

	mediaDataArr.Expenses = expenses
	mediaDataArr.Media = mData

	return mediaDataArr, nil
}

func (m *MediaService) GetMediaByExpId(id int) ([]string, error) {
	mediaArr, err := m.mediaRepo.GetMediaByExpId(id)
	if err != nil {
		return nil, err
	}

	return mediaArr, nil
}

func (m *MediaService) CreateMedia(fileNameArr []string, id int) error {
	return m.mediaRepo.CreateMedia(fileNameArr, id)
}
