package inrepository

import (
	"Go-PersonalFinanceTracker/config"
	model "Go-PersonalFinanceTracker/pkg/models"
	"context"
	"database/sql"
	"errors"
	"log"
)

var ErrIncomeNotFound = errors.New("FromRepository - Income not found")

type IncomeRepository struct{}

func (i *IncomeRepository) GetIncomes(id int) []model.Income {
	DB := config.NewDatabase()
	rows, err := DB.Query("SELECT * FROM incomes WHERE uid = ?", id)
	if err != nil {
		log.Println("Error querying the row", err)
	}
	defer rows.Close()

	var incomes []model.Income
	for rows.Next() {
		var income model.Income
		err := rows.Scan(&income.ID, &income.UserID, &income.Title, &income.Amount, &income.Description, &income.FileURL, &income.CreatedAt, &income.UpdatedAt)
		if err != nil {
			log.Println("Error scaning the row", err)
		}

		incomes = append(incomes, income)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return incomes
}

func (i *IncomeRepository) GetIncomeById(id int) (model.Income, error) {
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
	return income, nil
}

func (i *IncomeRepository) CreateIncome(incomes []model.Income) error {
	DB := config.NewDatabase()
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	insertQuery, err := transaction.PrepareContext(ctx, "INSERT INTO incomes (uid, amount, title, description, file_url, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)")

	for _, income := range incomes {
		_, err := insertQuery.Exec(income.UserID, income.Amount, income.Title, income.Description, income.FileURL, income.CreatedAt, income.UpdatedAt)
		if err != nil {
			if err = transaction.Rollback(); err != nil {
				log.Println("Error rolling back transaction on the insertion", err)
			}
		}
	}

	if err = transaction.Commit(); err != nil {
		return err
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
	ctx := context.Background()

	transaction, err := DB.BeginTx(ctx, nil)
	if err != nil {
		log.Panicln("Error beginning the transaction", err)
	}

	updateQuery, err := transaction.PrepareContext(ctx, "UPDATE incomes SET  uid=?, title=?, amount=?, description=?, file_url=? WHERE id=?")
	_, err = updateQuery.Exec(oldIncome.UserID, oldIncome.Title, oldIncome.Amount, oldIncome.Description, oldIncome.FileURL, income.ID)
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

func (i *IncomeRepository) GetAmountByDay(id int) []model.DailyAmount {
	DB := config.NewDatabase()
	query := `SELECT
				CASE
					WHEN dayname(created_at) = "Sunday" THEN "Sunday" 
					WHEN dayname(created_at) = "Monday" THEN "Monday" 
					WHEN dayname(created_at) = "Tuesday" THEN "Tuesday" 
					WHEN dayname(created_at) = "Wednesday" THEN "Wednesday"
					WHEN dayname(created_at) = "Thursday" THEN "Thursday" 
					WHEN dayname(created_at) = "Friday" THEN "Friday" 
					WHEN dayname(created_at) = "Saturday" THEN "Saturday"
				END AS DayName,
				SUM(CASE
					WHEN dayname(created_at) = "Sunday" THEN amount 
					WHEN dayname(created_at) = "Monday" THEN amount
					WHEN dayname(created_at) = "Tuesday" THEN amount
					WHEN dayname(created_at) = "Wednesday" THEN amount
					WHEN dayname(created_at) = "Thursday" THEN amount
					WHEN dayname(created_at) = "Friday" THEN amount
					WHEN dayname(created_at) = "Saturday" THEN amount
				END) AS Amount
			FROM incomes
			WHERE week(created_at) = week(now())
			AND uid = ?
			group by DayName;`

	rows, err := DB.Query(query, id)
	if err != nil {
		log.Println("Error querying the rows", err)
	}
	defer rows.Close()

	chartDataArr := []model.DailyAmount{}
	for rows.Next() {
		var chartData model.DailyAmount
		err = rows.Scan(&chartData.Day, &chartData.Amount)
		if err != nil {
			log.Println("Error scaning the row", err)
		}

		chartDataArr = append(chartDataArr, chartData)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return chartDataArr
}
