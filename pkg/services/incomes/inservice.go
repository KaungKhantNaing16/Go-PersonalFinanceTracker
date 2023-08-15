package inservice

import (
	model "Go-PersonalFinanceTracker/pkg/models"
	inrepository "Go-PersonalFinanceTracker/pkg/repository/incomes"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var ErrIDIsNotValid = errors.New("Id is not valid")
var ErrIncomeNotFound = errors.New("Income not found")

type IncomeService struct {
	repository inrepository.IncomeRepository
}

func (s *IncomeService) GetIncomes() []model.Income {
	return s.repository.GetIncomes()
}

func (s *IncomeService) GetIncomeById(request *http.Request) (model.Income, error) {
	params := mux.Vars(request)
	incomeId, err := strconv.Atoi(params["id"])
	fmt.Println("Income Id is :", incomeId)
	if err != nil {
		return model.Income{}, ErrIDIsNotValid
	}

	income, err := s.repository.GetIncomeById(incomeId)
	if err != nil {
		return income, ErrIncomeNotFound
	}

	return income, nil
}

func (s *IncomeService) CreateIncomes(incomes []model.Income) error {
	return s.repository.CreateIncome(incomes)
}

func (s *IncomeService) UpdateIncome(income model.Income) error {
	if income.ID == 0 {
		return ErrIDIsNotValid
	}

	return s.repository.UpdateIncome(income)
}
