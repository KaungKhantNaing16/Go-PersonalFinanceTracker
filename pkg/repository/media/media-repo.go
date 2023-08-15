package mediaRepository

import (
	"Go-PersonalFinanceTracker/config"
	"errors"
	"fmt"
	"log"
	"time"
)

type MediaRepository struct{}

var ErrIDIsNotValid = errors.New("Id is not valid")

func (m *MediaRepository) GetMediaByExpId(id int) ([]string, error) {
	if id == 0 {
		return nil, ErrIDIsNotValid
	}
	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT filename FROM slips WHERE exp_id = ?", id)
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
		return ErrIDIsNotValid
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
