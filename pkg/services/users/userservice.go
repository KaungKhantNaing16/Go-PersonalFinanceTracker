package userservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	userrepository "Go-PersonalFinanceTracker/pkg/repository/users"
)

type UserDetailService struct {
	detailRepo userrepository.UserRepository
}

func (d *UserDetailService) CreateUserDetails(formData model.UserDetail) error {
	return d.detailRepo.CreateUserDetails(formData)
}

func (d *UserDetailService) GetUserDetailByEmail(email string) (model.UserDetail, error) {
	return d.detailRepo.GetUserDetailByEmail(email)
}

func (d *UserDetailService) GetUserDetailByID(id int) (model.UserDetail, error) {
	return d.detailRepo.GetUserDetailByID(id)
}

func (d *UserDetailService) GetExpAmtByUserId(id int) (int, error) {
	return d.detailRepo.GetExpAmtByUserId(id)
}

func (d *UserDetailService) GetIncomesAmtByUserId(id int) (int, error) {
	return d.detailRepo.GetIncomesAmtByUserId(id)
}
