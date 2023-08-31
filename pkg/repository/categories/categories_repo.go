package cateRepositroy

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"context"
	"database/sql"
	"errors"
	"log"
)

var ErrInvalidUserID = errors.New("FromRepository - Invalid ID Value")
var ErrCateNotFound = errors.New("FromRepository - Category not found")

type CategoriesRepository struct{}

func (c *CategoriesRepository) CreateCategory(bid int, category string) error {
	defDesc := "これは、カテゴリのデフォルトの説明です。"
	ctx := context.Background()
	DB := config.NewDatabase()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	insertQuery, err := transaction.PrepareContext(ctx, "INSERT INTO categories (bid, title, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	_, err = insertQuery.Exec(bid, category, defDesc, config.CreateAt, config.UpdatedAt)
	if err != nil {
		if err = transaction.Rollback(); err != nil {
			log.Println("Error rolling back transaction on the insertion", err)
		}
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func (c *CategoriesRepository) GetCategories(userId int) ([]model.Category, error) {
	if userId == 0 {
		return []model.Category{}, ErrInvalidUserID
	}

	DB := config.NewDatabase()
	query := "SELECT categories.* FROM categories LEFT JOIN budget ON categories.bid = budget.id WHERE budget.uid = ?"
	rows, err := DB.Query(query, userId)
	if err != nil {
		log.Println("Error querying the row", err)
	}
	defer rows.Close()

	var categories []model.Category
	for rows.Next() {
		var category model.Category
		err := rows.Scan(&category.ID, &category.BID, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)
		if err != nil {
			log.Println("Error scaning the row", err)
		}
		categories = append(categories, category)

	}
	if err := rows.Err(); err != nil {
		return []model.Category{}, err
	}

	return categories, nil
}

func (c *CategoriesRepository) GetCategoryById(id int) (model.Category, error) {
	var category model.Category
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM categories WHERE id = ?", id)

	err := row.Scan(&category.ID, &category.BID, &category.Title, &category.Description, &category.CreatedAt, &category.UpdatedAt, &category.DeletedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return model.Category{}, ErrCateNotFound
		}
		return model.Category{}, ErrCateNotFound
	}

	return category, nil
}

func (c *CategoriesRepository) EditCategoryByTtl(bid int) error {
	DB := config.NewDatabase()
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	updateQuery, err := transaction.PrepareContext(ctx, "UPDATE categories SET deleted_at=? WHERE bid=?")
	_, err = updateQuery.Exec(config.Deletedat, bid)
	if err != nil {
		if err = transaction.Rollback(); err != nil {
			log.Println("Error rolling back transaction on the updating", err)
		}
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}
