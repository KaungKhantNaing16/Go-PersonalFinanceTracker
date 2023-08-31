package mediaRepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"context"
	"errors"
	"log"
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
		log.Println("Error querying the row", err)
	}

	defer rows.Close()
	var mediaDataArr []model.ExpenseMediaData
	for rows.Next() {
		var mediaData model.ExpenseMediaData
		err = rows.Scan(&mediaData.ID, &mediaData.UserID, &mediaData.CateID, &mediaData.Amount, &mediaData.Title, &mediaData.Description, &mediaData.Date, &mediaData.CreateAt, &mediaData.UpdatedAt, &mediaData.ImgURL)
		if err != nil {
			log.Println("Error scaning the row", err)
		}
		mediaDataArr = append(mediaDataArr, mediaData)
	}

	if err := rows.Err(); err != nil {
		return nil, err
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
		log.Println("Error querying the row", err)
	}

	defer rows.Close()
	var mediaArray []string
	for rows.Next() {
		var mediaName string
		err = rows.Scan(&mediaName)
		if err != nil {
			log.Println("Error scaning the row", err)
		}
		mediaArray = append(mediaArray, mediaName)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return mediaArray, nil
}

func (m *MediaRepository) CreateMedia(fileNameArr []string, id int) error {
	if id == 0 {
		return ErrExpIDIsNotValid
	}

	DB := config.NewDatabase()
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	insertQuery, err := transaction.PrepareContext(ctx, "INSERT INTO slips (exp_id, filename, created_at, updated_at) VALUES (?, ?, ?, ?)")

	for _, filename := range fileNameArr {
		_, err := insertQuery.Exec(id, filename, config.CreateAt, config.UpdatedAt)
		if err != nil {
			if err = transaction.Rollback(); err != nil {
				log.Println("Error rolling back transaction on the insertion", err)
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return err
	}
	return nil
}
