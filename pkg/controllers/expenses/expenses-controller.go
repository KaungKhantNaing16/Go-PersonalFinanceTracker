package expcontroller

import (
	expservice "Go-PersonalFinanceTracker/pkg/services/expenses"
	"net/http"
)

type expensesHandler struct {
	service expservice.IExpensesServices
}

func (eh *expensesHandler) GetExpenses(writer http.ResponseWriter, request *http.Request) {

}

func (eh *expensesHandler) GetExpensesById(writer http.ResponseWriter, request *http.Request) {

}

func (eh *expensesHandler) CreateExpenses(writer http.ResponseWriter, request *http.Request) {

}

func (eh *expensesHandler) UpdateExpenses(writer http.ResponseWriter, request *http.Request) {

}

func (eh *expensesHandler) DeleteExpenses(writer http.ResponseWriter, request *http.Request) {

}
