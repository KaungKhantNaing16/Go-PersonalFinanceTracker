package userrepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"errors"
	"time"
)

var ErrInvalidUserId = errors.New("FromRepository - Invalid User Id Value")
var ErrUserNotFound = errors.New("FromRepository - User Detail Not Found")

type UserRepository struct{}

func (u *UserRepository) CreateUserDetails(formData model.UserDetail) error {
	currentDateTime := time.Now()
	CreateAt := currentDateTime
	UpdatedAt := currentDateTime

	DB := config.NewDatabase()
	InsertQuery := "INSERT INTO user_details(status, name, email, password, profile_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err := DB.Exec(InsertQuery, formData.Status, formData.Name, formData.Email, formData.Password, formData.Profile, CreateAt, UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserDetailById(id int) (model.UserDetail, error) {
	if id == 0 {
		return model.UserDetail{}, ErrInvalidUserId
	}

	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM user_details WHERE id = ?", id)

	var userDetail model.UserDetail
	err := row.Scan(&userDetail.ID, &userDetail.Status, &userDetail.Name, &userDetail.Email, &userDetail.Password, &userDetail.Profile, &userDetail.CreatedAt, &userDetail.UpdatedAt, &userDetail.DeletedAt)
	if err != nil {
		return model.UserDetail{}, err
	}

	return userDetail, nil
}
