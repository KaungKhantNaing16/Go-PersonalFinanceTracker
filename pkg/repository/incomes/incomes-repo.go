package inrepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var ErrIncomeNotFound = errors.New("FromRepository - Income not found")

type IncomeRepository struct{}

// Retrieve the list of Incomes record from the database
func (i *IncomeRepository) GetIncomes() []model.Income {
	DB := config.NewDatabase()

	rows, err := DB.Query("SELECT * FROM incomes")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var incomes []model.Income
	for rows.Next() {
		var income model.Income
		err := rows.Scan(&income.ID, &income.UserID, &income.Title, &income.Amount, &income.Description, &income.FileURL, &income.CreatedAt, &income.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}

		incomes = append(incomes, income)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return incomes
}

func (i *IncomeRepository) GetIncomeById(id int) (model.Income, error) {
	fmt.Println(id)
	var income model.Income
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT * FROM incomes WHERE id=?", id)

	err := row.Scan(&income.ID, &income.UserID, &income.Title, &income.Amount, &income.Description, &income.FileURL, &income.CreatedAt, &income.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return income, ErrIncomeNotFound
		}
		return income, ErrIncomeNotFound
	}
	fmt.Print("In repository:")
	fmt.Println(income)
	return income, nil
}

func (i *IncomeRepository) CreateIncome(incomes []model.Income) error {
	DB := config.NewDatabase()
	insertQuery := "INSERT INTO incomes (uid, amount, title, description, file_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)"

	for _, income := range incomes {
		_, err := DB.Exec(insertQuery, income.UserID, income.Amount, income.Title, income.Description, income.FileURL, income.CreatedAt, income.UpdatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *IncomeRepository) UpdateIncome(income model.Income) error {
	oldIncome, err := i.GetIncomeById(income.ID)
	if err != nil {
		return ErrIncomeNotFound
	}

	oldIncome.UserID = income.UserID
	oldIncome.Amount = income.Amount
	oldIncome.Description = income.Description
	oldIncome.FileURL = income.FileURL

	DB := config.NewDatabase()
	updateQuery := "UPDATE incomes SET  uid=?, title=?, amount=?, description=?, file_url=? WHERE id=?"
	_, err = DB.Exec(updateQuery, oldIncome.UserID, oldIncome.Title, oldIncome.Amount, oldIncome.Description, oldIncome.FileURL, income.ID)
	if err != nil {
		return err
	}

	return nil
}

func (i *IncomeRepository) GetTotalAmount() (int, error) {
	var totalAmount int
	DB := config.NewDatabase()
	row := DB.QueryRow("SELECT SUM(amount) FROM incomes")

	if err := row.Scan(&totalAmount); err != nil {
		if err == sql.ErrNoRows {
			return totalAmount, err
		}
	}

	return totalAmount, nil
}
