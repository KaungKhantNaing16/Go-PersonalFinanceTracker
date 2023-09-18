package cateservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	cateRepositroy "Go-PersonalFinanceTracker/pkg/repository/categories"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type CategoriesService struct {
	cateRepo cateRepositroy.CategoriesRepository
}

func (c *CategoriesService) CreateCategory(bid int, category string) error {
	return c.cateRepo.CreateCategory(bid, category)
}

func (c *CategoriesService) GetCategories(userId int) ([]model.Category, error) {
	return c.cateRepo.GetCategories(userId)
}

func (c *CategoriesService) GetCategoryById(id int) (model.Category, error) {
	return c.cateRepo.GetCategoryById(id)
}

func (c *CategoriesService) EditCategoryByTtl(bid int) error {
	return c.cateRepo.EditCategoryByTtl(bid)
}
