package userrepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"database/sql"
	"errors"
	"time"
)

var ErrInvalidUserID = errors.New("FromRepository - Invalid ID Value")
var ErrInvalidEmail = errors.New("FromRepository - Invalid Email Value")
var ErrUserNotFound = errors.New("FromRepository - User Detail Not Found")

type UserRepository struct{}

func (u *UserRepository) CreateUserDetails(formData model.UserDetail) error {
	currentDateTime := time.Now()
	CreateAt := currentDateTime
	UpdatedAt := currentDateTime

	DB := config.NewDatabase()
	InsertQuery := "INSERT INTO user_details(status, name, email, job, password, profile_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := DB.Exec(InsertQuery, formData.Status, formData.Name, formData.Email, formData.Job, formData.Password, formData.Profile, CreateAt, UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) GetUserDetailByID(id int) (model.UserDetail, error) {
	if id == 0 {
		return model.UserDetail{}, ErrInvalidUserID
	}
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM user_details WHERE id = ?", id)

	var userDetail model.UserDetail
	err := row.Scan(&userDetail.ID, &userDetail.Status, &userDetail.Name, &userDetail.Email, &userDetail.Job, &userDetail.Password, &userDetail.Profile, &userDetail.CreatedAt, &userDetail.UpdatedAt, &userDetail.DeletedAt)
	if err != nil {
		return model.UserDetail{}, err
	}

	return userDetail, nil
}

func (u *UserRepository) GetUserDetailByEmail(email string) (model.UserDetail, error) {
	if email == "" {
		return model.UserDetail{}, ErrInvalidEmail
	}

	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM user_details WHERE email = ?", email)

	var userDetail model.UserDetail
	err := row.Scan(&userDetail.ID, &userDetail.Status, &userDetail.Name, &userDetail.Email, &userDetail.Job, &userDetail.Password, &userDetail.Profile, &userDetail.CreatedAt, &userDetail.UpdatedAt, &userDetail.DeletedAt)
	if err != nil {
		return model.UserDetail{}, err
	}

	return userDetail, nil
}

func (u *UserRepository) GetExpAmtByUserId(userId int) (int, error) {
	var totalAmount int
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT SUM(amount) FROM expenses WHERE uid = ?", userId)

	if err := row.Scan(&totalAmount); err != nil {
		if err == sql.ErrNoRows {
			return totalAmount, err
		}
	}

	return totalAmount, nil
}

func (u *UserRepository) GetIncomesAmtByUserId(userId int) (int, error) {
	var totalAmount int
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT SUM(amount) FROM incomes WHERE uid = ?", userId)

	if err := row.Scan(&totalAmount); err != nil {
		if err == sql.ErrNoRows {
			return totalAmount, err
		}
	}

	return totalAmount, nil
}
