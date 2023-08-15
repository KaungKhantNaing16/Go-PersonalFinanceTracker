package mediaservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	mediarepository "Go-PersonalFinanceTracker/pkg/repository/media"
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
)

type MediaService struct {
	mediaRepo mediarepository.MediaRepository
}

var expensesService = expservice.ExpensesService{}

func (m *MediaService) GetMedia() ([]model.Expenses, error) {
	expenses, err := expensesService.GetExpenses()
	if err != nil {
		return nil, err
	}

	return expenses, nil
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
