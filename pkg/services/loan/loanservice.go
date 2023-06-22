package loanservice

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	loanRepo "Go-PersonalFinanceTracker/pkg/repository/loan"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type ILoanServices interface {
	GetGiveLoan() ([]models.Loan, error)
	GetReceiveLoan() (models.Loan, error)
}

type LoanService struct {
	expRepo loanRepo.ILoanRepository
}

func (i *LoanService) GetGiveLoan() {

}

func (i *LoanService) GetReceiveLoan() {

}
