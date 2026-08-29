package repository

import (
	"github.com/Kotoninja/expense-tracker/pkg/expense"
)

type ExpenseRepository interface {
	Add(exp *expense.Expense) (int, error)
	Delete(id int) error
	Update(id int, exp *expense.Expense) error
	FindByID(id int) (*expense.Expense, error)
	FindAll() ([]expense.Expense, error)
	GetTotal(month int) int
}
