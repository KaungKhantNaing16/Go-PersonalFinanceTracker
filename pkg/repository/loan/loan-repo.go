package loanRepo

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	"errors"
)

var ErrLoanNotFound = errors.New("FromRepository - loan not found")

type ILoanRepository interface {
	GetGiveLoan() ([]models.Loan, error)
	GetReceiveLoan(id int) (models.Loan, error)
}

type LoanRepository struct {
	Loan []models.Loan
}

func (l *LoanRepository) GetGiveLoan() {
	// Retrieve the list of loan record from the database
}

func (l *LoanRepository) GetReceiveLoan() {
	// Retrieve loan record from the database
}
