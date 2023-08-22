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

func (d *UserDetailService) GetUserDetailById(id int) (model.UserDetail, error) {
	return d.detailRepo.GetUserDetailById(id)
}
