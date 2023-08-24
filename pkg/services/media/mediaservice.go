package mediaservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	mediarepository "Go-PersonalFinanceTracker/pkg/repository/media"
)

type MediaService struct {
	mediaRepo mediarepository.MediaRepository
}

func (m *MediaService) GetMedia(userID int) ([]model.ExpenseMediaData, error) {
	mediaDataArr, err := m.mediaRepo.GetMedia(userID)
	if err != nil {
		return nil, err
	}

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
