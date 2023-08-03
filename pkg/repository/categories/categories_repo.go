package cateRepositroy

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"log"
)

type CategoriesRepository struct{}

func (c *CategoriesRepository) GetCategories() ([]model.Category, error) {
	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT * FROM categories")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.Status, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)
		if err != nil {
			log.Fatal(err)
		}
		categories = append(categories, category)

	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return categories, err
}
