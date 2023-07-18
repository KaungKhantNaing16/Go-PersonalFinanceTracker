package inservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	inrepository "Go-PersonalFinanceTracker/pkg/repository/incomes"
	"errors"
)

var ErrIDIsNotValid = errors.New("Id is not valid")

type IncomeService struct {
	repository inrepository.IncomeRepository
}

func (s *IncomeService) GetIncomes() []model.Income {
	return s.repository.GetIncomes()
}

func (s *IncomeService) GetIncomeById(id int) {

}

func (s *IncomeService) CreateIncome(income model.Income) {

}

func (s *IncomeService) UpdateIncome(id int, income model.Income) {

}

func (s *IncomeService) DeleteIncome(id int) {

}
