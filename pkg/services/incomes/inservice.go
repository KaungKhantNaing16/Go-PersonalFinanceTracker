package inservice

import (
	models "Go-PersonalFinanceTracker/pkg/models"
	inrepo "Go-PersonalFinanceTracker/pkg/repository/incomes"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type IIncomeServices interface {
	GetIncomes() ([]models.Income, error)
	GetIncomeById() (models.Income, error)
	CreateIncome() error
	UpdateIncome() error
	DeleteIncome() error
}

type IncomeService struct {
	expRepo inrepo.IIncomeRepository
}

func (i *IncomeService) GetIncomes() {

}

func (i *IncomeService) GetIncomeById() {

}

func (i *IncomeService) CreateIncome() {

}

func (i *IncomeService) UpdateIncome() {

}

func (i *IncomeService) DeleteIncome() {

}
