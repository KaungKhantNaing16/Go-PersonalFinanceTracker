package cateRepositroy

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"log"
	"time"
)

type CategoriesRepository struct{}

func (c *CategoriesRepository) CreateCategory(category string) error {
	defStatus := 1
	defDesc := "これは、カテゴリのデフォルトの説明です。"
	currentDateTime := time.Now()
	CreateAt := currentDateTime
	UpdatedAt := currentDateTime

	DB := config.NewDatabase()
	query := "INSERT INTO categories (status, title, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	_, err := DB.Exec(query, defStatus, category, defDesc, CreateAt, UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoriesRepository) GetCategories() ([]model.Category, error) {
	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT * FROM categories WHERE status = ?", 1)
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

func (c *CategoriesRepository) GetCategoryById(id int) (model.Category, error) {
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM categories WHERE id = ?", id)
	var category model.Category
	err := row.Scan(&category.ID, &category.Status, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)
	if err != nil {
		return model.Category{}, err
	}

	return category, nil
}

func (c *CategoriesRepository) EditCategoryByTtl(category string) error {
	DB := config.NewDatabase()
	query := "UPDATE categories SET  status=? WHERE title=?"
	_, err := DB.Exec(query, 0, category)
	if err != nil {
		return err
	}

	return nil
}
