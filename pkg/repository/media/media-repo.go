package mediaRepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"errors"
	"fmt"
	"log"
	"time"
)

type MediaRepository struct{}

var ErrUserIDIsNotValid = errors.New("User Id is not valid")
var ErrExpIDIsNotValid = errors.New("Exp Id is not valid")

func (m *MediaRepository) GetMedia(userId int) ([]model.ExpenseMediaData, error) {
	if userId == 0 {
		return nil, ErrUserIDIsNotValid
	}
	DB := config.NewDatabase()
	query := "SELECT expenses.*, slips.filename AS image_url FROM expenses JOIN slips ON expenses.id = slips.exp_id WHERE expenses.uid = ?"
	rows, err := DB.Query(query, userId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var mediaDataArr []model.ExpenseMediaData
	for rows.Next() {
		var mediaData model.ExpenseMediaData
		err = rows.Scan(&mediaData.ID, &mediaData.UserID, &mediaData.CateID, &mediaData.Amount, &mediaData.Title, &mediaData.Description, &mediaData.Date, &mediaData.CreateAt, &mediaData.UpdatedAt, &mediaData.ImgURL)
		if err != nil {
			log.Fatal(err)
		}
		mediaDataArr = append(mediaDataArr, mediaData)
	}

	return mediaDataArr, nil
}

func (m *MediaRepository) GetMediaByExpId(expId int) ([]string, error) {
	if expId == 0 {
		return nil, ErrExpIDIsNotValid
	}
	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT filename FROM slips WHERE exp_id = ?", expId)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var mediaArray []string
	for rows.Next() {
		var mediaName string
		err = rows.Scan(&mediaName)
		if err != nil {
			log.Fatal(err)
		}
		mediaArray = append(mediaArray, mediaName)
	}

	return mediaArray, nil
}

func (m *MediaRepository) CreateMedia(fileNameArr []string, id int) error {
	if id == 0 {
		return ErrExpIDIsNotValid
	}

	currentDateTime := time.Now()
	CreateAt := currentDateTime
	UpdatedAt := currentDateTime

	DB := config.NewDatabase()
	query := "INSERT INTO slips (exp_id, filename, created_at, updated_at) VALUES (?, ?, ?, ?)"

	for _, filename := range fileNameArr {
		_, err := DB.Exec(query, id, filename, CreateAt, UpdatedAt)
		if err != nil {
			return err
		}
	}

	fmt.Println("Ceate Media successful.")
	return nil
}
