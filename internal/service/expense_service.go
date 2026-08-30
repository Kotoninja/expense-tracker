package service

import (
	"errors"
	"sort"
	"time"

	"github.com/Kotoninja/expense-tracker/internal/repository"
	"github.com/Kotoninja/expense-tracker/pkg/expense"
)

type ExpenseService struct {
	repo repository.ExpenseRepository
}

func NewExpenseService(repo repository.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) AddExpense(description string, amount int) (int, error) {
	if description == "" {
		return 0, errors.New("description cannot be empty")
	}
	if amount <= 0 {
		return 0, errors.New("amount must be positive")
	}

	exp := expense.NewExpense(description, amount)
	return s.repo.Add(exp)
}

func (s *ExpenseService) DeleteExpense(id int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}
	return s.repo.Delete(id)
}

func (s *ExpenseService) UpdateExpense(id int, description string, amount int) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	existing, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if description != "" {
		existing.Description = description
	}
	if amount > 0 {
		existing.Amount = amount
	}
	existing.Date = time.Now()

	return s.repo.Update(id, existing)
}

func (s *ExpenseService) ListExpenses() []ExpenseDTO {
	expenses, err := s.repo.FindAll()
	if err != nil {
		return []ExpenseDTO{}
	}

	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].ID < expenses[j].ID
	})

	dtos := make([]ExpenseDTO, len(expenses))
	for i, exp := range expenses {
		dtos[i] = ExpenseDTO{
			ID:          exp.ID,
			Description: exp.Description,
			Amount:      exp.Amount,
			Date:        exp.Date.Format(time.DateOnly),
		}
	}

	return dtos
}

func (s *ExpenseService) GetSummary(month int) int {
	return s.repo.GetTotal(month)
}

type ExpenseDTO struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
	Date        string `json:"date"`
}
