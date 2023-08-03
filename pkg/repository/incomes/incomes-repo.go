package inrepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"errors"
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

func (i *IncomeRepository) GetIncomeById(id int) {
	// Retrieve Income record from the database
}

func (i *IncomeRepository) CreateIncome(income model.Income) {
	// Insert the Income record into the database
}

func (i *IncomeRepository) UpdateIncome(id int, income model.Income) {
	// Update the Income record from the database
}

func (i *IncomeRepository) DeleteIncome(id int) {
	// Delete the Income record from the database
}
