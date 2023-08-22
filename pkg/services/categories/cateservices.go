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

func (c *CategoriesService) CreateCategory(category string) error {
	return c.cateRepo.CreateCategory(category)
}

func (c *CategoriesService) GetCategories() ([]model.Category, error) {
	return c.cateRepo.GetCategories()
}

func (c *CategoriesService) GetCategoryById(id int) (model.Category, error) {
	return c.cateRepo.GetCategoryById(id)
}

func (c *CategoriesService) EditCategoryByTtl(category string) error {
	return c.cateRepo.EditCategoryByTtl(category)
}
